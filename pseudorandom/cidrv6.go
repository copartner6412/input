package pseudorandom

import (
	"math/rand/v2"
	"net"
)

// CIDRv6 generates a deterministic pseudo-random IPv6 CIDR (IPv6 network) using the provided random source.
func CIDRv6(r *rand.Rand) net.IPNet {
	// Generate a random IPv6 address
	ip := make([]byte, 16)
	for i := range ip {
		ip[i] = byte(r.UintN(256))
	}

	// Generate a random subnet mask length (between 8 and 126)
	maskLen := r.IntN(119) + 8

	// Create the IPNet structure
	ipNet := net.IPNet{
		IP:   net.IP(ip).Mask(net.CIDRMask(maskLen, 128)),
		Mask: net.CIDRMask(maskLen, 128),
	}

	return ipNet
}
