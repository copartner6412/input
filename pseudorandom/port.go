package pseudorandom

import (
	"fmt"
	"math/rand/v2"
)

const (
	limitPorts           uint   = 1 << 16
	limitPortsWellKnown  uint   = 1 << 10
	limitPortsRegistered uint   = 49151 + 1
	minPortAllowed       uint16 = 0
	maxPortAllowed       uint16 = 65535
)

// Port generates a deterministic pseudo-random port number in the range 0–65535.
func Port(r *rand.Rand, minPort, maxPort uint16) (uint16, error) {
	if minPort == 0 && maxPort == 0 {
		maxPort = maxPortAllowed
	} else {
		if maxPort < minPort {
			return 0, fmt.Errorf("maxPort can not be less than minPort")
		}
	}

	port := uint16(r.UintN(uint(maxPort-minPort)+1)) + minPort

	return port, nil
}

// PortWellKnown generates a deterministic pseudo-random well-known port number [0–1023].
func PortWellKnown(r *rand.Rand) uint16 {
	return uint16(r.UintN(limitPortsWellKnown))
}

// PortNotWellKnown generates a deterministic pseudo-random port number in the range [1024–65535].
func PortNotWellKnown(r *rand.Rand) uint16 {
	return uint16(r.UintN(limitPorts-limitPortsWellKnown)) + uint16(limitPortsWellKnown)
}

// PortRegistered generates a deterministic pseudo-random registered port number [1024–49151].
func PortRegistered(r *rand.Rand) uint16 {
	return uint16(r.UintN(limitPortsRegistered-limitPortsWellKnown)) + uint16(limitPortsWellKnown)
}

// PortPrivate generates a deterministic pseudo-random private or dynamic port number [49152–65535].
func PortPrivate(r *rand.Rand) uint16 {
	return uint16(r.UintN(limitPorts-limitPortsRegistered)) + uint16(limitPortsRegistered)
}
