package event

import (
	"bytes"
	"encoding/gob"

	// jsoniter "github.com/json-iterator/go" // for golang version less than 1.10
	"encoding/json"
)

// MsgPacker is used to packs and unpacks messages.
type MsgPacker interface {
	Pack(topic, f string, params []interface{}) ([]byte, error)
	Unpack(data *[]byte) (*Msg, error)
}

// Msg represents a simple msg struct.
type Msg struct {
	Topic    string        `json:"tpc"`
	Func     string        `json:"func"`
	Callback string        `json:"cb"`
	Args     []interface{} `json:"args"`
}

// JsonPacker represents the normal json format.
type JsonPacker struct {
}

// GobPacker represents golang's Gob format.
type GobPacker struct {
}

func (this *JsonPacker) Pack(topic, f string, params []interface{}) ([]byte, error) {
	if msg, err := json.Marshal(&Msg{
		Topic:    topic,
		Func:     f,
		Callback: "",
		Args:     params,
	}); err == nil {
		return msg, nil
	}
	return nil, nil
}

func (this *JsonPacker) Unpack(data *[]byte) (*Msg, error) {
	msg := &Msg{}
	// decoder := json.NewDecoder(bytes.NewReader(*data))
	// decoder.UseNumber()
	if err := json.Unmarshal(*data, msg); err != nil {
		// if err := decoder.Decode(msg); err != nil {
		return nil, err
	}
	return msg, nil
}

func (this *GobPacker) Pack(topic, f string, params []interface{}) ([]byte, error) {
	buf := make([]byte, 0, 0)
	buffer := bytes.NewBuffer(buf)
	enc := gob.NewEncoder(buffer)
	if err := enc.Encode(&Msg{
		Topic:    topic,
		Func:     f,
		Callback: "",
		Args:     params,
	}); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func (this *GobPacker) Unpack(data *[]byte) (*Msg, error) {
	msg := &Msg{}
	if err := gob.NewDecoder(bytes.NewBuffer(*data)).Decode(msg); err != nil {
		return nil, err
	}
	return msg, nil
}
