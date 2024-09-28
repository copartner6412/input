package random

import (
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"math/big"
	"unicode"
)

const (
	minStringLength   = 1        // Minimum acceptable string length
	maxStringLength   = 8192     // Maximum acceptable string length
	asciiLowerBound   = 32       // Start of printable ASCII range (space)
	asciiUpperBound   = 126      // End of printable ASCII range (tilde ~)
	unicodeUpperBound = 0x10FFFF // Maximum Unicode character
)

// Behavior:
//  - If `justASCII` is true, only printable ASCII characters (between 32 and 126) will be included in the string.

// String generates a random string of a specified length, with options to include only ASCII characters and to ensure the presence of at least one space.
//
// Parameters:
//   - minLength: The minimum length of the generated string. Must be at least 1.
//   - maxLength: The maximum length of the generated string. Must be at most 8192.
//   - justASCII: A boolean indicating whether the generated string should include only printable ASCII characters.
//
// Returns:
//   - A random string of length between minLength and maxLength using the provided random source.
//   - An error if the length constraints are violated, such as when maxLength is less than minLength,
//     or if minLength is less than 1, or maxLength exceeds 8192.
func String(randomness io.Reader, minLength, maxLength uint, justASCII bool) (string, error) {
	if maxLength < minLength {
		return "", errors.New("maximum length can not be less than minimum length")
	}

	// Validate that the length requirements fall within acceptable system bounds.
	if minLength < minStringLength {
		return "", fmt.Errorf("minimum string length must not be less than %d characters", minStringLength)
	}

	if maxLength > maxStringLength {
		return "", fmt.Errorf("maximum string length must not exceed %d characters", maxStringLength)
	}

	// Determine the length of the string to generate.
	random1, err := rand.Int(randomness, big.NewInt(int64(maxLength-minLength+1)))
	if err != nil {
		return "", fmt.Errorf("error generating a random number for calculating string length: %w", err)
	}
	length := minLength + uint(random1.Int64())

	// Generate the string.
	var generated []rune

	for i := uint(0); i < length; i++ {
		var char rune
		if justASCII {
			// Generate a random printable ASCII character.
			random2, err := rand.Int(randomness, big.NewInt(int64(asciiUpperBound-asciiLowerBound+1)))
			if err != nil {
				return "", fmt.Errorf("error generating a random number for calculating ASCII rune: %w", err)
			}
			char = rune(asciiLowerBound + random2.Int64())
		} else {
			// Generate a random Unicode character.
			random3, err := rand.Int(randomness, big.NewInt(int64(unicodeUpperBound)))
			if err != nil {
				return "", fmt.Errorf("error generating a random number for calculating unicode rune: %w", err)
			}
			char = rune(random3.Int64())
			// If it's an invalid Unicode character, retry.
			if !unicode.IsPrint(char) {
				i--
				continue
			}
		}
		generated = append(generated, char)
	}

	return string(generated), nil
}
