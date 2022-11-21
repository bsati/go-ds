package go_ds

import "testing"

func TestSet(t *testing.T) {
	s := NewSet[int]()

	if s.Contains(10) {
		t.Error("Set should be empty and not contain the element `10`.")
	}

	new := s.Add(10)
	if !new {
		t.Error("Value `10` should have been newly added, but `new` was `false`.")
	}

	new = s.Add(10)
	if new {
		t.Error("Value `10` should not have been newly added, but `new` was `true`.")
	}

	if !s.Contains(10) {
		t.Error("Set should contain value `10` but returned `false` for Contains.")
	}

	if !s.Remove(10) {
		t.Error("Set should have removed value `10` but returned `false` for Remove.")
	}

	if s.Remove(10) {
		t.Error("Set should not have removed value `10` but returned `true` for Remove.")
	}
}
