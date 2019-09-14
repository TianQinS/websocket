package module

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/TianQinS/websocket/config"
	"github.com/TianQinS/websocket/database"
	"github.com/TianQinS/fastapi/basic"

	// for 10ms hook.
	_ "github.com/TianQinS/fastapi/timer"
)

var (
	// conf is a configuration for go-redis.
	conf = &config.Conf.Rdb
	// pid is the global uuid for pubsub.
	pid = config.Conf.Pid
	// Hook is a default global hookmgr.
	Hook = basic.HookMgr
)

// Msg represents a simple msg struct for remote call.
type Msg struct {
	Func          string        `json:"func"`
	Args          []interface{} `json:"args"`
	CallbackTopic string        `json:"ct"`
	Callback      string        `json:"cb"`
}

// PubMsg represents a simple msg struct for pubsub.
type PubMsg struct {
	Func string        `json:"func"`
	Args []interface{} `json:"args"`
	Pid  string        `json:"pid"`
}

type RPCModule struct {
	BaseModule
	topic string
	Rdb   *database.Rdb
	// RPCFunctions is RPCModule's exclusive as a gesture to security
	RPCFunctions map[string]interface{}
	tick         chan bool
}

func NewRPCModule(topic string, qSize uint64) *RPCModule {
	m := &RPCModule{}
	m.OnInit(topic, qSize)
	return m
}

// packMsg packing messages for redis's rpush.
func (this *RPCModule) packMsg(f, topic, callback string, params []interface{}) ([]byte, error) {
	msg, err := json.Marshal(&Msg{
		Func:          f,
		Args:          params,
		CallbackTopic: topic,
		Callback:      callback,
	})
	if err == nil {
		return msg, nil
	}
	basic.PackErrorMsg(err, f)
	return nil, err
}

func (this *RPCModule) unpackMsg(data []byte) (*Msg, error) {
	msg := &Msg{}
	if err := json.Unmarshal(data, msg); err != nil {
		basic.PackErrorMsg(err, string(data))
		return nil, err
	}
	return msg, nil
}

// packPubMsg packing messages for redis's publish.
func (this *RPCModule) packPubMsg(f string, params []interface{}) ([]byte, error) {
	msg, err := json.Marshal(&PubMsg{
		Func: f,
		Args: params,
		Pid:  pid,
	})
	if err == nil {
		return msg, nil
	}
	basic.PackErrorMsg(err, f)
	return nil, err
}

func (this *RPCModule) unpackPubMsg(data []byte) (*PubMsg, error) {
	msg := &PubMsg{}
	if err := json.Unmarshal(data, msg); err != nil {
		basic.PackErrorMsg(err, string(data))
		return nil, err
	}
	return msg, nil
}

// execute the functions in RPCFunctions, the mqtt's protocol is not allowed to these methods.
func (this *RPCModule) execute(f string, args []interface{}, callbackTopic, callback string) {
	if function, ok := this.RPCFunctions[f]; ok {
		defer func() {
			if e, ok := recover().(error); ok {
				basic.PackErrorMsg(e, f)
			}
		}()
		_f := reflect.ValueOf(function)
		in := make([]reflect.Value, len(args))
		for k := range in {
			in[k] = reflect.ValueOf(args[k])
		}

		res := _f.Call(in)
		ret := make([]interface{}, 0)
		for _, atom := range res {
			ret = append(ret, atom.Interface())
		}
		if callback != "" {
			this.RemoteCallNR(callbackTopic, callback, ret...)
		}
	}
}

// onSubscribe processing subscription messages for this module.
func (this *RPCModule) onSubscribe(channel, payload string) {
	msg, err := this.unpackPubMsg([]byte(payload))
	if err != nil {
		return
	}
	if msg.Pid != pid {
		this.execute(msg.Func, msg.Args, "", "")
	}
}

// CallOther call other module's function by RPCModule.
func (this *RPCModule) CallOther(f, otherTopic, callbackTopic, callback string, args ...interface{}) (err error) {
	if callback != "" {
		err = this.GetMgr().CallWithCallback(otherTopic, f, this.RemoteCallNR, []interface{}{callbackTopic, callback}, args)
	} else {
		err = this.GetMgr().Call(otherTopic, f, args)
	}
	return
}

func (this *RPCModule) CallOtherNR(f, otherTopic string, args ...interface{}) error {
	return this.GetMgr().Call(otherTopic, f, args)
}

// Publish a message, the msg's Pid is a publisher assertion.
func (this *RPCModule) RemoteCallNR(topic, f string, args ...interface{}) {
	msg, err := this.packPubMsg(f, args)
	if err != nil {
		return
	}
	this.Rdb.Publish(topic, string(msg))
}

func (this *RPCModule) onCall(data string) {
	msg, err := this.unpackMsg([]byte(data))
	if err != nil {
		return
	}
	this.execute(msg.Func, msg.Args, msg.CallbackTopic, msg.Callback)
}

func (this *RPCModule) RemoteCall(topic, f, callback string, args ...interface{}) {
	callbackTopic := ""
	if callback != "" {
		callbackTopic = this.topic
	}
	msg, err := this.packMsg(f, callbackTopic, callback, args)
	if err != nil {
		return
	}
	this.Rdb.Rpush(topic, string(msg))
}

// forTest is used for scripts's sample2.
func (this *RPCModule) forTest(arg1, arg2 string) string {
	fmt.Println(arg1, arg2)
	return fmt.Sprintf("args: %s, %s.", arg1, arg2)
}

func (this *RPCModule) GetTopic() string {
	return this.topic
}

func (this *RPCModule) OnInit(topic string, qSize uint64) {
	this.topic = topic
	this.BaseModule.Init(qSize)
	this.Rdb = database.NewRdb(conf)
	this.tick = make(chan bool, qSize)
	this.RPCFunctions = make(map[string]interface{}, 0)
	// for test script.
	this.RegisterRpc("test", this.forTest)
	this.RegisterRpc("callOther", this.CallOther)
	this.RegisterRpc("callOtherNR", this.CallOtherNR)
}

// RegisterRpc register a function only for remote call by redis.
func (this *RPCModule) RegisterRpc(id string, f interface{}) {
	if _, ok := this.RPCFunctions[id]; ok {
		fmt.Printf("RPCFunction id %v: already registered\n", id)
	}
	this.RPCFunctions[id] = f
}

// onTick called every ten ms.
func (this *RPCModule) onTick(args ...interface{}) {
	if this.IsRun {
		this.tick <- true
	}
}

func (this *RPCModule) closeTick() {
	this.IsRun = false
	close(this.tick)
}

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

func (this *RPCModule) OnDestroy() {
	this.BaseModule.Close()
}
