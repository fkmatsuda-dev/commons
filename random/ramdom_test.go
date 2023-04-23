package random

import "testing"

// TestRandom tests random package
func TestRandom(t *testing.T) {
	// Test Int
	t.Run("Test Int", func(t *testing.T) {
		min := 1
		max := 10
		r := Int(min, max)
		if r < min || r > max {
			t.Errorf("The random number is out of range")
		}
	})

	// Test String
	t.Run("Test String", func(t *testing.T) {
		length := 10
		r := String(length)
		if len(r) != length {
			t.Errorf("The random string length is not the same")
		}
	})
}
