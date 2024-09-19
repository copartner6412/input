package pseudorandom

import "math/rand/v2"

const (
	minLinuxHostnameLength uint = 1
	maxLinuxHostnameLength uint = 64
	maxLinuxHostnameWords  uint = 12
)

// LinuxHostname generates a deterministic pseudo-random Linux hostname.
// A valid hostname has a length between 1 and 64 characters, begins with letter and ends with an alphanumeric character, and may contain hyphens in the middle.
func LinuxHostname(r *rand.Rand, minLength, maxLength uint) (string, error) {
	length, err := checkLength(r, minLength, maxLength, minLinuxHostnameLength, maxLinuxHostnameLength)
	if err != nil {
		return "", err
	}

	// Define character sets for allowed hostname characters.
	allowedCharacters := append(lowerAlphanumericalRunes, '-') // Middle characters can include hyphens.

	// Create a rune slice to hold the generated hostname.

	hostname := make([]rune, length)

	// First character: must be alphanumeric.
	hostname[0] = lowerCaseRunes[r.IntN(len(lowerCaseRunes))]

	// Middle characters: can include hyphens.
	for i := 1; i < int(length)-1; i++ {
		hostname[i] = allowedCharacters[r.IntN(len(allowedCharacters))]
	}

	// Last character: must be alphanumeric (if the total length is greater than 1).
	if length > 1 {
		hostname[length-1] = lowerAlphanumericalRunes[r.IntN(len(lowerAlphanumericalRunes))]
	}

	// Convert the rune slice to a string and return the result.
	return string(hostname), nil
}
