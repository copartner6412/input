package random

import (
	"crypto/rand"
	"fmt"
	"io"
	"math/big"
)

const (
	limitPorts           uint = 1 << 16
	limitPortsWellKnown  uint = 1 << 10
	limitPortsRegistered uint = 49151 + 1
)

func Port(randomness io.Reader) (uint16, error) {
	random, err := rand.Int(randomness, big.NewInt(int64(limitPorts)))
	if err != nil {
		return 0, fmt.Errorf("error generating a random number for port: %w", err)
	}
	port := uint16(random.Int64())
	return port, nil
}

// PortWellKnown generates a cryptographically-secure random well-known port number [0–1023].
func PortWellKnown(randomness io.Reader) (uint16, error) {
	random, err := rand.Int(randomness, big.NewInt(int64(limitPortsWellKnown)))
	if err != nil {
		return 0, fmt.Errorf("error generating a random number for well-known port: %w", err)
	}
	port := uint16(random.Int64())
	return port, nil
}

// PortNotWellKnown generates a cryptographically-secure random port number in the range [1024–65535].
func PortNotWellKnown(randomness io.Reader) (uint16, error) {
	random, err := rand.Int(randomness, big.NewInt(int64(limitPorts-limitPortsWellKnown)))
	if err != nil {
		return 0, fmt.Errorf("error generating a random number for port outside of well-kown ports range: %w", err)
	}
	port := uint16(random.Int64()) + uint16(limitPortsWellKnown)
	return port, nil
}

// PortRegistered generates a cryptographically-secure random registered port number [1024–49151].
func PortRegistered(randomness io.Reader) (uint16, error) {
	random, err := rand.Int(randomness, big.NewInt(int64(limitPortsRegistered-limitPortsWellKnown)))
	if err != nil {
		return 0, fmt.Errorf("error generating a random number for port outside of registered ports range: %w", err)
	}
	port := uint16(random.Int64()) + uint16(limitPortsWellKnown)
	return port, nil
}

// PortPrivate generates a cryptographically-secure random private or dynamic port number [49152–65535].
func PortPrivate(randomness io.Reader) (uint16, error) {
	random, err := rand.Int(randomness, big.NewInt(int64(limitPorts-limitPortsRegistered)))
	if err != nil {
		return 0, fmt.Errorf("error generating a random number for port outside of private (dynamic) ports range: %w", err)
	}
	port := uint16(random.Int64()) + uint16(limitPortsRegistered)
	return port, nil
}
