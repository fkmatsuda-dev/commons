package errorex

import "fmt"

var (
	errorCodes = make(map[string]string)
)

// EX is a custom errorex type with additional information
type EX interface {
	error
	// Code is the errorex code
	Code() string
	// Message is the errorex message
	Message() string
	// Detail is the errorex detail
	Detail() string
}

// RegisterErrorCode registers errorex codes to prevent repeats
func RegisterErrorCode(code string, description string) {
	// Prevent repeats
	if _, ok := errorCodes[code]; ok {
		// Fatal errorex
		panic(fmt.Errorf("errorex code %s already exists", code))
	}
	errorCodes[code] = description
}

type ex struct {
	code    string
	message string
	detail  string
}

// Code is the errorex code
func (e *ex) Code() string {
	// Return the code
	return e.code
}

// Message is the errorex message
func (e *ex) Message() string {
	// Return the message
	return e.message
}

// Detail is the errorex detail
func (e *ex) Detail() string {
	// Return the detail
	return e.detail
}

// Error returns the errorex message
func (e *ex) Error() string {
	// Return the message
	return fmt.Sprintf("%s: %s", e.code, e.message)
}

// New returns a new errorex.EX
// Code is the errorex code.
// Message is the errorex message.
// Detail is the errorex detail.
func New(code string, message string, detail string) EX {
	// Check if the code exists
	if _, ok := errorCodes[code]; !ok {
		// Fatal errorex
		panic(fmt.Errorf("errorex code %s does not exist", code))
	}
	// Return the exerro
	return &ex{
		code:    code,
		message: message,
		detail:  detail,
	}
}

// IS checks if the errorex is of type EX and if the code matches
func IS(err error, code string) bool {
	// Check if the error code is registered
	if _, ok := errorCodes[code]; !ok {
		// Fatal errorex
		panic(fmt.Errorf("errorex code %s does not exist", code))
	}
	// Check if the error is nil
	if err == nil {
		return false
	}
	// Check if the error is of type EX
	ex, ok := err.(EX)
	if !ok {
		return false
	}
	// Check if the code matches
	return ex.Code() == code
}
