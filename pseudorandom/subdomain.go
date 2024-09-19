package pseudorandom

import "math/rand/v2"

const (
	minSubdomainLengthAllowed uint = 1
	maxSubdomainLengthAllowed uint = 63
)

// Subdomain generates a deterministic pseudo-random subdomain string with a length between 1 and 63.
// The first character must be a letter, the middle characters can include letters, digits, or hyphens, and the last character must be a letter or digit.
func Subdomain(r *rand.Rand, minLength, maxLength uint) (string, error) {
	length, err := checkLength(r, minLength, maxLength, minSubdomainLengthAllowed, maxSubdomainCount)
	if err != nil {
		return "", err
	}

	// Allowed characters: lowercase letters, digits, and hyphens (for middle characters).
	allowedCharacters := append(lowerAlphanumericalRunes, '-')

	// Preallocate the slice with the required subdomain length.
	subdomain := make([]rune, 0, length)

	// First character: must be a lowercase letter.
	subdomain = append(subdomain, lowerAlphanumericalRunes[r.IntN(len(lowerAlphanumericalRunes))])

	// Middle characters: can include lowercase letters, digits, or hyphens.
	for i := 1; i < int(length)-1; i++ {
		subdomain = append(subdomain, allowedCharacters[r.IntN(len(allowedCharacters))])
	}

	// Last character: must be a letter or digit (if subdomainLength > 1).
	if length > 1 {
		subdomain = append(subdomain, lowerAlphanumericalRunes[r.IntN(len(lowerAlphanumericalRunes))])
	}

	return string(subdomain), nil
}
