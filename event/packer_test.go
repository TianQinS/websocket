package event

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJson(t *testing.T) {
	packer := &JsonPacker{}
	b, err := packer.Pack("topic", "f1", []interface{}{1, 2})
	assert.Equal(t, nil, err)
	msg, err := packer.Unpack(&b)
	assert.Equal(t, nil, err)
	assert.Equal(t, "f1", msg.Func)
	assert.Equal(t, float64(1), msg.Args[0].(float64))
	assert.Equal(t, float64(2), msg.Args[1].(float64))
}

func TestGob(t *testing.T) {
	packer := &GobPacker{}
	b, err := packer.Pack("topic", "f1", []interface{}{1, 2})
	assert.Equal(t, nil, err)
	msg, err := packer.Unpack(&b)
	assert.Equal(t, nil, err)
	assert.Equal(t, "f1", msg.Func)
	assert.Equal(t, 1, msg.Args[0].(int))
	assert.Equal(t, 2, msg.Args[1].(int))
}
