package validate

import (
	"errors"
	"fmt"
	"strings"
)

const (
	minSubdomainLengthAllowed uint = 1
	maxSubdomainLengthAllowed uint = 63
)

// A subdomain must be 63 characters or less, start with a letter and end with a letter or digit, only contain letters, digits, or hyphens.
// https://developers.cloudflare.com/dns/manage-dns-records/reference/dns-record-types/

// Subdomain checks if a single subdomain (label) is valid according to DNS naming conventions.
// It ensures that:
//  1. The subdomain length is between 1 and 63 characters.
//  2. It contains only letters, digits, and hyphens.
//  3. It does not start or end with a hyphen.
//
// Each element of the internet domain name must be from 1 to 63 characters long and the entire domain, including the dots, can be at most 253 characters long. Valid characters for hostnames are ASCII letters from a to z, the digits from 0 to 9, and the hyphen (-). A hostname may not start with a hyphen.
// https://man7.org/linux/man-pages/man7/hostname.7.html
func Subdomain(subdomain string, minLength, maxLength uint) error {
	err := checkLength(len(subdomain), minLength, maxLength, minSubdomainLengthAllowed, maxSubdomainLengthAllowed, "characters")
	if err != nil {
		return err
	}

	var errs []error

	// Check if the subdomain starts with a hyphen.
	if strings.HasPrefix(subdomain, "-") {
		errs = append(errs, fmt.Errorf("subdomain '%s' starts with a hyphen", subdomain))
	}

	// Check if the subdomain ends with a hyphen.
	if strings.HasSuffix(subdomain, "-") {
		errs = append(errs, fmt.Errorf("subdomain '%s' ends with a hyphen", subdomain))
	}

	allowedCharacters := append(alphanumericalRunes, '-')

	// Check if the subdomain contains only valid characters.
	for i, char := range subdomain {
		if !strings.Contains(string(allowedCharacters), string(char)) {
			errs = append(errs, fmt.Errorf("subdomain '%s' contains invalid character '%c' at index %d", subdomain, char, i))
		}
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}
