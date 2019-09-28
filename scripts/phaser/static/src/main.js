'use strict';
var extend = require('./utils/inherits.js');
var MyScaleManager = require('./utils/MyScaleManager.js');
var BootState = require('./states/BootState.js');

window.onload = function () {
    var gameDiv = document.getElementById("game");
    Phaser.myScaleManager = new MyScaleManager(gameDiv);
    var width=1600;
    var scale = screen.width / screen.height;
    if (scale > 1) {
        scale = 1 / scale;
    }
    var game;;
	if (window.location.hostname == "") {
		game = new Phaser.Game(width, width * scale, Phaser.CANVAS, gameDiv);
	} else {
		game = new Phaser.Game(width, width * scale, Phaser.AUTO, gameDiv);
	}
    Phaser.myScaleManager.boot();
    game.state.add("BootState", BootState, true);
    game.state.start("BootState");
};

