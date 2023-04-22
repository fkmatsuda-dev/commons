# "ERROREX" Package
The "errorex" package is a Go package that provides a custom error type with additional information such as error code and details. It also includes a function to register error codes and a function to check if an error is of type "errorex" and if its code matches a particular value.

## Installation
To install the "errorex" package, run the following command:
```bash
go get github.com/fkmatsuda-dev/commons/errorex
```
## Usage
To use the "errorex" package, import it in your Go code:
```go
import "github.com/fkmatsuda-dev/commons/errorex"
```

First of all you need to ensure that all the error codes you intend to use are registered, this ensures that there will not be more than one type of error with the same code. My recommendation is that this be done in the init function of one of the .go files in each package you have in your project and preferably that this be named as errorcodes.go. 
```go
const (
    ERROR_CODE_EXAMPLE = "001.001"
)

func init() {
	errorex.RegisterErrorCode(ERROR_CODE_EXAMPLE, "This is an example")
}
```

To create a new "errorex" error, use the errorex.New() function:
```go
err := errorex.New(ERROR_CODE_EXAMPLE, "Error message", "Error details")
```
To check if an error is of type "errorex" and if its code matches a particular value, use the errorex.IS() function:
```go
if errorex.IS(err, ERROR_CODE_EXAMPLE) {
// Handle error
}
```
