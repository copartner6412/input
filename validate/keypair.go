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