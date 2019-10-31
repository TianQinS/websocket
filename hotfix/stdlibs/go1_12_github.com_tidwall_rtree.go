// Code generated by automatic for 'github.com/tidwall/rtree'. DO NOT EDIT.

// +build go1.12,!go1.13

package stdlibs

import (
	"github.com/tidwall/rtree"
	"reflect"
)

func init() {
	Symbols["github.com/tidwall/rtree"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"New": reflect.ValueOf(rtree.New),

		// type definitions
		"Item":     reflect.ValueOf((*rtree.Item)(nil)),
		"Iterator": reflect.ValueOf((*rtree.Iterator)(nil)),
		"RTree":    reflect.ValueOf((*rtree.RTree)(nil)),

		// interface wrapper definitions
		"_Item": reflect.ValueOf((*_github_com_tidwall_rtree_Item)(nil)),
	}
}

// _github_com_tidwall_rtree_Item is an interface wrapper for Item type
type _github_com_tidwall_rtree_Item struct {
	WRect func(ctx interface{}) (min []float64, max []float64)
}

func (W _github_com_tidwall_rtree_Item) Rect(ctx interface{}) (min []float64, max []float64) {
	return W.WRect(ctx)
}