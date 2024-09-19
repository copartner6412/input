package validate_test

import (
	"testing"

	"github.com/copartner6412/input/validate"
)

const (
	minPINLengthAllowed = 3
	maxPINLengthAllowed = 32
)

func TestPINSuccessfulForValidInput(t *testing.T) {
	testCases := map[string]struct {
		pin       string
		minLength uint
		maxLength uint
	}{
		// Edge case: minimal length, valid single-digit PIN.
		"valid_min_length": {
			pin:       "123",
			minLength: 3,
			maxLength: 3,
		},
		// Edge case: maximum length, valid PIN with max characters.
		"valid_max_length": {
			pin:       "12345678901234567890123456789012", // 32 digits
			minLength: 32,
			maxLength: 32,
		},
		// Valid PIN with length between minLength and maxLength.
		"valid_intermediate_length": {
			pin:       "12345",
			minLength: 3,
			maxLength: 6,
		},
		// Edge case: minimal and maximal lengths are equal, valid PIN with same length.
		"valid_min_max_equal": {
			pin:       "7890",
			minLength: 4,
			maxLength: 4,
		},
		// Valid PIN exactly at the minimum length limit.
		"valid_exact_min_length": {
			pin:       "456",
			minLength: 3,
			maxLength: 5,
		},
		// Valid PIN exactly at the maximum length limit.
		"valid_exact_max_length": {
			pin:       "123456",
			minLength: 4,
			maxLength: 6,
		},
		// Valid PIN with larger allowable range between min and max lengths.
		"valid_large_range": {
			pin:       "1234567",
			minLength: 3,
			maxLength: 32,
		},
		// Edge case: minimum PIN length.
		"valid_min_PIN_length": {
			pin:       "123",
			minLength: minPINLengthAllowed,
			maxLength: 6,
		},
		// Edge case: maximum PIN length.
		"valid_max_PIN_length": {
			pin:       "12345678901234567890123456789012", // 32 digits
			minLength: 3,
			maxLength: maxPINLengthAllowed,
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			err := validate.PIN(testCase.pin, testCase.minLength, testCase.maxLength)
			if err != nil {
				t.Errorf("expected no error for valid PIN %s, but got error: %v", testCase.pin, err)
			}
		})
	}
}

func TestPINFailsForInvalidInput(t *testing.T) {
	testCases := map[string]struct {
		pin       string
		minLength uint
		maxLength uint
	}{
		// Case where maxLength is less than minLength.
		"max_length_less_than_min_length": {
			pin:       "123",
			minLength: 5,
			maxLength: 3,
		},
		// Case where minLength is less than the minimum allowed PIN length.
		"min_length_less_than_allowed": {
			pin:       "123",
			minLength: 2, // minPINLength is 3
			maxLength: 5,
		},
		// Case where maxLength exceeds the maximum allowed PIN length.
		"max_length_exceeds_allowed": {
			pin:       "123456789012345678901234567890123", // 33 digits
			minLength: 3,
			maxLength: 33, // maxPINLength is 32
		},
		// Case where the PIN is shorter than the minLength.
		"pin_too_short": {
			pin:       "12", // 2 digits
			minLength: 3,
			maxLength: 6,
		},
		// Case where the PIN exceeds the maxLength.
		"pin_too_long": {
			pin:       "123456789", // 9 digits
			minLength: 3,
			maxLength: 8,
		},
		// Case where the PIN contains non-digit characters.
		"pin_contains_non_digit_characters": {
			pin:       "12a45", // 'a' is not a digit
			minLength: 3,
			maxLength: 6,
		},
		// Case where the PIN contains whitespace.
		"pin_contains_whitespace": {
			pin:       "123 45", // contains space
			minLength: 3,
			maxLength: 6,
		},
		// Case where the PIN is empty.
		"empty_pin": {
			pin:       "", // empty string
			minLength: 3,
			maxLength: 6,
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			err := validate.PIN(testCase.pin, testCase.minLength, testCase.maxLength)
			if err == nil {
				t.Errorf("expected error for invalid PIN %s, but got no error", testCase.pin)
			}
		})
	}
}
