package pseudorandom

import (
	"math/rand/v2"
)

const (
	minByteSliceLengthAllowed uint = 1
	maxByteSliceLengthAllowed uint = 8192
	maxByteNumber             uint = 256
)

// Bytes generates a deterministic pseudo-random byte slice of a length between the specified minLength and maxLength, using the provided random source.
//
// Parameters:
//   - r: A pointer to a random number generator (rand.Rand) used for generating random values.
//   - minLength: The minimum length of the byte slice to generate. This must be between 1 and 8192.
//   - maxLength: The maximum length of the byte slice to generate. This must be between 1 and 8192.
//
// Returns:
//   - A pseudo-random byte slice with a length between minLength and maxLength, inclusive.
//   - An error if minLength is less than the allowed minimum, maxLength exceeds the allowed maximum, or if maxLength is less than minLength.
func Bytes(r *rand.Rand, minLength, maxLength uint) ([]byte, error) {
	length, err := checkLength(r, minLength, maxLength, minByteSliceLengthAllowed, maxByteSliceLengthAllowed)
	if err != nil {
		return nil, err
	}

	// Allocate a byte slice with the specified length.
	byteSlice := make([]byte, length)

	// Generate random bytes using the provided random source and fill the slice
	for i := 0; i < int(length); i++ {
		byteSlice[i] = byte(uint8(r.UintN(maxByteNumber))) // Generate a random byte (0-255)
	}

	return byteSlice, nil
}
