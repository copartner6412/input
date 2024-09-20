package validate

import (
	"errors"
	"fmt"
	"strings"
)

const (
	minDomainLengthAllowed uint = 1
	maxDomainLengthAllowed uint = 253
	maxSubdomainCount      uint = 127 // (253 + 1) / 2
)

// https://developers.cloudflare.com/dns/manage-dns-records/reference/dns-record-types/

// Domain checks if the given domain name is valid according to DNS naming conventions.
// It ensures that:
//  1. The domain is not empty.
//  2. The total length of the domain is within the allowed limit.
//  3. Each subdomain (label) is valid.
//  4. The label "www" is only used as the first label.
//
// https://man7.org/linux/man-pages/man7/hostname.7.html
func Domain(domain string, minLength, maxLength uint) error {
	err := checkLength(len([]rune(domain)), minLength, maxLength, minDomainLengthAllowed, maxDomainLengthAllowed, "characters")
	if err != nil {
		return err
	}

	// Split the domain into subdomains.
	subdomains := strings.Split(domain, ".")

	// Validate each subdomain.
	var errs []error
	for i, subdomain := range subdomains {
		// Check if "www" is used as a subdomain and is not the first label.
		if subdomain == "www" && i != 0 {
			errs = append(errs, errors.New("'www' cannot be used as a subdomain"))
		}

		// Validate the subdomain.
		if err := Subdomain(subdomain, minSubdomainLengthAllowed, maxSubdomainLengthAllowed); err != nil {
			errs = append(errs, fmt.Errorf("invalid subdomain: %w", err))
		}
	}

	// Return combined errors if any.
	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}

// TLD checks the validity of the top-level domain (TLD) in a given domain name.
// It checks whether the TLD exists in valid TLDs from IANA: https://data.iana.org/TLD/tlds-alpha-by-domain.txt.
// Returns an error if the TLD is invalid, otherwise returns nil.
func DomainWithValidTLD(domain string, minLength, maxLength uint) error {
	err := Domain(domain, minLength, maxLength)
	if err != nil {
		return err
	}

	// Split the domain by dots to extract the TLD (last segment).
	parts := strings.Split(domain, ".")
	tldPart := parts[len(parts)-1]

	_, ok := TLDs[tldPart]
	if !ok {
		// If all checks pass, return nil indicating the TLD is valid
		return fmt.Errorf("tld \"%s\" not found in valid TLDs", tldPart)
	}

	return nil
}

// CountryTLD checks if the given domain has a valid Internet country code top-level domain (ccTLD).
// It returns nil if the tld is a valid ccTLD, otherwise it returns an error.
func DomainWithValidCCTLD(domain string, minLength, maxLength uint) error {
	err := Domain(domain, minLength, maxLength)
	if err != nil {
		return err
	}

	parts := strings.Split(domain, ".")
	tld := parts[len(parts)-1]

	// Prepopulate a map to hold all valid ccTLDs for fast lookup.
	validTLDs := make(map[string]struct{})
	for _, country := range Countries {
		validTLDs[country.CCTLD] = struct{}{}
	}

	// Check if the given ccTLD exists in the map of valid TLDs.
	if _, ok := validTLDs[tld]; !ok {
		return fmt.Errorf("invalid Internet country code top-level domain (CCTLD) %s", tld)
	}

	return nil
}
