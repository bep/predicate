// Copyright 2024 Bjørn Erik Pedersen
// SPDX-License-Identifier: MIT

package predicate_test

import (
	"fmt"
	"testing"

	"github.com/bep/predicate"

	qt "github.com/frankban/quicktest"
)

func TestP(t *testing.T) {
	c := qt.New(t)

	var p predicate.P[int] = intP1

	c.Assert(p(1), qt.IsTrue)
	c.Assert(p(2), qt.IsFalse)

	neg := p.Negate()
	c.Assert(neg(1), qt.IsFalse)
	c.Assert(neg(2), qt.IsTrue)

	and := p.And(intP2)
	c.Assert(and(1), qt.IsFalse)
	c.Assert(and(2), qt.IsFalse)
	c.Assert(and(10), qt.IsTrue)

	or := p.Or(intP2)
	c.Assert(or(1), qt.IsTrue)
	c.Assert(or(2), qt.IsTrue)
	c.Assert(or(10), qt.IsTrue)
	c.Assert(or(11), qt.IsFalse)

	var nilp predicate.P[int] = nil
	c.Assert(nilp.Or(intP1)(1), qt.IsTrue)
	c.Assert(nilp.And(intP1)(1), qt.IsTrue)

	var zerop predicate.P[int]
	c.Assert(zerop.Or(intP1)(1), qt.IsTrue)
	c.Assert(zerop.And(intP1)(1), qt.IsTrue)
}

func TestFilter(t *testing.T) {
	c := qt.New(t)

	var p predicate.P[int] = intP1
	p = p.Or(intP2)

	ints := []int{1, 2, 3, 4, 1, 6, 7, 8, 2}

	c.Assert(p.Filter(ints), qt.DeepEquals, []int{1, 2, 1, 2})
	c.Assert(ints, qt.DeepEquals, []int{1, 2, 1, 2, 1, 6, 7, 8, 2})
}

func TestFilterCopy(t *testing.T) {
	c := qt.New(t)

	var p predicate.P[int] = intP1
	p = p.Or(intP2)

	ints := []int{1, 2, 3, 4, 1, 6, 7, 8, 2}

	c.Assert(p.FilterCopy(ints), qt.DeepEquals, []int{1, 2, 1, 2})
	c.Assert(ints, qt.DeepEquals, []int{1, 2, 3, 4, 1, 6, 7, 8, 2})
}

var intP1 = func(i int) bool {
	if i == 10 {
		return true
	}
	return i == 1
}

var intP2 = func(i int) bool {
	if i == 10 {
		return true
	}
	return i == 2
}

func ExampleP() {
	var (
		pHello predicate.P[string] = func(s string) bool {
			return s == "hello"
		}
		pWorld predicate.P[string] = func(s string) bool {
			return s == "world"
		}
		pAny predicate.P[string] = func(s string) bool {
			return s != ""
		}
	)

	fmt.Println("Or (true):", pHello.Or(pWorld)("hello"))
	fmt.Println("Or (false):", pHello.Or(pWorld)("foo"))
	fmt.Println("And (false):", pHello.And(pWorld)("hello"))
	fmt.Println("And (true):", pHello.And(pAny)("hello"))
	fmt.Println("Negate (false):", pHello.Negate()("hello"))
	fmt.Println("Negate (true):", pHello.Negate()("world"))
	fmt.Println("Chained (true):", pHello.And(pAny.Or(pWorld))("hello"))
	fmt.Println("Chained (false):", pHello.And(pAny.Or(pWorld))("foo"))

	// Output:
	// Or (true): true
	// Or (false): false
	// And (false): false
	// And (true): true
	// Negate (false): false
	// Negate (true): true
	// Chained (true): true
	// Chained (false): false
}
