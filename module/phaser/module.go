package phaser

import (
	"github.com/TianQinS/websocket/event"
	"github.com/TianQinS/websocket/module"
)

type PhaserModule struct {
	topic string
	module.BaseModule
}

func NewPhaserModule(topic string, qSize uint64) *PhaserModule {
	m := &PhaserModule{}
	m.OnInit(topic, qSize)
	return m
}

func (this *PhaserModule) GetTopic() string {
	return this.topic
}

func (this *PhaserModule) OnInit(topic string, qSize uint64) {
	this.topic = topic
	this.BaseModule.Init(qSize)
	this.RegisterGo("phaserLogin", this.login)
}

func (this *PhaserModule) Run(closeSig chan bool) {
	<-closeSig
	return
}

func (this *PhaserModule) login(client *event.Client) (name string, lv int) {
	name = "测试名称"
	lv = 99
	return
}
