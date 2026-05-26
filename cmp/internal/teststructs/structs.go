// Copyright 2017, The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package teststructs

type InterfaceA interface {
	InterfaceA()
}

type (
	StructA struct{ X string } // Equal method on value receiver
	StructB struct{ X string } // Equal method on pointer receiver
	StructC struct{ X string } // Equal method (with interface argument) on value receiver
	StructD struct{ X string } // Equal method (with interface argument) on pointer receiver
	StructE struct{ X string } // Equal method (with interface argument on value receiver) on pointer receiver
	StructF struct{ X string } // Equal method (with interface argument on pointer receiver) on value receiver

	// These embed the above types as a value.
	StructA1 struct {
		StructA
		X string
	}
	StructB1 struct {
		StructB
		X string
	}
	StructC1 struct {
		StructC
		X string
	}
	StructD1 struct {
		StructD
		X string
	}
	StructE1 struct {
		StructE
		X string
	}
	StructF1 struct {
		StructF
		X string
	}

	// These embed the above types as a pointer.
	StructA2 struct {
		*StructA
		X string
	}
	StructB2 struct {
		*StructB
		X string
	}
	StructC2 struct {
		*StructC
		X string
	}
	StructD2 struct {
		*StructD
		X string
	}
	StructE2 struct {
		*StructE
		X string
	}
	StructF2 struct {
		*StructF
		X string
	}

	StructNo struct{ X string } // Equal method (with interface argument) on non-satisfying receiver

	AssignA func() int
	AssignB struct{ A int }
	AssignC chan bool
	AssignD <-chan bool
)

func (x StructA) Equal(y StructA) bool     { _ = "STUB: not implemented"; return false }
func (x *StructB) Equal(y *StructB) bool   { _ = "STUB: not implemented"; return false }
func (x StructC) Equal(y InterfaceA) bool  { _ = "STUB: not implemented"; return false }
func (x StructC) InterfaceA()              { _ = "STUB: not implemented"; return }
func (x *StructD) Equal(y InterfaceA) bool { _ = "STUB: not implemented"; return false }
func (x *StructD) InterfaceA()             { _ = "STUB: not implemented"; return }
func (x *StructE) Equal(y InterfaceA) bool { _ = "STUB: not implemented"; return false }
func (x StructE) InterfaceA()              { _ = "STUB: not implemented"; return }
func (x StructF) Equal(y InterfaceA) bool  { _ = "STUB: not implemented"; return false }
func (x *StructF) InterfaceA()             { _ = "STUB: not implemented"; return }
func (x StructNo) Equal(y InterfaceA) bool { _ = "STUB: not implemented"; return false }

func (x AssignA) Equal(y func() int) bool      { _ = "STUB: not implemented"; return false }
func (x AssignB) Equal(y struct{ A int }) bool { _ = "STUB: not implemented"; return false }
func (x AssignC) Equal(y chan bool) bool       { _ = "STUB: not implemented"; return false }
func (x AssignD) Equal(y <-chan bool) bool     { _ = "STUB: not implemented"; return false }

var _ = func(
	a StructA, b StructB, c StructC, d StructD, e StructE, f StructF,
	ap *StructA, bp *StructB, cp *StructC, dp *StructD, ep *StructE, fp *StructF,
	a1 StructA1, b1 StructB1, c1 StructC1, d1 StructD1, e1 StructE1, f1 StructF1,
	a2 StructA2, b2 StructB2, c2 StructC2, d2 StructD2, e2 StructE2, f2 StructF1,
) {
	a.Equal(a)
	b.Equal(&b)
	c.Equal(c)
	d.Equal(&d)
	e.Equal(e)
	f.Equal(&f)

	ap.Equal(*ap)
	bp.Equal(bp)
	cp.Equal(*cp)
	dp.Equal(dp)
	ep.Equal(*ep)
	fp.Equal(fp)

	a1.Equal(a1.StructA)
	b1.Equal(&b1.StructB)
	c1.Equal(c1)
	d1.Equal(&d1)
	e1.Equal(e1)
	f1.Equal(&f1)

	a2.Equal(*a2.StructA)
	b2.Equal(b2.StructB)
	c2.Equal(c2)
	d2.Equal(&d2)
	e2.Equal(e2)
	f2.Equal(&f2)
}

type (
	privateStruct struct{ Public, private int }
	PublicStruct  struct{ Public, private int }
	ParentStructA struct{ privateStruct }
	ParentStructB struct{ PublicStruct }
	ParentStructC struct {
		privateStruct
		Public, private int
	}
	ParentStructD struct {
		PublicStruct
		Public, private int
	}
	ParentStructE struct {
		privateStruct
		PublicStruct
	}
	ParentStructF struct {
		privateStruct
		PublicStruct
		Public, private int
	}
	ParentStructG struct {
		*privateStruct
	}
	ParentStructH struct {
		*PublicStruct
	}
	ParentStructI struct {
		*privateStruct
		*PublicStruct
	}
	ParentStructJ struct {
		*privateStruct
		*PublicStruct
		Public  PublicStruct
		private privateStruct
	}
)

func NewParentStructG() *ParentStructG { _ = "STUB: not implemented"; return nil }

func NewParentStructH() *ParentStructH { _ = "STUB: not implemented"; return nil }

func NewParentStructI() *ParentStructI { _ = "STUB: not implemented"; return nil }

func NewParentStructJ() *ParentStructJ { _ = "STUB: not implemented"; return nil }

func (s *privateStruct) SetPrivate(i int)              { _ = "STUB: not implemented"; return }
func (s *PublicStruct) SetPrivate(i int)               { _ = "STUB: not implemented"; return }
func (s *ParentStructC) SetPrivate(i int)              { _ = "STUB: not implemented"; return }
func (s *ParentStructD) SetPrivate(i int)              { _ = "STUB: not implemented"; return }
func (s *ParentStructF) SetPrivate(i int)              { _ = "STUB: not implemented"; return }
func (s *ParentStructA) PrivateStruct() *privateStruct { _ = "STUB: not implemented"; return nil }
func (s *ParentStructC) PrivateStruct() *privateStruct { _ = "STUB: not implemented"; return nil }
func (s *ParentStructE) PrivateStruct() *privateStruct { _ = "STUB: not implemented"; return nil }
func (s *ParentStructF) PrivateStruct() *privateStruct { _ = "STUB: not implemented"; return nil }
func (s *ParentStructG) PrivateStruct() *privateStruct { _ = "STUB: not implemented"; return nil }
func (s *ParentStructI) PrivateStruct() *privateStruct { _ = "STUB: not implemented"; return nil }
func (s *ParentStructJ) PrivateStruct() *privateStruct { _ = "STUB: not implemented"; return nil }
func (s *ParentStructJ) Private() *privateStruct       { _ = "STUB: not implemented"; return nil }
