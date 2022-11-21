package go_ds

// NewStack creates a new empty Stack for the given type
func NewStack[T any]() Stack[T] {
	return Stack[T]{size: 0}
}

// Stack implements a basic last-in-first-out datastructure with generic typing
type Stack[T any] struct {
	top  *node[T]
	size int
}

// node is used to implement a linked list which is the underlying data structure of the stack
type node[T any] struct {
	val  T
	prev *node[T]
}

// Push adds the given value to the stack
func (s *Stack[T]) Push(value T) {
	s.top = &node[T]{val: value, prev: s.top}
	s.size++
}

// Pop removes the top value of the stack and returns it.
// If the stack is empty the second return value will be false, otherwise true.
func (s *Stack[T]) Pop() (T, bool) {
	var result T
	if s.top == nil {
		return result, false
	}
	result = s.top.val
	s.top = s.top.prev
	s.size--
	return result, true
}

// Peek is similar to Pop but does not remove the top element from the stack.
// The second return value also indicates existence just like in Pop.
func (s *Stack[T]) Peek() (T, bool) {
	var result T
	if s.top == nil {
		return result, false
	}
	return s.top.val, true
}

// IsEmpty returns whether the stack has no elements
func (s *Stack[T]) IsEmpty() bool {
	return s.size == 0
}
