package validate

import (
	"fmt"
)

// TLD validates whether the provided string is a valid top-level domain (TLD)
// according to the list of TLDs recognized by IANA.
// Returns an error if the TLD is invalid.
func TLD(tld string, minLength, maxLength uint) error {
	err := checkLength(len(tld), minLength, maxLength, minTLDLengthAllowed, maxTLDLengthAllowed, "characters")
	if err != nil {
		return err
	}

	_, ok := TLDs[tld]
	if !ok {
		return fmt.Errorf("\"%s\" not a valid TLD recognized by IANA", tld)
	}
	return nil
}

// CCTLD validates whether the provided string is a valid country code top-level domain (ccTLD)
// based on the list of country codes recognized by IANA.
// Returns an error if the ccTLD is invalid.
func CCTLD(tld string) error {
	var tlds = map[string]struct{}{}
	for _, country := range Countries {
		tlds[country.CCTLD] = struct{}{}
	}
	_, ok := tlds[tld]
	if !ok {
		return fmt.Errorf("\"%s\" not a valid ccTLD recognized by IANA", tld)
	}
	return nil
}
