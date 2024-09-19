package pseudorandom

import (
	"math/rand/v2"
	"unicode"
)

const (
	minStringLengthAllowed   = 1        // Minimum acceptable string length
	maxStringLengthAllowed   = 8192     // Maximum acceptable string length
	asciiLowerBound   = 32       // Start of printable ASCII range (space)
	asciiUpperBound   = 126      // End of printable ASCII range (tilde ~)
	unicodeUpperBound = 0x10FFFF // Maximum Unicode character
)

// Behavior:
//  - If `justASCII` is true, only printable ASCII characters (between 32 and 126) will be included in the string.

// String generates a deterministic pseudo-random string of a specified length, with options to include only ASCII characters and to ensure the presence of at least one space.
//
// Parameters:
//   - r: A pointer to a rand.Rand object, used to generate the pseudo-random characters deterministically.
//   - minLength: The minimum length of the generated string. Must be at least 1.
//   - maxLength: The maximum length of the generated string. Must be at most 8192.
//   - justASCII: A boolean indicating whether the generated string should include only printable ASCII characters.
//
// Returns:
//   - A deterministic pseudo-random string of length between minLength and maxLength using the provided random source.
//   - An error if the length constraints are violated, such as when maxLength is less than minLength,
//     or if minLength is less than 1, or maxLength exceeds 8192.
func String(r *rand.Rand, minLength, maxLength uint, justASCII bool) (string, error) {
	length, err := checkLength(r, minLength, maxLength, minStringLengthAllowed, maxStringLengthAllowed)
	if err != nil {
		return "", err
	}

	// Generate the string.
	var generated []rune

	for i := 0; i < int(length); i++ {
		var char rune
		if justASCII {
			// Generate a random printable ASCII character.
			char = rune(asciiLowerBound + r.IntN(asciiUpperBound-asciiLowerBound+1))
		} else {
			// Generate a random Unicode character.
			char = rune(r.IntN(unicodeUpperBound))
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
