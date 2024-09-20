package pseudorandom_test

import (
	"math/rand/v2"
	"testing"

	"github.com/copartner6412/input/pseudorandom"
	"github.com/copartner6412/input/validate"
)

func FuzzPortWellKnown(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed1, seed2 uint64) {
		r1 := rand.New(rand.NewPCG(seed1, seed2))
		port1 := pseudorandom.PortWellKnown(r1)
		err := validate.PortWellKnown(port1)
		if err != nil {
			t.Errorf("expected no error for valid pseudo-random well-known port %d, but got error: %v", port1, err)
		}
		r2 := rand.New(rand.NewPCG(seed1, seed2))
		port2 := pseudorandom.PortWellKnown(r2)
		if port1 != port2 {
			t.Errorf("not deterministic, expected: %d, got: %d", port1, port2)
		}
	})
}

func FuzzPortNotWellKnown(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed1, seed2 uint64) {
		r1 := rand.New(rand.NewPCG(seed1, seed2))
		port1 := pseudorandom.PortNotWellKnown(r1)
		err := validate.PortNotWellKnown(port1)
		if err != nil {
			t.Errorf("expected no error for valid pseudo-random not-well-known port %d, but got error: %v", port1, err)
		}
		r2 := rand.New(rand.NewPCG(seed1, seed2))
		port2 := pseudorandom.PortNotWellKnown(r2)
		if port1 != port2 {
			t.Errorf("not deterministic, expected: %d, got: %d", port1, port2)
		}
	})
}

func FuzzPortRegistered(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed1, seed2 uint64) {
		r1 := rand.New(rand.NewPCG(seed1, seed2))
		port1 := pseudorandom.PortRegistered(r1)
		err := validate.PortRegistered(port1)
		if err != nil {
			t.Errorf("expected no error for valid pseudo-random registered port %d, but got error: %v", port1, err)
		}
		r2 := rand.New(rand.NewPCG(seed1, seed2))
		port2 := pseudorandom.PortRegistered(r2)
		if port1 != port2 {
			t.Errorf("not deterministic, expected: %d, got: %d", port1, port2)
		}
	})
}

func FuzzPortPrivate(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed1, seed2 uint64) {
		r1 := rand.New(rand.NewPCG(seed1, seed2))
		port1 := pseudorandom.PortPrivate(r1)
		err := validate.PortPrivate(port1)
		if err != nil {
			t.Fatalf("expected no error for valid pseudo-random private port %d, but got error: %v", port1, err)
		}
		r2 := rand.New(rand.NewPCG(seed1, seed2))
		port2 := pseudorandom.PortPrivate(r2)
		if port1 != port2 {
			t.Fatalf("not deterministic")
		}
	})
}
