// Copyright 2017, The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmpopts

import (
	"reflect"

	"github.com/google/go-cmp/cmp"
)

// IgnoreFields returns an [cmp.Option] that ignores fields of the
// given names on a single struct type. It respects the names of exported fields
// that are forwarded due to struct embedding.
// The struct type is specified by passing in a value of that type.
//
// The name may be a dot-delimited string (e.g., "Foo.Bar") to ignore a
// specific sub-field that is embedded or nested within the parent struct.
func IgnoreFields(typ interface{}, names ...string) cmp.Option {
	_ = "STUB: not implemented"
	return *new(cmp.Option)
}

// IgnoreTypes returns an [cmp.Option] that ignores all values assignable to
// certain types, which are specified by passing in a value of each type.
func IgnoreTypes(typs ...interface{}) cmp.Option {
	_ = "STUB: not implemented"
	return *new(cmp.Option)
}

type typeFilter []reflect.Type

func newTypeFilter(typs ...interface{}) (tf typeFilter) {
	_ = "STUB: not implemented"
	return *new(typeFilter)
}

// This occurs if someone tries to pass in sync.Locker(nil)

func (tf typeFilter) filter(p cmp.Path) bool { _ = "STUB: not implemented"; return false }

// IgnoreInterfaces returns an [cmp.Option] that ignores all values or references of
// values assignable to certain interface types. These interfaces are specified
// by passing in an anonymous struct with the interface types embedded in it.
// For example, to ignore [sync.Locker], pass in struct{sync.Locker}{}.
func IgnoreInterfaces(ifaces interface{}) cmp.Option {
	_ = "STUB: not implemented"
	return *new(cmp.Option)
}

type ifaceFilter []reflect.Type

func newIfaceFilter(ifaces interface{}) (tf ifaceFilter) {
	_ = "STUB: not implemented"
	return *new(ifaceFilter)
}

// This matches everything; why would you ever want this?

func (tf ifaceFilter) filter(p cmp.Path) bool { _ = "STUB: not implemented"; return false }

// IgnoreUnexported returns an [cmp.Option] that only ignores the immediate unexported
// fields of a struct, including anonymous fields of unexported types.
// In particular, unexported fields within the struct's exported fields
// of struct types, including anonymous fields, will not be ignored unless the
// type of the field itself is also passed to IgnoreUnexported.
//
// Avoid ignoring unexported fields of a type which you do not control (i.e. a
// type from another repository), as changes to the implementation of such types
// may change how the comparison behaves. Prefer a custom [cmp.Comparer] instead.
func IgnoreUnexported(typs ...interface{}) cmp.Option {
	_ = "STUB: not implemented"
	return *new(cmp.Option)
}

type unexportedFilter struct{ m map[reflect.Type]bool }

func newUnexportedFilter(typs ...interface{}) unexportedFilter {
	_ = "STUB: not implemented"
	return *new(unexportedFilter)
}

func (xf unexportedFilter) filter(p cmp.Path) bool { _ = "STUB: not implemented"; return false }

// isExported reports whether the identifier is exported.
func isExported(id string) bool { _ = "STUB: not implemented"; return false }

// IgnoreSliceElements returns an [cmp.Option] that ignores elements of []V.
// The discard function must be of the form "func(T) bool" which is used to
// ignore slice elements of type V, where V is assignable to T.
// Elements are ignored if the function reports true.
func IgnoreSliceElements(discardFunc interface{}) cmp.Option {
	_ = "STUB: not implemented"
	return *new(cmp.Option)
}

// IgnoreMapEntries returns an [cmp.Option] that ignores entries of map[K]V.
// The discard function must be of the form "func(T, R) bool" which is used to
// ignore map entries of type K and V, where K and V are assignable to T and R.
// Entries are ignored if the function reports true.
func IgnoreMapEntries(discardFunc interface{}) cmp.Option {
	_ = "STUB: not implemented"
	return *new(cmp.Option)
}
