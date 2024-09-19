package pseudorandom_test

import (
	"math/rand/v2"
	"testing"

	"github.com/copartner6412/input/pseudorandom"
	"github.com/copartner6412/input/validate"
)

func FuzzCountryName(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed1, seed2 uint64) {
		r1 := rand.New(rand.NewPCG(seed1, seed2))
		name1 := pseudorandom.CountryName(r1)
		err := validate.CountryName(name1)
		if err != nil {
			t.Fatalf("expected no error for valid pseudo-random country name \"%s\", but got error: %v", name1, err)
		}
		r2 := rand.New(rand.NewPCG(seed1, seed2))
		name2 := pseudorandom.CountryName(r2)
		if name1 != name2 {
			t.Fatalf("not deterministic")
		}
	})
}

func FuzzCountryCode2(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed1, seed2 uint64) {
		r1 := rand.New(rand.NewPCG(seed1, seed2))
		code21 := pseudorandom.CountryCode2(r1)
		err := validate.CountryCode2(code21)
		if err != nil {
			t.Fatalf("expected no error for valid pseudo-random country code-2 \"%s\", but got error: %v", code21, err)
		}
		r2 := rand.New(rand.NewPCG(seed1, seed2))
		code22 := pseudorandom.CountryCode2(r2)
		if code21 != code22 {
			t.Fatalf("not deterministic")
		}
	})
}

func FuzzCountryCode3(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed1, seed2 uint64) {
		r1 := rand.New(rand.NewPCG(seed1, seed2))
		code31 := pseudorandom.CountryCode3(r1)
		err := validate.CountryCode3(code31)
		if err != nil {
			t.Fatalf("expected no error for valid pseudo-random country code-3 \"%s\", but got error: %v", code31, err)
		}
		r2 := rand.New(rand.NewPCG(seed1, seed2))
		code32 := pseudorandom.CountryCode3(r2)
		if code31 != code32 {
			t.Fatal("not deterministic")
		}
	})
}
