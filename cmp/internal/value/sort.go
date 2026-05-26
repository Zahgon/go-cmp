// Copyright 2017, The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package value

import (
	"reflect"
)

// SortKeys sorts a list of map keys, deduplicating keys if necessary.
// The type of each value must be comparable.
func SortKeys(vs []reflect.Value) []reflect.Value { _ = "STUB: not implemented"; return nil }

// Sort the map keys.

// Deduplicate keys (fails for NaNs).

// isLess is a generic function for sorting arbitrary map keys.
// The inputs must be of the same type and must be comparable.
func isLess(x, y reflect.Value) bool { _ = "STUB: not implemented"; return false }

// NOTE: This does not sort -0 as less than +0
// since Go maps treat -0 and +0 as equal keys.

// This can happen in rare situations, so we fallback to just comparing
// the unique pointer for a reflect.Type. This guarantees deterministic
// ordering within a program, but it is obviously not stable.

// Must be Func, Map, or Slice; which are not comparable.
