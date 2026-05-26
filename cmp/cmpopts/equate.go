// Copyright 2017, The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package cmpopts provides common options for the cmp package.
package cmpopts

import (
	"reflect"
	"time"

	"github.com/google/go-cmp/cmp"
)

func equateAlways(_, _ interface{}) bool {
	_ = "STUB: not implemented"

	// EquateEmpty returns a [cmp.Comparer] option that determines all maps and slices
	// with a length of zero to be equal, regardless of whether they are nil.
	//
	// EquateEmpty can be used in conjunction with [SortSlices] and [SortMaps].
	return false
}

func EquateEmpty() cmp.Option { _ = "STUB: not implemented"; return *new(cmp.Option) }

func isEmpty(x, y interface{}) bool { _ = "STUB: not implemented"; return false }

// EquateApprox returns a [cmp.Comparer] option that determines float32 or float64
// values to be equal if they are within a relative fraction or absolute margin.
// This option is not used when either x or y is NaN or infinite.
//
// The fraction determines that the difference of two values must be within the
// smaller fraction of the two values, while the margin determines that the two
// values must be within some absolute margin.
// To express only a fraction or only a margin, use 0 for the other parameter.
// The fraction and margin must be non-negative.
//
// The mathematical expression used is equivalent to:
//
//	|x-y| ≤ max(fraction*min(|x|, |y|), margin)
//
// EquateApprox can be used in conjunction with [EquateNaNs].
func EquateApprox(fraction, margin float64) cmp.Option {
	_ = "STUB: not implemented"
	return *new(cmp.Option)
}

type approximator struct{ frac, marg float64 }

func areRealF64s(x, y float64) bool { _ = "STUB: not implemented"; return false }

func areRealF32s(x, y float32) bool { _ = "STUB: not implemented"; return false }

func (a approximator) compareF64(x, y float64) bool { _ = "STUB: not implemented"; return false }

func (a approximator) compareF32(x, y float32) bool { _ = "STUB: not implemented"; return false }

// EquateNaNs returns a [cmp.Comparer] option that determines float32 and float64
// NaN values to be equal.
//
// EquateNaNs can be used in conjunction with [EquateApprox].
func EquateNaNs() cmp.Option { _ = "STUB: not implemented"; return *new(cmp.Option) }

func areNaNsF64s(x, y float64) bool { _ = "STUB: not implemented"; return false }

func areNaNsF32s(x, y float32) bool { _ = "STUB: not implemented"; return false }

// EquateApproxTime returns a [cmp.Comparer] option that determines two non-zero
// [time.Time] values to be equal if they are within some margin of one another.
// If both times have a monotonic clock reading, then the monotonic time
// difference will be used. The margin must be non-negative.
func EquateApproxTime(margin time.Duration) cmp.Option {
	_ = "STUB: not implemented"
	return *new(cmp.Option)
}

func areNonZeroTimes(x, y time.Time) bool { _ = "STUB: not implemented"; return false }

type timeApproximator struct {
	margin time.Duration
}

func (a timeApproximator) compare(x, y time.Time) bool {
	_ = "STUB: not implemented"
	// Avoid subtracting times to avoid overflow when the
	// difference is larger than the largest representable duration.
	return false
}

// Ensure x is always before y

// We're within the margin if x+margin >= y.
// Note: time.Time doesn't have AfterOrEqual method hence the negation.

// AnyError is an error that matches any non-nil error.
var AnyError anyError

type anyError struct{}

func (anyError) Error() string { _ = "STUB: not implemented"; return "" }
func (anyError) Is(err error) bool {
	_ = "STUB: not implemented"

	// EquateErrors returns a [cmp.Comparer] option that determines errors to be equal
	// if [errors.Is] reports them to match. The [AnyError] error can be used to
	// match any non-nil error.
	return false
}

func EquateErrors() cmp.Option { _ = "STUB: not implemented"; return *new(cmp.Option) }

// areConcreteErrors reports whether x and y are types that implement error.
// The input types are deliberately of the interface{} type rather than the
// error type so that we can handle situations where the current type is an
// interface{}, but the underlying concrete types both happen to implement
// the error interface.
func areConcreteErrors(x, y interface{}) bool { _ = "STUB: not implemented"; return false }

func compareErrors(x, y interface{}) bool { _ = "STUB: not implemented"; return false }

// EquateComparable returns a [cmp.Option] that determines equality
// of comparable types by directly comparing them using the == operator in Go.
// The types to compare are specified by passing a value of that type.
// This option should only be used on types that are documented as being
// safe for direct == comparison. For example, [net/netip.Addr] is documented
// as being semantically safe to use with ==, while [time.Time] is documented
// to discourage the use of == on time values.
func EquateComparable(typs ...interface{}) cmp.Option {
	_ = "STUB: not implemented"
	return *new(cmp.Option)
}

type typesFilter map[reflect.Type]bool

func (tf typesFilter) filter(p cmp.Path) bool { _ = "STUB: not implemented"; return false }

func equateAny(x, y interface{}) bool { _ = "STUB: not implemented"; return false }
