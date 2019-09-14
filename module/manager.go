package module

import (
	"sync"

	"github.com/TianQinS/fastapi/basic"
)

type ModuleManager struct {
	mods map[string]*DefaultModule
}

func NewModuleManager() (m *ModuleManager) {
	m = &ModuleManager{
		mods: make(map[string]*DefaultModule, 0),
	}
	return
}

type DefaultModule struct {
	mi       Module
	closeSig chan bool
	wg       sync.WaitGroup
}

func run(m *DefaultModule) {
	defer func() {
		if e, ok := recover().(error); ok {
			basic.PackErrorMsg(e, nil)
		}
	}()
	m.mi.Run(m.closeSig)
	m.wg.Done()
}

func destroy(m *DefaultModule) {
	defer func() {
		if e, ok := recover().(error); ok {
			basic.PackErrorMsg(e, nil)
		}
	}()
	m.mi.OnDestroy()
}

// Register a module by a unique topic.
func (this *ModuleManager) Register(m Module) {
	md := new(DefaultModule)
	topic := m.GetTopic()
	md.mi = m
	md.closeSig = make(chan bool, 1)
	m.RegisterMgr(this)
	this.mods[topic] = md
}

func (this *ModuleManager) GetModule(topic string) Module {
	if m, ok := this.mods[topic]; ok {
		return m.mi
	}
	return nil
}

func (this *ModuleManager) Call(topic string, f interface{}, params []interface{}) (err error) {
	if m := this.GetModule(topic); m != nil {
		err = m.Call(f, params...)
	} else {
		err = Post.PutQueue(f, params...)
	}
	return
}

// CallWithCallback call a registered function and a callback function is invoked with the results.
func (this *ModuleManager) CallWithCallback(topic string, f, cb interface{}, cbParams, params []interface{}) (err error) {
	if m := this.GetModule(topic); m != nil {
		err = m.CallWithCallback(f, cb, cbParams, params)
	} else {
		err = Post.PutQueueWithCallback(f, cb, cbParams, params...)
	}
	return
}

// Run all the modules registered.
func (this *ModuleManager) Run() {
	for _, m := range this.mods {
		m.wg.Add(1)
		go run(m)
	}
}

func (this *ModuleManager) Destroy() {
	for _, m := range this.mods {
		m.closeSig <- true
		m.wg.Wait()
		destroy(m)
	}
}
