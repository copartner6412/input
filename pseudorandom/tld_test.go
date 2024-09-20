package pseudorandom_test

import (
	"math/rand/v2"
	"testing"

	"github.com/copartner6412/input/pseudorandom"
	"github.com/copartner6412/input/validate"
)

func FuzzTLD(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed1, seed2 uint64, min, max uint) {
		r1, r2, minLength, maxLength := randoms(seed1, seed2, min, max, minTLDLengthAllowed, maxTLDLengthAllowed)
		tld1, err := pseudorandom.TLD(r1, minLength, maxLength)
		if err != nil {
			t.Fatalf("error generating a pseudo-random gTLD: %v", err)
		}

		err = validate.TLD(tld1, minLength, maxLength)
		if err != nil {
			t.Fatalf("expected no error for valid pseudo-random TLD \"%s\", but got error: %v", tld1, err)
		}

		tld2, err := pseudorandom.TLD(r2, minLength, maxLength)
		if err != nil {
			t.Fatalf("error regenerating the pseudo-random gTLD: %v", err)
		}

		if tld1 != tld2 {
			t.Fatal("not deterministic")
		}
	})
}

func FuzzCCTLD(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed1, seed2 uint64) {
		r1 := rand.New(rand.NewPCG(seed1, seed2))
		tld1 := pseudorandom.CCTLD(r1)
		err := validate.CCTLD(tld1)
		if err != nil {
			t.Fatalf("expected no error for valid pseudo-random ccTLD \"%s\", but got error: %v", tld1, err)
		}
		r2 := rand.New(rand.NewPCG(seed1, seed2))
		tld2 := pseudorandom.CCTLD(r2)
		if tld1 != tld2 {
			t.Fatal("not deterministic")
		}
	})
}
