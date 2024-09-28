package random_test

import (
	"crypto/rand"
	"testing"

	"github.com/copartner6412/input/random"
	"github.com/copartner6412/input/validate"
)

const (
	minLinuxHostnameLengthAllowed uint = 1
	maxLinuxHostnameLengthAllowed uint = 64
)

func FuzzLinuxHostname(f *testing.F) {
	f.Fuzz(func(t *testing.T, min, max uint) {
		minLength, maxLength := randoms(min, max, minLinuxHostnameLengthAllowed, maxLinuxHostnameLengthAllowed)
		hostname, err := random.LinuxHostname(rand.Reader, minLength, maxLength)
		if err != nil {
			t.Fatalf("error generating a random Linux hostname: %v", err)
		}

		err = validate.LinuxHostname(hostname, minLength, maxLength)
		if err != nil {
			t.Fatalf("expected no error for valid random Linux hostname \"%s\": %v", hostname, err)
		}
	})
}
