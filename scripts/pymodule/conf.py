# -*- coding: utf8 -*-
u"""config."""


class Config(object):
	u"""基础配置."""
	
	# 调试模式
	DEBUG = True

class Development(Config):
	u"""调试配置."""

	RPC_TOPIC = "pyRpc"
	# Redis数据库
	REDIS_CONFIG = {"host": "127.0.0.1", "port": 6379, "db": 1, "password": ""}


class Production(Config):
	u"""线上配置."""

	RPC_TOPIC = "pyRpc"
	REDIS_CONFIG = {"host": "127.0.0.1", "port": 6379, "db": 1, "password": ""}


config = Production
if Config.DEBUG:
	config = Development