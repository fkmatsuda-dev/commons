package slices

import "testing"

func TestFindInSliceInt(t *testing.T) {
	slice := []int{1, 5, 7, 8}
	var (
		idx   int
		found bool
	)
	idx, found = FindInSlice(slice, 1)
	if !found {
		t.Error("\"1\" exists in slices")
	}
	if idx != 0 {
		t.Error("\"1\" must be in index 0")
	}
	idx, found = FindInSlice(slice, 7)
	if !found {
		t.Error("\"7\" exists in slices")
	}
	if idx != 2 {
		t.Error("\"7\" must be in index 2")
	}
	_, found = FindInSlice(slice, 9)
	if found {
		t.Error("\"9\" not exists in slices")
	}
}

func TestInterfaceSliceWithThreeElements(t *testing.T) {
	slice := []int{1, 2, 3}
	iSlice := InterfaceSlice(slice)
	if iSlice == nil {
		t.Error("InterfaceSlice() must return a non-nil slices")
	}
	if len(iSlice) != 3 {
		t.Error("InterfaceSlice() must return a slices with 3 elements")
	}
	if iSlice[0] != 1 {
		t.Error("InterfaceSlice() must return a slices with 1st element as 1")
	}
	if iSlice[1] != 2 {
		t.Error("InterfaceSlice() must return a slices with 2nd element as 2")
	}
	if iSlice[2] != 3 {
		t.Error("InterfaceSlice() must return a slices with 3rd element as 3")
	}
}

func TestInterfaceSlicePanic(t *testing.T) {
	var slice interface{}
	defer func() {
		if r := recover(); r == nil {
			t.Error("InterfaceSlice() should panic")
		}
	}()
	InterfaceSlice(slice)
}

func TestInterfaceSliceNil(t *testing.T) {
	var slice []int
	if InterfaceSlice(slice) != nil {
		t.Error("InterfaceSlice() should return nil")
	}
}

func TestFind(t *testing.T) {
	slice := []int{1, 5, 7, 8}
	var (
		idx   int
		found bool
	)
	idx, found = Find(slice, 1)
	if !found {
		t.Error("\"1\" exists in slices")
	}
	if idx != 0 {
		t.Error("\"1\" must be in index 0")
	}
	idx, found = Find(slice, 7)
	if !found {
		t.Error("\"7\" exists in slices")
	}
	if idx != 2 {
		t.Error("\"7\" must be in index 2")
	}
	_, found = Find(slice, 9)
	if found {
		t.Error("\"9\" not exists in slices")
	}
}

func TestFindFunc(t *testing.T) {
	slice := []int{1, 5, 7, 8}
	var (
		idx   int
		found bool
	)
	idx, found = FindFunc(slice, func(i int) bool { return i == 1 })
	if !found {
		t.Error("\"1\" exists in slices")
	}
	if idx != 0 {
		t.Error("\"1\" must be in index 0")
	}
	idx, found = FindFunc(slice, func(i int) bool { return i == 7 })
	if !found {
		t.Error("\"7\" exists in slices")
	}
	if idx != 2 {
		t.Error("\"7\" must be in index 2")
	}
	_, found = FindFunc(slice, func(i int) bool { return i == 9 })
	if found {
		t.Error("\"9\" not exists in slices")
	}
}

func TestIndexOf(t *testing.T) {
	slice := []int{1, 5, 7, 8}
	var idx int
	idx = IndexOf(slice, 1)
	if idx != 0 {
		t.Error("\"1\" must be in index 0")
	}
	idx = IndexOf(slice, 7)
	if idx != 2 {
		t.Error("\"7\" must be in index 2")
	}
	idx = IndexOf(slice, 9)
	if idx != -1 {
		t.Error("\"9\" not exists in slices")
	}
}

func TestIndexOfFunc(t *testing.T) {
	slice := []int{1, 5, 7, 8}
	var idx int
	idx = IndexOfFunc(slice, func(i int) bool { return i == 1 })
	if idx != 0 {
		t.Error("\"1\" must be in index 0")
	}
	idx = IndexOfFunc(slice, func(i int) bool { return i == 7 })
	if idx != 2 {
		t.Error("\"7\" must be in index 2")
	}
	idx = IndexOfFunc(slice, func(i int) bool { return i == 9 })
	if idx != -1 {
		t.Error("\"9\" not exists in slices")
	}
}

func TestLastIndexOf(t *testing.T) {
	slice := []int{1, 5, 7, 8, 5}
	var idx int
	idx = LastIndexOf(slice, 1)
	if idx != 0 {
		t.Error("\"1\" must be in index 0")
	}
	idx = LastIndexOf(slice, 5)
	if idx != 4 {
		t.Error("\"5\" must be in index 4")
	}
	idx = LastIndexOf(slice, 9)
	if idx != -1 {
		t.Error("\"9\" not exists in slices")
	}
}

func TestFilter(t *testing.T) {
	slice := []int{1, 5, 7, 8}
	filtered := Filter(slice, func(i int) bool { return i > 5 })
	if len(filtered) != 2 {
		t.Error("Filtered slices must have 2 elements")
	}
	if filtered[0] != 7 {
		t.Error("Filtered slices must have 7 as 1st element")
	}
	if filtered[1] != 8 {
		t.Error("Filtered slices must have 8 as 2nd element")
	}
}

func TestContains(t *testing.T) {
	slice := []int{1, 5, 7, 8}
	if !Contains(slice, 1) {
		t.Error("\"1\" exists in slices")
	}
	if !Contains(slice, 7) {
		t.Error("\"7\" exists in slices")
	}
	if Contains(slice, 9) {
		t.Error("\"9\" not exists in slices")
	}
}

func TestContainsFunc(t *testing.T) {
	slice := []int{1, 5, 7, 8}
	if !ContainsFunc(slice, func(i int) bool { return i == 1 }) {
		t.Error("\"1\" exists in slices")
	}
	if !ContainsFunc(slice, func(i int) bool { return i == 7 }) {
		t.Error("\"7\" exists in slices")
	}
	if ContainsFunc(slice, func(i int) bool { return i == 9 }) {
		t.Error("\"9\" not exists in slices")
	}
}

func TestFindString(t *testing.T) {
	slice := []string{"alpha", "beta", "gamma", "theta"}
	var (
		idx   int
		found bool
	)
	idx, found = Find(slice, "alpha")
	if !found {
		t.Error("\"alpha\" exists in slices")
	}
	if idx != 0 {
		t.Error("\"alpha\" must be in index 0")
	}
	idx, found = Find(slice, "gamma")
	if !found {
		t.Error("\"gamma\" exists in slices")
	}
	if idx != 2 {
		t.Error("\"gamma\" must be in index 2")
	}
	_, found = Find(slice, "delta")
	if found {
		t.Error("\"delta\" not exists in slices")
	}
}

type testStruct struct {
	str   string
	value int
}

func TestFindStruct(t *testing.T) {
	slice := []testStruct{{"alpha", 1}, {"beta", 5}, {"gamma", 7}, {"theta", 8}}
	var (
		idx   int
		found bool
	)
	idx, found = Find(slice, testStruct{"alpha", 1})
	if !found {
		t.Error("\"alpha\" exists in slices")
	}
	if idx != 0 {
		t.Error("\"alpha\" must be in index 0")
	}
	idx, found = Find(slice, testStruct{"gamma", 7})
	if !found {
		t.Error("\"gamma\" exists in slices")
	}
	if idx != 2 {
		t.Error("\"gamma\" must be in index 2")
	}
	_, found = Find(slice, testStruct{"delta", 9})
	if found {
		t.Error("\"delta\" not exists in slices")
	}
}
