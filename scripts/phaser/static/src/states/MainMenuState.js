"use strict";
var extend = require('../utils/inherits.js');
var GameState = require('./GameState.js');

module.exports = extend(function () {}, {
	ctor: function () {
		this.server = window.Client;
	},
	init: function() {
		this.physics.startSystem(Phaser.Physics.ARCADE);
		this.game.stage.backgroundColor = "#DFDFDF";
		this.game.time.desiredFps = 25;
	},
	preload: function() {
		this.game.load.image("sgBtn", "static/assets/sgBtn.png");
		this.game.load.image("bg", "static/assets/bg.jpg");
	},
	create: function() {
		this.game.scale.pageAlignHorizontally = true;
		this.game.scale.pageAlignVertically = true;
		// 缩放以显示全部
		this.game.scale.scaleMode = Phaser.ScaleManager.SHOW_ALL;
		var bgSprite = this.game.add.tileSprite(0, 0, 2*this.game.world.width, 2*this.game.world.height, "bg");
		this.singleBtn = this.game.add.button(0, 0, "sgBtn", this.startGame, this);
		this.singleBtn.scale.x = 1.5;
		this.singleBtn.scale.y = 1.5;
		this.singleBtn.reset((this.game.width - this.singleBtn.width) / 2, (this.game.height - this.singleBtn.height) / 2);
	},
	startGame: function() {
		var that = this;
		this.server.request("", "phaserLogin", [], function(guid, nick, lv){
			that.game.state.add("GameState", new GameState(guid, nick, lv), false);
			that.state.start("GameState");
		});
    }
});