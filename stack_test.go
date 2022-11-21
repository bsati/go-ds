package go_ds

import "testing"

func TestStack(t *testing.T) {
	stack := NewStack[int]()
	_, peek := stack.Peek()
	if peek {
		t.Error("Peeking an empty stack should return false")
	}

	stack.Push(10)
	elem, peek := stack.Peek()
	if elem != 10 || !peek {
		t.Errorf("Element should be 10 but was %d and peek should be `false` but was %v", elem, peek)
	}

	elem, res := stack.Pop()
	if elem != 10 || !res {
		t.Errorf("Element should be 10 but was %d and res should be `false` but was %v", elem, res)
	}

	if !stack.IsEmpty() {
		t.Error("Expected stack to be empty")
	}

	_, res = stack.Pop()
	if res {
		t.Error("Expected res to be `false` but was `true`")
	}

	stack.Push(11)
	stack.Push(12)

	elem, res = stack.Pop()
	if elem != 12 || !res {
		t.Errorf("Element should be 12 but was %d and res should be `false` but was %v", elem, res)
	}
}
