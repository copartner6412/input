package random_test

import (
	"crypto/rand"
	"testing"

	"github.com/copartner6412/input/random"
	"github.com/copartner6412/input/validate"
)

func FuzzIPv6(f *testing.F) {
	f.Fuzz(func(t *testing.T, r int) {
		cidr6, err := random.CIDRv6(rand.Reader)
		if err != nil {
			t.Fatalf("error generating a random CIDR v6: %v", err)
		}

		ipv6, err := random.IPv6(rand.Reader, cidr6.String())
		if err != nil {
			t.Fatalf("error generating a random IPv6: %v", err)
		}

		err = validate.IP(ipv6.String(), cidr6.String())
		if err != nil {
			t.Fatalf("expected no error for the random IPv6 \"%s\" within \"%s\" network: %v", ipv6.String(), cidr6.String(), err)
		}
	})
}
