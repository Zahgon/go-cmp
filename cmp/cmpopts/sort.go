// Copyright 2017, The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmpopts

import (
	"reflect"

	"github.com/google/go-cmp/cmp"
)

// SortSlices returns a [cmp.Transformer] option that sorts all []V.
// The lessOrCompareFunc function must be either
// a less function of the form "func(T, T) bool" or
// a compare function of the format "func(T, T) int"
// which is used to sort any slice with element type V that is assignable to T.
//
// A less function must be:
//   - Deterministic: less(x, y) == less(x, y)
//   - Irreflexive: !less(x, x)
//   - Transitive: if !less(x, y) and !less(y, z), then !less(x, z)
//
// A compare function must be:
//   - Deterministic: compare(x, y) == compare(x, y)
//   - Irreflexive: compare(x, x) == 0
//   - Transitive: if !less(x, y) and !less(y, z), then !less(x, z)
//
// The function does not have to be "total". That is, if x != y, but
// less or compare report inequality, their relative order is maintained.
//
// SortSlices can be used in conjunction with [EquateEmpty].
func SortSlices(lessOrCompareFunc interface{}) cmp.Option {
	_ = "STUB: not implemented"
	return *new(cmp.Option)
}

type sliceSorter struct {
	in  reflect.Type  // T
	fnc reflect.Value // func(T, T) bool
}

func (ss sliceSorter) filter(x, y interface{}) bool { _ = "STUB: not implemented"; return false }

// Check whether the slices are already sorted to avoid an infinite
// recursion cycle applying the same transform to itself.

func (ss sliceSorter) sort(x interface{}) interface{} { _ = "STUB: not implemented"; return nil }

func (ss sliceSorter) checkSort(v reflect.Value) {
	_ = "STUB: not implemented"
	// Start of a sequence of equal elements.
	return
}

// Check that first and last elements in v[start:i] are equal.

func (ss sliceSorter) less(v reflect.Value, i, j int) bool { _ = "STUB: not implemented"; return false }

// SortMaps returns a [cmp.Transformer] option that flattens map[K]V types to be
// a sorted []struct{K, V}. The lessOrCompareFunc function must be either
// a less function of the form "func(T, T) bool" or
// a compare function of the format "func(T, T) int"
// which is used to sort any map with key K that is assignable to T.
//
// Flattening the map into a slice has the property that [cmp.Equal] is able to
// use [cmp.Comparer] options on K or the K.Equal method if it exists.
//
// A less function must be:
//   - Deterministic: less(x, y) == less(x, y)
//   - Irreflexive: !less(x, x)
//   - Transitive: if !less(x, y) and !less(y, z), then !less(x, z)
//   - Total: if x != y, then either less(x, y) or less(y, x)
//
// A compare function must be:
//   - Deterministic: compare(x, y) == compare(x, y)
//   - Irreflexive: compare(x, x) == 0
//   - Transitive: if compare(x, y) < 0 and compare(y, z) < 0, then compare(x, z) < 0
//   - Total: if x != y, then compare(x, y) != 0
//
// SortMaps can be used in conjunction with [EquateEmpty].
func SortMaps(lessOrCompareFunc interface{}) cmp.Option {
	_ = "STUB: not implemented"
	return *new(cmp.Option)
}

type mapSorter struct {
	in  reflect.Type  // T
	fnc reflect.Value // func(T, T) bool
}

func (ms mapSorter) filter(x, y interface{}) bool { _ = "STUB: not implemented"; return false }

func (ms mapSorter) sort(x interface{}) interface{} { _ = "STUB: not implemented"; return nil }

func (ms mapSorter) checkSort(v reflect.Value) { _ = "STUB: not implemented"; return }

func (ms mapSorter) less(v reflect.Value, i, j int) bool { _ = "STUB: not implemented"; return false }
