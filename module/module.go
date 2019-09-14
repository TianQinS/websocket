package module

import (
	"fmt"

	"github.com/TianQinS/fastapi/post"
)

var (
	// Post is an alias for the global Post object.
	Post = post.GPost
)

type BaseModule struct {
	post.RpcObject
	mgr   *ModuleManager
	topic string
}

func NewBaseModule(topic string, qSize uint64) *BaseModule {
	m := &BaseModule{}
	m.OnInit(topic, qSize)
	return m
}

func (this *BaseModule) OnInit(topic string, qSize uint64) {
	this.topic = topic
	this.RpcObject.Init(qSize)
}

func (this *BaseModule) GetTopic() string {
	return this.topic
}

func (this *BaseModule) Run(closeSig chan bool) {
}

func (this *BaseModule) OnDestroy() {
	this.RpcObject.Close()
}

func (this *BaseModule) RegisterMgr(mgr *ModuleManager) {
	this.mgr = mgr
}

// GetMgr return the module manager.
func (this *BaseModule) GetMgr() *ModuleManager {
	return this.mgr
}

// GetFunc get a registed function by name.
func (this *BaseModule) GetFunc(f string) (interface{}, error) {
	if function, ok := this.Functions[f]; ok {
		return function, nil
	}
	return nil, fmt.Errorf("Unknonwn function=%s topic=%s", f, this.GetTopic())
}

func (this *BaseModule) RegisterGo(id string, f interface{}) {
	Post.Register(id, f)
}

// Call call a register function of this module.
func (this *BaseModule) Call(f interface{}, params ...interface{}) error {
	return this.PutQueueForPost(f, false, params)
}

func (this *BaseModule) CallWithCallback(f, cb interface{}, cbParams, params []interface{}) error {
	return this.PutQueueWithCallback(f, cb, false, cbParams, params)
}

// CallGo call a register function of the global Post in any gorountines.
func (this *BaseModule) CallGo(f interface{}, params ...interface{}) error {
	return Post.PutQueue(f, params...)
}

// CallSpec call a register function of the Post in the main event tick.
func (this *BaseModule) CallSpec(f interface{}, params ...interface{}) error {
	return Post.Object.PutQueueForPost(f, false, params)
}
