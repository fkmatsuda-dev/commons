package slices

import "reflect"

// FindInSlice takes a slices and looks for an element in it. If found it will
// return it's index, otherwise it will return -1 and a bool of false.
// original source https://golangcode.com/check-if-element-exists-in-slice
func FindInSlice[T any](slice []T, val T) (int, bool) {
	return FindInSliceFunc(slice, func(item T) bool {
		return reflect.DeepEqual(item, val)
	})
}

// FindInSliceFunc takes a slices and looks for an element in it using a
// function to compare. If found it will return it's index, otherwise it will
// return -1 and a bool of false.
func FindInSliceFunc[T any](slice []T, f func(T) bool) (int, bool) {

	for i, item := range slice {
		if f(item) {
			return i, true
		}
	}
	return -1, false
}

// InterfaceSlice cast any slices to []interface{}
// original source https://stackoverflow.com/questions/12753805/type-converting-slices-of-interfaces
func InterfaceSlice[T any](slice T) []interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("InterfaceSlice() given a non-slices type")
	}

	// Keep the distinction between nil and empty slices input
	if s.IsNil() {
		return nil
	}

	ret := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret
}
