package pseudorandom_test

import (
	"testing"

	"github.com/copartner6412/input/pseudorandom"
	"github.com/copartner6412/input/validate"
)

const (
	minBitSizeAllowed uint = 1
	maxBitSizeAllowed uint = 4096
)

func FuzzBigInteger(f *testing.F) {
	f.Fuzz(func (t *testing.T, seed1, seed2 uint64, min, max uint)  {
		r1, r2, minBitSize, maxBitSize := randoms(seed1, seed2, min, max, minBitSizeAllowed, maxBitSizeAllowed)
		number1, err := pseudorandom.BigInteger(r1, minBitSize, maxBitSize)
		if err != nil {
			t.Fatalf("error generating a pseudo-random big integer: %v", err)
		}

		err = validate.BigInteger(number1, minBitSize, maxBitSize)
		if err != nil {
			t.Fatal(err)
		}

		number2, err := pseudorandom.BigInteger(r2, minBitSize, maxBitSize)
		if err != nil {
			t.Fatalf("error regenerating the pseudo-random big integer: %v", err)
		}

		if difference := number1.Cmp(number2); difference != 0 {
			t.Fatal("not deterministic")
		}
	})
}