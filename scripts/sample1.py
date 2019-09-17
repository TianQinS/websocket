# -*- coding: utf8 -*-
u"""测试脚本."""

import sys
import json
import time
import websocket
try:
	import thread
except ImportError:
	import _thread as thread

def on_message(ws, message):
	print(message)

def on_error(ws, error):
	print(error)

def on_close(ws):
	print("### closed ###")

def on_open(ws):
	def run(*args):
		ws.send(json.dumps({"tpc":"login", "func":"login", "args": ["guid","md5"], "cb":""}))
		i = 0
		while i < 10:
			ws.send(json.dumps({"tpc":"unknown", "func":"test", "args": ["arg1","arg2"], "cb":"callback"}))
			ws.send(json.dumps({"tpc":"unknown", "func":"test", "args": ["arg1","arg2"], "cb":"callback"}))
			ws.send(json.dumps({"tpc":"unknown", "func":"test", "args": ["arg1","arg2"], "cb":"callback"}))
			ws.send(json.dumps({"tpc":"unknown", "func":"test", "args": ["arg1","arg2"], "cb":"callback"}))
			ws.send(json.dumps({"tpc":"unknown", "func":"test", "args": ["arg1","arg2"], "cb":"callback"}))
			time.sleep(1)
			i += 1
		ws.close()
		print("thread terminating...")
	thread.start_new_thread(run, ())


if __name__ == "__main__":
	websocket.enableTrace(True)
	ws = websocket.WebSocketApp("ws://127.0.0.1:23456",
							  on_message = on_message,
							  on_error = on_error,
							  on_close = on_close)
	ws.on_open = on_open
	ws.run_forever()