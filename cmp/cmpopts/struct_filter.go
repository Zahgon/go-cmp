// Copyright 2017, The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmpopts

import (
	"reflect"

	"github.com/google/go-cmp/cmp"
)

// filterField returns a new Option where opt is only evaluated on paths that
// include a specific exported field on a single struct type.
// The struct type is specified by passing in a value of that type.
//
// The name may be a dot-delimited string (e.g., "Foo.Bar") to select a
// specific sub-field that is embedded or nested within the parent struct.
func filterField(typ interface{}, name string, opt cmp.Option) cmp.Option {
	_ = "STUB: not implemented"
	// TODO: This is currently unexported over concerns of how helper filters
	// can be composed together easily.
	// TODO: Add tests for FilterField.
	return *new(cmp.Option)
}

type structFilter struct {
	t  reflect.Type // The root struct type to match on
	ft fieldTree    // Tree of fields to match on
}

func newStructFilter(typ interface{}, names ...string) structFilter {
	_ = "STUB: not implemented"
	// TODO: Perhaps allow * as a special identifier to allow ignoring any
	// number of path steps until the next field match?
	// This could be useful when a concrete struct gets transformed into
	// an anonymous struct where it is not possible to specify that by type,
	// but the transformer happens to provide guarantees about the names of
	// the transformed fields.
	return *new(structFilter)
}

func (sf structFilter) filter(p cmp.Path) bool { _ = "STUB: not implemented"; return false }

// fieldTree represents a set of dot-separated identifiers.
//
// For example, inserting the following selectors:
//
//	Foo
//	Foo.Bar.Baz
//	Foo.Buzz
//	Nuka.Cola.Quantum
//
// Results in a tree of the form:
//
//	{sub: {
//		"Foo": {ok: true, sub: {
//			"Bar": {sub: {
//				"Baz": {ok: true},
//			}},
//			"Buzz": {ok: true},
//		}},
//		"Nuka": {sub: {
//			"Cola": {sub: {
//				"Quantum": {ok: true},
//			}},
//		}},
//	}}
type fieldTree struct {
	ok  bool                 // Whether this is a specified node
	sub map[string]fieldTree // The sub-tree of fields under this node
}

// insert inserts a sequence of field accesses into the tree.
func (ft *fieldTree) insert(cname []string) { _ = "STUB: not implemented"; return }

// matchPrefix reports whether any selector in the fieldTree matches
// the start of path p.
func (ft fieldTree) matchPrefix(p cmp.Path) bool { _ = "STUB: not implemented"; return false }

// canonicalName returns a list of identifiers where any struct field access
// through an embedded field is expanded to include the names of the embedded
// types themselves.
//
// For example, suppose field "Foo" is not directly in the parent struct,
// but actually from an embedded struct of type "Bar". Then, the canonical name
// of "Foo" is actually "Bar.Foo".
//
// Suppose field "Foo" is not directly in the parent struct, but actually
// a field in two different embedded structs of types "Bar" and "Baz".
// Then the selector "Foo" causes a panic since it is ambiguous which one it
// refers to. The user must specify either "Bar.Foo" or "Baz.Foo".
func canonicalName(t reflect.Type, sel string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Type must be a struct or pointer to struct.

// Find the canonical name for this current field name.
// If the field exists in an embedded struct, then it will be expanded.

// Avoid using reflect.Type.FieldByName for unexported fields due to
// buggy behavior with regard to embeddeding and unexported fields.
// See https://golang.org/issue/4876 for details.
