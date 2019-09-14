'use strict';
var extend = require('./utils/inherits.js');
var MyScaleManager = require('./utils/MyScaleManager.js');
var BootState = require('./states/BootState.js');

window.onload = function () {
    var gameDiv = document.getElementById("game");
    Phaser.myScaleManager = new MyScaleManager(gameDiv);
    var width=800;
    var scale = screen.width / screen.height;
    if (scale > 1) {
        scale = 1 / scale;
    }
    var game = new Phaser.Game(width, width * scale, Phaser.AUTO, gameDiv);
    Phaser.myScaleManager.boot();
    game.state.add("BootState", BootState, true);
    game.state.start("BootState");
};

