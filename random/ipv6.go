package random

import (
	"crypto/rand"
	"fmt"
	"io"
	"math/big"
	"net"
)

// IPv6 generates a random IPv6 address of type net.IP and returns an error.
func IPv6(randomness io.Reader, cidr string) (net.IP, error) {
	if cidr == "" {

		// Generate a random IPv6 without CIDR restriction
		ip := make([]byte, net.IPv6len)
		for i := range ip {
			random1, err := rand.Int(randomness, big.NewInt(int64(maxByteNumber)))
			if err != nil {
				return net.IP{}, fmt.Errorf("error generating a random number for byte: %w", err)
			}

			ip[i] = byte(random1.Int64())
		}
		return ip, nil
	}

	// Parse the CIDR string
	_, ipNet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, fmt.Errorf("invalid CIDR: %w", err)
	}

	ip := make([]byte, net.IPv6len)
	copy(ip, ipNet.IP)

	// Calculate the number of host bits
	ones, _ := ipNet.Mask.Size()
	hostBits := 128 - ones

	// Generate random host part
	for i := 0; i < hostBits; i++ {
		byteIndex := i / 8
		bitIndex := i % 8
		random2, err := rand.Int(randomness, big.NewInt(2))
		if err != nil {
			return net.IP{}, fmt.Errorf("error generating a random number for chance of convert a zero bit to one: %w", err)
		}

		if random2.Int64() == 1 {
			ip[15-byteIndex] |= 1 << bitIndex
		}
	}

	return ip, nil
}
