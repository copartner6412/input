package pseudorandom_test

import (
	"math/rand/v2"
	"testing"

	"github.com/copartner6412/input/pseudorandom"
)

func FuzzCIDRv6(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed1, seed2 uint64) {
		r1 := rand.New(rand.NewPCG(seed1, seed2))
		cidrv61 := pseudorandom.CIDRv6(r1)

		r2 := rand.New(rand.NewPCG(seed1, seed2))
		cidrv62 := pseudorandom.CIDRv6(r2)

		if cidrv61.String() != cidrv62.String() {
			t.Fatal("not deterministic")
		}
	})
}
