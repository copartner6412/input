package validate_test

import (
	"math/big"
	"testing"

	"github.com/copartner6412/input/validate"
)

const (
	minBitSizeAllowed uint = 1
	maxBitSizeAllowed uint = 4096
)

func TestBigIntegerSuccessfulForValidInput(t *testing.T) {
	testCases := map[string]struct{
		number     *big.Int
		minBitSize uint
		maxBitSize uint
	}{
		"minBitSizeAtBoundary": {
			number:     big.NewInt(1), // Single bit number
			minBitSize: minBitSizeAllowed,
			maxBitSize: 128,
		},
		"maxBitSizeAtBoundary": {
			number:     new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 4096), big.NewInt(1)), // 4096 bits
			minBitSize: 1024,
			maxBitSize: maxBitSizeAllowed,
		},
		"minBitSizeEqualsMaxBitSize": {
			number:     big.NewInt(123456),
			minBitSize: 17,
			maxBitSize: 17,
		},
		"bitLengthInRange": {
			number:     new(big.Int).Lsh(big.NewInt(1), 256), // 256 bits
			minBitSize: 128,
			maxBitSize: 512,
		},
		"exactlyMaxBitSize": {
			number:     new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), maxBitSizeAllowed), big.NewInt(1)), // max bit size of 4096
			minBitSize: 2048,
			maxBitSize: maxBitSizeAllowed,
		},
		"bitLengthAtMinRequired": {
			number:     new(big.Int).SetBit(big.NewInt(0), 1, 1), // 1-bit number
			minBitSize: minBitSizeAllowed,
			maxBitSize: 64,
		},
		"bitLengthEqualToMaxBitSize": {
			number:     new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 2048), big.NewInt(1)), // 2048 bits
			minBitSize: 1024,
			maxBitSize: 2048,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			err := validate.BigInteger(tc.number, tc.minBitSize, tc.maxBitSize)
			if err != nil {
				t.Fatalf("expected no error for valid big integer number %v, but got error: %v", tc.number, err)
			}
		})
	}
}


func TestBigIntegerFailsForInvalidInput(t *testing.T) {
	testCases := map[string]struct{
		number     *big.Int
		minBitSize uint
		maxBitSize uint
	}{
		"maxBitSizeLessThanMinBitSize": {
			number:     big.NewInt(123456),
			minBitSize: 128,
			maxBitSize: 64, // maxBitSize is less than minBitSize
		},
		"minBitSizeLessThanMinAllowedBitSize": {
			number:     big.NewInt(123456),
			minBitSize: 0,  // Less than minBitSizeAllowed (1)
			maxBitSize: 128,
		},
		"maxBitSizeExceedsMaxAllowedBitSize": {
			number:     big.NewInt(123456),
			minBitSize: 128,
			maxBitSize: 5000,  // Exceeds maxBitSizeAllowed (4096)
		},
		"numberBitLengthExceedsMaxAllowed": {
			number:     new(big.Int).Exp(big.NewInt(2), big.NewInt(5000), nil), // Bit length exceeds maxBitSizeAllowed
			minBitSize: minBitSizeAllowed,
			maxBitSize: maxBitSizeAllowed,
		},
		"numberBitLengthLessThanMinRequired": {
			number:     big.NewInt(3), // Bit length less than required minBitSize
			minBitSize: 128,
			maxBitSize: 256,
		},
		"negativeBigIntValue": {
			number:     big.NewInt(-1000), // Negative value
			minBitSize: 128,
			maxBitSize: 256,
		},
		"zeroBigIntWithMinBitSizeGreaterThanZero": {
			number:     big.NewInt(0),  // Big integer is zero
			minBitSize: minBitSizeAllowed,              // minBitSize greater than zero
			maxBitSize: 128,
		},
		"negativeMinBitSize": {
			number:     big.NewInt(1000),
			minBitSize: ^uint(0),  // Effectively a negative value due to uint overflow
			maxBitSize: 128,
		},
		"negativeMaxBitSize": {
			number:     big.NewInt(1000),
			minBitSize: minBitSizeAllowed,
			maxBitSize: ^uint(0),  // Effectively a negative value due to uint overflow
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			err := validate.BigInteger(tc.number, tc.minBitSize, tc.maxBitSize)
			if err == nil {
				t.Fatalf("expected error for invalid big integer number %v, but got no error", tc.number)
			}
		})
	}
}
