package validate_test

import (
	"strings"
	"testing"

	"github.com/copartner6412/input/validate"
)

const (
	minLinuxHostnameLengthAllowed uint = 1
	maxLinuxHostnameLengthAllowed uint = 64
)

func TestLinuxHostnameSuccessfulForValidInput(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct{
		hostname string
		minLength uint
		maxLength uint
	}{
		// Test a hostname exactly at the minimum length allowed
		"hostnameAtMinLength": {
			hostname: "a",
			minLength: minLinuxHostnameLengthAllowed,
			maxLength: maxLinuxHostnameLengthAllowed,
		},

		// Test a hostname exactly at the maximum length allowed
		"hostnameAtMaxLength": {
			hostname: strings.Repeat("a", int(maxLinuxHostnameLengthAllowed)),
			minLength: minLinuxHostnameLengthAllowed,
			maxLength: maxLinuxHostnameLengthAllowed,
		},

		// Test a valid hostname with only letters and digits
		"hostnameWithLettersAndDigits": {
			hostname: "validhostname123",
			minLength: minLinuxHostnameLengthAllowed,
			maxLength: maxLinuxHostnameLengthAllowed,
		},

		// Test a valid hostname with hyphens but no leading or trailing hyphens
		"hostnameWithHyphens": {
			hostname: "valid-hostname",
			minLength: minLinuxHostnameLengthAllowed,
			maxLength: maxLinuxHostnameLengthAllowed,
		},

		// Test a valid hostname with numerical and alphabetical characters
		"hostnameWithNumbersAndLetters": {
			hostname: "hostname123",
			minLength: minLinuxHostnameLengthAllowed,
			maxLength: maxLinuxHostnameLengthAllowed,
		},

		// Test a valid hostname with a mix of letters, digits, and hyphens in the middle
		"hostnameWithMixedCharacters": {
			hostname: "host-name123",
			minLength: minLinuxHostnameLengthAllowed,
			maxLength: maxLinuxHostnameLengthAllowed,
		},

		// Test a valid hostname with a hyphen in the middle
		"hostnameWithHyphenInMiddle": {
			hostname: "example-host",
			minLength: minLinuxHostnameLengthAllowed,
			maxLength: maxLinuxHostnameLengthAllowed,
		},
	}

	// Run tests
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			err := validate.LinuxHostname(tc.hostname, tc.minLength, tc.maxLength)
			if err != nil {
				t.Errorf("expected no error for valid Linux hostname %q, but got error: %v", tc.hostname, err)
			}
		})
	}
}


func TestLinuxHostnameFailsForInvalidInput(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct{
		hostname string
		minLength uint
		maxLength uint
	}{
		// Test hostname shorter than the minimum length allowed
		"hostnameShorterThanMinLength": {
			hostname: "a", // Assuming minLength > 1
			minLength: 2,
			maxLength: maxLinuxHostnameLengthAllowed,
		},

		// Test hostname longer than the maximum length allowed
		"hostnameLongerThanMaxLength": {
			hostname: strings.Repeat("a", int(maxLinuxHostnameLengthAllowed+1)),
			minLength: minLinuxHostnameLengthAllowed,
			maxLength: maxLinuxHostnameLengthAllowed,
		},

		// Test hostname starting with a hyphen
		"hostnameStartsWithHyphen": {
			hostname: "-example",
			minLength: minLinuxHostnameLengthAllowed,
			maxLength: maxLinuxHostnameLengthAllowed,
		},

		// Test hostname starting with a number
		"hostnameStartsWithNumber": {
			hostname: "1example",
			minLength: minLinuxHostnameLengthAllowed,
			maxLength: maxLinuxHostnameLengthAllowed,
		},

		// Test hostname ending with a hyphen
		"hostnameEndsWithHyphen": {
			hostname: "example-",
			minLength: minLinuxHostnameLengthAllowed,
			maxLength: maxLinuxHostnameLengthAllowed,
		},

		// Test hostname containing invalid characters
		"hostnameContainsInvalidCharacter": {
			hostname: "example!@#",
			minLength: minLinuxHostnameLengthAllowed,
			maxLength: maxLinuxHostnameLengthAllowed,
		},

		// Test hostname containing non-printable characters
		"hostnameContainsNonPrintableCharacter": {
			hostname: "example\x00",
			minLength: minLinuxHostnameLengthAllowed,
			maxLength: maxLinuxHostnameLengthAllowed,
		},

		// Test hostname with mixed case characters (edge case, often valid but may need verification for specific systems)
		"hostnameMixedCase": {
			hostname: "ExAmPle",
			minLength: minLinuxHostnameLengthAllowed,
			maxLength: maxLinuxHostnameLengthAllowed,
		},

		// Test hostname with leading or trailing whitespace (usually trimmed, but can be an edge case)
		"hostnameWithLeadingOrTrailingWhitespace": {
			hostname: " example ",
			minLength: minLinuxHostnameLengthAllowed,
			maxLength: maxLinuxHostnameLengthAllowed,
		},
	}

	// Run tests
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			err := validate.LinuxHostname(tc.hostname, tc.minLength, tc.maxLength)
			if err == nil {
				t.Errorf("expected error for invalid Linux hostname %q, but got no error", tc.hostname)
			}
		})
	}
}

