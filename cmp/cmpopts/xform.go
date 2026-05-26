// Copyright 2018, The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmpopts

import (
	"github.com/google/go-cmp/cmp"
)

type xformFilter struct{ xform cmp.Option }

func (xf xformFilter) filter(p cmp.Path) bool { _ = "STUB: not implemented"; return false }

// AcyclicTransformer returns a [cmp.Transformer] with a filter applied that ensures
// that the transformer cannot be recursively applied upon its own output.
//
// An example use case is a transformer that splits a string by lines:
//
//	AcyclicTransformer("SplitLines", func(s string) []string{
//		return strings.Split(s, "\n")
//	})
//
// Had this been an unfiltered [cmp.Transformer] instead, this would result in an
// infinite cycle converting a string to []string to [][]string and so on.
func AcyclicTransformer(name string, xformFunc interface{}) cmp.Option {
	_ = "STUB: not implemented"
	return *new(cmp.Option)
}
