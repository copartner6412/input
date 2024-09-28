package random

import (
	"crypto/rand"
	"fmt"
	"io"
	"math/big"
)

const (
	minSubdomainLengthAllowed uint = 1
	maxSubdomainLengthAllowed uint = 63
)

// Subdomain generates a random subdomain string with a length between 1 and 63.
// The first character must be a letter, the middle characters can include letters, digits, or hyphens, and the last character must be a letter or digit.
func Subdomain(randomness io.Reader, minLength, maxLength uint) (string, error) {
	length, err := checkLength(randomness, minLength, maxLength, minSubdomainLengthAllowed, maxSubdomainLengthAllowed)
	if err != nil {
		return "", err
	}

	// Allowed characters: lowercase letters, digits, and hyphens (for middle characters).
	allowedCharacters := append(lowerAlphanumericalRunes, '-')

	// Preallocate the slice with the required subdomain length.
	subdomain := make([]rune, 0, length)

	// First character: must be a lowercase letter.
	random1, err := rand.Int(randomness, big.NewInt(int64(len(lowerAlphanumericalRunes))))
	if err != nil {
		return "", fmt.Errorf("error generating a random index for the first character: %w", err)
	}

	subdomain = append(subdomain, lowerAlphanumericalRunes[random1.Int64()])

	// Middle characters: can include lowercase letters, digits, or hyphens.
	for i := 1; i < int(length)-1; i++ {
		random2, err := rand.Int(randomness, big.NewInt(int64(len(allowedCharacters))))
		if err != nil {
			return "", fmt.Errorf("error generating a random index for a middle character: %w", err)
		}
		subdomain = append(subdomain, allowedCharacters[random2.Int64()])
	}

	// Last character: must be a letter or digit (if subdomainLength > 1).
	if length > 1 {
		random3, err := rand.Int(randomness, big.NewInt(int64(len(lowerAlphanumericalRunes))))
		if err != nil {
			return "", fmt.Errorf("error generating a random index for the last character: %w", err)
		}

		subdomain = append(subdomain, lowerAlphanumericalRunes[random3.Int64()])
	}

	return string(subdomain), nil
}
