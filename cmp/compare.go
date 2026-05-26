// Copyright 2017, The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package cmp determines equality of values.
//
// This package is intended to be a more powerful and safer alternative to
// [reflect.DeepEqual] for comparing whether two values are semantically equal.
// It is intended to only be used in tests, as performance is not a goal and
// it may panic if it cannot compare the values. Its propensity towards
// panicking means that its unsuitable for production environments where a
// spurious panic may be fatal.
//
// The primary features of cmp are:
//
//   - When the default behavior of equality does not suit the test's needs,
//     custom equality functions can override the equality operation.
//     For example, an equality function may report floats as equal so long as
//     they are within some tolerance of each other.
//
//   - Types with an Equal method (e.g., [time.Time.Equal]) may use that method
//     to determine equality. This allows package authors to determine
//     the equality operation for the types that they define.
//
//   - If no custom equality functions are used and no Equal method is defined,
//     equality is determined by recursively comparing the primitive kinds on
//     both values, much like [reflect.DeepEqual]. Unlike [reflect.DeepEqual],
//     unexported fields are not compared by default; they result in panics
//     unless suppressed by using an [Ignore] option
//     (see [github.com/google/go-cmp/cmp/cmpopts.IgnoreUnexported])
//     or explicitly compared using the [Exporter] option.
package cmp

import (
	"reflect"

	"github.com/google/go-cmp/cmp/internal/diff"
)

// TODO(≥go1.18): Use any instead of interface{}.

// Equal reports whether x and y are equal by recursively applying the
// following rules in the given order to x and y and all of their sub-values:
//
//   - Let S be the set of all [Ignore], [Transformer], and [Comparer] options that
//     remain after applying all path filters, value filters, and type filters.
//     If at least one [Ignore] exists in S, then the comparison is ignored.
//     If the number of [Transformer] and [Comparer] options in S is non-zero,
//     then Equal panics because it is ambiguous which option to use.
//     If S contains a single [Transformer], then use that to transform
//     the current values and recursively call Equal on the output values.
//     If S contains a single [Comparer], then use that to compare the current values.
//     Otherwise, evaluation proceeds to the next rule.
//
//   - If the values have an Equal method of the form "(T) Equal(T) bool" or
//     "(T) Equal(I) bool" where T is assignable to I, then use the result of
//     x.Equal(y) even if x or y is nil. Otherwise, no such method exists and
//     evaluation proceeds to the next rule.
//
//   - Lastly, try to compare x and y based on their basic kinds.
//     Simple kinds like booleans, integers, floats, complex numbers, strings,
//     and channels are compared using the equivalent of the == operator in Go.
//     Functions are only equal if they are both nil, otherwise they are unequal.
//
// Structs are equal if recursively calling Equal on all fields report equal.
// If a struct contains unexported fields, Equal panics unless an [Ignore] option
// (e.g., [github.com/google/go-cmp/cmp/cmpopts.IgnoreUnexported]) ignores that field
// or the [Exporter] option explicitly permits comparing the unexported field.
//
// Slices are equal if they are both nil or both non-nil, where recursively
// calling Equal on all non-ignored slice or array elements report equal.
// Empty non-nil slices and nil slices are not equal; to equate empty slices,
// consider using [github.com/google/go-cmp/cmp/cmpopts.EquateEmpty].
//
// Maps are equal if they are both nil or both non-nil, where recursively
// calling Equal on all non-ignored map entries report equal.
// Map keys are equal according to the == operator.
// To use custom comparisons for map keys, consider using
// [github.com/google/go-cmp/cmp/cmpopts.SortMaps].
// Empty non-nil maps and nil maps are not equal; to equate empty maps,
// consider using [github.com/google/go-cmp/cmp/cmpopts.EquateEmpty].
//
// Pointers and interfaces are equal if they are both nil or both non-nil,
// where they have the same underlying concrete type and recursively
// calling Equal on the underlying values reports equal.
//
// Before recursing into a pointer, slice element, or map, the current path
// is checked to detect whether the address has already been visited.
// If there is a cycle, then the pointed at values are considered equal
// only if both addresses were previously visited in the same path step.
func Equal(x, y interface{}, opts ...Option) bool { _ = "STUB: not implemented"; return false }

// Diff returns a human-readable report of the differences between two values:
// y - x. It returns an empty string if and only if Equal returns true for the
// same input values and options.
//
// The output is displayed as a literal in pseudo-Go syntax.
// At the start of each line, a "-" prefix indicates an element removed from x,
// a "+" prefix to indicates an element added from y, and the lack of a prefix
// indicates an element common to both x and y. If possible, the output
// uses fmt.Stringer.String or error.Error methods to produce more humanly
// readable outputs. In such cases, the string is prefixed with either an
// 's' or 'e' character, respectively, to indicate that the method was called.
//
// Do not depend on this output being stable. If you need the ability to
// programmatically interpret the difference, consider using a custom Reporter.
func Diff(x, y interface{}, opts ...Option) string {
	_ = "STUB: not implemented"

	// Optimization: If there are no other reporters, we can optimize for the
	// common case where the result is equal (and thus no reported difference).
	// This avoids the expensive construction of a difference tree.
	return ""
}

// Reset results

// rootStep constructs the first path step. If x and y have differing types,
// then they are stored within an empty interface type.
func rootStep(x, y interface{}) PathStep { _ = "STUB: not implemented"; return *new(PathStep) }

// If the inputs are different types, auto-wrap them in an empty interface
// so that they have the same parent type.

type state struct {
	// These fields represent the "comparison state".
	// Calling statelessCompare must not result in observable changes to these.
	result    diff.Result // The current result of comparison
	curPath   Path        // The current path in the value tree
	curPtrs   pointerPath // The current set of visited pointers
	reporters []reporter  // Optional reporters

	// recChecker checks for infinite cycles applying the same set of
	// transformers upon the output of itself.
	recChecker recChecker

	// dynChecker triggers pseudo-random checks for option correctness.
	// It is safe for statelessCompare to mutate this value.
	dynChecker dynChecker

	// These fields, once set by processOption, will not change.
	exporters []exporter // List of exporters for structs with unexported fields
	opts      Options    // List of all fundamental and filter options
}

func newState(opts []Option) *state {
	_ = "STUB: not implemented"
	// Always ensure a validator option exists to validate the inputs.
	return nil
}

func (s *state) processOption(opt Option) { _ = "STUB: not implemented"; return }

// statelessCompare compares two values and returns the result.
// This function is stateless in that it does not alter the current result,
// or output to any registered reporters.
func (s *state) statelessCompare(step PathStep) diff.Result {
	_ = "STUB: not implemented"
	// We do not save and restore curPath and curPtrs because all of the
	// compareX methods should properly push and pop from them.
	// It is an implementation bug if the contents of the paths differ from
	// when calling this function to when returning from it.
	return *new(diff.Result)
}

// Reset result
// Remove reporters to avoid spurious printouts

func (s *state) compareAny(step PathStep) {
	_ = "STUB: not implemented"
	// Update the path stack.
	return
}

// Cycle-detection for slice elements (see NOTE in compareSlice).

// Rule 1: Check whether an option applies on this node in the value tree.

// Rule 2: Check whether the type has a valid Equal method.

// Rule 3: Compare based on the underlying kind.

func (s *state) tryOptions(t reflect.Type, vx, vy reflect.Value) bool {
	_ = "STUB: not implemented"
	// Evaluate all filters and apply the remaining options.
	return false
}

func (s *state) tryMethod(t reflect.Type, vx, vy reflect.Value) bool {
	_ = "STUB: not implemented"
	// Check if this type even has an Equal method.
	return false
}

func (s *state) callTRFunc(f, v reflect.Value, step Transform) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}

// Run the function twice and ensure that we get the same results back.
// We run in goroutines so that the race detector (if enabled) can detect
// unsafe mutations to the input.

// To avoid false-positives with non-reflexive equality operations,
// we sanity check whether a value is equal to itself.

func (s *state) callTTBFunc(f, x, y reflect.Value) bool { _ = "STUB: not implemented"; return false }

// Swapping the input arguments is sufficient to check that
// f is symmetric and deterministic.
// We run in goroutines so that the race detector (if enabled) can detect
// unsafe mutations to the input.

func detectRaces(c chan<- reflect.Value, f reflect.Value, vs ...reflect.Value) {
	_ = "STUB: not implemented"
	return
}

// Ignore panics, let the other call to f panic instead

func (s *state) compareStruct(t reflect.Type, vx, vy reflect.Value) {
	_ = "STUB: not implemented"
	return
}

// Addressable versions of vx and vy

// Defer checking of unexported fields until later to give an
// Ignore a chance to ignore the field.

// For retrieveUnexportedField to work, the parent struct must
// be addressable. Create a new copy of the values if
// necessary to make them addressable.

func (s *state) compareSlice(t reflect.Type, vx, vy reflect.Value) {
	_ = "STUB: not implemented"
	return
}

// NOTE: It is incorrect to call curPtrs.Push on the slice header pointer
// since slices represents a list of pointers, rather than a single pointer.
// The pointer checking logic must be handled on a per-element basis
// in compareAny.
//
// A slice header (see reflect.SliceHeader) in Go is a tuple of a starting
// pointer P, a length N, and a capacity C. Supposing each slice element has
// a memory size of M, then the slice is equivalent to the list of pointers:
//	[P+i*M for i in range(N)]
//
// For example, v[:0] and v[:1] are slices with the same starting pointer,
// but they are clearly different values. Using the slice pointer alone
// violates the assumption that equal pointers implies equal values.

// Ignore options are able to ignore missing elements in a slice.
// However, detecting these reliably requires an optimal differencing
// algorithm, for which diff.Difference is not.
//
// Instead, we first iterate through both slices to detect which elements
// would be ignored if standing alone. The index of non-discarded elements
// are stored in a separate slice, which diffing is then performed on.

// Compute an edit-script for slices vx and vy (excluding ignored elements).

// Replay the ignore-scripts and the edit-script.

func (s *state) compareMap(t reflect.Type, vx, vy reflect.Value) { _ = "STUB: not implemented"; return }

// Cycle-detection for maps.

// We combine and sort the two map keys so that we can perform the
// comparisons in a deterministic order.

// It is possible for both vx and vy to be invalid if the
// key contained a NaN value in it.
//
// Even with the ability to retrieve NaN keys in Go 1.12,
// there still isn't a sensible way to compare the values since
// a NaN key may map to multiple unordered values.
// The most reasonable way to compare NaNs would be to compare the
// set of values. However, this is impossible to do efficiently
// since set equality is provably an O(n^2) operation given only
// an Equal function. If we had a Less function or Hash function,
// this could be done in O(n*log(n)) or O(n), respectively.
//
// Rather than adding complex logic to deal with NaNs, make it
// the user's responsibility to compare such obscure maps.

func (s *state) comparePtr(t reflect.Type, vx, vy reflect.Value) { _ = "STUB: not implemented"; return }

// Cycle-detection for pointers.

func (s *state) compareInterface(t reflect.Type, vx, vy reflect.Value) {
	_ = "STUB: not implemented"
	return
}

func (s *state) report(eq bool, rf resultFlags) { _ = "STUB: not implemented"; return }

// recChecker tracks the state needed to periodically perform checks that
// user provided transformers are not stuck in an infinitely recursive cycle.
type recChecker struct{ next int }

// Check scans the Path for any recursive transformers and panics when any
// recursive transformers are detected. Note that the presence of a
// recursive Transformer does not necessarily imply an infinite cycle.
// As such, this check only activates after some minimal number of path steps.
func (rc *recChecker) Check(p Path) { _ = "STUB: not implemented"; return }

// Check whether the same transformer has appeared at least twice.

// Transformer was used exactly once before

// dynChecker tracks the state needed to periodically perform checks that
// user provided functions are symmetric and deterministic.
// The zero value is safe for immediate use.
type dynChecker struct{ curr, next int }

// Next increments the state and reports whether a check should be performed.
//
// Checks occur every Nth function call, where N is a triangular number:
//
//	0 1 3 6 10 15 21 28 36 45 55 66 78 91 105 120 136 153 171 190 ...
//
// See https://en.wikipedia.org/wiki/Triangular_number
//
// This sequence ensures that the cost of checks drops significantly as
// the number of functions calls grows larger.
func (dc *dynChecker) Next() bool { _ = "STUB: not implemented"; return false }

// makeAddressable returns a value that is always addressable.
// It returns the input verbatim if it is already addressable,
// otherwise it creates a new value and returns an addressable copy.
func makeAddressable(v reflect.Value) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}
