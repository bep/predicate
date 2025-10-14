// Copyright 2024 Bjørn Erik Pedersen
// SPDX-License-Identifier: MIT

package predicate

// Match represents the result of a predicate evaluation.
type Match interface {
	OK() bool
}

var (
	// Predefined Match values for common cases.
	True  = BoolMatch(true)
	False = BoolMatch(false)
)

// BoolMatch is a simple Match implementation based on a boolean value.
type BoolMatch bool

func (b BoolMatch) OK() bool {
	return bool(b)
}

// breakMatch is a Match implementation that always returns false for OK() and signals to break evaluation.
type breakMatch struct{}

func (b breakMatch) OK() bool {
	return false
}

var matchBreak = breakMatch{}

// P is a predicate function that tests whether a value of type T satisfies some condition.
type P[T any] func(T) bool

// PR is a predicate function that tests whether a value of type T satisfies some condition and returns a Match result.
type PR[T any] func(T) Match

// BoolFunc returns a P[T] version of this predicate.
func (p PR[T]) BoolFunc() P[T] {
	return func(v T) bool {
		if p == nil {
			return false
		}
		return p(v).OK()
	}
}

// And returns a predicate that is a short-circuiting logical AND of this and the given predicates.
func (p PR[T]) And(ps ...PR[T]) PR[T] {
	return func(v T) Match {
		if p != nil {
			m := p(v)
			if !m.OK() || shouldBreak(m) {
				return matchBreak
			}
		}
		for _, pp := range ps {
			m := pp(v)
			if !m.OK() || shouldBreak(m) {
				return matchBreak
			}
		}
		return BoolMatch(true)
	}
}

// Or returns a predicate that is a short-circuiting logical OR of this and the given predicates.
func (p PR[T]) Or(ps ...PR[T]) PR[T] {
	return func(v T) Match {
		if p != nil {
			m := p(v)
			if m.OK() {
				return m
			}
			if shouldBreak(m) {
				return matchBreak
			}
		}
		for _, pp := range ps {
			m := pp(v)
			if m.OK() {
				return m
			}
			if shouldBreak(m) {
				return matchBreak
			}
		}
		return BoolMatch(false)
	}
}

func shouldBreak(m Match) bool {
	_, ok := m.(breakMatch)
	return ok
}

// Filter returns a new slice holding only the elements of s that satisfy p.
// Filter modifies the contents of the slice s and returns the modified slice, which may have a smaller length.
func (p PR[T]) Filter(s []T) []T {
	var n int
	for _, v := range s {
		if p(v).OK() {
			s[n] = v
			n++
		}
	}
	return s[:n]
}

// FilterCopy returns a new slice holding only the elements of s that satisfy p.
func (p PR[T]) FilterCopy(s []T) []T {
	var result []T
	for _, v := range s {
		if p(v).OK() {
			result = append(result, v)
		}
	}
	return result
}
