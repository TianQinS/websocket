# -*- coding: utf8 -*-
u"""Main."""

import gevent
import gevent.monkey
gevent.monkey.patch_all()
import time
import json
import traceback
from gevent.event import Event
from module import ModuleManager
from rpcmodule import RPCModule
from conf import config


def test(arg1, arg2):
	return arg1+arg2

def callback(arg):
	print(arg)

if __name__ == "__main__":
	app = ModuleManager()
	app.run(
		RPCModule(**config.REDIS_CONFIG)
	)
	app.register("test", test)
	app.register("callback", callback)
	
	rpc = app[config.RPC_TOPIC]
	
	rpc.call("pyRpc", "test", "callback", "arg1", "arg2")
	rpc.call("pyRpc", "test", "callback", "arg1", "arg2")
	time.sleep(10)