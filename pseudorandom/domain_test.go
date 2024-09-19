package pseudorandom_test

import (
	"testing"

	"github.com/copartner6412/input/pseudorandom"
	"github.com/copartner6412/input/validate"
)

const (
	minDomainLengthAllowed   uint = 1
	maxDomainLengthAllowed   uint = 253
	minTLDLengthAllowed uint = 2
	maxTLDLengthAllowed uint = 2
	ccTLDLength uint = 2
	minDomainWithValidTLDLengthAllowed uint = minTLDLengthAllowed + 2
	minDomainWithValidCCTLDLengthAllowed uint = ccTLDLength + 2
)

func FuzzDomainFunc(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed1, seed2 uint64, min, max uint) {
		r1, r2, minLength, maxLength := randoms(seed1, seed2, min, max, minDomainLengthAllowed, maxDomainLengthAllowed)
		
		domain1, err := pseudorandom.Domain(r1, minLength, maxLength)
		if err != nil {
			t.Fatalf("error generating a pseudo-random domain: %v", err)
		}

		err = validate.Domain(domain1, minLength, maxLength)
		if err != nil {
			t.Fatalf("expected no error for valid pseudo-random domain \"%s\", but got error: %v", domain1, err)
		}

		domain2, err := pseudorandom.Domain(r2, minLength, maxLength)
		if err != nil {
			t.Fatalf("error regenerating the pseudo-random domain: %v", err)
		}

		if domain1 != domain2 {
			t.Fatalf("not deterministic")
		}
	})
}


func FuzzDomainWithValidTLD(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed1, seed2 uint64, min, max uint) {
		r1, r2, minLength, maxLength := randoms(seed1, seed2, min, max, minDomainWithValidTLDLengthAllowed, maxDomainLengthAllowed)

		domain1, err := pseudorandom.DomainWithValidTLD(r1, minLength, maxLength)
		if err != nil {
			t.Fatalf("error generating a pseudo-random domain with valid TLD: %v", err)
		}

		err = validate.DomainWithValidTLD(domain1, minLength, maxLength)
		if err != nil {
			t.Fatalf("expected no error for valid pseudo-random domain \"%s\", but got error: %v", domain1, err)
		}

		domain2, err := pseudorandom.DomainWithValidTLD(r2, minLength, maxLength)
		if err != nil {
			t.Fatalf("error regenerating the pseudo-random domain with valid TLD: %v", err)
		}

		if domain1 != domain2 {
			t.Fatal("not deterministic")
		}
	})
}

func FuzzDomainWithValidCCTLD(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed1, seed2 uint64, min, max uint) {
		r1, r2, minLength, maxLength := randoms(seed1, seed2, min, max, minDomainWithValidCCTLDLengthAllowed, maxDomainLengthAllowed)

		domain1, err := pseudorandom.DomainWithValidCCTLD(r1, minLength, maxLength)
		if err != nil {
			t.Fatalf("error generating a pseudo-random domain with valid ccTLD: %v", err)
		}

		err = validate.DomainWithValidCCTLD(domain1, minLength, maxLength)
		if err != nil {
			t.Fatalf("expected no error for valid pseudo-random domain \"%s\", but got error: %v", domain1, err)
		}

		domain2, err := pseudorandom.DomainWithValidCCTLD(r2, minLength, maxLength)
		if err != nil {
			t.Fatalf("error regenerating the pseudo-random domain with valid ccTLD: %v", err)
		}

		if domain1 != domain2 {
			t.Fatal("not deterministic")
		}
	})
}