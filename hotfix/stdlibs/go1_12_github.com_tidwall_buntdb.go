// Code generated by automatic for 'github.com/tidwall/buntdb'. DO NOT EDIT.

// +build go1.12,!go1.13

package stdlibs

import (
	"github.com/tidwall/buntdb"
	"reflect"
)

func init() {
	Symbols["github.com/tidwall/buntdb"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Always":                 reflect.ValueOf(buntdb.Always),
		"Desc":                   reflect.ValueOf(buntdb.Desc),
		"ErrDatabaseClosed":      reflect.ValueOf(&buntdb.ErrDatabaseClosed).Elem(),
		"ErrIndexExists":         reflect.ValueOf(&buntdb.ErrIndexExists).Elem(),
		"ErrInvalid":             reflect.ValueOf(&buntdb.ErrInvalid).Elem(),
		"ErrInvalidOperation":    reflect.ValueOf(&buntdb.ErrInvalidOperation).Elem(),
		"ErrInvalidSyncPolicy":   reflect.ValueOf(&buntdb.ErrInvalidSyncPolicy).Elem(),
		"ErrNotFound":            reflect.ValueOf(&buntdb.ErrNotFound).Elem(),
		"ErrPersistenceActive":   reflect.ValueOf(&buntdb.ErrPersistenceActive).Elem(),
		"ErrShrinkInProcess":     reflect.ValueOf(&buntdb.ErrShrinkInProcess).Elem(),
		"ErrTxClosed":            reflect.ValueOf(&buntdb.ErrTxClosed).Elem(),
		"ErrTxIterating":         reflect.ValueOf(&buntdb.ErrTxIterating).Elem(),
		"ErrTxNotWritable":       reflect.ValueOf(&buntdb.ErrTxNotWritable).Elem(),
		"EverySecond":            reflect.ValueOf(buntdb.EverySecond),
		"IndexBinary":            reflect.ValueOf(buntdb.IndexBinary),
		"IndexFloat":             reflect.ValueOf(buntdb.IndexFloat),
		"IndexInt":               reflect.ValueOf(buntdb.IndexInt),
		"IndexJSON":              reflect.ValueOf(buntdb.IndexJSON),
		"IndexJSONCaseSensitive": reflect.ValueOf(buntdb.IndexJSONCaseSensitive),
		"IndexRect":              reflect.ValueOf(buntdb.IndexRect),
		"IndexString":            reflect.ValueOf(buntdb.IndexString),
		"IndexUint":              reflect.ValueOf(buntdb.IndexUint),
		"Match":                  reflect.ValueOf(buntdb.Match),
		"Never":                  reflect.ValueOf(buntdb.Never),
		"Open":                   reflect.ValueOf(buntdb.Open),
		"Point":                  reflect.ValueOf(buntdb.Point),
		"Rect":                   reflect.ValueOf(buntdb.Rect),

		// type definitions
		"Config":       reflect.ValueOf((*buntdb.Config)(nil)),
		"DB":           reflect.ValueOf((*buntdb.DB)(nil)),
		"IndexOptions": reflect.ValueOf((*buntdb.IndexOptions)(nil)),
		"SetOptions":   reflect.ValueOf((*buntdb.SetOptions)(nil)),
		"SyncPolicy":   reflect.ValueOf((*buntdb.SyncPolicy)(nil)),
		"Tx":           reflect.ValueOf((*buntdb.Tx)(nil)),
	}
}
