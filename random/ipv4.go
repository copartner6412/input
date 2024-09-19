package random

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"net"
)

// IPv4 generates a random IPv4 address of type net.IP and returns an error.
func IPv4(cidr string) (net.IP, error) {
	if cidr == "" {
		// Generate a random IPv4 without CIDR restriction
		ip := make([]byte, 4)
		for i := range ip {
			random1, err := rand.Int(rand.Reader, big.NewInt(int64(maxByteNumber)))
			if err != nil {
				return net.IP{}, fmt.Errorf("error generating a random number for byte: %w", err)
			}

			ip[i] = byte(random1.Int64())
		}

		return net.IPv4(ip[0], ip[1], ip[2], ip[3]), nil
	}

	// Parse the CIDR string
	ip, ipNet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, fmt.Errorf("invalid CIDR: %v", err)
	}

	// Convert IP to uint32 for easy arithmetic
	ipUint32 := ipToUint32(ip)
	
	// Calculate the size of the network
	ones, bits := ipNet.Mask.Size()
	size := uint32(1 << (bits - ones))
	
	// Generate a random number within the network size
	random2, err := rand.Int(rand.Reader, big.NewInt(int64(size)))
	
	// Add the random number to the IP
	result := uint32ToIP(ipUint32 + uint32(random2.Int64()))

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