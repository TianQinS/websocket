# -*- coding: utf8 -*-
u"""redis测试Py脚本."""

import time
import json
import traceback
from redis import ConnectionPool, StrictRedis
from redis.exceptions import ConnectionError, TimeoutError


class Config(object):
	u"""基础配置."""
	
	# 调试模式
	DEBUG = True

class Development(Config):
	u"""调试配置."""

	# Redis数据库
	REDIS_CONFIG = {"host": "127.0.0.1", "port": 6379, "db": 1, "password": ""}


class Production(Config):
	u"""线上配置."""
	
	# Redis数据库，用于Redis模式直接发布缓存同步消息使用
	REDIS_CONFIG = {"host": "127.0.0.1", "port": 6379, "db": 1, "password": ""}


config = Production
if Config.DEBUG:
	config = Development


class RDB(object):
	u"""Redis连接对象."""

	def __init__(self, **conf):
		super(RDB, self).__init__()
		self.cmd_tables = {}
		self.lua_scripts = {}
		self.conf = conf
		self.alive = True
		self.topic = "test"
		self.pubsub = None
		self.reconnect()
	
	def register(self, f, func):
		u"""注册处理函数"""
		self.cmd_tables[f] = func
	
	def dispatch(self, f, args):
		u"""处理注册函数调用."""
		func = self.cmd_tables.get(f, None)
		if func:
			func(*args)

	def reconnect(self):
		u"""重连."""
		try:
			self.pool = ConnectionPool(socket_connect_timeout=3, **self.conf) 
			self.db = StrictRedis(connection_pool=self.pool, socket_timeout=1, socket_keepalive=60, socket_connect_timeout=1, **self.conf)
		except Exception, e:
			self.alive = False
			print traceback.format_exc()

	def call(self, topic, f, callback, *args):
		u"""调用测试."""
		data = json.dumps({
			"func": f,
			"ct": self.topic,
			"cb": callback,
			"args": args,
		})
		self.db.rpush(topic, data)
	
	def hotfix(self, topic, id, fix_content, fix_function):
		u"""在线更新测试1"""
		data = json.dumps({
			"func": "eval",
			"ct": "",
			"cb": "",
			"args": [topic, id, fix_content, fix_function],
		})
		self.db.publish("rpc", data)
	
	def callNR(self, func, *args):
		u"""实时调用，无回调."""
		data = json.dumps({
			"func": func,
			"ct": "",
			"cb": "",
			"args": args,
		})
		self.db.publish("rpc", data)

	def patch(self, content):
		u"""在线更新测试2"""
		self.callNR("exec", content)
	
	def subscribe(self, tpoic):
		u""""订阅."""
		pub = self.db.pubsub()
		pub.subscribe(self.topic)
		return pub
	
	def test(self, *args):
		u"""测试."""
		print args
	
	def processData(self, data):
		u""""处理数据."""
		print data
		data = json.loads(data)
		f = data.get("func", "")
		args = data.get("args", [])
		self.dispatch(f, args)
	
	def loop(self):
		u"""接收消息测试."""
		while self.alive:
			data = self.db.lpop(self.topic)
			if not data:
				time.sleep(0.3)
			else:
				self.processData(data)
	
	def read(self):
		u"""阻塞接收回调消息."""
		if not self.pubsub:
			self.pubsub = self.subscribe(self.topic)
		data = self.pubsub.parse_response()
		if data[0] == "message":
			self.processData(data[2])
		else:
			print data



if __name__ == "__main__":
	rdb = RDB(**config.REDIS_CONFIG)
	rdb.register("callback", rdb.test)
	rdb.call("rpc", "test", "callback", "arg1", "arg2")
	rdb.read()
	rdb.patch("""package patch

import (
	"fmt"
	"github.com/TianQinS/websocket/event"
	"github.com/TianQinS/websocket/module"
)

func Process(ev *event.EventMgr) error {
	if mod := ev.GetModule("rpc"); mod != nil {
		rmod := mod.(*module.RPCModule)
		rmod.RegisterRpc("test", func(arg1, arg2 string) string {
			return arg1 + arg2 + arg1 + arg2
		})
	}
	fmt.Println("Patch process finished.")
	return nil
}
	
""")
	rdb.call("rpc", "test", "callback", "arg1", "arg2")
	rdb.read()
	rdb.hotfix("rpc", "test", """package hotfix

func Test(arg1, arg2 string) string {
	return arg2 + arg1
}""", "hotfix.Test")
	rdb.call("rpc", "test", "callback", "arg1", "arg2")
	rdb.read()
	# rdb.loop()