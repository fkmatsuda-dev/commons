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

To find an element index in a slice, use the slices.FindInSlice() function:
```go
slice := []string{"a", "b", "c"}
index, found := slices.FindInSlice(slice, "b")
if found {
    fmt.Println(index) // 1
}
```

You can also find an element index in a slice using a custom comparator function:
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

If you need to convert a slice of some type to a slice of interface{}, use the slices.InterfaceSlice() function:
```go
slice := []string{"a", "b", "c"}
interfaceSlice := slices.InterfaceSlice(slice)
fmt.Println(interfaceSlice) // [a b c]
```
