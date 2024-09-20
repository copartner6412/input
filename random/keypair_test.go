package random_test

import (
	"testing"

	"github.com/copartner6412/input/random"
	"github.com/copartner6412/input/validate"
)

func FuzzKeyPair(f *testing.F) {
	f.Fuzz(func (t *testing.T, a uint)  {
		algorithm := random.Algorithm(int(a % 9))

		publicKey1, privateKey1, err := random.KeyPair(algorithm)
		if err != nil {
			t.Fatalf("error generating a pseudo-random key pair of type %s: %v", algorithm.String(), err)
		}

		err = validate.KeyPair(validate.Algorithm(algorithm), publicKey1, privateKey1)
		if err != nil {
			t.Fatalf("invalid key pair: %v", err)
		}
	})
}