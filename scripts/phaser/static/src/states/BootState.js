'use strict';
var MainMenuState = require('./MainMenuState.js');

module.exports = {
	preload: function() {
		this.game.load.image("compass", "static/assets/touch1.png");
		this.game.load.image("touch", "static/assets/touch2.png");
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
					self.game.state.add("MainMenuState", new MainMenuState(), false);
					self.state.start("MainMenuState");
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

