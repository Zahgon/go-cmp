// Copyright 2017, The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package function provides functionality for identifying function types.
package function

import (
	"reflect"
	"regexp"
)

type funcType int

const (
	_ funcType = iota

	tbFunc  // func(T) bool
	ttbFunc // func(T, T) bool
	ttiFunc // func(T, T) int
	trbFunc // func(T, R) bool
	tibFunc // func(T, I) bool
	trFunc  // func(T) R

	Equal             = ttbFunc // func(T, T) bool
	EqualAssignable   = tibFunc // func(T, I) bool; encapsulates func(T, T) bool
	Transformer       = trFunc  // func(T) R
	ValueFilter       = ttbFunc // func(T, T) bool
	Less              = ttbFunc // func(T, T) bool
	Compare           = ttiFunc // func(T, T) int
	ValuePredicate    = tbFunc  // func(T) bool
	KeyValuePredicate = trbFunc // func(T, R) bool
)

var boolType = reflect.TypeOf(true)
var intType = reflect.TypeOf(0)

// IsType reports whether the reflect.Type is of the specified function type.
func IsType(t reflect.Type, ft funcType) bool { _ = "STUB: not implemented"; return false }

// func(T) bool

// func(T, T) bool

// func(T, T) int

// func(T, R) bool

// func(T, I) bool

// func(T) R

var lastIdentRx = regexp.MustCompile(`[_\p{L}][_\p{L}\p{N}]*$`)

// NameOf returns the name of the function value.
func NameOf(v reflect.Value) string { _ = "STUB: not implemented"; return "" }

// e.g., "long/path/name/mypkg.(*MyType).(long/path/name/mypkg.myMethod)-fm"

// Method closures have a "-fm" suffix.
