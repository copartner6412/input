package random_test

import (
	"crypto/rand"
	"testing"

	"github.com/copartner6412/input/random"
	"github.com/copartner6412/input/validate"
)

const (
	minSubdomainLengthAllowed uint = 1
	maxSubdomainLengthAllowed uint = 63
)

func FuzzSubdomain(f *testing.F) {
	f.Fuzz(func(t *testing.T, min, max uint) {
		minLength, maxLength := randoms(min, max, minSubdomainLengthAllowed, maxSubdomainLengthAllowed)
		subdomain, err := random.Subdomain(rand.Reader, minLength, maxLength)
		if err != nil {
			t.Fatalf("error generating a random subdomain: %v", err)
		}

		err = validate.Subdomain(subdomain, minLength, maxLength)
		if err != nil {
			t.Fatalf("expected no error for valid random subdomain \"%s\", but got error: %v", subdomain, err)
		}
	})
}
