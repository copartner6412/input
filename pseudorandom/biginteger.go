package pseudorandom

import (
	"math/big"
	"math/rand/v2"
)

const (
	minBitSizeAllowed uint = 1
	maxBitSizeAllowed uint = 4096
)

func BigInteger(r *rand.Rand, minBitSize, maxBitSize uint) (*big.Int, error) {
	bitSize, err := checkLength(r, minBitSize, maxBitSize, minBitSizeAllowed, maxBitSizeAllowed)
	if err != nil {
		return nil, err
	}

	max := new(big.Int).Lsh(big.NewInt(1), uint(bitSize))
	number := new(big.Int)

	// Calculate how many uint64 values we need to generate
	numUint64s := (bitSize + 63) / 64

	// Generate random uint64 values and combine them into a big.Int
	for i := uint(0); i < numUint64s; i++ {
		randomUint64 := r.Uint64()
		part := new(big.Int).SetUint64(randomUint64)
		part.Lsh(part, 64*i)
		number.Or(number, part)
	}

	// Ensure the number is within the correct range
	number.Mod(number, max)

	// Ensure the number has at least minBitSize bits
	number.SetBit(number, int(minBitSize-1), 1)

	return number, nil
}