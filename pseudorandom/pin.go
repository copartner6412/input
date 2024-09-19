package pseudorandom

import (
	"math/rand/v2"
	"strconv"
)

// Constants for PIN length limits.
const (
	minPINLengthAllowed uint = 3
	maxPINLengthAllowed uint = 32
)

// PIN generates a deterministic pseudo-random random numeric Personal Identification Number (PIN) of a length between the provided minLength and maxLength.
//
// Parameters:
//   - r: A random number generator of type *rand.Rand.
//   - minLength: The minimum length of the PIN to generate. It must be between 3 and 32.
//   - maxLength: The maximum length of the PIN to generate. It must be between 3 and 32.
//
// Returns:
//   - string: A pseudo-randomly generated numeric PIN of the chosen length.
//   - error: Returns an error if the minLength is less than minPINLengthAllowed, if maxLength exceeds maxPINLengthAllowed,
//     or if maxLength is less than minLength.
func PIN(r *rand.Rand, minLength, maxLength uint) (string, error) {
	length, err := checkLength(r, minLength, maxLength, minPINLengthAllowed, maxPINLengthAllowed)
	if err != nil {
		return "", err
	}

	// Generate a random PIN.
	var pin string
	for i := 0; i < int(length); i++ {
		// Generate a random digit (0-9)
		pin += strconv.Itoa(r.IntN(10))
	}

	return pin, nil
}
