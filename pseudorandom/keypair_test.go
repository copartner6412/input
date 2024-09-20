package pseudorandom_test

import (
	"crypto/ed25519"
	"math/rand/v2"
	"testing"

	"github.com/copartner6412/input/pseudorandom"
	"github.com/copartner6412/input/validate"
)

func FuzzKeyPair(f *testing.F) {
	f.Fuzz(func (t *testing.T, seed1, seed2 uint64)  {
		t.Parallel()
		r1 := rand.New(rand.NewPCG(seed1, seed2))
		reader1 := pseudorandom.New(r1)
		publicKey1, privateKey1, err := ed25519.GenerateKey(reader1)
		if err != nil {
			t.Fatalf("error generating a pseudo-random ED25519 key pair: %v", err)
		}

		err = validate.KeyPair(validate.AlgorithmED25519, publicKey1, privateKey1)
		if err != nil {
			t.Fatalf("invalid key pair: %v", err)
		}

		r2 := rand.New(rand.NewPCG(seed1, seed2))
		reader2 := pseudorandom.New(r2)

		publicKey2, privateKey2, err := ed25519.GenerateKey(reader2)
		if err != nil {
			t.Fatalf("error regenerating the pseudo-random key pair: %v", err)
		}

		if !publicKey1.Equal(publicKey2) || !privateKey1.Equal(privateKey2) {
			t.Fatal("not deterministic")
		}
	})
}