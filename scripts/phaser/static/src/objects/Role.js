'use strict';

var extend = require('../utils/inherits.js');
module.exports = extend(Phaser.Sprite, {
	//ctor 可以省略  省略以后会继续执行其父构造函数 如 this._super.apply(this,arguments);
	ctor: function ctor(game, x, y, key, frame, properties) {
		properties = properties || {};
		this._super(game, x, y, key, frame);
		this.rid = properties.rid; //系统角色
	},
	dead: function dead() {
		if(this.alive){
			this.kill();
		}
	}
});