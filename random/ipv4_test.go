package random_test

import (
	"crypto/rand"
	"testing"

	"github.com/copartner6412/input/random"
	"github.com/copartner6412/input/validate"
)

func FuzzIPv4(f *testing.F) {
	f.Fuzz(func(t *testing.T, r int) {
		cidr4, err := random.CIDRv4(rand.Reader)
		if err != nil {
			t.Fatalf("error generating a random CIDR v4: %v", err)
		}

		ipv4, err := random.IPv4(rand.Reader, cidr4.String())
		if err != nil {
			t.Fatalf("error generating a random IPv4: %v", err)
		}

		err = validate.IP(ipv4.String(), cidr4.String())
		if err != nil {
			t.Fatalf("unexpected error for the pseudo-random IPv4 \"%s\" within \"%s\" network: %v", ipv4.String(), cidr4.String(), err)
		}
	})
}
