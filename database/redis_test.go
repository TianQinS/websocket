package database

import (
	"fmt"
	"testing"
	"time"

	"github.com/TianQinS/websocket/config"
	"github.com/stretchr/testify/assert"
)

func TestRedis(t *testing.T) {
	conf := &config.Conf.Rdb
	rdb := NewRdb(conf)
	pubData := ""
	rdb.Rpush("k", "v")
	v, _ := rdb.Lpop("k")
	assert.Equal(t, "v", v)
	go rdb.ListenForSubscribe("p", func(channel, payload string) {
		pubData = payload
		fmt.Println(channel, payload)
	})
	time.Sleep(time.Second)
	rdb.Publish("p", "t")
	time.Sleep(time.Second)
	assert.Equal(t, "t", pubData)
}
