package validate_test

import (
	"math/rand/v2"
	"testing"

	"github.com/copartner6412/input/pseudorandom"
	"github.com/copartner6412/input/validate"
)

func FuzzTLDSuccessfulForValidPseudorandomInput(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed1, seed2 uint64, min, max uint) {
		r1, r2, minLength, maxLength := randoms(seed1, seed2, min, max, minTLDLengthAllowed, maxTLDLengthAllowed)
		tld1, err := pseudorandom.TLD(r1, minLength, maxLength)
		if err != nil {
			t.Fatalf("error generating a pseudo-random gTLD: %v", err)
		}

		err = validate.TLD(tld1, minLength, maxLength)
		if err != nil {
			t.Fatalf("expected no error for valid pseudo-random TLD \"%s\", but got error: %v", tld1, err)
		}
		
		tld2, err := pseudorandom.TLD(r2, minLength, maxLength)
		if err != nil {
			t.Fatalf("error regenerating the pseudo-random gTLD: %v", err)
		}
		
		if tld1 != tld2 {
			t.Fatal("not deterministic")
		}
	})
}

func TestTLDSuccessfulForValidTLD(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct{
		tld string
		minLength uint
		maxLength uint
	}{
		// Test case with a valid TLD exactly at the minimum length
		"validTLDAtMinLength": {
			tld: "uk", // Example TLD, assuming minLength allows this
			minLength: minTLDLengthAllowed,
			maxLength: maxTLDLengthAllowed,
		},

		// Test case with a valid TLD exactly at the maximum length
		"validTLDAtMaxLength": {
			tld: "international", // Example TLD, assuming maxLength allows this
			minLength: minTLDLengthAllowed,
			maxLength: maxTLDLengthAllowed,
		},

		// Test case with a valid TLD within the valid range
		"validTLDWithinRange": {
			tld: "com",
			minLength: minTLDLengthAllowed,
			maxLength: maxTLDLengthAllowed,
		},

		// Test case with a valid TLD exactly at a boundary condition
		"validTLDAtBoundaryCondition": {
			tld: "xyz", // Example TLD, ensure this is a valid TLD in TLDs map
			minLength: 3,
			maxLength: 3,
		},

		// Test case with a valid TLD at the upper boundary of length
		"validTLDAtUpperBoundary": {
			tld: "museum", // Example TLD, assuming it’s a valid TLD
			minLength: minTLDLengthAllowed,
			maxLength: 7,
		},

		// Test case with a valid TLD with exactly the allowed maximum length
		"validTLDExactlyMaxLength": {
			tld: "technology", // Example TLD, assuming it’s a valid TLD
			minLength: minTLDLengthAllowed,
			maxLength: 10,
		},

		// Test case with a valid TLD within allowed range with varying lengths
		"validTLDWithVariedLength": {
			tld: "info",
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
		tld        string
		minLength  uint
		maxLength  uint
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


func FuzzCCTLDSuccessfulForValidPseudorandomCCTLD(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed1, seed2 uint64) {
		r1 := rand.New(rand.NewPCG(seed1, seed2))
		tld1 := pseudorandom.CCTLD(r1)
		err := validate.CCTLD(tld1)
		if err != nil {
			t.Fatalf("expected no error for valid pseudo-random ccTLD \"%s\", but got error: %v", tld1, err)
		}
		r2 := rand.New(rand.NewPCG(seed1, seed2))
		tld2 := pseudorandom.CCTLD(r2)
		if tld1 != tld2 {
			t.Fatal("not deterministic")
		}
	})
}

func TestCCTLDSuccessfulForValidTLD(t *testing.T) {
	t.Parallel()

	testCases := map[string]string{
		// Example of a valid ccTLD from the list
		"validCCtldForUnitedKingdom": "uk",  // Assuming "uk" is in Countries

		// Example of another valid ccTLD
		"validCCtldForGermany": "de",  // Assuming "de" is in Countries

		// Example of a valid ccTLD for a large country
		"validCCtldForBrazil": "br",  // Assuming "br" is in Countries

		// Example of a valid ccTLD for a small country
		"validCCtldForSanMarino": "sm",  // Assuming "sm" is in Countries
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
