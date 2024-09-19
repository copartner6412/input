package validate

import (
	"fmt"
)

// CountryName checks if the given ISO 3166-1 country name is valid.
// It returns nil if the code is valid, otherwise it returns an error.
func CountryName(name string) error {
	// Prepopulate a map to hold all valid names.
	validNames := make(map[string]struct{}, len(Countries))
	for name := range Countries {
		validNames[name] = struct{}{}
	}

	// Check if the name exists in the map of valid names.
	if _, ok := validNames[name]; !ok {
		return fmt.Errorf("invalid ISO 3166-1 country name %s", name)
	}

	return nil
}

// CountryCode2 checks if the given ISO 3166-1 alpha-2 code is valid.
// It returns nil if the code is valid, otherwise it returns an error.
//
// For example US is valid but USA invalid.
//
// For more information about ISO 3166-1 alpha-2 code see https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2.
func CountryCode2(code string) error {
	if len(code) != 2 {
		return fmt.Errorf("country code %s doesn't have 2 letters", code)
	}

	// Prepopulate a map to hold all valid ISO 3166-1 alpha-2 codes for fast lookup.
	validCodes := make(map[string]struct{}, len(Countries))
	for _, country := range Countries {
		validCodes[country.Code2] = struct{}{}
	}

	// Check if the given code exists in the map of valid codes.
	if _, ok := validCodes[code]; !ok {
		return fmt.Errorf("invalid ISO 3166-1 alpha-2 conutry code %s", code)
	}

	return nil
}

// CountryCode3 checks if the given ISO 3166-1 alpha-3 code is valid.
// It returns nil if the code is valid, otherwise it returns an error.
//
// For example USA is valid but US invalid.
//
// For more information about ISO 3166-1 alpha-3 code see https://en.wikipedia.org/wiki/ISO_3166-1_alpha-3.
func CountryCode3(code string) error {
	if len(code) != 3 {
		return fmt.Errorf("country code %s doesn't have 3 letters", code)
	}

	// Prepopulate a map to hold all valid ISO 3166-1 alpha-3 codes for fast lookup.
	validCodes := make(map[string]struct{}, len(Countries))
	for _, country := range Countries {
		validCodes[country.Code3] = struct{}{}
	}

	// Check if the given code exists in the map of valid codes.
	if _, ok := validCodes[code]; !ok {
		return fmt.Errorf("invalid ISO 3166-1 alpha-3 code %s", code)
	}

	return nil
}
