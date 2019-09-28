"use strict";
var extend = require('../utils/inherits.js');
var Hero = require('../objects/Player.js');

module.exports = extend(function () {}, {
	ctor: function (guid, nick, lv) {
		this.guid = guid;
		this.nick = nick;
		this.lv = lv;
		
		this.frame = 0;
	},
	init: function() {
		this.physics.startSystem(Phaser.Physics.ARCADE);
		this.game.stage.backgroundColor = "#DFDFDF";
		this.game.time.desiredFps = 15;
		this.game.touchControl = this.game.plugins.add(Phaser.Plugin.TouchControl);
		this.game.touchControl.settings.singleDirection = false;
		this.game.touchControl.inputEnable();
	},
	preload: function() {
		this.game.load.image("arrow", "static/assets/arrow.png");
		this.game.load.image("ball", "static/assets/ball.png");
	},
	create: function() {
		this.game.scale.pageAlignHorizontally = true;
		this.game.scale.pageAlignVertically = true;
		this.game.scale.scaleMode = Phaser.ScaleManager.SHOW_ALL;
		this.game.world.resize(2560, 2560);
		this.game.world.setBounds(0, 0, 2560, 2560);

		var bgSprite = this.game.add.tileSprite(0, 0, 2*this.game.world.width, 2*this.game.world.height, "bg");
		// bgSprite.autoScroll(-50,50);
		/* �Ӿ��۲��� */
		this.arrows = this.game.add.group();
		this.heros = this.game.add.group();
		this.names = this.game.add.group();
		/* ͼ����Ⱦ�� */
		this.gicGroup = this.game.add.group();
		this.gicBlood = this.game.add.graphics(0, 0, this.gicGroup); // ���Ѫ��
		this.initMap();
		
		this.hero = this.joinHero({x:100, y:100, lv: this.lv, nick: this.nick, guid: this.guid});
		this.game.camera.deadzone = new Phaser.Rectangle(400, 400, this.game.stage.width - 400, this.game.stage.height - 400); //��ͷ���津������dead zone
		this.game.camera.follow(this.hero); 
	},
	initMap: function(){
		this.gicMap = this.game.add.graphics(0, 0, this.gicGroup); // С��ͼ
		this.gicMap.lineStyle(1, 0xdfdfdf);
		this.gicMap.alpha = 0.4;
		this.gicMap.fixedToCamera = true;
		this.gicMap.beginFill(0x99CC66, 1);
		this.gicMap.drawCircle(200, 200, 360);
		this.gicMap.endFill();
	},
	joinHero: function(role) {
		var hero = new Hero(this.game, role.x, role.y, this.heros, {
			arrowsGroup: this.arrows,
			namesGroup: this.names,
			lv: role.lv, // �ȼ�
			nick: role.nick, // ����ǳ�
			guid: role.guid,
			hp: 1600, // ��ǰ��Ѫ
		});
		hero.scale.x = 2;
		hero.scale.y = 2;
		return hero;
	},
	updateBlood: function(player){
		// �̶Ȼ���
		var blood = player.hp / 35;
		var x = player.x - blood/2;
		var y = player.y - 60;
		var deltaX = 6;
		if (player.hp >0) {
			if (player.alpha == 0.3)
			{
				player.alpha = 1;
			}
		} else if (player.alpha == 1) {
			// ����Ч��
			player.alpha = 0.3;
			player.rotateCd=(new Date()).getTime() + 3000;
		}
		if (blood >= 0) {
			if (player.blood == 0 && blood > 0){
				// ��ʼ��Ѫ��
				this.game.add.tween(player).to( { blood: blood }, 500, "Linear", true);
			} else if (player.hp == player.lastHp) {
				// �����仯
				this.hpChangeTag = false;
			} else if (this.hpChangeTag == false) {
				// �״α䶯ʱ��Ѫ����ʾ
				var hpDelta = player.hp - player.lastHp;
				player.lastHp = player.hp;
				this.hpChangeTag = true;
				if (this.hpTips != undefined)
				{
					this.hpTips.kill();
				}
				// Ѫ���䶯����
				var hpStr = hpDelta.toString();
				var style = { font: "20px Arial", fill: "#ff0000", align: "center" };
				if (hpDelta > 0)
				{
					hpStr = "+" + hpStr;
					style = { font: "20px Arial", fill: "#00ff00", align: "center" };
				}
				this.hpTips = player.addChild(this.game.make.text(0, -10, hpStr, style));
				this.hpTips.anchor.set(0.5);
				this.hpTips.alpha = 0.1;
				this.game.add.tween(this.hpTips).to( { alpha: 1, y: -50 }, 500, "Linear", true);
				this.game.add.tween(player).to( { blood: blood }, 300, "Linear", true);	
			}
		}
		if (this.hpTips != undefined && this.hpTips.alpha == 1) {
			// �������Ч��
			this.hpTips.kill();
		}

		var gic = this.gicBlood;
		var colorBlood = 0xFF0000; // �з���Ӫ
		if (this.guid == player.guid) {
			colorBlood = 0x00EE00; // ����Ӫ
		}
		gic.beginFill(colorBlood, 0.7);
		gic.drawRoundedRect(x, y, player.blood, 14, 4);
		gic.beginFill(0x333333, 0.5);
		// Ѫ���̶�
		for (var i=deltaX; i<player.blood; i+=deltaX)
		{
			gic.drawRect(x+i, y, 2, 10);
		}
	},
	Fire: function() {
		var speed = this.game.touchControl.speed;
		var x = speed.x;
		var y = speed.y;
		if (x == 0 && y == 0) {
			return;
		};
		this.hero.Fire(x, y);
	},
	update: function(){
		this.frame++;
		this.gicBlood.clear(); // ���������������к��ŵĶ�̬����Ч����#todoһ�������ļ�����Ч
		this.heros.forEachAlive(function(player){
			player.Rotate();
			player.Move(); //�ƶ�
			this.updateBlood(player);
		}, this);
		this.Fire();
		if (this.frame % 10 == 0) {
		}
	},
	resize: function(){
		this.game.world.resize(2560, 2560);
		this.game.world.setBounds(0, 0, 2560, 2560);
		this.game.camera.follow(this.hero);
	},
	paused: function() {
	},
	resumed: function() {
	}
});