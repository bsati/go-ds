package go_ds

import (
	"testing"
)

func TestArrayListAdd(t *testing.T) {
	list := NewArrayList[int]()
	list.Add(4)
	list.Add(3)

	if list.elements[0] != 4 {
		t.Error("Expected first element of the list to be value 4 but got: ", list.elements[0])
	}

	if list.elements[1] != 3 {
		t.Error("Expected second element of the list to be value 3 but got: ", list.elements[1])
	}
}

func TestArrayListRemove(t *testing.T) {
	list := NewArrayList[int]()
	list.Add(4)
	list.Add(3)

	removed := list.Remove(1)
	if !removed {
		t.Error("Expected list to remove element at index 1")
	}
	if list.size != 1 {
		t.Error("Expected list size to be 1, but got: ", list.size)
	}
	if list.elements[0] != 4 {
		t.Error("Expected first element of the list to be 4 but got: ", list.elements[0])
	}

	list.Add(3)
	list.Add(5)

	removed = list.Remove(1)

	if !removed {
		t.Error("Expected list to remove element at index 1")
	}
	if list.size != 2 {
		t.Error("Expected list size to be 2, but got: ", list.size)
	}
	if list.elements[1] != 5 {
		t.Error("Expected second element to be moved element 5, but got: ", list.elements[1])
	}
}

func TestArrayListFind(t *testing.T) {
	list := NewArrayList[int]()
	list.Add(4)
	list.Add(5)
	list.Add(6)

	if index := list.Find(4); index != 0 {
		t.Error("Expected element 4 to be found at index 0, but got: ", index)
	}

	if index := list.Find(6); index != 2 {
		t.Error("Expected element 6 to be found at index 2, but got: ", index)
	}

	if index := list.Find(7); index != -1 {
		t.Error("Expected element 7 not to be found, but got index: ", index)
	}
}
