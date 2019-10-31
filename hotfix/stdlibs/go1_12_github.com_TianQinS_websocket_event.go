// Code generated by automatic for 'github.com/TianQinS/websocket/event'. DO NOT EDIT.

// +build go1.12,!go1.13

package stdlibs

import (
	"github.com/TianQinS/websocket/event"
	"reflect"
)

func init() {
	Symbols["github.com/TianQinS/websocket/event"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Db":             reflect.ValueOf(&event.Db).Elem(),
		"Hook":           reflect.ValueOf(&event.Hook).Elem(),
		"MaxSleepTime":   reflect.ValueOf(&event.MaxSleepTime).Elem(),
		"Mdb":            reflect.ValueOf(&event.Mdb).Elem(),
		"NewClient":      reflect.ValueOf(event.NewClient),
		"NewEventMgr":    reflect.ValueOf(event.NewEventMgr),
		"OpText":         reflect.ValueOf(&event.OpText).Elem(),
		"Packer":         reflect.ValueOf(&event.Packer).Elem(),
		"Post":           reflect.ValueOf(&event.Post).Elem(),
		"ST_CONNECTING":  reflect.ValueOf(event.ST_CONNECTING),
		"ST_ESTABLISHED": reflect.ValueOf(event.ST_ESTABLISHED),
		"ST_STOP":        reflect.ValueOf(event.ST_STOP),
		"UserTable":      reflect.ValueOf(&event.UserTable).Elem(),

		// type definitions
		"Client":     reflect.ValueOf((*event.Client)(nil)),
		"EventMgr":   reflect.ValueOf((*event.EventMgr)(nil)),
		"GobPacker":  reflect.ValueOf((*event.GobPacker)(nil)),
		"JsonPacker": reflect.ValueOf((*event.JsonPacker)(nil)),
		"Msg":        reflect.ValueOf((*event.Msg)(nil)),
		"MsgPacker":  reflect.ValueOf((*event.MsgPacker)(nil)),

		// interface wrapper definitions
		"_MsgPacker": reflect.ValueOf((*_github_com_TianQinS_websocket_event_MsgPacker)(nil)),
	}
}

// _github_com_TianQinS_websocket_event_MsgPacker is an interface wrapper for MsgPacker type
type _github_com_TianQinS_websocket_event_MsgPacker struct {
	WPack   func(topic string, f string, params []interface{}) ([]byte, error)
	WUnpack func(data *[]byte) (*event.Msg, error)
}

func (W _github_com_TianQinS_websocket_event_MsgPacker) Pack(topic string, f string, params []interface{}) ([]byte, error) {
	return W.WPack(topic, f, params)
}
func (W _github_com_TianQinS_websocket_event_MsgPacker) Unpack(data *[]byte) (*event.Msg, error) {
	return W.WUnpack(data)
}