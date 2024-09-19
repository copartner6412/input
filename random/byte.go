package random

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const (
	minByteSliceLengthAllowed uint = 1
	maxByteSliceLengthAllowed uint = 8192
	maxByteNumber uint = 256
)

// Bytes generates a random byte slice of a length between the specified minLength and maxLength.
//
// Parameters:
//   - minLength: The minimum length of the byte slice to generate. This must be between 1 and 8192.
//   - maxLength: The maximum length of the byte slice to generate. This must be between 1 and 8192.
//
// Returns:
//   - A random byte slice with a length between minLength and maxLength, inclusive.
//   - An error if minLength is less than the allowed minimum, maxLength exceeds the allowed maximum, or if maxLength is less than minLength.
func Bytes(minLength, maxLength uint) ([]byte, error) {
	length, err := checkLength(minLength, maxLength, minByteSliceLengthAllowed, maxByteSliceLengthAllowed)
	if err != nil {
		return nil, err
	}

	// Allocate a byte slice with the specified length.
	byteSlice := make([]byte, length)

	// Generate random bytes using the provided rand source and fill the slice
	for i := 0; i < int(length); i++ {
		random2, err := rand.Int(rand.Reader, big.NewInt(256))
		if err != nil {
			return nil, fmt.Errorf("error generating a random number for byte value: %v", err)
		}
		byteSlice[i] = byte(uint8(random2.Int64())) // Generate a random byte (0-255)
	}

	return byteSlice, nil
}
