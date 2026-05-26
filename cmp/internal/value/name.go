// Copyright 2020, The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package value

import (
	"reflect"
)

var anyType = reflect.TypeOf((*interface{})(nil)).Elem()

// TypeString is nearly identical to reflect.Type.String,
// but has an additional option to specify that full type names be used.
func TypeString(t reflect.Type, qualified bool) string { _ = "STUB: not implemented"; return "" }

func appendTypeName(b []byte, t reflect.Type, qualified, elideFunc bool) []byte {
	_ = "STUB: not implemented"
	// BUG: Go reflection provides no way to disambiguate two named types
	// of the same name and within the same package,
	// but declared within the namespace of different functions.
	return nil
}

// Use the "any" alias instead of "interface{}" for better readability.

// Named type.

// Unnamed type.

// Do nothing
