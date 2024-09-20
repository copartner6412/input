package validate_test

import (
	"testing"

	"github.com/copartner6412/input/validate"
)

func TestTLDSuccessfulForValidTLD(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		tld       string
		minLength uint
		maxLength uint
	}{
		// Test case with a valid TLD exactly at the minimum length
		"validTLDAtMinLength": {
			tld:       "uk", // Example TLD, assuming minLength allows this
			minLength: minTLDLengthAllowed,
			maxLength: maxTLDLengthAllowed,
		},

		// Test case with a valid TLD exactly at the maximum length
		"validTLDAtMaxLength": {
			tld:       "international", // Example TLD, assuming maxLength allows this
			minLength: minTLDLengthAllowed,
			maxLength: maxTLDLengthAllowed,
		},

		// Test case with a valid TLD within the valid range
		"validTLDWithinRange": {
			tld:       "com",
			minLength: minTLDLengthAllowed,
			maxLength: maxTLDLengthAllowed,
		},

		// Test case with a valid TLD exactly at a boundary condition
		"validTLDAtBoundaryCondition": {
			tld:       "xyz", // Example TLD, ensure this is a valid TLD in TLDs map
			minLength: 3,
			maxLength: 3,
		},

		// Test case with a valid TLD at the upper boundary of length
		"validTLDAtUpperBoundary": {
			tld:       "museum", // Example TLD, assuming it’s a valid TLD
			minLength: minTLDLengthAllowed,
			maxLength: 7,
		},

		// Test case with a valid TLD with exactly the allowed maximum length
		"validTLDExactlyMaxLength": {
			tld:       "technology", // Example TLD, assuming it’s a valid TLD
			minLength: minTLDLengthAllowed,
			maxLength: 10,
		},

		// Test case with a valid TLD within allowed range with varying lengths
		"validTLDWithVariedLength": {
			tld:       "info",
			minLength: 2,
			maxLength: 4,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			err := validate.TLD(tc.tld, tc.minLength, tc.maxLength)
			if err != nil {
				t.Errorf("expected no error for valid TLD \"%s\", but got error: %v", tc.tld, err)
			}
		})
	}
}

func TestTLDFailsForInvalidTLD(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		tld       string
		minLength uint
		maxLength uint
	}{
		// Case where the TLD is empty
		"emptyTLD": {
			tld:       "",
			minLength: minTLDLengthAllowed,
			maxLength: maxTLDLengthAllowed,
		},

		// Case where the TLD is shorter than the minimum allowed length
		"shortTLD": {
			tld:       "a",
			minLength: 2, // assuming the minimum length is greater than 1
			maxLength: maxTLDLengthAllowed,
		},

		// Case where the TLD is longer than the maximum allowed length
		"longTLD": {
			tld:       "thisisaverylongtldvaluewhichshouldfail",
			minLength: minTLDLengthAllowed,
			maxLength: 10, // assuming the maximum length is less than this
		},

		// Case where the TLD contains invalid characters (e.g., numbers)
		"invalidCharacterTLD": {
			tld:       "tld123", // Assuming TLDs should only contain letters
			minLength: minTLDLengthAllowed,
			maxLength: maxTLDLengthAllowed,
		},

		// Case where TLD contains spaces (invalid characters)
		"spaceInTLD": {
			tld:       "t ld", // Assuming TLDs should not contain spaces
			minLength: minTLDLengthAllowed,
			maxLength: maxTLDLengthAllowed,
		},

		// Case where TLD contains special characters (e.g., @, #)
		"specialCharacterTLD": {
			tld:       "tld@#", // Assuming TLDs should not contain special characters
			minLength: minTLDLengthAllowed,
			maxLength: maxTLDLengthAllowed,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			err := validate.TLD(tc.tld, tc.minLength, tc.maxLength)
			if err == nil {
				t.Errorf("expected error for invalid TLD \"%s\", but got no error", tc.tld)
			}
		})
	}
}

func TestCCTLDSuccessfulForValidTLD(t *testing.T) {
	t.Parallel()

	testCases := map[string]string{
		// Example of a valid ccTLD from the list
		"validCCtldForUnitedKingdom": "uk", // Assuming "uk" is in Countries

		// Example of another valid ccTLD
		"validCCtldForGermany": "de", // Assuming "de" is in Countries

		// Example of a valid ccTLD for a large country
		"validCCtldForBrazil": "br", // Assuming "br" is in Countries

		// Example of a valid ccTLD for a small country
		"validCCtldForSanMarino": "sm", // Assuming "sm" is in Countries
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			err := validate.CCTLD(tc)
			if err != nil {
				t.Errorf("expected no error for valid ccTLD \"%s\", but got error: %v", tc, err)
			}
		})
	}
}

func TestCCTLDFailsForInvalidTLD(t *testing.T) {
	t.Parallel()

	testCases := map[string]string{
		// Case where the TLD is empty
		"emptyTLD": "",

		// Case where the TLD is too short (less than typical valid length)
		"shortTLD": "x", // Assuming valid ccTLDs are longer

		// Case where the TLD is too long (more than typical valid length)
		"longTLD": "toolongtldtoolongtldtoolongtld", // Assuming valid ccTLDs are shorter

		// Case where the TLD contains invalid characters
		"invalidCharacterTLD": "tld$", // Assuming TLDs are alphabetic

		// Case where the TLD is not in the list of valid ccTLDs
		"nonexistentTLD": "zzz", // Assuming "zzz" is not in Countries

		// Case where the TLD contains numeric characters (assuming TLDs should be alphabetic)
		"numericTLD": "t1d", // Assuming valid TLDs do not contain numbers

		// Case with a TLD that is valid in length but not recognized
		"validLengthButInvalidTLD": "xy", // Assuming "xy" is not in Countries

		// Case where TLD contains spaces
		"spaceInTLD": "t ld", // Assuming TLDs should not contain spaces

		// Case where the TLD is a known but invalid ccTLD for this application
		"knownInvalidTLD": "xyz", // Assuming "xyz" is not in Countries
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			err := validate.CCTLD(tc)
			if err == nil {
				t.Errorf("expected error for invalid ccTLD \"%s\", but got no error", tc)
			}
		})
	}
}
