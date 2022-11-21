package go_ds

// NewSet creates a new set with the given datatype
func NewSet[T comparable]() Set[T] {
	return make(map[T]struct{})
}

// Set is an abstraction of the built-in map to allow a more user-friendly API.
// The underlying map uses the set items as keys and struct{}{} as values since this is a zero-sized type.
type Set[T comparable] map[T]struct{}

// Add inserts the element into the set and returns if the value was newly inserted (`true`) or overriden (`false`).
func (s Set[T]) Add(value T) bool {
	new := true
	if _, ok := s[value]; ok {
		new = false
	}
	s[value] = struct{}{}
	return new
}

// Remove deletes the element from the set and returns whether it existed.
func (s Set[T]) Remove(value T) bool {
	if _, ok := s[value]; ok {
		delete(s, value)
		return true
	}
	return false
}

// Contains checks whether the value is in the set and returns a corresponding boolean.
func (s Set[T]) Contains(value T) bool {
	_, ok := s[value]
	return ok
}
