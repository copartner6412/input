package random

import (
	"crypto/rand"
	"fmt"
	"io"
	"math/big"
	"net"
)

func CIDRv4(randomness io.Reader) (net.IPNet, error) {
	// Generate a random IPv4 address
	ip := make([]byte, 4)
	for i := range ip {
		random1, err := rand.Int(randomness, big.NewInt(int64(maxByteNumber)))
		if err != nil {
			return net.IPNet{}, fmt.Errorf("error generating a random number for byte: %w", err)
		}
		ip[i] = byte(random1.Int64())
	}

	// Generate a random subnet mask length (between 8 and 30)
	random2, err := rand.Int(randomness, big.NewInt(23))
	if err != nil {
		return net.IPNet{}, fmt.Errorf("error generating a random number for calculating mask length: %w", err)
	}

	maskLen := int(random2.Int64()) + 8

	// Create the IPNet structure
	ipNet := net.IPNet{
		IP:   net.IPv4(ip[0], ip[1], ip[2], ip[3]).Mask(net.CIDRMask(maskLen, 32)),
		Mask: net.CIDRMask(maskLen, 32),
	}

	return ipNet, nil
}
