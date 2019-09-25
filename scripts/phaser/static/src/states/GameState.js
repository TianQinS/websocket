"use strict";
var extend = require('../utils/inherits.js');
var Hero = require('../objects/Player.js');

module.exports = extend(function () {}, {
	ctor: function (nick, lv) {
		this.nick = nick;
		this.lv = lv;
	},
	init: function() {
		this.physics.startSystem(Phaser.Physics.ARCADE);
		this.game.stage.backgroundColor = "#DFDFDF";
		this.game.time.desiredFps = 25;
		this.game.touchControl = this.game.plugins.add(Phaser.Plugin.TouchControl);
		this.game.touchControl.settings.singleDirection = false;
		this.game.touchControl.inputEnable();
	},
	preload: function() {
		this.game.load.image("arrow", "static/assets/arrow.png");
		this.game.load.image("ball", "static/assets/ball.png");
	},
	create: function() {
		var bgSprite = this.game.add.tileSprite(0, 0, 2*this.game.world.width, 2*this.game.world.height, "bg");
		bgSprite.autoScroll(-50,50);
		this.arrows = this.game.add.group(); //视觉观察组
		this.heros = this.game.add.group();
		this.names = this.game.add.group();
		this.gicGroup = this.game.add.group();
		this.gicBlood = this.game.add.graphics(0, 0, this.gicGroup); // 玩家血条组
		this.hero = this.joinHero({x:100, y:100, lv: this.lv, nick: this.nick});
		this.game.camera.deadzone = new Phaser.Rectangle(200, 200, this.game.stage.width - 200, this.game.stage.height - 200); //镜头跟随触发区域dead zone
		this.game.camera.follow(this.hero); 
	},
	joinHero: function(role) {
		var hero = new Hero(this.game, role.x, role.y, this.heros, {
			arrowsGroup: this.arrows,
			namesGroup: this.names,
			lv: role.lv, // 等级
			nick: role.nick, // 玩家昵称
			hp: 800, // 当前气血
		});
		return hero;
	},
	updateBlood: function(player){
		// 刻度换算
		var blood = player.hp / 35;
		var x = player.x - blood / 2;
		var y = player.y - 30;
		var deltaX = 5;
		if (player.hp >0) {
			if (player.alpha == 0.3)
			{
				player.alpha = 1;
			}
		} else if (player.alpha == 1) {
			// 死亡效果
			player.alpha = 0.3;
			player.rotateCd=(new Date()).getTime() + 3000;
		}
		if (blood >= 0) {
			if (player.blood == 0 && blood > 0){
				// 初始化血条
				this.game.add.tween(player).to( { blood: blood }, 500, "Linear", true);
			} else if (player.hp == player.lastHp) {
				// 监听变化
				this.hpChangeTag = false;
			} else if (this.hpChangeTag == false) {
				// 首次变动时的血量提示
				var hpDelta = player.hp - player.lastHp;
				player.lastHp = player.hp;
				this.hpChangeTag = true;
				if (this.hpTips != undefined)
				{
					this.hpTips.kill();
				}
				// 血量变动动画
				var hpStr = hpDelta.toString();
				var style = { font: "14px Arial", fill: "#ff0000", align: "center" };
				if (hpDelta > 0)
				{
					hpStr = "+" + hpStr;
					style = { font: "14px Arial", fill: "#00ff00", align: "center" };
				}
				this.hpTips = player.addChild(this.game.make.text(0, -10, hpStr, style));
				this.hpTips.anchor.set(0.5);
				this.hpTips.alpha = 0.1;
				this.game.add.tween(this.hpTips).to( { alpha: 1, y: -50 }, 400, "Linear", true);
				this.game.add.tween(player).to( { blood: blood }, 300, "Linear", true);	
			}
		}
		if (this.hpTips != undefined && this.hpTips.alpha == 1)
		{
			// 清理残余效果
			this.hpTips.kill();
		}

		var gic = this.gicBlood;
		var colorBlood = 0xFF0000; // 敌方阵营
		// colorBlood = 0x00EE00; // 本阵营
		gic.beginFill(colorBlood, 0.7);
		gic.drawRoundedRect(x, y, player.blood, 7, 2);
		gic.beginFill(0x333333, 0.5);
		// 血条刻度
		for (var i=deltaX; i<player.blood; i+=deltaX)
		{
			gic.drawRect(x+i, y, 1, 4);
		}
	},
	update: function(){
		this.gicBlood.clear(); // 这里如果不清理会有很炫的动态残留效果，#todo一个这样的技能特效
		this.heros.forEachAlive(function(player){
			player.Rotate();
			player.Move(); //移动
			this.updateBlood(player);
		}, this);
	}
});