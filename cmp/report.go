// Copyright 2017, The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmp

// defaultReporter implements the reporter interface.
//
// As Equal serially calls the PushStep, Report, and PopStep methods, the
// defaultReporter constructs a tree-based representation of the compared value
// and the result of each comparison (see valueNode).
//
// When the String method is called, the FormatDiff method transforms the
// valueNode tree into a textNode tree, which is a tree-based representation
// of the textual output (see textNode).
//
// Lastly, the textNode.String method produces the final report as a string.
type defaultReporter struct {
	root *valueNode
	curr *valueNode
}

func (r *defaultReporter) PushStep(ps PathStep) { _ = "STUB: not implemented"; return }

func (r *defaultReporter) Report(rs Result) { _ = "STUB: not implemented"; return }

func (r *defaultReporter) PopStep() { _ = "STUB: not implemented"; return }

// String provides a full report of the differences detected as a structured
// literal in pseudo-Go syntax. String may only be called after the entire tree
// has been traversed.
func (r *defaultReporter) String() string { _ = "STUB: not implemented"; return "" }

func assert(ok bool) { _ = "STUB: not implemented"; return }
