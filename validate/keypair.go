package validate

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/rsa"
	"errors"
	"fmt"
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

var algorithmString = map[Algorithm]string{
	AlgorithmUntyped:   "untyped",
	AlgorithmED25519:   "ED25519",
	AlgorithmECDSAP521: "ECDSA P521",
	AlgorithmECDSAP384: "ECDSA P384",
	AlgorithmECDSAP256: "ECDSA P256",
	AlgorithmECDSAP224: "ECDSA P224",
	AlgorithmRSA4096:   "RSA 4096",
	AlgorithmRSA2048:   "RSA 2048",
	AlgorithmRSA1024:   "RSA 1024",
}

func (a Algorithm) String() string {
	return algorithmString[a]
}

func KeyPair(algorithm Algorithm, publicKey crypto.PublicKey, privateKey crypto.PrivateKey) error {
	var nilErrs []error
	if publicKey == nil {
		nilErrs = append(nilErrs, errors.New("nil public key"))
	}

	if privateKey == nil {
		nilErrs = append(nilErrs, errors.New("nil private key"))
	}

	if len(nilErrs) > 0 {
		return errors.Join(nilErrs...)
	}

	switch algorithm {
	case AlgorithmUntyped, AlgorithmED25519:
		privateKey, ok := privateKey.(ed25519.PrivateKey)
		if !ok {
			return fmt.Errorf("different algorithm type for private key, expected ED25519 but it's %v", algorithm.String())
		}

		publicKey, ok := publicKey.(ed25519.PublicKey)
		if !ok {
			return fmt.Errorf("different algorithm type for public key, expected ED25519 but it's %v", algorithm.String())
		}

		match := privateKey.Public().(ed25519.PublicKey).Equal(publicKey)
		if !match {
			return fmt.Errorf("private and public key don't match with each other")
		}
	case AlgorithmECDSAP521, AlgorithmECDSAP384, AlgorithmECDSAP256, AlgorithmECDSAP224:
		privateKey, ok := privateKey.(*ecdsa.PrivateKey)
		if !ok {
			return fmt.Errorf("different algorithm type for private key, expected ED25519 but it's %v", algorithm.String())
		}

		publicKey, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			return fmt.Errorf("different algorithm type for public key, expected ED25519 but it's %v", algorithm.String())
		}

		match := privateKey.PublicKey.Equal(publicKey)
		if !match {
			return fmt.Errorf("private and public key don't match with each other")
		}

	case AlgorithmRSA4096, AlgorithmRSA2048, AlgorithmRSA1024:
		privateKey, ok := privateKey.(*rsa.PrivateKey)
		if !ok {
			return fmt.Errorf("different algorithm type for private key, expected ED25519 but it's %v", algorithm.String())
		}

		publicKey, ok = publicKey.(*rsa.PublicKey)
		if !ok {
			return fmt.Errorf("different algorithm type for public key, expected ED25519 but it's %v", algorithm.String())
		}

		match := privateKey.PublicKey.Equal(publicKey)
		if !match {
			return fmt.Errorf("private and public key don't match with each other")
		}

	default:
		return fmt.Errorf("unsupported algorithm type")
	}

	return nil
}
