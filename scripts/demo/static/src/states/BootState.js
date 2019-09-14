'use strict';
var MainMenuState = require('./MainMenuState.js');

module.exports = {
    preload: function() {

    },
    create: function() {
        this.server = window.mqant;
    },
    login: function(guid, pwd) {
        var self = this;
        var useSSL = 'https:' == document.location.protocol ? true : false;
        try{
			this.server.init({
				host: window.location.hostname,
				port: 23456,
				client_id: "demo",
				userName: guid,
				password: pwd,
				useSSL:useSSL,
				onSuccess:function() {
					//alert("游戏连接成功!");
					self.server.requestNR("demo", "login", []);

					self.server.on("onLogin", function(data) {
                        State.nick = data.nick;
                        State.lv = data.lv;
                        self.game.state.add("MainMenuState", new MainMenuState(nick, lv), false);
						self.state.start("MainMenuState");
					});
				},
				onConnectionLost:function(code,reason) {
					console.log(code)
					alert("连接断开了:"+code);
				}
			});
		}catch (e){
			alert(e);
		}
    },
};

