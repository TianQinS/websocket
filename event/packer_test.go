package event

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJson(t *testing.T) {
	var packer MsgPacker
	var f string
	var args []interface{}
	packer = &JsonPacker{}
	b, err := packer.Pack("f1", []interface{}{1, 2})
	assert.Equal(t, nil, err)
	f, _, args, err = packer.Unpack(b)
	assert.Equal(t, nil, err)
	assert.Equal(t, "f1", f)
	assert.Equal(t, float64(1), args[0].(float64))
	assert.Equal(t, float64(2), args[1].(float64))
}

func TestGob(t *testing.T) {
	var packer MsgPacker
	var f string
	var args []interface{}
	packer = &GobPacker{}
	b, err := packer.Pack("f1", []interface{}{1, 2})
	assert.Equal(t, nil, err)
	f, _, args, err = packer.Unpack(b)
	assert.Equal(t, nil, err)
	assert.Equal(t, "f1", f)
	assert.Equal(t, 1, args[0].(int))
	assert.Equal(t, 2, args[1].(int))
}
