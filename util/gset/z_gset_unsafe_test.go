package gset

import "testing"

func makeUnsafeSet(ints []int) Set {
	set := NewSet()
	for _, i := range ints {
		set.Add(i)
	}
	return set
}

func Test_AddUnsafeSet(t *testing.T) {
	a := makeUnsafeSet([]int{1, 2, 3})

	if a.Size() != 3 {
		t.Error("AddSet does not have a size of 3 even though 3 items were added to a new set")
	}
}

func Test_AddUnsafeSetNoDuplicate(t *testing.T) {
	a := makeUnsafeSet([]int{7, 5, 3, 7})

	if a.Size() != 3 {
		t.Error("AddSetNoDuplicate set should have 3 elements since 7 is a duplicate")
	}

	if !(a.Contains(7) && a.Contains(5) && a.Contains(3)) {
		t.Error("AddSetNoDuplicate set should have a 7, 5, and 3 in it.")
	}
}

func Test_RemoveUnsafeSet(t *testing.T) {
	a := makeUnsafeSet([]int{6, 3, 1})

	a.Remove(3)

	if a.Size() != 2 {
		t.Error("RemoveSet should only have 2 items in the set")
	}

	if !(a.Contains(6) && a.Contains(1)) {
		t.Error("RemoveSet should have only items 6 and 1 in the set")
	}

	a.Remove(6)
	a.Remove(1)

	if a.Size() != 0 {
		t.Error("RemoveSet should be an empty set after removing 6 and 1")
	}
}

func Test_ContainsMultipleUnsafeSet(t *testing.T) {
	a := makeUnsafeSet([]int{8, 6, 7, 5, 3, 0, 9})

	if !a.Contains(8, 6, 7, 5, 3, 0, 9) {
		t.Error("ContainsAll should contain Jenny's phone number")
	}

	if a.Contains(8, 6, 11, 5, 3, 0, 9) {
		t.Error("ContainsAll should not have all of these numbers")
	}
}

func Test_ClearUnsafeSet(t *testing.T) {
	a := makeUnsafeSet([]int{2, 5, 9, 10})

	a.Clear()

	if a.Size() != 0 {
		t.Error("ClearSet should be an empty set")
	}
}

func Test_ToSliceUnthreadsafe(t *testing.T) {
	s := makeUnsafeSet([]int{1, 2, 3})
	setAsSlice := s.ToSlice()

	if len(setAsSlice) != s.Size() {
		t.Errorf("Set length is incorrect: %v", len(setAsSlice))
	}

	for _, i := range setAsSlice {
		if !s.Contains(i) {
			t.Errorf("Set is missing element: %v", i)
		}
	}
}
