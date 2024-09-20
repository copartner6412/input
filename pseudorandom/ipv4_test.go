package pseudorandom_test

import (
	"math/rand/v2"
	"testing"

	"github.com/copartner6412/input/pseudorandom"
	"github.com/copartner6412/input/validate"
)

func FuzzIPv4(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed1, seed2 uint64) {
		r1 := rand.New(rand.NewPCG(seed1, seed2))
		cidr41 := pseudorandom.CIDRv4(r1)
		ipv41, err := pseudorandom.IPv4(r1, cidr41.String())
		if err != nil {
			t.Fatalf("error generating a pseudo-random IPv4: %v", err)
		}

		err = validate.IP(ipv41.String(), cidr41.String())
		if err != nil {
			t.Fatalf("unexpected error for the pseudo-random IPv4 \"%s\" within \"%s\" network: %v", ipv41.String(), cidr41.String(), err)
		}

		r2 := rand.New(rand.NewPCG(seed1, seed2))
		cidr42 := pseudorandom.CIDRv4(r2)
		ipv42, err := pseudorandom.IPv4(r2, cidr42.String())

		if ipv41.String() != ipv42.String() {
			t.Fatal("not deterministic")
		}
	})
}
