package database

import (
	"time"

	"github.com/TianQinS/websocket/config"
	// jsoniter "github.com/json-iterator/go"
	"encoding/json"

	"github.com/tidwall/buntdb"
	"github.com/tidwall/gjson"
)

var (
	// DB is a global kvdb object for normal use.
	DB *buntdb.DB
)

// CreateIndex create indexes for JSON documents.
func CreateIndex(name string, indexes ...string) {
	var less []func(a, b string) bool
	for _, index := range indexes {
		less = append(less, buntdb.IndexJSON(index))
	}
	DB.CreateIndex(name, "*", less...)
}

// Put set a value with a read/write transaction,
// there can only be one read/write transaction running at a time.
func Put(key string, val interface{}) (err error) {
	var dat []byte
	dat, err = json.Marshal(val)
	if err == nil {
		return DB.Update(func(tx *buntdb.Tx) error {
			_, _, err = tx.Set(key, string(dat), nil)
			return err
		})
	}
	return
}

// PutExpire put a mapping data with expiration time.
func PutExpire(key string, val interface{}, expire time.Duration) (err error) {
	var dat []byte
	dat, err = json.Marshal(val)
	if err == nil {
		return DB.Update(func(tx *buntdb.Tx) error {
			_, _, err = tx.Set(key, string(dat), &buntdb.SetOptions{Expires: true, TTL: expire})
			return err
		})
	}
	return
}

func Delete(key string) error {
	return DB.Update(func(tx *buntdb.Tx) error {
		_, err := tx.Delete(key)
		return err
	})
}

// Get grab the JSON data with type conversion, this read-only transaction should be used when you don't need to make changes to the data.
// the advantage of a read-only transaction is that there can be many running concurrently.
func Get(key string) map[string]interface{} {
	var dat map[string]interface{}
	DB.View(func(tx *buntdb.Tx) error {
		val, err := tx.Get(key)
		if err == nil {
			if m, ok := gjson.Parse(val).Value().(map[string]interface{}); ok {
				dat = m
			}
		}
		return nil
	})
	return dat
}

// GetValue read a field in JSON data.
func GetValue(key string, path string) gjson.Result {
	var dat gjson.Result
	DB.View(func(tx *buntdb.Tx) error {
		val, err := tx.Get(key)
		if err == nil {
			dat = gjson.Get(val, path)
		}
		return nil
	})
	return dat
}

func Close() {
	if DB != nil {
		DB.Close()
	}
	DB = nil
}

func Initialize(fpath string) error {
	var err error
	Close()
	// buntdb.Open(":memory:")
	DB, err = buntdb.Open(fpath)
	return err
}

func init() {
	conf := config.Conf.Kdb
	Initialize(conf.Path)
}
