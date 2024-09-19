package validate

import (
	"errors"
	"fmt"
	"strings"
)

const (
	minLinuxHostnameLengthAllowed uint = 1
	maxLinuxHostnameLengthAllowed uint = 64
)

// LinuxHostname validates if the provided hostname is a valid DNS hostname for Linux systems.
//
// The hostname must:
//   - Not be empty.
//   - Be at most 64 characters long.
//   - Contain only lower-case ASCII letters, digits, or hyphens.
//   - Not start with a number or hyphen or end with a hyphen.
//
// If the hostname is invalid, an appropriate error is returned.
//
// A Linux hostname should be composed of up to 64 7-bit ASCII lower-case alphanumeric characters or hyphens forming a valid DNS domain name. It is recommended that this name contains only a single label, i.e. without any dots. https://www.freedesktop.org/software/systemd/man/latest/hostname.html.
func LinuxHostname(hostname string, minLength, maxLength uint) error {
	err := checkLength(len([]rune(hostname)), minLength, maxLength, minLinuxHostnameLengthAllowed, maxLinuxHostnameLengthAllowed, "characters")
	if err != nil {
		return err
	}

	var firstChar rune = []rune(hostname)[0]

	// Hostname cannot start with a hyphen or number.
	if !strings.ContainsAny(string(lowerCaseRunes), string(firstChar)) {
		return fmt.Errorf("hostname '%s' starts with %s", hostname, string([]rune(hostname)[0]))
	}

	// Hostname cannot end with a hyphen.
	if strings.HasSuffix(hostname, "-") {
		return fmt.Errorf("hostname '%s' ends with a hyphen", hostname)
	}

	var charErrs []error

	allowedCharacters := append(lowerAlphanumericalRunes, '-')

	// Iterate over each character in the hostname to ensure it is valid.
	for i, char := range hostname {
		if !strings.ContainsAny(string(allowedCharacters), string(char)) {
			charErrs = append(charErrs, fmt.Errorf("hostname '%s' contains invalid character '%c' at index %d", hostname, char, i))
		}
	}

	// If there are any character errors, return them joined together.
	if len(charErrs) > 0 {
		return fmt.Errorf("invalid hostname: %w", errors.Join(charErrs...))
	}

	// Return nil if the hostname is valid.
	return nil
}
