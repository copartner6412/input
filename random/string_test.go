package random_test

import (
	"crypto/rand"
	"testing"

	"github.com/copartner6412/input/random"
	"github.com/copartner6412/input/validate"
)

func FuzzString(f *testing.F) {
	f.Fuzz(func(t *testing.T, min, max uint, justASCII bool) {
		minLength := min%(8192) + 1
		maxLength := minLength + max%(8192-minLength+1)
		str, err := random.String(rand.Reader, minLength, maxLength, justASCII)
		if err != nil {
			t.Fatalf("error generating a pseudo-random string: %v", err)
		}
		err = validate.String(str, minLength, maxLength, justASCII)
		if err != nil {
			t.Fatalf("expected no error for valid pseudo-random string:\n%s\nbut got error: %v", str, err)
		}
	})
}
