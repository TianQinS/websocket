// Code generated by automatic for 'github.com/TianQinS/websocket/config'. DO NOT EDIT.

// +build go1.12,!go1.13

package stdlibs

import (
	"github.com/TianQinS/websocket/config"
	"reflect"
)

func init() {
	Symbols["github.com/TianQinS/websocket/config"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"CONF_FILE": reflect.ValueOf(config.CONF_FILE),
		"Conf":      reflect.ValueOf(&config.Conf).Elem(),

		// type definitions
		"Config": reflect.ValueOf((*config.Config)(nil)),
		"Hotfix": reflect.ValueOf((*config.Hotfix)(nil)),
		"Kdb":    reflect.ValueOf((*config.Kdb)(nil)),
		"Mdb":    reflect.ValueOf((*config.Mdb)(nil)),
		"Rdb":    reflect.ValueOf((*config.Rdb)(nil)),
		"Web":    reflect.ValueOf((*config.Web)(nil)),
	}
}
