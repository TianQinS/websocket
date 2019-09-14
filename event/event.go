package event

import (
	"time"

	"github.com/TianQinS/websocket/config"
	"github.com/TianQinS/websocket/module"
	"github.com/TianQinS/evio"
	"github.com/TianQinS/fastapi/post"
)

var (
	// MaxSleepTime is max sleep time in ms.
	MaxSleepTime = 10 * time.Millisecond
	// Post represents the global worker for job.
	Post = post.GPost
)

type EventMgr struct {
	events evio.Events
	// apps is a default module manager.
	apps *module.ModuleManager
	// users
	Users map[string]*Client
}

func NewEventMgr() *EventMgr {
	return &EventMgr{
		events: evio.Events{},
		apps:   module.NewModuleManager(),
		Users:  make(map[string]*Client),
	}
}

// GetModule get a module for hotfix.
func (this *EventMgr) GetModule(topic string) module.Module {
	return this.apps.GetModule(topic)
}

// Run register modules and start them.
func (this *EventMgr) Run(apps ...module.Module) {
	for _, app := range apps {
		this.apps.Register(app)
	}
	this.apps.Run()
}

// Serve starts handling events for the specified addresses.
func (this *EventMgr) Serve(addr ...string) error {
	this.events.NumLoops = config.Conf.NumLoop
	this.events.LoadBalance = evio.RoundRobin

	this.events.Serving = func(srv evio.Server) (action evio.Action) {
		return
	}
	this.events.Opened = func(ec evio.Conn) (out []byte, opts evio.Options, action evio.Action) {
		cc := NewClient(&ec, this, this.apps)
		ec.SetContext(cc)
		return
	}
	this.events.Closed = func(ec evio.Conn, err error) (action evio.Action) {
		cc := ec.Context().(*Client)
		Post.PutQueueSpec(cc.OnClose, false)
		return
	}

	this.events.Data = func(ec evio.Conn, in []byte) (out []byte, action evio.Action) {
		cc := ec.Context().(*Client)
		out, action = cc.OnData(&in)
		return
	}
	if err := evio.Serve(this.events, addr...); err != nil {
		return err
	}
	return nil
}

// Close called when receipt signal, such as syscall.SIGINT and syscall.SIGTERM.
func (this *EventMgr) Close() {
	this.apps.Destroy()
}
