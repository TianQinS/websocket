"use strict";
var extend = require('../utils/inherits.js');

module.exports = extend(function () {}, {
    ctor: function (nick, lv) {

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
		var bgSprite = this.game.add.tileSprite(0, 0, 2*this.game.world.width, 2*this.game.world.height, "bg");
		bgSprite.autoScroll(-50,50);
	},
    startGame: function() {
        // this.state.start("GameState");
    }
});