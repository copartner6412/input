package pseudorandom

import (
	"math/rand/v2"
	"net"
)

// CIDRv4 generates a deterministic pseudo-random IPv4 CIDR (IPv4 network) using the provided random source.
func CIDRv4(r *rand.Rand) net.IPNet {
	// Generate a random IPv4 address
	ip := make([]byte, 4)
	for i := range ip {
		ip[i] = byte(r.UintN(256))
	}

	// Generate a random subnet mask length (between 8 and 30)
	maskLen := r.IntN(23) + 8

	// Create the IPNet structure
	ipNet := net.IPNet{
		IP:   net.IPv4(ip[0], ip[1], ip[2], ip[3]).Mask(net.CIDRMask(maskLen, 32)),
		Mask: net.CIDRMask(maskLen, 32),
	}

	return ipNet
}