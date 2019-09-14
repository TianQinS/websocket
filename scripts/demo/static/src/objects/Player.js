"use strict";

var extend = require('../utils/inherits.js');
var GameRole = require('./Role.js');
var mqant=window.mqant
module.exports = extend(GameRole, {
	ctor: function ctor(game, x, y, group, properties) {
		this.roleType = "hero";
		properties = properties || {};
		this.game = game;
		this._super(this.game, x, y, "ball", null, properties);
		this.game.physics.arcade.enable(this);
		this.anchor.setTo(0.5, 0.5);
		this.checkWorldBounds = true;
		this.outOfBoundsKill = true;
		this.body.collideWorldBounds = true; //与世界边境进行物理检测
		this.inputEnabled = true;
		this.input.useHandCursor = true; //当鼠标移动到其上面时显示小手
		if (group) {
			group.add(this);
		}
	},
	dead: function () {
		if(this.alive){
			this._super();
		}
	},
});