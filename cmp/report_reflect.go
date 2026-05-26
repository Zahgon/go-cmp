// Copyright 2019, The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmp

import (
	"reflect"
)

var (
	anyType    = reflect.TypeOf((*interface{})(nil)).Elem()
	stringType = reflect.TypeOf((*string)(nil)).Elem()
	bytesType  = reflect.TypeOf((*[]byte)(nil)).Elem()
	byteType   = reflect.TypeOf((*byte)(nil)).Elem()
)

type formatValueOptions struct {
	// AvoidStringer controls whether to avoid calling custom stringer
	// methods like error.Error or fmt.Stringer.String.
	AvoidStringer bool

	// PrintAddresses controls whether to print the address of all pointers,
	// slice elements, and maps.
	PrintAddresses bool

	// QualifiedNames controls whether FormatType uses the fully qualified name
	// (including the full package path as opposed to just the package name).
	QualifiedNames bool

	// VerbosityLevel controls the amount of output to produce.
	// A higher value produces more output. A value of zero or lower produces
	// no output (represented using an ellipsis).
	// If LimitVerbosity is false, then the level is treated as infinite.
	VerbosityLevel int

	// LimitVerbosity specifies that formatting should respect VerbosityLevel.
	LimitVerbosity bool
}

// FormatType prints the type as if it were wrapping s.
// This may return s as-is depending on the current type and TypeMode mode.
func (opts formatOptions) FormatType(t reflect.Type, s textNode) textNode {
	_ = "STUB: not implemented"
	// Check whether to emit the type or not.
	return *new(textNode)
}

// elide type for identical nodes

// Determine the type label, applying special handling for unnamed types.

// According to Go grammar, certain type literals contain symbols that
// do not strongly bind to the next lexicographical token (e.g., *T).

// wrapParens wraps s with a set of parenthesis, but avoids it if the
// wrapped node itself is already surrounded by a pair of parenthesis or braces.
// It handles unwrapping one level of pointer-reference nodes.
func wrapParens(s textNode) textNode { _ = "STUB: not implemented"; return *new(textNode) }

// Unwrap a single pointer reference node.

// Already has delimiters that make parenthesis unnecessary.

// FormatValue prints the reflect.Value, taking extra care to avoid descending
// into pointers already in ptrs. As pointers are visited, ptrs is also updated.
func (opts formatOptions) FormatValue(v reflect.Value, parentKind reflect.Kind, ptrs *pointerReferences) (out textNode) {
	_ = "STUB: not implemented"
	return *new(textNode)
}

// Check slice element for cycles.

// Check whether there is an Error or String method to call.

// Avoid calling Error or String methods on nil receivers since many
// implementations crash when doing so.

// Swallow and ignore any panics from String or Error.

// Check whether to explicitly wrap the result with the type.

// needed for retrieveUnexportedField

// 0, 4, 8, 16, 32, etc...

// Elide fields with zero values

// Check whether this is a []byte of text data.

// 0, 4, 8, 16, 32, etc...

// Check pointer for cycles.

// 0, 4, 8, 16, 32, etc...

// Check pointer for cycles.

// Skip the name only if this is an unnamed pointer type.
// Otherwise taking the address of a value does not reproduce
// the named pointer type.

// Let the underlying value print the type instead

// Interfaces accept different concrete types,
// so configure the underlying value to explicitly print the type.

func (opts formatOptions) formatString(prefix, s string) textNode {
	_ = "STUB: not implemented"
	return *new(textNode)
}

// 32, 64, 128, 256, etc...
//  4, 8, 16, 32, 64, etc...

// For multiline strings, use the triple-quote syntax,
// but only use it when printing removed or inserted nodes since
// we only want the extra verbosity for those cases.

// trim leading/trailing carriage returns for legacy Windows endline support

// specially treat tab as printable

// Format the string as a single-line quoted string.

// formatMapKey formats v as if it were a map key.
// The result is guaranteed to be a single line.
func formatMapKey(v reflect.Value, disambiguate bool, ptrs *pointerReferences) string {
	_ = "STUB: not implemented"
	return ""
}

// formatString prints s as a double-quoted or backtick-quoted string.
func formatString(s string) string {
	_ = "STUB: not implemented"
	// Use quoted string if it the same length as a raw string literal.
	// Otherwise, attempt to use the raw string form.
	return ""
}

// Disallow newlines to ensure output is a single line.
// Only allow printable runes for readability purposes.

// formatHex prints u as a hexadecimal integer in Go notation.
func formatHex(u uint64) string { _ = "STUB: not implemented"; return "" }
