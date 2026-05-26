// Copyright 2017, The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmp

import (
	"reflect"
	"regexp"
)

// Option configures for specific behavior of [Equal] and [Diff]. In particular,
// the fundamental Option functions ([Ignore], [Transformer], and [Comparer]),
// configure how equality is determined.
//
// The fundamental options may be composed with filters ([FilterPath] and
// [FilterValues]) to control the scope over which they are applied.
//
// The [github.com/google/go-cmp/cmp/cmpopts] package provides helper functions
// for creating options that may be used with [Equal] and [Diff].
type Option interface {
	// filter applies all filters and returns the option that remains.
	// Each option may only read s.curPath and call s.callTTBFunc.
	//
	// An Options is returned only if multiple comparers or transformers
	// can apply simultaneously and will only contain values of those types
	// or sub-Options containing values of those types.
	filter(s *state, t reflect.Type, vx, vy reflect.Value) applicableOption
}

// applicableOption represents the following types:
//
//	Fundamental: ignore | validator | *comparer | *transformer
//	Grouping:    Options
type applicableOption interface {
	Option

	// apply executes the option, which may mutate s or panic.
	apply(s *state, vx, vy reflect.Value)
}

// coreOption represents the following types:
//
//	Fundamental: ignore | validator | *comparer | *transformer
//	Filters:     *pathFilter | *valuesFilter
type coreOption interface {
	Option
	isCore()
}

type core struct{}

func (core) isCore() {
	_ = "STUB: not implemented"

	// Options is a list of [Option] values that also satisfies the [Option] interface.
	// Helper comparison packages may return an Options value when packing multiple
	// [Option] values into a single [Option]. When this package processes an Options,
	// it will be implicitly expanded into a flat list.
	//
	// Applying a filter on an Options is equivalent to applying that same filter
	// on all individual options held within.
	return
}

type Options []Option

func (opts Options) filter(s *state, t reflect.Type, vx, vy reflect.Value) (out applicableOption) {
	_ = "STUB: not implemented"
	return *new(applicableOption)
}

// Only ignore can short-circuit evaluation

// Takes precedence over comparer or transformer

// Keep validator

// Conflicting comparers or transformers

func (opts Options) apply(s *state, _, _ reflect.Value) { _ = "STUB: not implemented"; return }

func (opts Options) String() string { _ = "STUB: not implemented"; return "" }

// FilterPath returns a new [Option] where opt is only evaluated if filter f
// returns true for the current [Path] in the value tree.
//
// This filter is called even if a slice element or map entry is missing and
// provides an opportunity to ignore such cases. The filter function must be
// symmetric such that the filter result is identical regardless of whether the
// missing value is from x or y.
//
// The option passed in may be an [Ignore], [Transformer], [Comparer], [Options], or
// a previously filtered [Option].
func FilterPath(f func(Path) bool, opt Option) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

type pathFilter struct {
	core
	fnc func(Path) bool
	opt Option
}

func (f pathFilter) filter(s *state, t reflect.Type, vx, vy reflect.Value) applicableOption {
	_ = "STUB: not implemented"
	return *new(applicableOption)
}

func (f pathFilter) String() string { _ = "STUB: not implemented"; return "" }

// FilterValues returns a new [Option] where opt is only evaluated if filter f,
// which is a function of the form "func(T, T) bool", returns true for the
// current pair of values being compared. If either value is invalid or
// the type of the values is not assignable to T, then this filter implicitly
// returns false.
//
// The filter function must be
// symmetric (i.e., agnostic to the order of the inputs) and
// deterministic (i.e., produces the same result when given the same inputs).
// If T is an interface, it is possible that f is called with two values with
// different concrete types that both implement T.
//
// The option passed in may be an [Ignore], [Transformer], [Comparer], [Options], or
// a previously filtered [Option].
func FilterValues(f interface{}, opt Option) Option { _ = "STUB: not implemented"; return *new(Option) }

type valuesFilter struct {
	core
	typ reflect.Type  // T
	fnc reflect.Value // func(T, T) bool
	opt Option
}

func (f valuesFilter) filter(s *state, t reflect.Type, vx, vy reflect.Value) applicableOption {
	_ = "STUB: not implemented"
	return *new(applicableOption)
}

func (f valuesFilter) String() string { _ = "STUB: not implemented"; return "" }

// Ignore is an [Option] that causes all comparisons to be ignored.
// This value is intended to be combined with [FilterPath] or [FilterValues].
// It is an error to pass an unfiltered Ignore option to [Equal].
func Ignore() Option { _ = "STUB: not implemented"; return *new(Option) }

type ignore struct{ core }

func (ignore) isFiltered() bool { _ = "STUB: not implemented"; return false }
func (ignore) filter(_ *state, _ reflect.Type, _, _ reflect.Value) applicableOption {
	_ = "STUB: not implemented"
	return *new(applicableOption)
}
func (ignore) apply(s *state, _, _ reflect.Value) { _ = "STUB: not implemented"; return }
func (ignore) String() string                     { _ = "STUB: not implemented"; return "" }

// validator is a sentinel Option type to indicate that some options could not
// be evaluated due to unexported fields, missing slice elements, or
// missing map entries. Both values are validator only for unexported fields.
type validator struct{ core }

func (validator) filter(_ *state, _ reflect.Type, vx, vy reflect.Value) applicableOption {
	_ = "STUB: not implemented"
	return *new(applicableOption)
}

func (validator) apply(s *state, vx, vy reflect.Value) {
	_ = "STUB: not implemented"
	// Implies missing slice element or map entry.
	return
}

// Unable to Interface implies unexported field without visibility access.

// Named type with unexported fields.
// e.g., "path/to/package".MyType

// Unnamed type with unexported fields. Derive PkgPath from field.

// e.g., "path/to/package".(struct { a int })

// identRx represents a valid identifier according to the Go specification.
const identRx = `[_\p{L}][_\p{L}\p{N}]*`

var identsRx = regexp.MustCompile(`^` + identRx + `(\.` + identRx + `)*$`)

// Transformer returns an [Option] that applies a transformation function that
// converts values of a certain type into that of another.
//
// The transformer f must be a function "func(T) R" that converts values of
// type T to those of type R and is implicitly filtered to input values
// assignable to T. The transformer must not mutate T in any way.
//
// To help prevent some cases of infinite recursive cycles applying the
// same transform to the output of itself (e.g., in the case where the
// input and output types are the same), an implicit filter is added such that
// a transformer is applicable only if that exact transformer is not already
// in the tail of the [Path] since the last non-[Transform] step.
// For situations where the implicit filter is still insufficient,
// consider using [github.com/google/go-cmp/cmp/cmpopts.AcyclicTransformer],
// which adds a filter to prevent the transformer from
// being recursively applied upon itself.
//
// The name is a user provided label that is used as the [Transform.Name] in the
// transformation [PathStep] (and eventually shown in the [Diff] output).
// The name must be a valid identifier or qualified identifier in Go syntax.
// If empty, an arbitrary name is used.
func Transformer(name string, f interface{}) Option { _ = "STUB: not implemented"; return *new(Option) }

// Lambda-symbol as placeholder name

type transformer struct {
	core
	name string
	typ  reflect.Type  // T
	fnc  reflect.Value // func(T) R
}

func (tr *transformer) isFiltered() bool { _ = "STUB: not implemented"; return false }

func (tr *transformer) filter(s *state, t reflect.Type, _, _ reflect.Value) applicableOption {
	_ = "STUB: not implemented"
	return *new(applicableOption)
}

// Hit most recent non-Transform step

// Cannot directly use same Transform

func (tr *transformer) apply(s *state, vx, vy reflect.Value) { _ = "STUB: not implemented"; return }

func (tr transformer) String() string { _ = "STUB: not implemented"; return "" }

// Comparer returns an [Option] that determines whether two values are equal
// to each other.
//
// The comparer f must be a function "func(T, T) bool" and is implicitly
// filtered to input values assignable to T. If T is an interface, it is
// possible that f is called with two values of different concrete types that
// both implement T.
//
// The equality function must be:
//   - Symmetric: equal(x, y) == equal(y, x)
//   - Deterministic: equal(x, y) == equal(x, y)
//   - Pure: equal(x, y) does not modify x or y
func Comparer(f interface{}) Option { _ = "STUB: not implemented"; return *new(Option) }

type comparer struct {
	core
	typ reflect.Type  // T
	fnc reflect.Value // func(T, T) bool
}

func (cm *comparer) isFiltered() bool { _ = "STUB: not implemented"; return false }

func (cm *comparer) filter(_ *state, t reflect.Type, _, _ reflect.Value) applicableOption {
	_ = "STUB: not implemented"
	return *new(applicableOption)
}

func (cm *comparer) apply(s *state, vx, vy reflect.Value) { _ = "STUB: not implemented"; return }

func (cm comparer) String() string { _ = "STUB: not implemented"; return "" }

// Exporter returns an [Option] that specifies whether [Equal] is allowed to
// introspect into the unexported fields of certain struct types.
//
// Users of this option must understand that comparing on unexported fields
// from external packages is not safe since changes in the internal
// implementation of some external package may cause the result of [Equal]
// to unexpectedly change. However, it may be valid to use this option on types
// defined in an internal package where the semantic meaning of an unexported
// field is in the control of the user.
//
// In many cases, a custom [Comparer] should be used instead that defines
// equality as a function of the public API of a type rather than the underlying
// unexported implementation.
//
// For example, the [reflect.Type] documentation defines equality to be determined
// by the == operator on the interface (essentially performing a shallow pointer
// comparison) and most attempts to compare *[regexp.Regexp] types are interested
// in only checking that the regular expression strings are equal.
// Both of these are accomplished using [Comparer] options:
//
//	Comparer(func(x, y reflect.Type) bool { return x == y })
//	Comparer(func(x, y *regexp.Regexp) bool { return x.String() == y.String() })
//
// In other cases, the [github.com/google/go-cmp/cmp/cmpopts.IgnoreUnexported]
// option can be used to ignore all unexported fields on specified struct types.
func Exporter(f func(reflect.Type) bool) Option { _ = "STUB: not implemented"; return *new(Option) }

type exporter func(reflect.Type) bool

func (exporter) filter(_ *state, _ reflect.Type, _, _ reflect.Value) applicableOption {
	_ = "STUB: not implemented"
	return *

	// AllowUnexported returns an [Option] that allows [Equal] to forcibly introspect
	// unexported fields of the specified struct types.
	//
	// See [Exporter] for the proper use of this option.
	new(applicableOption)
}

func AllowUnexported(types ...interface{}) Option { _ = "STUB: not implemented"; return *new(Option) }

// Result represents the comparison result for a single node and
// is provided by cmp when calling Report (see [Reporter]).
type Result struct {
	_     [0]func() // Make Result incomparable
	flags resultFlags
}

// Equal reports whether the node was determined to be equal or not.
// As a special case, ignored nodes are considered equal.
func (r Result) Equal() bool { _ = "STUB: not implemented"; return false }

// ByIgnore reports whether the node is equal because it was ignored.
// This never reports true if [Result.Equal] reports false.
func (r Result) ByIgnore() bool { _ = "STUB: not implemented"; return false }

// ByMethod reports whether the Equal method determined equality.
func (r Result) ByMethod() bool { _ = "STUB: not implemented"; return false }

// ByFunc reports whether a [Comparer] function determined equality.
func (r Result) ByFunc() bool { _ = "STUB: not implemented"; return false }

// ByCycle reports whether a reference cycle was detected.
func (r Result) ByCycle() bool { _ = "STUB: not implemented"; return false }

type resultFlags uint

const (
	_ resultFlags = (1 << iota) / 2

	reportEqual
	reportUnequal
	reportByIgnore
	reportByMethod
	reportByFunc
	reportByCycle
)

// Reporter is an [Option] that can be passed to [Equal]. When [Equal] traverses
// the value trees, it calls PushStep as it descends into each node in the
// tree and PopStep as it ascend out of the node. The leaves of the tree are
// either compared (determined to be equal or not equal) or ignored and reported
// as such by calling the Report method.
func Reporter(r interface {
	// PushStep is called when a tree-traversal operation is performed.
	// The PathStep itself is only valid until the step is popped.
	// The PathStep.Values are valid for the duration of the entire traversal
	// and must not be mutated.
	//
	// Equal always calls PushStep at the start to provide an operation-less
	// PathStep used to report the root values.
	//
	// Within a slice, the exact set of inserted, removed, or modified elements
	// is unspecified and may change in future implementations.
	// The entries of a map are iterated through in an unspecified order.
	PushStep(PathStep)

	// Report is called exactly once on leaf nodes to report whether the
	// comparison identified the node as equal, unequal, or ignored.
	// A leaf node is one that is immediately preceded by and followed by
	// a pair of PushStep and PopStep calls.
	Report(Result)

	// PopStep ascends back up the value tree.
	// There is always a matching pop call for every push call.
	PopStep()
}) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

type reporter struct{ reporterIface }
type reporterIface interface {
	PushStep(PathStep)
	Report(Result)
	PopStep()
}

func (reporter) filter(_ *state, _ reflect.Type, _, _ reflect.Value) applicableOption {
	_ = "STUB: not implemented"
	return *

	// normalizeOption normalizes the input options such that all Options groups
	// are flattened and groups with a single element are reduced to that element.
	// Only coreOptions and Options containing coreOptions are allowed.
	new(applicableOption)
}

func normalizeOption(src Option) Option { _ = "STUB: not implemented"; return *new(Option) }

// flattenOptions copies all options in src to dst as a flat list.
// Only coreOptions and Options containing coreOptions are allowed.
func flattenOptions(dst, src Options) Options { _ = "STUB: not implemented"; return *new(Options) }
