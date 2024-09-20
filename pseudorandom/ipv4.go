package pseudorandom

import (
	"fmt"
	"math/rand/v2"
	"net"
)

// IPv4 generates a deterministic pseudo-random IPv4 address within the specified IP network using the provided random source.
func IPv4(r *rand.Rand, cidr string) (net.IP, error) {
	if cidr == "" {
		// Generate a random IPv4 without CIDR restriction
		ip := make([]byte, 4)
		for i := range ip {
			ip[i] = byte(r.UintN(maxByteNumber))
		}
		return net.IPv4(ip[0], ip[1], ip[2], ip[3]), nil
	}

	// Parse the CIDR string
	ip, ipNet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, fmt.Errorf("invalid CIDR: %v", err)
	}

	// Convert IP to uint32 for easy arithmetic
	ipInt := ipToUint32(ip)

	// Calculate the size of the network
	ones, bits := ipNet.Mask.Size()
	size := uint32(1 << (bits - ones))

	// Generate a random number within the network size
	random := r.Uint32N(size)

	// Add the random number to the IP
	result := uint32ToIP(ipInt + random)

	return result, nil
}

// Helper function to convert IP to uint32
func ipToUint32(ip net.IP) uint32 {
	ip = ip.To4()
	return uint32(ip[0])<<24 | uint32(ip[1])<<16 | uint32(ip[2])<<8 | uint32(ip[3])
}

// Helper function to convert uint32 to IP
func uint32ToIP(ipInt uint32) net.IP {
	return net.IPv4(
		byte(ipInt>>24),
		byte(ipInt>>16),
		byte(ipInt>>8),
		byte(ipInt),
	)
}
