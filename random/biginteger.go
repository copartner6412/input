package random

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const (
	minBitSizeAllowed uint = 1
	maxBitSizeAllowed uint = 4096
)

func BigInteger(minBitSize, maxBitSize uint) (*big.Int, error) {
	bitSize, err := checkLength(minBitSize, maxBitSize, minBitSizeAllowed, maxBitSizeAllowed)
	if err != nil {
		return nil, err
	}

	// Generate a random big.Int with the chosen bit size
	number, err := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), uint(bitSize)))
	if err != nil {
		return nil, fmt.Errorf("failed to generate random serial number: %w", err)
	}

	// Ensure the number is positive and has at least minBitSize bits
	number.SetBit(number, int(minBitSize-1), 1)

	return number, nil
}