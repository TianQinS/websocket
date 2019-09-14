package event

import (
	"fmt"

	"github.com/TianQinS/websocket/config"
	"github.com/TianQinS/websocket/database"
)

var (
	// Mdb represents a mongo connection for data storage.
	Mdb *database.MDB
	// Db represents a local kvdb object.
	Db = database.DB
	// conf is the mongo database configuration.
	conf = config.Conf.Mdb
)

func init() {
	if len(conf.Url) > 0 {
		var err error
		Mdb, err = database.NewMongoDB(conf.Url, conf.Db)
		if err != nil {
			fmt.Printf("Mongo config error url=%s db=%s err=%s\n", conf.Url, conf.Db, err)
		}
	}
}
