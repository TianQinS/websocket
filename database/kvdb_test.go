package database

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestKvdb(t *testing.T) {
	Initialize("./test.db")
	Put("test", map[string]int{"k": 1})
	fmt.Println(Get("test"))
	assert.Equal(t, int64(1), GetValue("test", "k").Int())
	PutExpire("test", map[string]int{"k": 2}, time.Second)
	assert.Equal(t, int64(2), GetValue("test", "k").Int())
	time.Sleep(time.Second)
	assert.Equal(t, int64(0), GetValue("test", "k").Int())
	Close()
}
