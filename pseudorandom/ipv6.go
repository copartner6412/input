package pseudorandom

import (
	"fmt"
	"math/rand/v2"
	"net"
)

// IPv6 generates a deterministic pseudo-random IPv6 address within the specified IP network using the provided random source.
func IPv6(r *rand.Rand, cidr string) (net.IP, error) {
	if cidr == "" {
		
		// Generate a random IPv6 without CIDR restriction
		ip := make([]byte, net.IPv6len)
		for i := range ip {
			ip[i] = byte(r.UintN(maxByteNumber))
		}
		return ip, nil
	}

	// Parse the CIDR string
	_, ipNet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, fmt.Errorf("invalid CIDR: %v", err)
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
		if r.IntN(2) == 1 {
			ip[15-byteIndex] |= 1 << bitIndex
		}
	}

	return ip, nil
}
