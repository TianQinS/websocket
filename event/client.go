package event

import (
	"fmt"
	"sync"
	"time"

	// "github.com/eclipse/paho.mqtt.golang/packets"
	"github.com/TianQinS/evio"
	"github.com/TianQinS/fastapi/basic"
	"github.com/TianQinS/websocket/config"
	kdb "github.com/TianQinS/websocket/database"
	"github.com/TianQinS/websocket/module"
	"github.com/gobwas/ws"
)

const (
	// ST_CONNECTING represents network connection state.
	ST_CONNECTING = iota
	ST_ESTABLISHED
	ST_STOP
)

var (
	Packer MsgPacker

	// UserTable represents the user table in mongo.
	UserTable = config.Conf.Mdb.User
	// Hook is a default global hookmgr.
	Hook = basic.HookMgr

	// StateServerSide = ws.StateServerSide
	OpText = ws.OpText
)

func init() {
	Packer = &JsonPacker{}
}

// Client represents a user object for one connection.
type Client struct {
	Index int
	// Guid(GUID) represents user's globally unique identifier.
	Guid  string
	Addr  string
	state int
	apps  *module.ModuleManager
	event *EventMgr
	conn  *evio.Conn
	// runVar store runtime variable.
	runVar map[string]interface{}
	// runInt store runtime variable only for integer type.
	runInt map[string]int
	// permVar represents user's local data.
	permVar map[string]interface{}
	// mdbVar represents user's mongo data.
	mdbVar   map[string]interface{}
	runLock  *sync.RWMutex
	permLock *sync.RWMutex
}

// NewClient make a new client object.
func NewClient(ec *evio.Conn, ev *EventMgr, apps *module.ModuleManager) *Client {
	return &Client{
		Index:    (*ec).AddrIndex(),
		Guid:     "",
		Addr:     (*ec).RemoteAddr().String(),
		state:    ST_CONNECTING,
		conn:     ec,
		apps:     apps,
		event:    ev,
		runVar:   make(map[string]interface{}),
		runInt:   make(map[string]int),
		permVar:  make(map[string]interface{}),
		mdbVar:   make(map[string]interface{}),
		runLock:  new(sync.RWMutex),
		permLock: new(sync.RWMutex),
	}
}

func (this *Client) updateMap(mp1, mp2 map[string]interface{}) {
	for k, v := range mp2 {
		mp1[k] = v
	}
}

func (this *Client) Stat() int {
	return this.state
}

// GetRunInt get an integer temporary variable.
func (this *Client) GetRunInt(key string) int {
	defer this.runLock.RUnlock()
	this.runLock.RLock()
	if val, ok := this.runInt[key]; ok {
		return val
	}
	return 0
}

// SetRunInt set an integer temporary variable.
func (this *Client) SetRunInt(key string, val int) {
	defer this.runLock.Unlock()
	this.runLock.Lock()
	this.runInt[key] = val
}

func (this *Client) DelRunInt(key string) {
	defer this.runLock.Unlock()
	this.runLock.Lock()
	delete(this.runInt, key)
}

// GetRunVar get a generic data variable.
func (this *Client) GetRunVar(key string) interface{} {
	defer this.runLock.RUnlock()
	this.runLock.RLock()
	if val, ok := this.runVar[key]; ok {
		return val
	}
	return nil
}

// SetRunVar set a generic data variable.
func (this *Client) SetRunVar(key string, val interface{}) {
	defer this.runLock.Unlock()
	this.runLock.Lock()
	this.runVar[key] = val
}

// GetPermVar get a local variable in local storage.
func (this *Client) GetPermVar(key string) interface{} {
	defer this.permLock.RUnlock()
	this.permLock.RLock()
	if val, ok := this.permVar[key]; ok {
		return val
	}
	return nil
}

func (this *Client) GetPermFloat64(key string) float64 {
	if val := this.GetPermVar(key); val != nil {
		return val.(float64)
	}
	return 0
}

// GetPermInt parse the interface{} to int with losing accuracy.
func (this *Client) GetPermInt(key string) int {
	if val := this.GetPermVar(key); val != nil {
		switch val.(type) {
		case float64:
			return int(val.(float64))
		case int:
			return val.(int)
		}
	}
	return 0
}

// SetPermVar set a local variable and save the whole permVar mapping in local storage.
func (this *Client) SetPermVar(key string, val interface{}) {
	defer this.permLock.Unlock()
	this.permLock.Lock()
	this.permVar[key] = val
	this.KSave()
}

func (this *Client) GetDayCounterData() map[string]interface{} {
	var mpDat map[string]interface{}
	today := time.Now().Format("2006-01-02")
	if mxDat := this.GetPermVar("dCnt"); mxDat != nil {
		mpDat = mxDat.(map[string]interface{})
		if day := mpDat["day"]; day != today {
			mpDat = map[string]interface{}{"day": today}
			this.SetPermVar("dCnt", mpDat)
		}
		return mpDat
	} else {
		mpDat = map[string]interface{}{"day": today}
	}
	return mpDat
}

func (this *Client) SetDayCounter(key string, val float64) {
	mpDat := this.GetDayCounterData()
	mpDat[key] = val
	this.SetPermVar("dCnt", mpDat)
}

// GerDayCounter get a daily variable.
func (this *Client) GetDayCounter(key string) float64 {
	mpDat := this.GetDayCounterData()
	if val, ok := mpDat[key]; ok {
		return val.(float64)
	}
	return 0.0
}

// GetMdbVar get a variable in the Mongo collection for user.
func (this *Client) GetMdbVar(key string) interface{} {
	defer this.permLock.RUnlock()
	this.permLock.RLock()
	if val, ok := this.mdbVar[key]; ok {
		return val
	}
	return nil
}

func (this *Client) GetMdbInt(key string) int {
	if val := this.GetMdbVar(key); val != nil {
		switch val.(type) {
		case float64:
			return int(val.(float64))
		case int:
			return val.(int)
		}
	}
	return 0
}

// SetMdbVar set a variable for mdbVar and update it in database.
func (this *Client) SetMdbVar(key string, val interface{}) {
	defer this.permLock.Unlock()
	this.permLock.Lock()
	this.mdbVar[key] = val
	this.MSave(map[string]interface{}{key: val})
}

// KSave save the permVar to local file system.
func (this *Client) KSave() error {
	if len(this.permVar) > 0 {
		return kdb.Put(this.Guid, this.permVar)
	}
	return nil
}

// KLoad load the permVar after login.
func (this *Client) KLoad() error {
	dat := kdb.Get(this.Guid)
	if dat != nil {
		this.updateMap(this.permVar, dat)
	}
	return nil
}

func (this *Client) MSave(data map[string]interface{}) bool {
	if Mdb != nil {
		if data == nil {
			data = this.mdbVar
		}
		if len(data) > 0 {
			query := map[string]interface{}{"guid": this.Guid}
			Mdb.UpsertOne(UserTable, query, data, nil)
			return true
		}
	}
	return false
}

// MLoad load data from database according to guid, this function will be called before KLoad function.
func (this *Client) MLoad(callback interface{}) bool {
	if Mdb != nil {
		query := map[string]interface{}{"guid": this.Guid}
		Mdb.FindOne(UserTable, query, callback)
		return true
	}
	return false
}

// onMLoad is a callback function attached to the login query.
func (this *Client) onMLoad(err error, res map[string]interface{}) {
	if err == nil {
		this.updateMap(this.mdbVar, res)
		pwdMd5 := this.mdbVar["md5"]
		if pwdMd5 != nil && pwdMd5.(string) == this.runVar["md5"].(string) {
			this.KLoad()
			this.state = ST_ESTABLISHED
			Post.PutQueueSpec(this.OnLogin)
			return
		}
	}
	this.state = ST_STOP
}

// auth process the login authentication of user.
func (this *Client) auth(msg *Msg) error {
	defer func() {
		if e, ok := recover().(error); ok {
			basic.PackErrorMsg(e, msg)
		}
	}()
	args := msg.Args
	if msg.Func == "login" && len(args) == 2 {
		this.Guid = args[0].(string)
		this.runVar["md5"] = args[1]
		this.state = ST_ESTABLISHED
		if !this.MLoad(this.onMLoad) {
			this.KLoad()
			Post.PutQueueSpec(this.OnLogin)
			return nil
		}
	} else {
		this.state = ST_STOP
	}
	return nil
}

// OnData called when the event's Data coming.
func (this *Client) OnData(in *[]byte) (out []byte, action evio.Action) {
	msg, err := Packer.Unpack(in)
	// fmt.Printf("%+v\n", msg)
	switch this.state {
	case ST_ESTABLISHED:
		if err == nil {
			args := msg.Args
			callback := msg.Callback
			// The first argument is a pointer to the client by default.
			args = append([]interface{}{this}, args...)
			if callback != "" {
				err = this.apps.CallWithCallback(msg.Topic, msg.Func, this.ClientCallback, []interface{}{callback}, args)
			} else {
				err = this.apps.Call(msg.Topic, msg.Func, args)
			}
		} else {
			basic.PackErrorMsg(err, nil)
		}
	case ST_CONNECTING:
		// authentication for new client.
		if err := this.auth(msg); err != nil {
			basic.PackErrorMsg(err, nil)
			action = evio.Close
		}
	case ST_STOP:
		action = evio.Close
	default:
		action = evio.Close
	}
	return
}

func (this *Client) OnLogin() {
	this.event.Users[this.Guid] = this
	Hook.Fire("afterLogin", this)
}

// OnClose called when the client connection is closed.
func (this *Client) OnClose() {
	err := this.KSave()
	if err == nil {
		Hook.Fire("beforeQuit", this)
		this.MSave(nil)
		delete(this.event.Users, this.Guid)
		this.conn = nil
		this.event = nil
		this.apps = nil
	}
}

// WriteMsg pushes an message to client.
func (this *Client) WriteMsg(body []byte) (err error) {
	if this.state == ST_ESTABLISHED {
		// wsutil.WriteMessage(*this.conn, StateServerSide, OpText, body)
		frame := ws.NewFrame(OpText, true, body)
		err = ws.WriteFrame(*this.conn, frame)
	}
	return
}

// CallNR call a registration function in remote client without callback.
func (this *Client) CallNR(topic, f string, params ...interface{}) error {
	body, err := Packer.Pack(topic, f, params)
	if err == nil {
		err = this.WriteMsg(body)
	}
	return err
}

func (this *Client) ClientCallback(f string, params ...interface{}) error {
	// fmt.Println(f, params)
	return this.CallNR("", f, params...)
}

func init() {
	Post.Register("test", func(client *Client, arg1, arg2 string) string {
		fmt.Println(arg1, arg2)
		return arg1 + arg2
	})
}
