package database

import (
	"fmt"

	"github.com/TianQinS/websocket/config"
	"github.com/TianQinS/fastapi/basic"
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
)

type Rdb struct {
	redis.UniversalClient
}

// NewRdb get a redis object in redis pool.
func NewRdb(cfg *config.Rdb) *Rdb {
	client, e := NewRedisClientByConfig(cfg)
	return RedisMust(client, e)
}

func RedisMust(client *Rdb, err error) *Rdb {
	if err != nil {
		panic(err)
	}
	return client
}

func NewRedisClientByConfig(cfg *config.Rdb) (*Rdb, error) {
	switch cfg.Type {
	case "", "simple":
		return NewRedisClientSimple(cfg)
	case "cluster":
		return NewRedisClientCluster(cfg)
	case "sentinel":
		return NewRedisClientSentinel(cfg)
	default:
		return nil, fmt.Errorf("unsupport redis type: %v", cfg.Type)
	}
}

func NewRedisClientSimple(cfg *config.Rdb) (*Rdb, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Addrs[0],
		DB:       cfg.DB,
		Password: cfg.Password,
	})
	if cfg.FailFast {
		err := client.Ping().Err()
		if err != nil {
			return nil, errors.Wrap(err, "redis simple connect fail")
		}
	}
	return &Rdb{client}, nil
}

func NewRedisClientSentinel(cfg *config.Rdb) (*Rdb, error) {
	client := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    cfg.MasterName,
		SentinelAddrs: cfg.Addrs,
		Password:      cfg.Password,
		DB:            cfg.DB,
	})
	if cfg.FailFast {
		err := client.Ping().Err()
		if err != nil {
			return nil, errors.Wrap(err, "redis sentinel connect fail")
		}
	}
	return &Rdb{client}, nil
}

func NewRedisClientCluster(cfg *config.Rdb) (*Rdb, error) {
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    cfg.Addrs,
		Password: cfg.Password,
	})
	if cfg.FailFast {
		err := client.Ping().Err()
		if err != nil {
			return nil, errors.Wrap(err, "redis cluster connect fail")
		}
	}
	return &Rdb{client}, nil
}

// Rpush push an item on a queue.
func (this *Rdb) Rpush(key string, value string) (err error) {
	_, err = this.RPush(key, value).Result()
	return
}

// Lpop.
func (this *Rdb) Lpop(key string) (value string, err error) {
	value, err = this.LPop(key).Result()
	return
}

// BLPop for chan.
func (this *Rdb) BLpop(key string) chan string {
	result := make(chan string, 1)
	go func() {
		for {
			if res, err := this.BLPop(0, key).Result(); err == nil {
				result <- res[1]
			} else {
				basic.PackErrorMsg(err, key)
			}
		}
	}()
	return result
}

// ListenForSubscribe creates a subscription by key with a process function.
// A useage of publish is this.Publish(key, data).
func (this *Rdb) ListenForSubscribe(key string, onMessage func(channal, payload string)) error {
	pubsub := this.Subscribe(key)
	if _, err := pubsub.Receive(); err != nil {
		return err
	}
	ch := pubsub.Channel()
	for msg := range ch {
		onMessage(msg.Channel, msg.Payload)
	}
	return nil
}
