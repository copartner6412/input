package random_test

import (
	"crypto/rand"
	"testing"

	"github.com/copartner6412/input/random"
	"github.com/copartner6412/input/validate"
)

const (
	minPINLength = 3
	maxPINLength = 32
)

func FuzzPIN(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed1, seed2 uint64, min, max uint) {
		minLength := min%(maxPINLength-minPINLength+1) + minPINLength
		maxLength := max%(maxPINLength-minLength+1) + minLength
		pin, err := random.PIN(rand.Reader, minLength, maxLength)
		if err != nil {
			t.Fatalf("error generating a random PIN: %v", err)
		}
		err = validate.PIN(pin, minLength, maxLength)
		if err != nil {
			t.Fatalf("expected no error for valid random PIN \"%s\", but got error: %v", pin, err)
		}
	})
}
