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
		return p(v)
	}
}

// Negate returns a predicate that is a logical negation of this predicate.
func (p P[T]) Negate() P[T] {
	return func(v T) bool {
		return !p(v)
	}
}
