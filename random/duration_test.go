package random_test

import (
	"crypto/rand"
	"testing"
	"time"

	"github.com/copartner6412/input/random"
	"github.com/copartner6412/input/validate"
)

func FuzzDuration(f *testing.F) {
	f.Fuzz(func(t *testing.T, min, max uint) {
		minDuration := time.Duration(min)
		maxDuration := minDuration + time.Duration(max)
		duration, err := random.Duration(rand.Reader, minDuration, maxDuration)
		if err != nil {
			t.Fatalf("error generating a random duration: %v", err)
		}
		err = validate.Duration(duration, minDuration, maxDuration)
		if err != nil {
			t.Fatalf("expected no erro for valid random duration %v between %v and %v, but got error: %v", duration, minDuration, maxDuration, err)
		}
	})
}
