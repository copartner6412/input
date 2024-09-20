package pseudorandom_test

import (
	"bytes"
	"testing"

	"github.com/copartner6412/input/pseudorandom"
	"github.com/copartner6412/input/validate"
)

const (
	minByteSliceLengthAllowed = 1
	maxByteSliceLengthAllowed = 8192
)

func FuzzBytes(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed1, seed2 uint64, min, max uint) {
		r1, r2, minLength, maxLength := randoms(seed1, seed2, min, max, minByteSliceLengthAllowed, maxByteSliceLengthAllowed)

		bytes1, err := pseudorandom.Bytes(r1, minLength, maxLength)
		if err != nil {
			t.Fatalf("error generating a pseudo-random byte slice: %v", err)
		}

		err = validate.Bytes(bytes1, minLength, maxLength)
		if err != nil {
			t.Fatalf("expected no error for valid pseudo-random byte slice, but got error: %v", err)
		}

		bytes2, err := pseudorandom.Bytes(r2, minLength, maxLength)
		if err != nil {
			t.Fatalf("error regenerating the pseudo-random byte slice: %v", err)
		}

		if !bytes.Equal(bytes1, bytes2) {
			t.Fatal("not deterministic")
		}
	})
}
