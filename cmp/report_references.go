// Copyright 2020, The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmp

import (
	"reflect"

	"github.com/google/go-cmp/cmp/internal/value"
)

const (
	pointerDelimPrefix = "⟪"
	pointerDelimSuffix = "⟫"
)

// formatPointer prints the address of the pointer.
func formatPointer(p value.Pointer, withDelims bool) string { _ = "STUB: not implemented"; return "" }

// Only used for stable testing purposes

// pointerReferences is a stack of pointers visited so far.
type pointerReferences [][2]value.Pointer

func (ps *pointerReferences) PushPair(vx, vy reflect.Value, d diffMode, deref bool) (pp [2]value.Pointer) {
	_ = "STUB: not implemented"
	return nil
}

func (ps *pointerReferences) Push(v reflect.Value) (p value.Pointer, seen bool) {
	_ = "STUB: not implemented"
	return *new(value.Pointer), false
}

func (ps *pointerReferences) Pop() { _ = "STUB: not implemented"; return }

// trunkReferences is metadata for a textNode indicating that the sub-tree
// represents the value for either pointer in a pair of references.
type trunkReferences struct{ pp [2]value.Pointer }

// trunkReference is metadata for a textNode indicating that the sub-tree
// represents the value for the given pointer reference.
type trunkReference struct{ p value.Pointer }

// leafReference is metadata for a textNode indicating that the value is
// truncated as it refers to another part of the tree (i.e., a trunk).
type leafReference struct{ p value.Pointer }

func wrapTrunkReferences(pp [2]value.Pointer, s textNode) textNode {
	_ = "STUB: not implemented"
	return *new(textNode)
}

func wrapTrunkReference(p value.Pointer, printAddress bool, s textNode) textNode {
	_ = "STUB: not implemented"
	return *new(textNode)
}

func makeLeafReference(p value.Pointer, printAddress bool) textNode {
	_ = "STUB: not implemented"
	return *new(textNode)
}

// resolveReferences walks the textNode tree searching for any leaf reference
// metadata and resolves each against the corresponding trunk references.
// Since pointer addresses in memory are not particularly readable to the user,
// it replaces each pointer value with an arbitrary and unique reference ID.
func resolveReferences(s textNode) { _ = "STUB: not implemented"; return }

// Collect all trunks and leaves with reference metadata.

// No leaf references to resolve.

// Collect the set of all leaf references to resolve.

// Collect the set of trunk pointers that are always paired together.
// This allows us to assign a single ID to both pointers for brevity.
// If a pointer in a pair ever occurs by itself or as a different pair,
// then the pair is broken.

// invalidate other half

// invalidate this half

// standalone pointer cannot be part of a pair

// Register the newly seen pair.

// Exact pair already seen; do nothing.

// Pair conflicts with some other pair; break all pairs.

// Correlate each pointer referenced by leaves to a unique identifier,
// and print the IDs for each trunk that matches those pointers.

// must be seen together or not at all

// must have the same ID

// Update all leaf references with the unique identifier.

func formatReference(id uint) string { _ = "STUB: not implemented"; return "" }

func updateReferencePrefix(prefix, ref string) string { _ = "STUB: not implemented"; return "" }
