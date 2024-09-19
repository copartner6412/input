package pseudorandom_test

import (
	"testing"

	"github.com/copartner6412/input/pseudorandom"
	"github.com/copartner6412/input/validate"
)

const (
	minStringLengthAllowed = 1
	maxStringLengthAllowed = 8192
)

func FuzzString(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed1, seed2 uint64, min, max uint, justASCII bool) {
		r1, r2, minLength, maxLength := randoms(seed1, seed2, min, max, minStringLengthAllowed, maxStringLengthAllowed)

		str1, err := pseudorandom.String(r1, minLength, maxLength, justASCII)
		if err != nil {
			t.Fatalf("error generating a pseudo-random string: %v", err)
		}

		err = validate.String(str1, minLength, maxLength, justASCII)
		if err != nil {
			t.Fatalf("expected no error for valid pseudo-random string:\n%s\nbut got error: %v", str1, err)
		}

		str2, err := pseudorandom.String(r2, minLength, maxLength, justASCII)
		if err != nil {
			t.Fatalf("error generating a pseudo-random string: %v", err)
		}

		if str1 != str2 {
			t.Fatal("not deterministic")
		}
	})
}
