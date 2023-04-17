package errorex

import (
	"errors"
	"testing"
)

// TestErrorEX tests errorex package
func TestErrorEX(t *testing.T) {

	// Register an error code
	const TEST_ERROR_CODE = "001"
	const UNREGISTERED_TEST_ERROR_CODE = "002"
	RegisterErrorCode(TEST_ERROR_CODE, "Test Error")

	// try to register the same code again
	t.Run("Try to register an existing code", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code was not registered")
			}
		}()
		RegisterErrorCode(TEST_ERROR_CODE, "Test Error")
	})

	// try to create an unregistered error
	t.Run("Try to create an unregistered error", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code was registered")
			}
		}()
		_ = New(UNREGISTERED_TEST_ERROR_CODE, "Unregistered error code", "Unregistered error code error")
	})

	// Test IS
	t.Run("Test IS", func(t *testing.T) {
		err := New(TEST_ERROR_CODE, "Test Error", "Test Error")
		if !IS(err, TEST_ERROR_CODE) {
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
		err := New(TEST_ERROR_CODE, "Unregistered error code", "Unregistered error code error")
		_ = IS(err, UNREGISTERED_TEST_ERROR_CODE)
	})

	// Test IS with a GO error
	t.Run("Test IS with a GO error", func(t *testing.T) {
		err := errors.New("GO error")
		if IS(err, TEST_ERROR_CODE) {
			t.Errorf("Must be false")
		}
	})

	// Test IS with nil
	t.Run("Test IS with nil", func(t *testing.T) {
		if IS(nil, TEST_ERROR_CODE) {
			t.Errorf("Must be false")
		}
	})

}
