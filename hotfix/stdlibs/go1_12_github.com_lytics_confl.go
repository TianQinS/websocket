// Code generated by automatic for 'github.com/lytics/confl'. DO NOT EDIT.

// +build go1.12,!go1.13

package stdlibs

import (
	"github.com/lytics/confl"
	"reflect"
)

func init() {
	Symbols["github.com/lytics/confl"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Decode":        reflect.ValueOf(confl.Decode),
		"DecodeFile":    reflect.ValueOf(confl.DecodeFile),
		"DecodeReader":  reflect.ValueOf(confl.DecodeReader),
		"IdentityChars": reflect.ValueOf(&confl.IdentityChars).Elem(),
		"Marshal":       reflect.ValueOf(confl.Marshal),
		"NewDecoder":    reflect.ValueOf(confl.NewDecoder),
		"NewEncoder":    reflect.ValueOf(confl.NewEncoder),
		"Parse":         reflect.ValueOf(confl.Parse),
		"Unmarshal":     reflect.ValueOf(confl.Unmarshal),

		// type definitions
		"Decoder":         reflect.ValueOf((*confl.Decoder)(nil)),
		"Encoder":         reflect.ValueOf((*confl.Encoder)(nil)),
		"Key":             reflect.ValueOf((*confl.Key)(nil)),
		"MetaData":        reflect.ValueOf((*confl.MetaData)(nil)),
		"Primitive":       reflect.ValueOf((*confl.Primitive)(nil)),
		"TextMarshaler":   reflect.ValueOf((*confl.TextMarshaler)(nil)),
		"TextUnmarshaler": reflect.ValueOf((*confl.TextUnmarshaler)(nil)),

		// interface wrapper definitions
		"_TextMarshaler":   reflect.ValueOf((*_github_com_lytics_confl_TextMarshaler)(nil)),
		"_TextUnmarshaler": reflect.ValueOf((*_github_com_lytics_confl_TextUnmarshaler)(nil)),
	}
}

// _github_com_lytics_confl_TextMarshaler is an interface wrapper for TextMarshaler type
type _github_com_lytics_confl_TextMarshaler struct {
	WMarshalText func() (text []byte, err error)
}

func (W _github_com_lytics_confl_TextMarshaler) MarshalText() (text []byte, err error) {
	return W.WMarshalText()
}

// _github_com_lytics_confl_TextUnmarshaler is an interface wrapper for TextUnmarshaler type
type _github_com_lytics_confl_TextUnmarshaler struct {
	WUnmarshalText func(text []byte) error
}

func (W _github_com_lytics_confl_TextUnmarshaler) UnmarshalText(text []byte) error {
	return W.WUnmarshalText(text)
}
