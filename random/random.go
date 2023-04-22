package random

import (
	"math/rand"
	"time"
)

var (
	rnd         = rand.New(rand.NewSource(time.Now().UnixNano()))
	letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
)

// Int generates a random integer between min and max
func Int(min, max int) int {
	return rnd.Intn(max-min+1) + min
}

// String generates a random string of a given length
func String(length int) string {
	// Create a new random string
	randomString := ""
	// Generate a random string
	for i := 0; i < length; i++ {
		// Generate a random character
		randomCharacter := Character()
		// Append the random character to the random string
		randomString += string(randomCharacter)
	}
	// Return the random string
	return randomString
}

// Character generates a random character
func Character() rune {
	return letterRunes[Int(0, len(letterRunes)-1)]
}
