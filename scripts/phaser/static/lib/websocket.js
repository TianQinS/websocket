'use strict';

var hashmap = function () {
}
hashmap.prototype = {
	constructor: hashmap,
	add: function (k, v) {
		if (!this.hasOwnProperty(k)) {
			this[k] = v;
		}
	},
	remove: function (k) {
		if (this.hasOwnProperty(k)) {
			delete this[k];
		}
	},
	update: function (k, v) {
		this[k] = v;
	},
	has: function (k) {
		var type = typeof k;
		if (type === 'string' || type === 'number') {
			return this.hasOwnProperty(k);
		} else if (type === 'function' && this.some(k)) {
			return true;
		}
		return false;
	},
	clear: function () {
		for (var k in this) {
			if (this.hasOwnProperty(k)) {
				delete this[k];
			}
		}
	},
	empty: function () {
		for (var k in this) {
			if (this.hasOwnProperty(k)) {
				return false;
			}
		}
		return true;
	},
	each: function (fn) {
		for (var k in this) {
			if (this.hasOwnProperty(k)) {
				fn.call(this, this[k], k, this);
			}
		}
	},
	map: function (fn) {
		var hash = new Hash;
		for (var k in this) {
			if (this.hasOwnProperty(k)) {
				hash.add(k, fn.call(this, this[k], k, this));
			}
		}
		return hash;
	},
	filter: function (fn) {
		var hash = new Hash;
		for (var k in this) {

		}
	},
	join: function (split) {
		split = split !== undefined ? split : ',';
		var rst = [];
		this.each(function (v) {
			rst.push(v);
		});
		return rst.join(split);
	},
	every: function (fn) {
		for (var k in this) {
			if (this.hasOwnProperty(k)) {
				if (!fn.call(this, this[k], k, this)) {
					return false;
				}
			}
		}
		return true;
	},
	some: function (fn) {
		for (var k in this) {
			if (this.hasOwnProperty(k)) {
				if (fn.call(this, this[k], k, this)) {
					return true;
				}
			}
		}
		return false;
	},
	find: function (k) {
		var type = typeof k;
		if (type === 'string' || type === 'number' && this.has(k)) {
			return this[k];
		} else if (type === 'function') {
			for (var _k in this) {
				if (this.hasOwnProperty(_k) && k.call(this, this[_k], _k, this)) {
					return this[_k];
				}
			}
		}
		return null;
	}
};

var client = function () {
}
client.prototype = {
	curr_id: 0,
	waiting_queue:new hashmap(),
	/**
	  * 建立websocket连接
	  * @param opts.host
	  * @param opts.port
	  * @param opts.onError
	  * @param opts.onConnect
	  * @param opts.onClose
	  */
	connect:function(opts) {
		var endpoint = "ws://" + opts.host + ":" + opts.port;
		var that = this;
		if (window.WebSocket) {
			this._sock = new WebSocket(endpoint);
		} else if (window.MozWebSocket) {
			this._sock = MozWebSocket(endpoint);
		} else if (opts.onError){
			opts.onError("Not Supported");
			return;
		}
		// 连接成功
		this._sock.onopen = function() {
			if (opts.onConnect) {
				opts.onConnect();
			}
		};
		// 连接关闭
		this._sock.onclose = function(event) {
			if (opts.onClose) {
				opts.onClose();
			} else {
				console.log("远端已经关闭");
			}
		};
		// 消息处理
		this._sock.onmessage = function(event)
		{
			// console.log(event);
			var data = JSON.parse(event.data.toString());
			if (data.error != undefined && opts.onError) {
			    opts.onError(data.error); 
			 } else {
				that.process_msg(data);
			 }
		};
	},
	 /**
	  * 处理服务端发过来的消息
	  * @param data
	  */
	process_msg: function(data) {
		try{
			var callback=this.waiting_queue.find(data.func);
			if(callback!=null){
				// 约定"callback/id"格式的回调函数只调用一次
				var h=data.func.split("/")
				if(h.length>=2){
					this.waiting_queue.remove(data.func)
				}
				callback.apply(null, data.args);
			}
		}catch(e) {
			console.log(e);
		}
	},
	/**
	 * 向服务器发送一条消息
	 * @param topic
	 * @param func
	 * @param args
	 * @param callback
	 */
	request: function(topic,func,args,callback){
		this.curr_id=this.curr_id+1
		var cb = func+"/"+this.curr_id;
		var payload=JSON.stringify({
			tpc: topic,
			func: func,
			args: args,
			cb: cb, 
		})
		this.on(cb,callback)
		// console.log(payload);
		this._sock.send(payload);
	},
	/**
	 * 向服务器发送一条消息,但不要求服务器返回结果
	 * @param topic
	 * @param func
	 * @param args
	 */
	requestNR:function(topic,func,args){
		var payload=JSON.stringify({
			tpc: topic,
			func: func,
			args: args,
			cb: "",
		})
		this._sock.send(payload);
	},
	/**
	 * 监听指定类型的topic消息
	 * @param topic
	 * @param callback
	 */
	on:function(topic, callback){
		//服务器不会返回结果
		this.waiting_queue.add(topic, callback) //添加这条消息到等待队列
	}
}

window.Client = new client();