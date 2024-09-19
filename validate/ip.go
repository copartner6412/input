package validate

import (
	"fmt"
	"net"
)

// IP validates if the provided IP address is within the specified CIDR range.
//
// Parameters:
//   - ip: A string representation of an IP address (IPv4 or IPv6).
//   - cidr: A string representation of a CIDR block (e.g., "192.168.0.0/24").
//
// Returns:
//   - An error if the IP is invalid, the CIDR is invalid, or the IP is not within the CIDR range.
//   - nil if the IP is valid and within the specified CIDR range.
func IP(ip string, cidr string) error {
	parsedIP := net.ParseIP(ip)
	// Check if the IP address is nil (invalid).
	if parsedIP == nil {
		return fmt.Errorf("invalid textual representation of an IP address: %s", ip)
	}

	if cidr != "" {
		_, ipNet, err := net.ParseCIDR(cidr)
		if err != nil {
			return fmt.Errorf("invalid CIDR notation: %s", cidr)
		}

		// Check if the parsed IP is contained within the IP network.
		if !ipNet.Contains(parsedIP) {
			return fmt.Errorf("IP address %s is not within the specified CIDR range %s", ip, cidr)
		}
	}

	// Return nil to indicate that the IP is a valid IP in the specified IP network.
	return nil
}
