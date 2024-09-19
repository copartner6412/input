package pseudorandom_test

import (
	"testing"

	"github.com/copartner6412/input/pseudorandom"
	"github.com/copartner6412/input/validate"
)

const (
	minPINLengthAllowed = 3
	maxPINLengthAllowed = 32
)

func FuzzPIN(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed1, seed2 uint64, min, max uint) {
		r1, r2, minLength, maxLength := randoms(seed1, seed2, min, max, minPINLengthAllowed, maxPINLengthAllowed)

		pin1, err := pseudorandom.PIN(r1, minLength, maxLength)
		if err != nil {
			t.Fatalf("error generating a pseudo-random PIN: %v", err)
		}

		err = validate.PIN(pin1, minLength, maxLength)
		if err != nil {
			t.Fatalf("expected no error for valid pseudo-random PIN \"%s\", but got error: %v", pin1, err)
		}

		pin2, err := pseudorandom.PIN(r2, minLength, maxLength)
		if err != nil {
			t.Fatalf("error regenerating the pseudo-random PIN: %v", err)
		}

		if pin1 != pin2 {
			t.Fatal("not deterministic")
		}
	})
}
