package pseudorandom_test

import (
	"math/rand/v2"
	"testing"

	"github.com/copartner6412/input/pseudorandom"
)

func FuzzCIDRv4(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed1, seed2 uint64) {
		r1 := rand.New(rand.NewPCG(seed1, seed2))
		cidrv41 := pseudorandom.CIDRv4(r1)

		r2 := rand.New(rand.NewPCG(seed1, seed2))
		cidrv42 := pseudorandom.CIDRv4(r2)

		if cidrv41.String() != cidrv42.String() {
			t.Fatal("not deterministic")
		}
	})
}
