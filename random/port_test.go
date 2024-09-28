package random_test

import (
	"crypto/rand"
	"testing"

	"github.com/copartner6412/input/random"
	"github.com/copartner6412/input/validate"
)

func FuzzPortWellKnown(f *testing.F) {
	f.Fuzz(func(t *testing.T, a int) {
		port, err := random.PortWellKnown(rand.Reader)
		if err != nil {
			t.Fatalf("error generating a valid random well-know port: %v", err)
		}
		err = validate.PortWellKnown(port)
		if err != nil {
			t.Fatalf("expected no error for valid random well-known port %d, but got error: %v", port, err)
		}
	})
}

func FuzzPortNotWellKnown(f *testing.F) {
	f.Fuzz(func(t *testing.T, a int) {
		port, err := random.PortNotWellKnown(rand.Reader)
		if err != nil {
			t.Fatalf("error generating a valid random not-well-know port: %v", err)
		}
		err = validate.PortNotWellKnown(port)
		if err != nil {
			t.Fatalf("expected no error for valid random not-well-known port %d, but got error: %v", port, err)
		}
	})
}

func FuzzPortRegistered(f *testing.F) {
	f.Fuzz(func(t *testing.T, a int) {
		port, err := random.PortRegistered(rand.Reader)
		if err != nil {
			t.Fatalf("error generating a valid random registered port: %v", err)
		}
		err = validate.PortRegistered(port)
		if err != nil {
			t.Fatalf("expected no error for valid random registered port %d, but got error: %v", port, err)
		}
	})
}

func FuzzPortPrivate(f *testing.F) {
	f.Fuzz(func(t *testing.T, a int) {
		port, err := random.PortPrivate(rand.Reader)
		if err != nil {
			t.Fatalf("error generating a valid random private port: %v", err)
		}
		err = validate.PortPrivate(port)
		if err != nil {
			t.Fatalf("expected no error for valid random private port %d, but got error: %v", port, err)
		}
	})
}
