// Code generated by automatic for 'github.com/TianQinS/websocket/module'. DO NOT EDIT.

// +build go1.12,!go1.13

package stdlibs

import (
	"github.com/TianQinS/websocket/module"
	"reflect"
)

func init() {
	Symbols["github.com/TianQinS/websocket/module"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Hook":             reflect.ValueOf(&module.Hook).Elem(),
		"NewBaseModule":    reflect.ValueOf(module.NewBaseModule),
		"NewModuleManager": reflect.ValueOf(module.NewModuleManager),
		"NewRPCModule":     reflect.ValueOf(module.NewRPCModule),
		"Post":             reflect.ValueOf(&module.Post).Elem(),

		// type definitions
		"BaseModule":    reflect.ValueOf((*module.BaseModule)(nil)),
		"DefaultModule": reflect.ValueOf((*module.DefaultModule)(nil)),
		"Module":        reflect.ValueOf((*module.Module)(nil)),
		"ModuleManager": reflect.ValueOf((*module.ModuleManager)(nil)),
		"Msg":           reflect.ValueOf((*module.Msg)(nil)),
		"PubMsg":        reflect.ValueOf((*module.PubMsg)(nil)),
		"RPCModule":     reflect.ValueOf((*module.RPCModule)(nil)),

		// interface wrapper definitions
		"_Module": reflect.ValueOf((*_github_com_TianQinS_websocket_module_Module)(nil)),
	}
}

// _github_com_TianQinS_websocket_module_Module is an interface wrapper for Module type
type _github_com_TianQinS_websocket_module_Module struct {
	WCall             func(f interface{}, params []interface{}) error
	WCallGo           func(f interface{}, params []interface{}) error
	WCallSpec         func(f interface{}, params []interface{}) error
	WCallWithCallback func(f interface{}, cb interface{}, cbParams []interface{}, params []interface{}) error
	WExecuteEvent     func() uint64
	WGetFunc          func(f string) (interface{}, error)
	WGetMgr           func() *module.ModuleManager
	WGetTopic         func() string
	WOnDestroy        func()
	WOnInit           func(topic string, qSize uint64)
	WRegister         func(id string, f interface{})
	WRegisterGo       func(id string, f interface{})
	WRegisterMgr      func(mgr *module.ModuleManager)
	WRun              func(closeSig chan bool)
}

func (W _github_com_TianQinS_websocket_module_Module) Call(f interface{}, params []interface{}) error {
	return W.WCall(f, params)
}
func (W _github_com_TianQinS_websocket_module_Module) CallGo(f interface{}, params []interface{}) error {
	return W.WCallGo(f, params)
}
func (W _github_com_TianQinS_websocket_module_Module) CallSpec(f interface{}, params []interface{}) error {
	return W.WCallSpec(f, params)
}
func (W _github_com_TianQinS_websocket_module_Module) CallWithCallback(f interface{}, cb interface{}, cbParams []interface{}, params []interface{}) error {
	return W.WCallWithCallback(f, cb, cbParams, params)
}
func (W _github_com_TianQinS_websocket_module_Module) ExecuteEvent() uint64 {
	return W.WExecuteEvent()
}
func (W _github_com_TianQinS_websocket_module_Module) GetFunc(f string) (interface{}, error) {
	return W.WGetFunc(f)
}
func (W _github_com_TianQinS_websocket_module_Module) GetMgr() *module.ModuleManager {
	return W.WGetMgr()
}
func (W _github_com_TianQinS_websocket_module_Module) GetTopic() string { return W.WGetTopic() }
func (W _github_com_TianQinS_websocket_module_Module) OnDestroy()       { W.WOnDestroy() }
func (W _github_com_TianQinS_websocket_module_Module) OnInit(topic string, qSize uint64) {
	W.WOnInit(topic, qSize)
}
func (W _github_com_TianQinS_websocket_module_Module) Register(id string, f interface{}) {
	W.WRegister(id, f)
}
func (W _github_com_TianQinS_websocket_module_Module) RegisterGo(id string, f interface{}) {
	W.WRegisterGo(id, f)
}
func (W _github_com_TianQinS_websocket_module_Module) RegisterMgr(mgr *module.ModuleManager) {
	W.WRegisterMgr(mgr)
}
func (W _github_com_TianQinS_websocket_module_Module) Run(closeSig chan bool) { W.WRun(closeSig) }
