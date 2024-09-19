package random_test

import (
	"testing"

	"github.com/copartner6412/input/random"
	"github.com/copartner6412/input/validate"
)

func FuzzTLD(f *testing.F) {
	f.Fuzz(func(t *testing.T, min, max uint) {
		minLength, maxLength := randoms(min, max, minTLDLengthAllowed, maxTLDLengthAllowed)
		tld1, err := random.TLD(minLength, maxLength)
		if err != nil {
			t.Fatalf("error generating a random gTLD: %v", err)
		}

		err = validate.TLD(tld1, minLength, maxLength)
		if err != nil {
			t.Fatalf("expected no error for valid random TLD \"%s\", but got error: %v", tld1, err)
		}
	})
}

func FuzzCCTLD(f *testing.F) {
	f.Fuzz(func(t *testing.T, r int) {
		tld, err := random.CCTLD()
		if err != nil {
			t.Fatalf("error generating a random ccTLD: %v", err)
		}
		err = validate.CCTLD(tld)

		if err != nil {
			t.Fatalf("expected no error for valid ccTLD \"%s\", but got error: %v", tld, err)
		}
	})
}