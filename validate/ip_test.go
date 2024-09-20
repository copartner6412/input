package validate_test

import (
	"testing"

	"github.com/copartner6412/input/validate"
)

func TestIPSuccessfulForValidInput(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		ip   string
		cidr string
	}{
		// Test a valid IPv4 address without CIDR
		"validIPv4WithoutCIDR": {
			ip:   "192.168.1.1",
			cidr: "", // No CIDR provided, should pass as valid IP
		},

		// Test a valid IPv6 address without CIDR
		"validIPv6WithoutCIDR": {
			ip:   "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
			cidr: "", // No CIDR provided, should pass as valid IP
		},

		// Test a valid IPv4 address within a CIDR range
		"validIPv4WithinCIDR": {
			ip:   "192.168.1.10",
			cidr: "192.168.1.0/24", // IP within range
		},

		// Test a valid IPv6 address within a CIDR range
		"validIPv6WithinCIDR": {
			ip:   "2001:db8::1",
			cidr: "2001:db8::/32", // IP within range
		},

		// Test a valid IPv4 address at the boundary of a CIDR range (network address)
		"validIPv4AtBoundaryNetworkAddress": {
			ip:   "10.0.0.0",
			cidr: "10.0.0.0/8", // Network address, valid boundary case
		},

		// Test a valid IPv4 address at the boundary of a CIDR range (broadcast address)
		"validIPv4AtBoundaryBroadcastAddress": {
			ip:   "10.255.255.255",
			cidr: "10.0.0.0/8", // Broadcast address, valid boundary case
		},

		// Test a valid IPv6 address at the boundary of a CIDR range
		"validIPv6AtBoundaryNetworkAddress": {
			ip:   "2001:db8::",
			cidr: "2001:db8::/32", // Network address, valid boundary case
		},

		// Test a valid IPv4 CIDR range with exact match
		"validIPv4ExactMatchInCIDR": {
			ip:   "10.0.0.1",
			cidr: "10.0.0.1/32", // IP matches exactly with CIDR
		},

		// Test a valid IPv6 CIDR range with exact match
		"validIPv6ExactMatchInCIDR": {
			ip:   "2001:db8::1",
			cidr: "2001:db8::1/128", // IP matches exactly with CIDR
		},

		// Test a valid IPv4 with classless inter-domain routing (CIDR block)
		"validIPv4WithCIDRBlock": {
			ip:   "172.16.5.4",
			cidr: "172.16.0.0/12", // Class B private network range
		},

		// Test a valid IPv6 with classless inter-domain routing (CIDR block)
		"validIPv6WithCIDRBlock": {
			ip:   "fe80::1",
			cidr: "fe80::/10", // Link-local IPv6 range
		},
	}

	// Run tests
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			err := validate.IP(tc.ip, tc.cidr)
			if err != nil {
				t.Errorf("expected no error for valid input %v, but got error: %v", tc.ip, err)
			}
		})
	}
}

func TestIPFailsForInvalidInput(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		ip   string
		cidr string
	}{
		// Test an empty IP string
		"emptyIP": {
			ip:   "",
			cidr: "", // No CIDR, should fail due to invalid IP
		},

		// Test an invalid IP address format (not an IP)
		"invalidIPFormat": {
			ip:   "invalid_ip",
			cidr: "", // Invalid IP format
		},

		// Test an IPv4 address with an invalid CIDR notation
		"invalidCIDRNotation": {
			ip:   "192.168.1.1",
			cidr: "192.168.1.0/abc", // Invalid CIDR notation
		},

		// Test an IP address not within the given CIDR range
		"ipOutsideCIDRRange": {
			ip:   "10.0.0.1",
			cidr: "192.168.1.0/24", // IP is outside the specified CIDR range
		},

		// Test a valid IPv4 address but invalid CIDR format
		"validIPInvalidCIDR": {
			ip:   "192.168.1.1",
			cidr: "192.168.1.0/33", // Invalid CIDR range (CIDR mask too large)
		},

		// Test an empty CIDR with an invalid IP address
		"invalidIPWithEmptyCIDR": {
			ip:   "invalid_ip",
			cidr: "", // Invalid IP with no CIDR should fail
		},

		// Test a valid IPv6 address outside of the given CIDR range
		"ipv6OutsideCIDRRange": {
			ip:   "2001:db7::2",
			cidr: "2001:db8::/64", // IP outside the given CIDR range
		},

		// Test an invalid IPv6 format
		"invalidIPv6Format": {
			ip:   "2001:db8:::",
			cidr: "", // Invalid IPv6 format
		},

		// Test a valid CIDR but empty IP
		"emptyIPWithValidCIDR": {
			ip:   "",
			cidr: "192.168.1.0/24", // Empty IP with a valid CIDR should fail
		},
	}

	// Run tests
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			err := validate.IP(tc.ip, tc.cidr)
			if err == nil {
				t.Errorf("expected error for IP %v and CIDR %v, but got no error", tc.ip, tc.cidr)
			}
		})
	}
}
