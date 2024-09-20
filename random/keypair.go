package random
/*
import (
	"crypto"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"errors"
	"fmt"
)

const (
	rsa1024 int = 1024
	rsa2048 int = 2048
	rsa4096 int = 4096
)

// Algorithm defines the supported key generation algorithms.
type Algorithm int

// List of supported algorithms for key generation.
const (
	AlgorithmUntyped Algorithm = iota
	AlgorithmED25519
	AlgorithmECDSAP521
	AlgorithmECDSAP384
	AlgorithmECDSAP256
	AlgorithmECDSAP224
	AlgorithmRSA4096
	AlgorithmRSA2048
	AlgorithmRSA1024
)

func (a Algorithm) String() string {
	switch a {
	case AlgorithmUntyped:
		return "untyped"
	case AlgorithmED25519:
		return "ED25519"
	case AlgorithmECDSAP521:
		return "ECDSA P521"
	case AlgorithmECDSAP384:
		return "ECDSA P384"
	case AlgorithmECDSAP256:
		return "ECDSA P256"
	case AlgorithmECDSAP224:
		return "ECDSA P224"
	case AlgorithmRSA4096:
		return "RSA 4096"
	case AlgorithmRSA2048:
		return "RSA 2048"
	case AlgorithmRSA1024:
		return "RSA 1024"
	default:
		return "unsupported"
	}
}

func maxAlgorithm() Algorithm {
	return AlgorithmRSA1024
}

// GenerateKeyPair creates a public-private key pair based on the specified algorithm.
// If you don't know what algorithm to use, insert zero to use the default (ED25519) key generation algorithm.
func GenerateKeyPair(algorithm Algorithm) (crypto.PublicKey, crypto.PrivateKey, error) {
	switch algorithm {
	case AlgorithmUntyped, AlgorithmED25519:
		return generateED25519KeyPair()
	case AlgorithmECDSAP521:
		return generateECDSAKeyPair(elliptic.P521())
	case AlgorithmECDSAP384:
		return generateECDSAKeyPair(elliptic.P384())
	case AlgorithmECDSAP256:
		return generateECDSAKeyPair(elliptic.P256())
	case AlgorithmECDSAP224:
		return generateECDSAKeyPair(elliptic.P224())
	case AlgorithmRSA4096:
		return generateRSAKeyPair(rsa4096)
	case AlgorithmRSA2048:
		return generateRSAKeyPair(rsa2048)
	case AlgorithmRSA1024:
		return generateRSAKeyPair(rsa1024)
	default:
		return nil, nil, errors.New("unsupported key generation algorithm")
	}
}

// generateRSAKeyPair creates an RSA public-private key pair with the specified bit size.
func generateRSAKeyPair(bits int) (*rsa.PublicKey, *rsa.PrivateKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, fmt.Errorf("error generating RSA key pair: %w", err)
	}
	return &privateKey.PublicKey, privateKey, nil

}

// generateECDSAKeyPair creates an ECDSA public-private key pair based on the specified elliptic curve.
func generateECDSAKeyPair(curve elliptic.Curve) (*ecdsa.PublicKey, *ecdsa.PrivateKey, error) {
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		return nil, nil, fmt.Errorf("error generating ECDSA key pair: %w", err)
	}
	return &privateKey.PublicKey, privateKey, nil
}

// generateED25519KeyPair creates an ED25519 public-private key pair.
func generateED25519KeyPair() (ed25519.PublicKey, ed25519.PrivateKey, error) {
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, nil, fmt.Errorf("error generating ED25519 key pair: %w", err)
	}
	return publicKey, privateKey, nil
}
*/