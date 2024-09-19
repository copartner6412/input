package pseudorandom_test
import (
	"math/rand/v2"
	"testing"

	"github.com/copartner6412/input/pseudorandom"
	"github.com/copartner6412/input/validate"
)

func FuzzIPv6(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed1, seed2 uint64) {
		r1 := rand.New(rand.NewPCG(seed1, seed2))
		cidr61 := pseudorandom.CIDRv6(r1)
		ipv61, err := pseudorandom.IPv6(r1, cidr61.String())
		if err != nil {
			t.Errorf("error generating a pseudo-random IPv6: %v", err)
		}

		err = validate.IP(ipv61.String(), cidr61.String())
		if err != nil {
			t.Errorf("expected no error for the pseudo-random IPv6 \"%s\" within \"%s\" network, but got error: %v", ipv61.String(), cidr61.String(), err)
		}

		r2 := rand.New(rand.NewPCG(seed1, seed2))
		cidr62 := pseudorandom.CIDRv6(r2)
		ipv62, err := pseudorandom.IPv6(r2, cidr62.String())
		if err != nil {
			t.Errorf("error generating a pseudo-random IPv6: %v", err)
		}
		
		if ipv61.String() != ipv62.String() {
			t.Error("not deterministic")
		}
	})
}