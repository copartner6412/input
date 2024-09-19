package pseudorandom_test

import (
	"testing"

	"github.com/copartner6412/input/pseudorandom"
	"github.com/copartner6412/input/validate"
)

const (
	minLinuxHostnameLengthAllowed uint = 1
	maxLinuxHostnameLengthAllowed uint = 64
)

func FuzzLinuxHostname(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed1, seed2 uint64, min, max uint) {
		r1, r2, minLength, maxLength := randoms(seed1, seed2, min, max, minLinuxHostnameLengthAllowed, maxLinuxHostnameLengthAllowed)

		hostname1, err := pseudorandom.LinuxHostname(r1, minLength, maxLength)
		if err != nil {
			t.Fatalf("error generating a pseudo-random Linux hostname: %v", err)
		}

		err = validate.LinuxHostname(hostname1, minLength, maxLength)
		if err != nil {
			t.Fatalf("expected no error for the valid pseudo-random Linux hostname \"%s\", but got error: %v", hostname1, err)
		}

		hostname2, err := pseudorandom.LinuxHostname(r2, minLength, maxLength)
		if err != nil {
			t.Fatalf("error generating a pseudo-random Linux hostname: %v", err)
		}

		if hostname1 != hostname2 {
			t.Fatal("not deterministic")
		}
	})
}
