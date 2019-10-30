// Code generated by automatic for 'math/big'. DO NOT EDIT.

// +build go1.12,!go1.13

package stdlibs

import (
	"math/big"
	"reflect"
)

func init() {
	Symbols["math/big"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Above":         reflect.ValueOf(big.Above),
		"AwayFromZero":  reflect.ValueOf(big.AwayFromZero),
		"Below":         reflect.ValueOf(big.Below),
		"Exact":         reflect.ValueOf(big.Exact),
		"Jacobi":        reflect.ValueOf(big.Jacobi),
		"MaxBase":       reflect.ValueOf(big.MaxBase),
		"MaxExp":        reflect.ValueOf(big.MaxExp),
		"MaxPrec":       reflect.ValueOf(uint32(big.MaxPrec)),
		"MinExp":        reflect.ValueOf(big.MinExp),
		"NewFloat":      reflect.ValueOf(big.NewFloat),
		"NewInt":        reflect.ValueOf(big.NewInt),
		"NewRat":        reflect.ValueOf(big.NewRat),
		"ParseFloat":    reflect.ValueOf(big.ParseFloat),
		"ToNearestAway": reflect.ValueOf(big.ToNearestAway),
		"ToNearestEven": reflect.ValueOf(big.ToNearestEven),
		"ToNegativeInf": reflect.ValueOf(big.ToNegativeInf),
		"ToPositiveInf": reflect.ValueOf(big.ToPositiveInf),
		"ToZero":        reflect.ValueOf(big.ToZero),

		// type definitions
		"Accuracy":     reflect.ValueOf((*big.Accuracy)(nil)),
		"ErrNaN":       reflect.ValueOf((*big.ErrNaN)(nil)),
		"Float":        reflect.ValueOf((*big.Float)(nil)),
		"Int":          reflect.ValueOf((*big.Int)(nil)),
		"Rat":          reflect.ValueOf((*big.Rat)(nil)),
		"RoundingMode": reflect.ValueOf((*big.RoundingMode)(nil)),
		"Word":         reflect.ValueOf((*big.Word)(nil)),
	}
}
