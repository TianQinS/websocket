"use strict";

var extend = require('../utils/inherits.js');
var GameRole = require('./Role.js');
module.exports = extend(GameRole, {
	ctor: function ctor(game, x, y, group, properties) {
		this.roleType = "hero";
		properties = properties || {};
		this.game = game;
		this.guid = properties.guid;
		this.lv = properties.lv;
		this.nick = properties.nick;
		this.hp = properties.hp;
		this.arrows = properties.arrowsGroup; //视觉观察组
		this.names = properties.namesGroup;

		this.friction = 0.9;
		this.fireSpeed = -10;
		this.rotateSpeed = 5;

		this.blood = 0;
		this.rotateCd = 0;
		this.xSpeed = 0;
		this.ySpeed = 0;
		this.angle = 0;
		this.rotateDirection = 1;
		this.rotateCd = 0;

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
			this.arrow.kill();
		}
	},
	getArrow: function getArrow() {
		if (this.arrow == null) {
			var arrow = this.arrows.getFirstExists(false);
			if (arrow) {
				this.arrow = arrow;
				this.arrow.reset(this.x, this.y);
			} else {
				//设置一个观察器
				this.arrow = this.game.add.sprite(this.game.world.centerX,this.game.world.centerY, "arrow");
				this.arrow.anchor.x = -1;
				this.arrow.anchor.y = 0.5;
			}
		}
		return this.arrow;
	},
	getName: function() {
		if (this.nameObj == null) {
			var color1 = "#FFFFFF", color2 = "#777777";
			// color1 = "red"; color2 = "#111111"; // 敌方阵营
			var nameObj = this.names.getFirstExists(false);
			if (nameObj) {
				this.nameObj = nameObj;
				this.nameObj.reset(this.x, this.y);
			} else {
				var nameStyle = {font: "10px '微软雅黑'", fill: color1, align: "center",};
				this.nameObj = this.game.add.text(this.game.world.centerX, this.game.world.centerY, this.name, nameStyle);
				this.nameObj.anchor.x = 0.5;
				this.nameObj.anchor.y = -1;
				this.nameObj.setShadow(1, 1, color2, 0);
				this.nameObj.alpha = 0.7;
			}
			if (this.nameObj.text != this.nick)
			{
				this.nameObj.setText(this.nick, 1);
			}
		}
		return this.nameObj;
	},
	Fire: function(x, y) {
		var curTime = new Date().getTime();
		var angle = Math.atan2(-y, -x) / Math.PI * 180;

		// 自转冷却
		this.rotateCd = (new Date()).getTime() + 250;
		if (angle > this.getArrow().angle) {
			if (this.rotateDirection < 0) {
				this.rotateDirection *= -1;
			}
		} else {
			if (this.rotateDirection > 0) {
				this.rotateDirection *= -1;
			}
		}
		this.getArrow().angle = angle;
		this.xSpeed = x / this.fireSpeed;
		this.ySpeed = y / this.fireSpeed;
	},
	Rotate: function () {
		var curTime = new Date().getTime();
		if (curTime > this.rotateCd) {
			// 自转特效
			this.getArrow().angle += this.rotateSpeed * this.rotateDirection;
		}
	},
	Move: function () {
		this.xSpeed *= this.friction;
		this.ySpeed *= this.friction
		this.x += this.xSpeed;
		this.y += this.ySpeed;
		this.getArrow().x = this.x;
		this.getArrow().y = this.y;
		this.getName().x = this.x;
		this.getName().y = this.y;
	}
});