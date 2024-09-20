package pseudorandom_test

import (
	"math/rand/v2"
	"testing"
	"time"

	"github.com/copartner6412/input/pseudorandom"
	"github.com/copartner6412/input/validate"
)

func FuzzDuration(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed1, seed2 uint64, min, max uint) {
		r1 := rand.New(rand.NewPCG(seed1, seed2))
		minDuration := time.Duration(min)
		maxDuration := minDuration + time.Duration(max)
		duration1, err := pseudorandom.Duration(r1, minDuration, maxDuration)
		if err != nil {
			t.Fatalf("error generating a pseudo-random duration: %v", err)
		}
		err = validate.Duration(duration1, minDuration, maxDuration)
		if err != nil {
			t.Fatalf("expected no erro for valid pseudo-random duration %v between %v and %v, but got error: %v", duration1, minDuration, maxDuration, err)
		}
		r2 := rand.New(rand.NewPCG(seed1, seed2))
		duration2, err := pseudorandom.Duration(r2, minDuration, maxDuration)
		if err != nil {
			t.Fatalf("error regenerating the pseudo-random duration: %v", err)
		}

		if duration1 != duration2 {
			t.Fatal("not deterministic")
		}
	})
}
