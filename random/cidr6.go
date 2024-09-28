package random

import (
	"crypto/rand"
	"fmt"
	"io"
	"math/big"
	"net"
)

// CIDRv6 generates a deterministic pseudo-random IPv6 CIDR (IPv6 network) using the provided random source.
func CIDRv6(randomness io.Reader) (net.IPNet, error) {
	// Generate a random IPv6 address
	ip := make([]byte, 16)
	for i := range ip {
		random1, err := rand.Int(randomness, big.NewInt(int64(maxByteNumber)))
		if err != nil {
			return net.IPNet{}, fmt.Errorf("error generating random number for calculating byte: %w", err)
		}
		ip[i] = byte(random1.Int64())
	}

	// Generate a random subnet mask length (between 8 and 126)
	random2, err := rand.Int(randomness, big.NewInt(119))
	if err != nil {
		return net.IPNet{}, fmt.Errorf("error generating random number for calculating mask length: %w", err)
	}
	maskLen := int(random2.Int64()) + 8

	// Create the IPNet structure
	ipNet := net.IPNet{
		IP:   net.IP(ip).Mask(net.CIDRMask(maskLen, 128)),
		Mask: net.CIDRMask(maskLen, 128),
	}

	return ipNet, nil
}
