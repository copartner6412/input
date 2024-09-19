package validate_test

import (
	"strings"
	"testing"

	"github.com/copartner6412/input/validate"
)

const (
	minSubdomainLengthAllowed uint = 1
	maxSubdomainLengthAllowed uint = 63
)

func TestSubdomainSuccessfulForValidInput(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct{
		subdomain string
		minLength uint
		maxLength uint
	}{
		// Test a subdomain exactly at the minimum length allowed
		"subdomainAtMinLength": {
			subdomain: "a",
			minLength: minSubdomainLengthAllowed,
			maxLength: maxSubdomainLengthAllowed,
		},

		// Test a subdomain exactly at the maximum length allowed
		"subdomainAtMaxLength": {
			subdomain: strings.Repeat("a", int(maxSubdomainLengthAllowed)),
			minLength: minSubdomainLengthAllowed,
			maxLength: maxSubdomainLengthAllowed,
		},

		// Test a valid subdomain with only letters and digits
		"subdomainWithLettersAndDigits": {
			subdomain: "validsubdomain123",
			minLength: minSubdomainLengthAllowed,
			maxLength: maxSubdomainLengthAllowed,
		},

		// Test a valid subdomain with hyphens but not starting or ending with a hyphen
		"subdomainWithHyphens": {
			subdomain: "valid-subdomain",
			minLength: minSubdomainLengthAllowed,
			maxLength: maxSubdomainLengthAllowed,
		},

		// Test a valid subdomain with mixed case letters
		"subdomainWithMixedCaseLetters": {
			subdomain: "ValidSubDomain",
			minLength: minSubdomainLengthAllowed,
			maxLength: maxSubdomainLengthAllowed,
		},

		// Test a valid subdomain with numbers and letters
		"subdomainWithNumbersAndLetters": {
			subdomain: "subdomain123",
			minLength: minSubdomainLengthAllowed,
			maxLength: maxSubdomainLengthAllowed,
		},

		// Test a valid subdomain with hyphens in the middle
		"subdomainWithHyphenInMiddle": {
			subdomain: "sub-domain",
			minLength: minSubdomainLengthAllowed,
			maxLength: maxSubdomainLengthAllowed,
		},
	}

	// Run tests
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			err := validate.Subdomain(tc.subdomain, tc.minLength, tc.maxLength)
			if err != nil {
				t.Errorf("expected no error for valid subdomain \"%s\", but got error: %v", tc.subdomain, err)
			}
		})
	}
}


func TestSubdomainFailsForInvalidInput(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct{
		subdomain string
		minLength uint
		maxLength uint
	}{
		// Test case with empty subdomain (below minimum length)
		"emptySubdomain": {
			subdomain: "",
			minLength: minSubdomainLengthAllowed,
			maxLength: maxSubdomainLengthAllowed,
		},

		// Test case with subdomain shorter than the minimum length allowed
		"subdomainShorterThanMinLength": {
			subdomain: "a",
			minLength: 2,
			maxLength: maxSubdomainLengthAllowed,
		},

		// Test case with subdomain longer than the maximum length allowed
		"subdomainLongerThanMaxLength": {
			subdomain: strings.Repeat("a", int(maxSubdomainLengthAllowed)+1),
			minLength: minSubdomainLengthAllowed,
			maxLength: maxSubdomainLengthAllowed,
		},

		// Test case with subdomain that starts with a hyphen
		"subdomainStartsWithHyphen": {
			subdomain: "-invalid",
			minLength: minSubdomainLengthAllowed,
			maxLength: maxSubdomainLengthAllowed,
		},

		// Test case with subdomain that ends with a hyphen
		"subdomainEndsWithHyphen": {
			subdomain: "invalid-",
			minLength: minSubdomainLengthAllowed,
			maxLength: maxSubdomainLengthAllowed,
		},

		// Test case with subdomain containing invalid characters (e.g., special characters)
		"subdomainWithInvalidCharacters": {
			subdomain: "invalid@domain",
			minLength: minSubdomainLengthAllowed,
			maxLength: maxSubdomainLengthAllowed,
		},

		// Test case with subdomain containing a space (which is invalid)
		"subdomainWithSpace": {
			subdomain: "invalid domain",
			minLength: minSubdomainLengthAllowed,
			maxLength: maxSubdomainLengthAllowed,
		},

		// Test case with subdomain containing non-ASCII characters
		"subdomainWithNonASCIICharacters": {
			subdomain: "ñäçtive",
			minLength: minSubdomainLengthAllowed,
			maxLength: maxSubdomainLengthAllowed,
		},

		// Test case with subdomain where minLength > maxLength
		"minLengthGreaterThanMaxLength": {
			subdomain: "valid",
			minLength: 10,
			maxLength: 5,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			err := validate.Subdomain(tc.subdomain, tc.minLength, tc.maxLength)
			if err == nil {
				t.Errorf("expected error for invalid subdomain \"%s\", but got no error", tc.subdomain)
			}
		})
	}
}

