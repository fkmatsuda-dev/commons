# "SLICES" Package
The "slices" package is a Go package that provides functions to manipulate slices.

## Installation
To install the "slices" package, run the following command:
```bash
go get github.com/fkmatsuda-dev/commons/slices
```

## Usage
To use the "slices" package, import it in your Go code:
```go
import "github.com/fkmatsuda-dev/commons/slices"
```

## Functions

### InterfaceSlice
InterfaceSlice cast any slices to []interface{}
#### Original source https://stackoverflow.com/questions/12753805/type-converting-slices-of-interfaces
```go
slice := []string{"a", "b", "c"}
interfaceSlice := slices.InterfaceSlice(slice)
fmt.Println(interfaceSlice) // [a b c]
```

### Filter
Filter takes a slices and filters it based on the function passed in.
```go
slice := []string{"a", "b", "c"}
filteredSlice := slices.Filter(slice, func(item string) bool {
    return item != "b"
})
fmt.Println(filteredSlice) // [a c]
```

### Find
Find takes a slices and looks for an element in it. If found it will return it's index, otherwise it will return -1 and a bool of false.
```go
slice := []string{"a", "b", "c"}
index, found := slices.Find(slice, "b)
if found {
    fmt.Println(index) // 1
}
```

### FindFunc
FindFunc takes a slices and looks for an element in it using a function to compare. If found it will return it's index, otherwise it will return -1 and a bool of false.
```go
slice := []string{"a", "b", "c"}
value := "b"

index, found := slices.FindFunc(slice, func(item string) bool {
    return item == value
})
if found {
    fmt.Println(index) // 1
}
```

### IndexOf
IndexOf returns the index of the first instance of the given value, or -1 if not found.
```go
slice := []string{"a", "b", "c"}
index := slices.IndexOf(slice, "b")
fmt.Println(index) // 1
```

### IndexOfFunc
IndexOfFunc returns the index of the first instance using a function to compare, or -1 if not found.
```go
slice := []string{"a", "b", "c"}
value := "b"

index := slices.IndexOfFunc(slice, func(item string) bool {
    return item == value
})
fmt.Println(index) // 1
```

### Contains
Contains returns true if the given value exists in the slice.
```go
slice := []string{"a", "b", "c"}
contains := slices.Contains(slice, "b")
fmt.Println(contains) // true
```

### ContainsFunc
ContainsFunc returns true if the comparison function finds an instance among the elements.
```go
slice := []string{"a", "b", "c"}
value := "b"

contains := slices.ContainsFunc(slice, func(item string) bool {
    return item == value
})
fmt.Println(contains) // true
```

### LastIndexOf
LastIndexOf returns the index of the last instance of the given value, or -1 if not found.
```go
slice := []string{"a", "b", "c", "b"}
index := slices.LastIndexOf(slice, "b")
fmt.Println(index) // 3
```

### LastIndexOfFunc
LastIndexOfFunc returns the index of the last instance using a function to compare, or -1 if not found.
```go
slice := []string{"a", "b", "c", "b"}
value := "b"

index := slices.LastIndexOfFunc(slice, func(item string) bool {
    return item == value
})
fmt.Println(index) // 3
```

### FindInSlice
FindInSlice takes a slices and looks for an element in it. If found it will return it's index, otherwise it will return -1 and a bool of false.
#### Consider using the find function described above
```go
slice := []string{"a", "b", "c"}
index, found := slices.FindInSlice(slice, "b")
if found {
    fmt.Println(index) // 1
}
```

### FindInSliceFunc
FindInSliceFunc takes a slices and looks for an element in it using a function to compare. If found it will return it's index, otherwise it will
#### Consider using the find function described above
```go
slice := []string{"a", "b", "c"}
value := "b"

index, found := slices.FindInSliceFunc(slice, func(item string) bool {
    return item == value
})
if found {
    fmt.Println(index) // 1
}
```
