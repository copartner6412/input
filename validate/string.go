package validate

import (
	"errors"
	"fmt"
	"unicode"
)

const (
	minStringLengthAllowed = 1    // Minimum acceptable string length
	maxStringLengthAllowed = 8192 // Maximum acceptable string length
	asciiLowerBound = 32   // Start of printable ASCII range (space)
	asciiUpperBound = 126  // End of printable ASCII range (tilde ~)
)

// String validates that a given string meets certain criteria related to length, character set (ASCII or Unicode).
//
// Parameters:
//   - str: The input string to be validated. Must not be empty.
//   - minLength: The minimum acceptable length of the string in rune characters between 1 and 8192.
//   - maxLength: The maximum acceptable length of the string in rune characters between 1 and 8192.
//   - justASCII: A boolean flag that determines whether only ASCII characters are allowed.
//     If true, the string must only contain characters from the printable ASCII range (32-126).
//
// Returns:
//   - An error if the string fails to meet any of the criteria.
//     The error will indicate if the string is too short, too long, contains non-ASCII
//     characters when `justASCII` is true, or contains spaces when `hasSpace` is false.
func String(str string, minLength, maxLength uint, justASCII bool) error {
	err := checkLength(len([]rune(str)), minLength, maxLength, minStringLengthAllowed, maxStringLengthAllowed, "characters")
	if err != nil {
		return err
	}
	
	// Check if the string contains only ASCII characters, if required.
	var unicodeErrs []error
	if justASCII {
		for i, char := range str {
			if char < asciiLowerBound || char > asciiUpperBound {
				unicodeErrs = append(unicodeErrs, fmt.Errorf("non-ASCII character \"%s\" at index %d", string(char), i))
			}
		}
		if len(unicodeErrs) > 0 {
			return errors.Join(unicodeErrs[0], fmt.Errorf("%d non-ASCII characters in the string", len(unicodeErrs)))
		}

		// Check for non-printable Unicode characters, if justASCII is false.
	} else {
		for i, char := range str {
			if !unicode.IsPrint(char) {
				unicodeErrs = append(unicodeErrs, fmt.Errorf("non-printable Unicode character at index %d", i))
			}
		}
		if len(unicodeErrs) > 0 {
			return errors.Join(unicodeErrs[0], fmt.Errorf("%d non-printable Unicode characters in the string", len(unicodeErrs)))
		}
	}

	return nil
}
