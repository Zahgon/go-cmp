// Copyright 2017, The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmp

import (
	"reflect"

	"github.com/google/go-cmp/cmp/internal/value"
)

// Path is a list of [PathStep] describing the sequence of operations to get
// from some root type to the current position in the value tree.
// The first Path element is always an operation-less [PathStep] that exists
// simply to identify the initial type.
//
// When traversing structs with embedded structs, the embedded struct will
// always be accessed as a field before traversing the fields of the
// embedded struct themselves. That is, an exported field from the
// embedded struct will never be accessed directly from the parent struct.
type Path []PathStep

// PathStep is a union-type for specific operations to traverse
// a value's tree structure. Users of this package never need to implement
// these types as values of this type will be returned by this package.
//
// Implementations of this interface:
//   - [StructField]
//   - [SliceIndex]
//   - [MapIndex]
//   - [Indirect]
//   - [TypeAssertion]
//   - [Transform]
type PathStep interface {
	String() string

	// Type is the resulting type after performing the path step.
	Type() reflect.Type

	// Values is the resulting values after performing the path step.
	// The type of each valid value is guaranteed to be identical to Type.
	//
	// In some cases, one or both may be invalid or have restrictions:
	//   - For StructField, both are not interface-able if the current field
	//     is unexported and the struct type is not explicitly permitted by
	//     an Exporter to traverse unexported fields.
	//   - For SliceIndex, one may be invalid if an element is missing from
	//     either the x or y slice.
	//   - For MapIndex, one may be invalid if an entry is missing from
	//     either the x or y map.
	//
	// The provided values must not be mutated.
	Values() (vx, vy reflect.Value)
}

var (
	_ PathStep = StructField{}
	_ PathStep = SliceIndex{}
	_ PathStep = MapIndex{}
	_ PathStep = Indirect{}
	_ PathStep = TypeAssertion{}
	_ PathStep = Transform{}
)

func (pa *Path) push(s PathStep) { _ = "STUB: not implemented"; return }

func (pa *Path) pop() { _ = "STUB: not implemented"; return }

// Last returns the last [PathStep] in the Path.
// If the path is empty, this returns a non-nil [PathStep]
// that reports a nil [PathStep.Type].
func (pa Path) Last() PathStep {
	_ = "STUB: not implemented"
	return *

	// Index returns the ith step in the Path and supports negative indexing.
	// A negative index starts counting from the tail of the Path such that -1
	// refers to the last step, -2 refers to the second-to-last step, and so on.
	// If index is invalid, this returns a non-nil [PathStep]
	// that reports a nil [PathStep.Type].
	new(PathStep)
}

func (pa Path) Index(i int) PathStep { _ = "STUB: not implemented"; return *new(PathStep) }

// String returns the simplified path to a node.
// The simplified path only contains struct field accesses.
//
// For example:
//
//	MyMap.MySlices.MyField
func (pa Path) String() string { _ = "STUB: not implemented"; return "" }

// GoString returns the path to a specific node using Go syntax.
//
// For example:
//
//	(*root.MyMap["key"].(*mypkg.MyStruct).MySlices)[2][3].MyField
func (pa Path) GoString() string { _ = "STUB: not implemented"; return "" }

// Next step is indirection, so let them batch up

// Automatic indirection on struct fields

// Last step; no need for parenthesis

type pathStep struct {
	typ    reflect.Type
	vx, vy reflect.Value
}

func (ps pathStep) Type() reflect.Type { _ = "STUB: not implemented"; return *new(reflect.Type) }
func (ps pathStep) Values() (vx, vy reflect.Value) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), *new(reflect.Value)
}
func (ps pathStep) String() string { _ = "STUB: not implemented"; return "" }

// Type too simple or complex to print

// StructField is a [PathStep] that represents a struct field access
// on a field called [StructField.Name].
type StructField struct{ *structField }
type structField struct {
	pathStep
	name string
	idx  int

	// These fields are used for forcibly accessing an unexported field.
	// pvx, pvy, and field are only valid if unexported is true.
	unexported bool
	mayForce   bool                // Forcibly allow visibility
	paddr      bool                // Was parent addressable?
	pvx, pvy   reflect.Value       // Parent values (always addressable)
	field      reflect.StructField // Field information
}

func (sf StructField) Type() reflect.Type { _ = "STUB: not implemented"; return *new(reflect.Type) }
func (sf StructField) Values() (vx, vy reflect.Value) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), *new(reflect.Value)
}

// CanInterface reports true

// Forcibly obtain read-write access to an unexported struct field.

// CanInterface reports true

// CanInterface reports false

func (sf StructField) String() string { _ = "STUB: not implemented"; return "" }

// Name is the field name.
func (sf StructField) Name() string {
	_ = "STUB: not implemented"

	// Index is the index of the field in the parent struct type.
	// See [reflect.Type.Field].
	return ""
}

func (sf StructField) Index() int {
	_ = "STUB: not implemented"

	// SliceIndex is a [PathStep] that represents an index operation on
	// a slice or array at some index [SliceIndex.Key].
	return 0
}

type SliceIndex struct{ *sliceIndex }
type sliceIndex struct {
	pathStep
	xkey, ykey int
	isSlice    bool // False for reflect.Array
}

func (si SliceIndex) Type() reflect.Type { _ = "STUB: not implemented"; return *new(reflect.Type) }
func (si SliceIndex) Values() (vx, vy reflect.Value) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), *new(reflect.Value)
}
func (si SliceIndex) String() string { _ = "STUB: not implemented"; return "" }

// [5->?] means "I don't know where X[5] went"

// [?->3] means "I don't know where Y[3] came from"

// [5->3] means "X[5] moved to Y[3]"

// Key is the index key; it may return -1 if in a split state
func (si SliceIndex) Key() int { _ = "STUB: not implemented"; return 0 }

// SplitKeys are the indexes for indexing into slices in the
// x and y values, respectively. These indexes may differ due to the
// insertion or removal of an element in one of the slices, causing
// all of the indexes to be shifted. If an index is -1, then that
// indicates that the element does not exist in the associated slice.
//
// [SliceIndex.Key] is guaranteed to return -1 if and only if the indexes
// returned by SplitKeys are not the same. SplitKeys will never return -1 for
// both indexes.
func (si SliceIndex) SplitKeys() (ix, iy int) {
	_ = "STUB: not implemented"
	return 0,

		// MapIndex is a [PathStep] that represents an index operation on a map at some index Key.
		0
}

type MapIndex struct{ *mapIndex }
type mapIndex struct {
	pathStep
	key reflect.Value
}

func (mi MapIndex) Type() reflect.Type { _ = "STUB: not implemented"; return *new(reflect.Type) }
func (mi MapIndex) Values() (vx, vy reflect.Value) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), *new(reflect.Value)
}
func (mi MapIndex) String() string { _ = "STUB: not implemented"; return "" }

// Key is the value of the map key.
func (mi MapIndex) Key() reflect.Value {
	_ = "STUB: not implemented"

	// Indirect is a [PathStep] that represents pointer indirection on the parent type.
	return *new(reflect.Value)
}

type Indirect struct{ *indirect }
type indirect struct {
	pathStep
}

func (in Indirect) Type() reflect.Type { _ = "STUB: not implemented"; return *new(reflect.Type) }
func (in Indirect) Values() (vx, vy reflect.Value) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), *new(reflect.Value)
}
func (in Indirect) String() string { _ = "STUB: not implemented"; return "" }

// TypeAssertion is a [PathStep] that represents a type assertion on an interface.
type TypeAssertion struct{ *typeAssertion }
type typeAssertion struct {
	pathStep
}

func (ta TypeAssertion) Type() reflect.Type { _ = "STUB: not implemented"; return *new(reflect.Type) }
func (ta TypeAssertion) Values() (vx, vy reflect.Value) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), *new(reflect.Value)
}
func (ta TypeAssertion) String() string { _ = "STUB: not implemented"; return "" }

// Transform is a [PathStep] that represents a transformation
// from the parent type to the current type.
type Transform struct{ *transform }
type transform struct {
	pathStep
	trans *transformer
}

func (tf Transform) Type() reflect.Type { _ = "STUB: not implemented"; return *new(reflect.Type) }
func (tf Transform) Values() (vx, vy reflect.Value) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), *new(reflect.Value)
}
func (tf Transform) String() string { _ = "STUB: not implemented"; return "" }

// Name is the name of the [Transformer].
func (tf Transform) Name() string { _ = "STUB: not implemented"; return "" }

// Func is the function pointer to the transformer function.
func (tf Transform) Func() reflect.Value {
	_ = "STUB: not implemented"

	// Option returns the originally constructed [Transformer] option.
	// The == operator can be used to detect the exact option used.
	return *new(reflect.Value)
}

func (tf Transform) Option() Option {
	_ = "STUB: not implemented"

	// pointerPath represents a dual-stack of pointers encountered when
	// recursively traversing the x and y values. This data structure supports
	// detection of cycles and determining whether the cycles are equal.
	// In Go, cycles can occur via pointers, slices, and maps.
	//
	// The pointerPath uses a map to represent a stack; where descension into a
	// pointer pushes the address onto the stack, and ascension from a pointer
	// pops the address from the stack. Thus, when traversing into a pointer from
	// reflect.Ptr, reflect.Slice element, or reflect.Map, we can detect cycles
	// by checking whether the pointer has already been visited. The cycle detection
	// uses a separate stack for the x and y values.
	//
	// If a cycle is detected we need to determine whether the two pointers
	// should be considered equal. The definition of equality chosen by Equal
	// requires two graphs to have the same structure. To determine this, both the
	// x and y values must have a cycle where the previous pointers were also
	// encountered together as a pair.
	//
	// Semantically, this is equivalent to augmenting Indirect, SliceIndex, and
	// MapIndex with pointer information for the x and y values.
	// Suppose px and py are two pointers to compare, we then search the
	// Path for whether px was ever encountered in the Path history of x, and
	// similarly so with py. If either side has a cycle, the comparison is only
	// equal if both px and py have a cycle resulting from the same PathStep.
	//
	// Using a map as a stack is more performant as we can perform cycle detection
	// in O(1) instead of O(N) where N is len(Path).
	return *new(Option)
}

type pointerPath struct {
	// mx is keyed by x pointers, where the value is the associated y pointer.
	mx map[value.Pointer]value.Pointer
	// my is keyed by y pointers, where the value is the associated x pointer.
	my map[value.Pointer]value.Pointer
}

func (p *pointerPath) Init() { _ = "STUB: not implemented"; return }

// Push indicates intent to descend into pointers vx and vy where
// visited reports whether either has been seen before. If visited before,
// equal reports whether both pointers were encountered together.
// Pop must be called if and only if the pointers were never visited.
//
// The pointers vx and vy must be a reflect.Ptr, reflect.Slice, or reflect.Map
// and be non-nil.
func (p pointerPath) Push(vx, vy reflect.Value) (equal, visited bool) {
	_ = "STUB: not implemented"
	return false, false
}

// Pointers paired together

// Pop ascends from pointers vx and vy.
func (p pointerPath) Pop(vx, vy reflect.Value) { _ = "STUB: not implemented"; return }

// isExported reports whether the identifier is exported.
func isExported(id string) bool { _ = "STUB: not implemented"; return false }
