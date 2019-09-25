"use strict";
var extend = require('../utils/inherits.js');
var GameState = require('./GameState.js');

module.exports = extend(function () {}, {
    ctor: function (nick, lv) {
		this.nick = nick;
		this.lv = lv;
    },
	init: function() {
		this.physics.startSystem(Phaser.Physics.ARCADE);
		this.game.stage.backgroundColor = "#DFDFDF";
		this.game.time.desiredFps = 25;
	},
	preload: function() {
		this.game.load.image("bg", "static/assets/bg.jpg");
	},
	create: function() {
		this.startGame();
	},
    startGame: function() {
    	this.game.state.add("GameState", new GameState(this.nick, this.lv), false);
		this.state.start("GameState");
    }
});