package web

import (
	"github.com/TianQinS/websocket/config"
	"github.com/TianQinS/websocket/module"
	"github.com/kataras/iris"
)

var (
	conf = config.Conf.Web
)

type WebModule struct {
	topic string
	app   *iris.Application
	module.BaseModule
}

func NewWebModule(topic string, qSize uint64) *WebModule {
	m := &WebModule{}
	m.OnInit(topic, qSize)
	return m
}

func (this *WebModule) GetTopic() string {
	return this.topic
}

func (this *WebModule) OnInit(topic string, qSize uint64) {
	this.topic = topic
	this.BaseModule.Init(qSize)
	this.app = NewApp(conf.Charset)
}

func (this *WebModule) Run(closeSig chan bool) {
	this.app.StaticWeb("/static", conf.StaticPath)
	this.app.StaticWeb("/js", conf.JsPath)
	this.app.RegisterView(iris.HTML(conf.HtmlPath, ".html"))

	this.app.Run(iris.Addr(conf.Port))
	// <-closeSig
}
