// Copyright 2017, The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package teststructs

import (
	"time"

	pb "github.com/google/go-cmp/cmp/internal/testprotos"
)

// This is an sanitized example of equality from a real use-case.
// The original equality function was as follows:
/*
func equalCartel(x, y Cartel) bool {
	if !(equalHeadquarter(x.Headquarter, y.Headquarter) &&
		x.Source() == y.Source() &&
		x.CreationDate().Equal(y.CreationDate()) &&
		x.Boss() == y.Boss() &&
		x.LastCrimeDate().Equal(y.LastCrimeDate())) {
		return false
	}
	if len(x.Poisons()) != len(y.Poisons()) {
		return false
	}
	for i := range x.Poisons() {
		if !equalPoison(*x.Poisons()[i], *y.Poisons()[i]) {
			return false
		}
	}
	return true
}
func equalHeadquarter(x, y Headquarter) bool {
	xr, yr := x.Restrictions(), y.Restrictions()
	return x.ID() == y.ID() &&
		x.Location() == y.Location() &&
		reflect.DeepEqual(x.SubDivisions(), y.SubDivisions()) &&
		x.IncorporatedDate().Equal(y.IncorporatedDate()) &&
		pb.Equal(x.MetaData(), y.MetaData()) &&
		bytes.Equal(x.PrivateMessage(), y.PrivateMessage()) &&
		bytes.Equal(x.PublicMessage(), y.PublicMessage()) &&
		x.HorseBack() == y.HorseBack() &&
		x.Rattle() == y.Rattle() &&
		x.Convulsion() == y.Convulsion() &&
		x.Expansion() == y.Expansion() &&
		x.Status() == y.Status() &&
		pb.Equal(&xr, &yr) &&
		x.CreationTime().Equal(y.CreationTime())
}
func equalPoison(x, y Poison) bool {
	return x.PoisonType() == y.PoisonType() &&
		x.Expiration().Equal(y.Expiration()) &&
		x.Manufacturer() == y.Manufacturer() &&
		x.Potency() == y.Potency()
}
*/

type Cartel struct {
	Headquarter
	source        string
	creationDate  time.Time
	boss          string
	lastCrimeDate time.Time
	poisons       []*Poison
}

func (p Cartel) Source() string           { _ = "STUB: not implemented"; return "" }
func (p Cartel) CreationDate() time.Time  { _ = "STUB: not implemented"; return *new(time.Time) }
func (p Cartel) Boss() string             { _ = "STUB: not implemented"; return "" }
func (p Cartel) LastCrimeDate() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }
func (p Cartel) Poisons() []*Poison       { _ = "STUB: not implemented"; return nil }

func (p *Cartel) SetSource(x string)           { _ = "STUB: not implemented"; return }
func (p *Cartel) SetCreationDate(x time.Time)  { _ = "STUB: not implemented"; return }
func (p *Cartel) SetBoss(x string)             { _ = "STUB: not implemented"; return }
func (p *Cartel) SetLastCrimeDate(x time.Time) { _ = "STUB: not implemented"; return }
func (p *Cartel) SetPoisons(x []*Poison)       { _ = "STUB: not implemented"; return }

type Headquarter struct {
	id               uint64
	location         string
	subDivisions     []string
	incorporatedDate time.Time
	metaData         *pb.MetaData
	privateMessage   []byte
	publicMessage    []byte
	horseBack        string
	rattle           string
	convulsion       bool
	expansion        uint64
	status           pb.HoneyStatus
	restrictions     pb.Restrictions
	creationTime     time.Time
}

func (hq Headquarter) ID() uint64             { _ = "STUB: not implemented"; return 0 }
func (hq Headquarter) Location() string       { _ = "STUB: not implemented"; return "" }
func (hq Headquarter) SubDivisions() []string { _ = "STUB: not implemented"; return nil }
func (hq Headquarter) IncorporatedDate() time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}
func (hq Headquarter) MetaData() *pb.MetaData { _ = "STUB: not implemented"; return nil }
func (hq Headquarter) PrivateMessage() []byte { _ = "STUB: not implemented"; return nil }
func (hq Headquarter) PublicMessage() []byte  { _ = "STUB: not implemented"; return nil }
func (hq Headquarter) HorseBack() string      { _ = "STUB: not implemented"; return "" }
func (hq Headquarter) Rattle() string         { _ = "STUB: not implemented"; return "" }
func (hq Headquarter) Convulsion() bool       { _ = "STUB: not implemented"; return false }
func (hq Headquarter) Expansion() uint64      { _ = "STUB: not implemented"; return 0 }
func (hq Headquarter) Status() pb.HoneyStatus {
	_ = "STUB: not implemented"
	return *new(pb.HoneyStatus)
}
func (hq Headquarter) Restrictions() pb.Restrictions {
	_ = "STUB: not implemented"
	return *new(pb.Restrictions)
}
func (hq Headquarter) CreationTime() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (hq *Headquarter) SetID(x uint64)                    { _ = "STUB: not implemented"; return }
func (hq *Headquarter) SetLocation(x string)              { _ = "STUB: not implemented"; return }
func (hq *Headquarter) SetSubDivisions(x []string)        { _ = "STUB: not implemented"; return }
func (hq *Headquarter) SetIncorporatedDate(x time.Time)   { _ = "STUB: not implemented"; return }
func (hq *Headquarter) SetMetaData(x *pb.MetaData)        { _ = "STUB: not implemented"; return }
func (hq *Headquarter) SetPrivateMessage(x []byte)        { _ = "STUB: not implemented"; return }
func (hq *Headquarter) SetPublicMessage(x []byte)         { _ = "STUB: not implemented"; return }
func (hq *Headquarter) SetHorseBack(x string)             { _ = "STUB: not implemented"; return }
func (hq *Headquarter) SetRattle(x string)                { _ = "STUB: not implemented"; return }
func (hq *Headquarter) SetConvulsion(x bool)              { _ = "STUB: not implemented"; return }
func (hq *Headquarter) SetExpansion(x uint64)             { _ = "STUB: not implemented"; return }
func (hq *Headquarter) SetStatus(x pb.HoneyStatus)        { _ = "STUB: not implemented"; return }
func (hq *Headquarter) SetRestrictions(x pb.Restrictions) { _ = "STUB: not implemented"; return }
func (hq *Headquarter) SetCreationTime(x time.Time)       { _ = "STUB: not implemented"; return }

type Poison struct {
	poisonType   pb.PoisonType
	expiration   time.Time
	manufacturer string
	potency      int
}

func (p Poison) PoisonType() pb.PoisonType { _ = "STUB: not implemented"; return *new(pb.PoisonType) }
func (p Poison) Expiration() time.Time     { _ = "STUB: not implemented"; return *new(time.Time) }
func (p Poison) Manufacturer() string      { _ = "STUB: not implemented"; return "" }
func (p Poison) Potency() int              { _ = "STUB: not implemented"; return 0 }

func (p *Poison) SetPoisonType(x pb.PoisonType) { _ = "STUB: not implemented"; return }
func (p *Poison) SetExpiration(x time.Time)     { _ = "STUB: not implemented"; return }
func (p *Poison) SetManufacturer(x string)      { _ = "STUB: not implemented"; return }
func (p *Poison) SetPotency(x int)              { _ = "STUB: not implemented"; return }
