package random_test

import (
	"crypto/rand"
	"testing"

	"github.com/copartner6412/input/random"
	"github.com/copartner6412/input/validate"
)

func FuzzCountryName(f *testing.F) {
	f.Fuzz(func(t *testing.T, a int) {
		name, err := random.CountryName(rand.Reader)
		if err != nil {
			t.Fatalf("error generating a random country name: %v", err)
		}

		err = validate.CountryName(name)
		if err != nil {
			t.Fatalf("expected no error for valid random country name \"%s\": %v", name, err)
		}
	})
}

func FuzzCountryCode2(f *testing.F) {
	f.Fuzz(func(t *testing.T, a int) {
		code, err := random.CountryCode2(rand.Reader)
		if err != nil {
			t.Fatalf("error generating a random country code-2: %v", err)
		}

		err = validate.CountryCode2(code)
		if err != nil {
			t.Fatalf("expected no error for valid random country code-2 \"%s\": %v", code, err)
		}
	})
}

func FuzzCountryCode3(f *testing.F) {
	f.Fuzz(func(t *testing.T, a int) {
		code, err := random.CountryCode3(rand.Reader)
		if err != nil {
			t.Fatalf("error generating a random country code-3: %v", err)
		}

		err = validate.CountryCode3(code)
		if err != nil {
			t.Fatalf("unexpected error for valid random country code-3 \"%s\": %v", code, err)
		}
	})
}
