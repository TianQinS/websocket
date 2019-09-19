# -*- coding: utf8 -*-
u"""模块基类."""


import gevent
# import gevent.monkey
# gevent.monkey.patch_all()
import traceback
from gevent import Greenlet
from gevent.pool import Group
from gevent.event import Event


class Module(Greenlet):
	
	def __init__(self, topic=""):
		Greenlet.__init__(self)
		self._stop_evt = Event()
		self.app = None
		self.closeSig = False
		self.topic = topic
		self.settings = None

	def _run(self):
		try:
			self.Run()
		except Exception:
			print(traceback.format_exc())
		self._stop_evt.set()

	def stop(self):
		self.closeSig = True
		self.kill()
	
	def Run(self):
		pass

	def sleep(self,seconds=0):
		gevent.sleep(seconds)
	
	def onInit(self, app, settings):
		self.app = app
		self.settings = settings
	
	def onDestroy(self):
		self.app = None


class ModuleManager(object):
		
	def __init__(self):
		self.mods = {}
		self.group = Group()
		self.cmd_tables = {}
		self._tasks = []

	def __getitem__(self, topic):
		u"""获取某个模块对象."""
		return self.mods.get(topic, None)

	def _close(self):
		u"""清理异步任务."""
		gevent.killall(self._tasks, block=True)
		self._tasks = []

	def spawn(self ,f):
		u"""加入异步任务."""
		task = gevent.spawn(f)
		self._tasks.append(task)
		task.join()

	def register(self, f, func):
		u"""注册处理函数"""
		self.cmd_tables[f] = func
	
	def dispatch(self, f, args):
		u"""处理注册函数调用,结果返回用于执行callback逻辑."""
		ret = None
		func = self.cmd_tables.get(f, None)
		if func:
			ret = func(*args)
		return ret
		
	def run(self, *apps):
		for app in apps:
			self.mods[app.topic] = app
		for (k, m) in self.mods.items():
			m.onInit(self, m.settings)
			m.start()

	def close(self):
		self._close()
		for (k, m) in self.mods.items():
			m.stop()
			m._stop_evt.wait()
			m.onDestroy()