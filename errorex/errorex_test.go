package errorex

import (
	"errors"
	"testing"
)

// TestErrorEX tests errorex package
func TestErrorEX(t *testing.T) {

	// Register an error code
	const TestErrorCode = "001"
	const UnregisteredTestErrorCode = "002"
	const description = "Test Error"
	RegisterErrorCode(TestErrorCode, description)

	// try to register the same code again
	t.Run("Try to register an existing code", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code was not registered")
			}
		}()
		RegisterErrorCode(TestErrorCode, description)
	})

	// try to create an unregistered error
	t.Run("Try to create an unregistered error", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code was registered")
			}
		}()
		_ = New(UnregisteredTestErrorCode, "Unregistered error code", "Unregistered error code error")
	})

	// Test IS
	t.Run("Test IS", func(t *testing.T) {
		err := New(TestErrorCode, description, description)
		if !IS(err, TestErrorCode) {
			t.Errorf("The error code is not the same")
		}
	})

	// Test IS with unregistered error code
	t.Run("Test IS with unregistered error code", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code was registered")
			}
		}()
		err := New(TestErrorCode, "Unregistered error code", "Unregistered error code error")
		_ = IS(err, UnregisteredTestErrorCode)
	})

	// Test IS with a GO error
	t.Run("Test IS with a GO error", func(t *testing.T) {
		err := errors.New("GO error")
		if IS(err, TestErrorCode) {
			t.Errorf("Must be false")
		}
	})

	// Test IS with nil
	t.Run("Test IS with nil", func(t *testing.T) {
		if IS(nil, TestErrorCode) {
			t.Errorf("Must be false")
		}
	})

}
