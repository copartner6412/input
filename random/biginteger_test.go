package random_test

import (
	"testing"

	"github.com/copartner6412/input/random"
)

const (
	minBitSizeAllowed uint = 1
	maxBitSizeAllowed uint = 4096
)

func FuzzBigInteger(f *testing.F) {
	f.Fuzz(func (t *testing.T, min, max uint)  {
		minBitSize, maxBitSize := randoms(min, max, minBitSizeAllowed, maxBitSizeAllowed)
		_, err := random.BigInteger(minBitSize, maxBitSize)
		if err != nil {
			t.Fatalf("error generating a cryptographically-secure random big integer: %v", err)
		}
	})
}