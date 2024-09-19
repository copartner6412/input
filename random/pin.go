package random

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"strconv"
)

// Constants for PIN length limits.
const (
	minPINLength uint = 3
	maxPINLength uint = 32
)

// PIN generates a cryptographically-secure random numeric Personal Identification Number (PIN) of a length between the provided minLength and maxLength.
//
// Parameters:
//   - minLength: The minimum length of the PIN to generate. It must be between 3 and 32.
//   - maxLength: The maximum length of the PIN to generate. It must be between 3 and 32.
//
// Returns:
//   - string: A randomly generated numeric PIN of the chosen length.
//   - error: Returns an error if the minLength is less than minPINLength, if maxLength exceeds maxPINLength,
//     or if maxLength is less than minLength.
func PIN(minLength, maxLength uint) (string, error) {
	// Ensure that maxLength is not less than minLength.
	if maxLength < minLength {
		return "", errors.New("maximum length can not be less than minimum length")
	}

	// Validate that the length requirements fall within acceptable system bounds.
	if minLength < minPINLength {
		return "", fmt.Errorf("minimum PIN length must not be less than %d characters", minPINLength)
	}

	if maxLength > maxPINLength {
		return "", fmt.Errorf("maximum PIN length must not exceed %d characters", maxPINLength)
	}

    random1, err := rand.Int(rand.Reader, big.NewInt(int64(maxLength-minLength+1)))
    if err != nil {
        return "", fmt.Errorf("error generating a random number for calculating PIN length: %w", err)
    }
	length := uint(random1.Int64()) + minLength

	// Generate a random PIN.
	var pin string
	for i := 0; i < int(length); i++ {
		// Generate a random digit (0-9)
        random2, err := rand.Int(rand.Reader, big.NewInt(int64(10)))
        if err != nil {
            return "", fmt.Errorf("error generating a random digit for PIN: %W", err)
        }
		pin += strconv.Itoa(int(random2.Int64()))
	}

	return pin, nil
}