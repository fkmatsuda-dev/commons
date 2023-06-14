package slices

import "reflect"

// FindInSlice takes a slices and looks for an element in it. If found it will
// return it's index, otherwise it will return -1 and a bool of false.
// Deprecated: Use Find instead.
func FindInSlice[T any](slice []T, val T) (int, bool) {
	return Find(slice, val)
}

// FindInSliceFunc takes a slices and looks for an element in it using a
// function to compare. If found it will return it's index, otherwise it will return -1 and a bool of false.
// Deprecated: Use FindFunc instead.
func FindInSliceFunc[T any](slice []T, f func(T) bool) (int, bool) {
	return FindFunc(slice, f)
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

// Filter takes a slices and filters it based on the function passed in.
func Filter[T any](s []T, f func(T) bool) []T {
	var ret []T
	for _, item := range s {
		if f(item) {
			ret = append(ret, item)
		}
	}
	return ret
}

// Find takes a slices and looks for an element in it. If found it will return it's index, otherwise it will return -1 and a bool of false.
func Find[T any](s []T, val T) (int, bool) {
	return FindFunc(s, func(item T) bool {
		return reflect.DeepEqual(item, val)
	})
}

// FindFunc takes a slices and looks for an element in it using a function to compare. If found it will return it's index, otherwise it will return -1 and a bool of false.
func FindFunc[T any](s []T, f func(T) bool) (int, bool) {
	for i, item := range s {
		if f(item) {
			return i, true
		}
	}
	return -1, false
}

// IndexOf returns the index of the first instance of the given value, or -1 if not found.
func IndexOf[T any](s []T, val T) int {
	i, _ := Find(s, val)
	return i
}

// IndexOfFunc returns the index of the first instance using a function to compare, or -1 if not found.
func IndexOfFunc[T any](s []T, f func(T) bool) int {
	i, _ := FindFunc(s, f)
	return i
}

// Contains returns true if the given value exists in the slice.
func Contains[T any](s []T, val T) bool {
	_, exists := Find(s, val)
	return exists
}

// ContainsFunc returns true if the comparison function finds an instance among the elements.
func ContainsFunc[T any](s []T, f func(T) bool) bool {
	_, exists := FindFunc(s, f)
	return exists
}

// LastIndexOf returns the index of the last instance of the given value, or -1 if not found.
func LastIndexOf[T any](s []T, val T) int {
	return LastIndexOfFunc(s, func(item T) bool {
		return reflect.DeepEqual(item, val)
	})
}

// LastIndexOfFunc returns the index of the last instance using a function to compare, or -1 if not found.
func LastIndexOfFunc[T any](s []T, f func(T) bool) int {
	for i := len(s) - 1; i >= 0; i-- {
		if f(s[i]) {
			return i
		}
	}
	return -1
}
