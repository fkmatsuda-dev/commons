# "RANDOM" Package
The "random" package is a Go package that provides functions to generate random strings, characters and integers.

## Installation
To install the "random" package, run the following command:
```bash
go get github.com/fkmatsuda-dev/commons/random
```

## Usage
To use the "random" package, import it in your Go code:
```go
import "github.com/fkmatsuda-dev/commons/random"
```

## Random Strings
The "random" package provides a function to generate a random string of a specified length.

To generate a random string, use the random.String() function:
```go
randomString, err := random.String(32)
if err != nil {
    // Handle error
}
```

## Random Characters
The "random" package provides a function to generate a random character.

To generate a random character, use the random.Char() function:
```go
randomChar, err := random.Char()
if err != nil {
    // Handle error
}
```

## Random Integers
The "random" package provides a function to generate a random integer.

To generate a random integer, use the random.Int() function:
```go
randomInt, err := random.Int()
if err != nil {
    // Handle error
}
```
