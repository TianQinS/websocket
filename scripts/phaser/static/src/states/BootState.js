'use strict';
var MainMenuState = require('./MainMenuState.js');

module.exports = {
	preload: function() {
	},
	create: function() {
		this.server = window.Client;
		this.login();
	},
	login: function(guid, pwd) {
		var self = this;
		try{
			var host = window.location.hostname;
			if (host == "")
			{
				host = "127.0.0.1";
			}
			this.server.connect({
				host: host,
				port: 23456,
				onConnect: function() {
					//alert("游戏连接成功!");
					self.server.requestNR("", "login", ["guid", "md5"]);
					self.server.request("", "phaserLogin", [], function(nick, lv){
						self.game.state.add("MainMenuState", new MainMenuState(nick, lv), false);
						self.state.start("MainMenuState");
					});
				},
				onClose: function() {
					console.log("closed.");
				},
				onError: function(e) {
					console.log(e);
				}
			});
		}catch (e){
			alert(e);
		}
    },
};

