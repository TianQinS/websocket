package hotfix

import (
	"github.com/TianQinS/websocket/config"
	"github.com/TianQinS/websocket/event"
	"github.com/TianQinS/websocket/hotfix/stdlibs"
	"github.com/TianQinS/websocket/module"
	"github.com/TianQinS/fastapi/basic"
	hot "github.com/TianQinS/fastapi/hotfix"
	"github.com/containous/yaegi/interp"
)

var (
	EMgr *event.EventMgr
	conf = config.Conf.Hotfix
)

// Update hotfix's dependencies for configuration modules.
func Update() {
	hot.NewHotFix(
		conf.StdOutput,
		conf.ModulePrefix,
		conf.Modules...,
	)
}

// RegisterApp register a event manager for hotfix by rpc module.
func RegisterApp(app *event.EventMgr) {
	EMgr = app
	if mod := EMgr.GetModule("rpc"); mod != nil {
		rpc := mod.(*module.RPCModule)
		rpc.RegisterRpc("eval", Eval)
		rpc.RegisterRpc("exec", Exec)
	}
}

// Exec a Process function for event manager.
func Exec(content string) (err error) {
	i := interp.New(interp.Options{})
	i.Use(stdlibs.Symbols)
	_, err = i.Eval(content)
	if err == nil {
		v, e := i.Eval("patch.Process")
		err = e
		if err == nil {
			process := v.Interface().(func(ev *event.EventMgr) error)
			err = process(EMgr)
		}
	}
	if err != nil {
		basic.PackErrorMsg(err, "patch.Process")
	}
	return
}

// Eval evaluate the script in the context of interpreter.
func Eval(topic, id, content, function string) (err error) {
	i := interp.New(interp.Options{})
	i.Use(stdlibs.Symbols)

	_, err = i.Eval(content)
	if err == nil {
		v, e := i.Eval(function)
		err = e
		if err == nil {
			if mod := EMgr.GetModule(topic); mod != nil {
				switch mod.(type) {
				case *module.RPCModule:
					mod.(*module.RPCModule).RegisterRpc(id, v.Interface())
				default:
					mod.Register(id, v.Interface())
				}
			} else {
				event.Post.Register(id, v.Interface())
			}
		}
	}
	if err != nil {
		basic.PackErrorMsg(err, content)
	}

	return
}
