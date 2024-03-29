// Code generated by automatic for 'github.com/TianQinS/fastapi/basic'. DO NOT EDIT.

// +build go1.12,!go1.13

package stdlibs

import (
	"github.com/TianQinS/fastapi/basic"
	"reflect"
)

func init() {
	Symbols["github.com/TianQinS/fastapi/basic"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Catch":              reflect.ValueOf(basic.Catch),
		"CatchFunc":          reflect.ValueOf(basic.CatchFunc),
		"CatchWithParams":    reflect.ValueOf(basic.CatchWithParams),
		"CatchWithReflect":   reflect.ValueOf(basic.CatchWithReflect),
		"Exec":               reflect.ValueOf(basic.Exec),
		"HookMgr":            reflect.ValueOf(&basic.HookMgr).Elem(),
		"NewFile":            reflect.ValueOf(basic.NewFile),
		"NewQueue":           reflect.ValueOf(basic.NewQueue),
		"OVERFLOW_CHECK_NUM": reflect.ValueOf(int64(basic.OVERFLOW_CHECK_NUM)),
		"PackErrorMsg":       reflect.ValueOf(basic.PackErrorMsg),
		"TIME_FORMAT":        reflect.ValueOf(basic.TIME_FORMAT),
		"Throw":              reflect.ValueOf(basic.Throw),
		"UINT64_MAX_NUM":     reflect.ValueOf(uint64(basic.UINT64_MAX_NUM)),

		// type definitions
		"BasicHook":    reflect.ValueOf((*basic.BasicHook)(nil)),
		"EsQueue":      reflect.ValueOf((*basic.EsQueue)(nil)),
		"Func":         reflect.ValueOf((*basic.Func)(nil)),
		"FuncCallback": reflect.ValueOf((*basic.FuncCallback)(nil)),
		"Hook":         reflect.ValueOf((*basic.Hook)(nil)),
		"Hooks":        reflect.ValueOf((*basic.Hooks)(nil)),

		// interface wrapper definitions
		"_Hook": reflect.ValueOf((*_github_com_TianQinS_fastapi_basic_Hook)(nil)),
	}
}

// _github_com_TianQinS_fastapi_basic_Hook is an interface wrapper for Hook type
type _github_com_TianQinS_fastapi_basic_Hook struct {
	WFire    func(args []interface{})
	WTimeout func() bool
}

func (W _github_com_TianQinS_fastapi_basic_Hook) Fire(args []interface{}) { W.WFire(args) }
func (W _github_com_TianQinS_fastapi_basic_Hook) Timeout() bool           { return W.WTimeout() }
