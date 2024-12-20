// Copyright 2024 Bjørn Erik Pedersen
// SPDX-License-Identifier: MIT

package predicate

// P is a predicate function that tests whether a value of type T satisfies some condition.
type P[T any] func(T) bool

// And returns a predicate that is a short-circuiting logical AND of this and the given predicates.
func (p P[T]) And(ps ...P[T]) P[T] {
	return func(v T) bool {
		for _, pp := range ps {
			if !pp(v) {
				return false
			}
		}
		if p == nil {
			return true
		}
		return p(v)
	}
}

// Or returns a predicate that is a short-circuiting logical OR of this and the given predicates.
func (p P[T]) Or(ps ...P[T]) P[T] {
	return func(v T) bool {
		for _, pp := range ps {
			if pp(v) {
				return true
			}
		}
		if p == nil {
			return false
		}
		return p(v)
	}
}

// Negate returns a predicate that is a logical negation of this predicate.
func (p P[T]) Negate() P[T] {
	return func(v T) bool {
		return !p(v)
	}
}

// Filter returns a new slice holding only the elements of s that satisfy p.
// Filter modifies the contents of the slice s and returns the modified slice, which may have a smaller length.
func (p P[T]) Filter(s []T) []T {
	var n int
	for _, v := range s {
		if p(v) {
			s[n] = v
			n++
		}
	}
	return s[:n]
}

// FilterCopy returns a new slice holding only the elements of s that satisfy p.
func (p P[T]) FilterCopy(s []T) []T {
	var result []T
	for _, v := range s {
		if p(v) {
			result = append(result, v)
		}
	}
	return result
}
