package validate_test

import (
	"math/rand/v2"
	"testing"

	"github.com/copartner6412/input/pseudorandom"
	"github.com/copartner6412/input/validate"
)

func FuzzIPSuccessfulForValidPseudorandomV4Input(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed1, seed2 uint64) {
		r1 := rand.New(rand.NewPCG(seed1, seed2))
		cidr41 := pseudorandom.CIDRv4(r1)
		ipv41, err := pseudorandom.IPv4(r1, cidr41.String())
		if err != nil {
			t.Fatalf("error generating a pseudo-random IPv4: %v", err)
		}

		err = validate.IP(ipv41.String(), cidr41.String())
		if err != nil {
			t.Fatalf("unexpected error for the pseudo-random IPv4 \"%s\" within \"%s\" network: %v", ipv41.String(), cidr41.String(), err)
		}

		r2 := rand.New(rand.NewPCG(seed1, seed2))
		cidr42 := pseudorandom.CIDRv4(r2)
		ipv42, err := pseudorandom.IPv4(r2, cidr42.String())

		if ipv41.String() != ipv42.String() {
			t.Fatal("not deterministic")
		}
	})
}

func FuzzIPvSuccessfulForValidPseudorandomV6Input(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed1, seed2 uint64) {
		r1 := rand.New(rand.NewPCG(seed1, seed2))
		cidr61 := pseudorandom.CIDRv6(r1)
		ipv61, err := pseudorandom.IPv6(r1, cidr61.String())
		if err != nil {
			t.Errorf("error generating a pseudo-random IPv6: %v", err)
		}

		err = validate.IP(ipv61.String(), cidr61.String())
		if err != nil {
			t.Errorf("expected no error for the pseudo-random IPv6 \"%s\" within \"%s\" network, but got error: %v", ipv61.String(), cidr61.String(), err)
		}

		r2 := rand.New(rand.NewPCG(seed1, seed2))
		cidr62 := pseudorandom.CIDRv6(r2)
		ipv62, err := pseudorandom.IPv6(r2, cidr62.String())
		if err != nil {
			t.Errorf("error generating a pseudo-random IPv6: %v", err)
		}
		
		if ipv61.String() != ipv62.String() {
			t.Error("not deterministic")
		}
	})
}

func TestIPSuccessfulForValidInput(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct{
		ip   string
		cidr string
	}{
		// Test a valid IPv4 address without CIDR
		"validIPv4WithoutCIDR": {
			ip:   "192.168.1.1",
			cidr: "",  // No CIDR provided, should pass as valid IP
		},

		// Test a valid IPv6 address without CIDR
		"validIPv6WithoutCIDR": {
			ip:   "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
			cidr: "",  // No CIDR provided, should pass as valid IP
		},

		// Test a valid IPv4 address within a CIDR range
		"validIPv4WithinCIDR": {
			ip:   "192.168.1.10",
			cidr: "192.168.1.0/24",  // IP within range
		},

		// Test a valid IPv6 address within a CIDR range
		"validIPv6WithinCIDR": {
			ip:   "2001:db8::1",
			cidr: "2001:db8::/32",  // IP within range
		},

		// Test a valid IPv4 address at the boundary of a CIDR range (network address)
		"validIPv4AtBoundaryNetworkAddress": {
			ip:   "10.0.0.0",
			cidr: "10.0.0.0/8",  // Network address, valid boundary case
		},

		// Test a valid IPv4 address at the boundary of a CIDR range (broadcast address)
		"validIPv4AtBoundaryBroadcastAddress": {
			ip:   "10.255.255.255",
			cidr: "10.0.0.0/8",  // Broadcast address, valid boundary case
		},

		// Test a valid IPv6 address at the boundary of a CIDR range
		"validIPv6AtBoundaryNetworkAddress": {
			ip:   "2001:db8::",
			cidr: "2001:db8::/32",  // Network address, valid boundary case
		},

		// Test a valid IPv4 CIDR range with exact match
		"validIPv4ExactMatchInCIDR": {
			ip:   "10.0.0.1",
			cidr: "10.0.0.1/32",  // IP matches exactly with CIDR
		},

		// Test a valid IPv6 CIDR range with exact match
		"validIPv6ExactMatchInCIDR": {
			ip:   "2001:db8::1",
			cidr: "2001:db8::1/128",  // IP matches exactly with CIDR
		},

		// Test a valid IPv4 with classless inter-domain routing (CIDR block)
		"validIPv4WithCIDRBlock": {
			ip:   "172.16.5.4",
			cidr: "172.16.0.0/12",  // Class B private network range
		},

		// Test a valid IPv6 with classless inter-domain routing (CIDR block)
		"validIPv6WithCIDRBlock": {
			ip:   "fe80::1",
			cidr: "fe80::/10",  // Link-local IPv6 range
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

	testCases := map[string]struct{
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

