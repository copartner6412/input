package pseudorandom_test

import (
	"testing"

	"github.com/copartner6412/input/pseudorandom"
	"github.com/copartner6412/input/validate"
)

const (
	minSubdomainLengthAllowed uint = 1
	maxSubdomainLengthAllowed uint = 63
)

func FuzzSubdomain(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed1, seed2 uint64, min, max uint) {
		r1, r2, minLength, maxLength := randoms(seed1, seed2, min, max, minSubdomainLengthAllowed, maxSubdomainLengthAllowed)
		subdomain1, err := pseudorandom.Subdomain(r1, minLength, maxLength)
		if err != nil {
			t.Fatalf("error generating a pseudo-random subdomain: %v", err)
		}

		err = validate.Subdomain(subdomain1, minLength, maxLength)
		if err != nil {
			t.Fatalf("expected no error for valid pseudo-random subdomain \"%s\", but got error: %v", subdomain1, err)
		}

		subdomain2, err := pseudorandom.Subdomain(r2, minLength, maxLength)
		if err != nil {
			t.Fatalf("error regenerating the pseudo-random subdomain: %v", err)
		}

		if subdomain1 != subdomain2 {
			t.Fatal("not deterministic")
		}
	})
}
