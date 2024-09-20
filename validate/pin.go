package validate

import (
	"fmt"
	"strings"
)

// Constants for PIN length limits.
const (
	minPINLengthAllowed = 3
	maxPINLengthAllowed = 32
)

// PIN validates a given Personal Identification Number (PIN) string to ensure it meets the specified
// length and character requirements.
//
// Parameters:
//   - pin: The PIN string to validate. It should contain only numeric characters (0-9).
//   - minLength: The minimum length that the PIN must meet. It must be between 3 and 32.
//   - maxLength: The maximum length that the PIN can have. It must be between 3 and 32.
//
// Returns:
//   - error: Returns an error if the PIN fails any of the following conditions:
//   - The PIN length is smaller than the specified minLength.
//   - The PIN Length exceeds the specified maxLength.
//   - The PIN contains non-numeric characters.
func PIN(pin string, minLength, maxLength uint) error {
	err := checkLength(len([]rune(pin)), minLength, maxLength, minPINLengthAllowed, maxPINLengthAllowed, "characters")
	if err != nil {
		return err
	}

	for i, char := range pin {
		if !strings.ContainsRune(string(digitRunes), char) {
			return fmt.Errorf("not-digit character %s at index %d", string(digitRunes), i)
		}
	}

	return nil
}
