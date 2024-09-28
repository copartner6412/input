package random_test

import (
	"crypto/rand"
	"testing"

	"github.com/copartner6412/input/random"
	"github.com/copartner6412/input/validate"
)

const (
	minDomainLengthAllowed               uint = 1
	maxDomainLengthAllowed               uint = 253
	minTLDLengthAllowed                  uint = 2
	maxTLDLengthAllowed                  uint = 16
	ccTLDLength                          uint = 2
	minDomainWithValidTLDLengthAllowed   uint = minTLDLengthAllowed + 2
	minDomainWithValidCCTLDLengthAllowed uint = ccTLDLength + 2
)

func FuzzDomainFunc(f *testing.F) {
	f.Fuzz(func(t *testing.T, min, max uint) {
		minLength, maxLength := randoms(min, max, minDomainLengthAllowed, maxDomainLengthAllowed)
		domain, err := random.Domain(rand.Reader, minLength, maxLength)
		if err != nil {
			t.Fatalf("error generating a random domain: %v", err)
		}
		err = validate.Domain(domain, minLength, maxLength)
		if err != nil {
			t.Fatalf("expected no error for valid random domain \"%s\": %v", domain, err)
		}
	})
}

func FuzzDomainWithValidTLD(f *testing.F) {
	f.Fuzz(func(t *testing.T, min, max uint) {
		minLength, maxLength := randoms(min, max, minDomainWithValidTLDLengthAllowed, maxDomainLengthAllowed)
		domain, err := random.DomainWithValidTLD(rand.Reader, minLength, maxLength)
		if err != nil {
			t.Fatalf("error generating a random domain with valid TLD: %v", err)
		}
		err = validate.DomainWithValidTLD(domain, minLength, maxLength)
		if err != nil {
			t.Fatalf("expected no error for valid random domain \"%s\", but got err: %v", domain, err)
		}
	})
}

func FuzzDomainWithValidCCTLD(f *testing.F) {
	f.Fuzz(func(t *testing.T, min, max uint) {
		minLength, maxLength := randoms(min, max, minDomainWithValidCCTLDLengthAllowed, maxDomainLengthAllowed)
		domain, err := random.DomainWithValidCCTLD(rand.Reader, minLength, maxLength)
		if err != nil {
			t.Fatalf("error generating a random domain with valid ccTLD: %v", err)
		}
		err = validate.DomainWithValidCCTLD(domain, minLength, maxLength)
		if err != nil {
			t.Fatalf("expected no error for random domain with valid ccTLD \"%s\", but got err: %v", domain, err)
		}
	})
}
