package random

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const (
	minLinuxHostnameLengthAllowed uint = 1
	maxLinuxHostnameLengthAllowed uint = 64
)

// Future development: LinuxHostname(word bool, length uint) (string, error)

// LinuxHostname generates a random Linux hostname.
// A valid hostname has a length between 1 and 64 characters, begins with letter and ends with an alphanumeric character, and may contain hyphens in the middle.
func LinuxHostname(minLength, maxLength uint) (string, error) {
	length, err := checkLength(minLength, maxLength, minLinuxHostnameLengthAllowed, maxLinuxHostnameLengthAllowed)
	if err != nil {
		return "", err
	}

	// Define character sets for allowed hostname characters.
	allowedCharacters := append(lowerAlphanumericalRunes, '-') // Middle characters can include hyphens.

	// Create a rune slice to hold the generated hostname.
	hostname := make([]rune, length)

	random1, err := rand.Int(rand.Reader, big.NewInt(int64(len(lowerCaseRunes))))
	if err != nil {
		return "", fmt.Errorf("error generating a random index for the first character: %w", err)
	}

	// First character: must be alphanumeric.
	hostname[0] = lowerCaseRunes[random1.Int64()]

	// Middle characters: can include hyphens.
	for i := 1; i < int(length)-1; i++ {
		random2, err := rand.Int(rand.Reader, big.NewInt(int64(len(allowedCharacters))))
		if err != nil {
			return "", fmt.Errorf("error generating a random index for a middle character: %w", err)
		}
		hostname[i] = allowedCharacters[random2.Int64()]
	}

	// Last character: must be alphanumeric (if the total length is greater than 1).
	if length > 1 {
		random3, err := rand.Int(rand.Reader, big.NewInt(int64(len(lowerAlphanumericalRunes))))
		if err != nil {
			return "", fmt.Errorf("error generating a random index for the last character: %w", err)
		}
		hostname[length-1] = lowerAlphanumericalRunes[random3.Int64()]
	}

	// Convert the rune slice to a string and return the result.
	return string(hostname), nil
}