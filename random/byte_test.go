package random_test

import (
	"testing"

	"github.com/copartner6412/input/random"
	"github.com/copartner6412/input/validate"
)

func FuzzBytes(f *testing.F) {
	f.Fuzz(func(t *testing.T, min, max uint) {
		minLength := min%(8192) + 1
		maxLength := minLength + max%(8192-minLength+1)
		bytes, err := random.Bytes(minLength, maxLength)
		if err != nil {
			t.Fatalf("error generating a random byte slice: %v", err)
		}
		err = validate.Bytes(bytes, minLength, maxLength)
		if err != nil {
			t.Fatalf("expected no error for valid random byte slice, but got error: %v", err)
		}
	})
}