# Websocket
[![GoDoc](https://godoc.org/github.com/TianQinS/websocket?status.svg)](https://godoc.org/github.com/TianQinS/websocket)

**简介**

基于evio事件库，模块化功能添加，便捷编写，支持在线更新。
>1. module：模块化结构，方便添加各种自定义模块。
>2. hotfix：提供对register functions和public variable的在线更新功能。
>3. client：提供临时变量，本地变量和数据库变量的存取接口。


**基础模块**

---------------------------------------
  * [basic module](#websocket)
	* [Hotfix](#hotfix)
	* [Module](#module)
	* [RPCModule](#rpcmodule)
---------------------------------------

### Hotfix

在解释器上下文中动态解释patch模块及其Process函数，并将当前运行环境中的变量对象指针传递给Process函数进行处理。
```golang
/* github.com/TianQinS/websocket/hotfix/hotfix.go */
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
```
- 需要预先执行hotfix.Update函数进行相关依赖模块更新。
- 这里将eventMgr对象指针传递到解释代码上下文中，通过该对象可以获取所有启用的module，更新相关注册函数。

在程序外部将相关代码发给程序进行在线更新，相关样例参见scripts/sample2.py。
```python
u"""github.com/TianQinS/websocket/scripts/sample2.py."""
rdb.call("rpc", "test", "callback", "arg1", "arg2")
rdb.patch("""package patch

import (
	"fmt"
	"github.com/TianQinS/websocket/event"
	"github.com/TianQinS/websocket/module"
)

func Process(ev *event.EventMgr) error {
	if mod := ev.GetModule("rpc"); mod != nil {
		rmod := mod.(*module.RPCModule)
		rmod.RegisterRpc("test", func(arg1, arg2 string) string {
			return arg1 + arg2 + arg1 + arg2
		})
	}
	fmt.Println("Patch process finished.")
	return nil
}
	
""")
rdb.call("rpc", "test", "callback", "arg1", "arg2")
```

- 通过预置的RPCModule将上面所示的patch代码发布到相关服务进程中进行动态解释执行。
- 首先，远程调用服务端rpc模块注册的test函数，传入参数arg1和arg2；函数返回结果会被callback函数进行展示。
- 然后，通过patch函数将RPCModule的test注册函数的逻辑动态变更。
- 最后，再次远程调用test函数进行处理逻辑变更结果验证。

### Module

可以继承BaseModule实现新的模块，然后执行EventMgr的Run函数注册并启动这些模块。

```golang
/* github.com/TianQinS/websocket/module/define.go */
type Module interface {
	// OnInit initialize the module's mqtt topic and the lock-free queue's capacity.
	OnInit(topic string, qSize uint64)
	OnDestroy()
	// Run the module in goroutine.
	Run(closeSig chan bool)
	GetTopic() string
	RegisterMgr(mgr *ModuleManager)
	// GetMgr return the module manager.
	GetMgr() *ModuleManager
	// RegisterGo register a function for the global Post.
	RegisterGo(id string, f interface{})
	GetFunc(f string) (interface{}, error)
	// CallGo call a register function of the global Post in any gorountines.
	CallGo(f interface{}, params ...interface{}) error
	// CallSpec call a register function of the Post in the main event tick.
	CallSpec(f interface{}, params ...interface{}) error
	// Register register a function for this module.
	Register(id string, f interface{})
	// Call call a register function of this module.
	Call(f interface{}, params ...interface{}) error
	CallWithCallback(f, cb interface{}, cbParams, params []interface{}) error
	// ExecuteEvent process function calls on this module itself.
	ExecuteEvent() uint64
}

/* github.com/TianQinS/websocket/module/manager.go */
// Run register modules and start them.
func (ev *EventMgr) Run(apps ...module.Module) {
	for _, app := range apps {
		ev.apps.Register(app)
	}
	ev.apps.Run()
}
```

- BaseModule默认内置了无锁队列及一个主循环函数（选用），封装了基础的协程池，可以依据使用场景在全局event事件主循环/随机一个协程/模块主循环中执行相关函数。
- RPCModule模块结合redis提供了分布式扩展和热更新支持；EventMgr模块封装了高性能evio事件库。
- Client类提供了基础的临时数据、本地数据和数据库数据的存取接口及协议数据的封装与解析功能。

### RPCModule

RPCModule是一个基于BaseModule的自定义扩展模块，满足`define.go`中声明的接口约束，借助redis的队列和发布订阅进行跨进程模块间消息传递与共享。

```golang
/* github.com/TianQinS/websocket/module/rpcmodule.go */
func (this *RPCModule) Run(closeSig chan bool) {
	defer this.closeTick()
	pubsub := this.Rdb.Subscribe(this.topic)
	if _, err := pubsub.Receive(); err != nil {
		basic.PackErrorMsg(err, this.topic)
		return
	}
	ch := pubsub.Channel()
	// hook in the main loop for timer.
	Hook.Register("10ms", "2006-01-02 15:04:05", "3006-01-02 15:04:05", this.onTick)
	ldat := this.Rdb.BLpop(this.topic)

	for {
		select {
		case pmsg := <-ch:
			this.onSubscribe(pmsg.Channel, pmsg.Payload)
		case lmsg := <-ldat:
			this.onCall(lmsg)
		case <-this.tick:
			// for normal register.
			this.ExecuteEvent()
		case <-closeSig:
			return
		}
	}
}
```
- 继承BaseModule的Register和RegisterGo函数，注册的函数可以直接被客户端协议调用。
- 可以使用CallGo, CallSpec, CallWithCallback或Call函数直接调用某个模块的所有公有函数和注册函数；一般情况下，被调用的函数只会在被调用模块的单个（唯一）协程中进行调用。
- 可以使用RemoteCall（带回调BLpop）和RemoteCallNR（无回调PubSub）函数执行远程模块调用。
- 新增了RegisterRpc函数（注册的函数只有该模块可以远程调用），可以通过默认注册的exec和eval函数进行运行期模块数据和相关注册函数的在线更新。