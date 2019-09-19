# -*- coding: utf8 -*-
u"""RPCModule."""

import time
import json
import traceback
from module import Module
from redis import ConnectionPool, StrictRedis
from redis.exceptions import ConnectionError, TimeoutError
from conf import config


class RPCModule(Module):

	def __init__(self, **conf):
		# super(RPCModule, self).__init__()
		Module.__init__(self)
		self.lua_scripts = {}
		self.conf = conf
		self.alive = True
		self.topic = config.RPC_TOPIC
		self.db = None
		self.app = None

	def reconnect(self):
		u"""重连."""
		try:
			self.pool = ConnectionPool(socket_connect_timeout=3, **self.conf) 
			self.db = StrictRedis(connection_pool=self.pool, socket_timeout=1, socket_keepalive=60, socket_connect_timeout=1, **self.conf)
		except Exception:
			self.alive = False
			print(traceback.format_exc())

	def call(self, topic, f, callback, *args):
		u"""调用测试."""
		_data = json.dumps({
			"func": f,
			"ct": self.topic,
			"cb": callback,
			"args": args,
		})
		self.db.rpush(topic, _data)

	def Run(self):
		while self.closeSig == False:
			topic, data = self.db.blpop(self.topic)
			data = json.loads(data)
			f = data.get("func", None)
			if f:
				ret = self.app.dispatch(f, data["args"])
				cb = data.get("cb", None)
				ct = data.get("ct", None)
				if cb and ct:
					self.call(ct, cb, "", ret)
	
	def onInit(self, app, setting):
		u"""初始化."""
		Module.onInit(self, app, setting)
		self.reconnect()