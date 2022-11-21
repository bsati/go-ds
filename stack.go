package go_ds

type Node[T any] struct {
	val  T
	prev *Node[T]
}

type Stack[T any] struct {
	top  *Node[T]
	size int
}

func NewStack[T any]() Stack[T] {
	return Stack[T]{size: 0}
}

func (s *Stack[T]) Push(value T) {
	s.top = &Node[T]{val: value, prev: s.top}
	s.size++
}

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

func (s *Stack[T]) Peek() (T, bool) {
	var result T
	if s.top == nil {
		return result, false
	}
	return s.top.val, true
}

func (s *Stack[T]) IsEmpty() bool {
	return s.size == 0
}
