package validate

import (
	"errors"
	"fmt"
	"net"
	"strings"
	"unicode"
)

const (
	minEmailLocalPartLengthAllowed  uint = 1
	maxEmailLocalPartLengthAllowed  uint = 64
	minEmailDomainPartLengthAllowed uint = minDomainLengthAllowed
	maxEmailDomainPartLengthAllowed uint = maxDomainLengthAllowed
	minEmailLengthAllowed           uint = minEmailLocalPartLengthAllowed + 1 + minEmailDomainPartLengthAllowed // 1 for @
	maxEmailLengthAllowed           uint = maxEmailLocalPartLengthAllowed + 1 + maxEmailDomainPartLengthAllowed
)

// Email validates if the provided email address has a valid format.
// It checks if the email has a valid local part and a valid domain or IP address in the domain part.
// If any part is invalid, it returns a detailed error.
func Email(email string, minLength, maxLength uint, quotedLocalPart, ipDomainPart bool) error {
	err := checkLength(len([]rune(email)), minLength, maxLength, minEmailLengthAllowed, maxEmailLengthAllowed, "characters")
	if err != nil {
		return err
	}
	// Split the email into local part and domain by the "@" symbol.
	parts := strings.Split(email, "@")
	numParts := len(parts)

	// If the email does not contain exactly one '@', return an error.
	if numParts < 2 {
		return errors.New("invalid email format: missing '@' symbol")
	}

	// Extract the local part and domain
	localPart := parts[0]
	domain := parts[numParts-1]

	// If there are more than two parts, check if the local part is quoted.
	if numParts > 2 {
		joinedLocalParts := strings.Join(parts[:numParts-1], "@")
		if quotedLocalPart {
			localPart = joinedLocalParts // Quoted local parts can contain '@', so we re-join it.
		} else {
			return fmt.Errorf("invalid email format: multiple '@' symbols not allowed unless quoted, local part: %s", joinedLocalParts)
		}
	}

	var errs []error // Collect multiple validation errors for better debugging.

	// Validate the local part of the email.
	if err := validateEmailLocalPart(localPart, minEmailLocalPartLengthAllowed, maxEmailLocalPartLengthAllowed, quotedLocalPart); err != nil {
		errs = append(errs, fmt.Errorf("invalid local part: %w", err))
	}

	// Check if the domain part is an IP address enclosed in square brackets.
	if ipDomainPart {
		// Parse and validate the IP address within the domain.
		ip := parseEmailDomainIP(domain)
		if ip == nil {
			errs = append(errs, fmt.Errorf("invalid IP address in domain part: %s", domain))
		}
	} else {
		// Validate the domain name.
		if err := Domain(domain, minDomainLengthAllowed, maxDomainLengthAllowed); err != nil {
			errs = append(errs, fmt.Errorf("invalid domain part: %w", err))
		}
	}

	// If there are any accumulated errors, return them as a single error.
	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}

// validateEmailLocalPart checks if the provided local part of an email address is valid.
// It ensures the local part follows these rules:
//  1. It must not be empty.
//  2. It must not exceed 64 characters.
//  3. If unquoted, it must not start or end with a dot, contain spaces, consecutive dots, or special characters like "\"(),:;<>@[\\]".
//  4. All characters must be printable.
func validateEmailLocalPart(localPart string, minLength, maxLength uint, quotedLocalPart bool) error {
	err := checkLength(len(localPart), minLength, maxLength, minEmailLocalPartLengthAllowed, maxEmailLocalPartLengthAllowed, "characters")
	if err != nil {
		return err
	}

	// If the local part is quoted, we only check if all characters are printable.
	if quotedLocalPart {
		for i, char := range localPart {
			if !unicode.IsPrint(char) {
				return fmt.Errorf("quoted local part contains non-printable character '%c' at index %d", char, i)
			}
		}
		return nil
	}

	// For unquoted local parts, we apply more strict rules:
	// Rule 3: Unquoted local part must not start or end with a dot.
	if strings.HasPrefix(localPart, ".") || strings.HasSuffix(localPart, ".") {
		return errors.New("unquoted local part cannot start or end with a dot")
	}

	// Rule 4: Unquoted local part must not contain spaces.
	if strings.Contains(localPart, " ") {
		return errors.New("unquoted local part cannot contain spaces")
	}

	// Rule 5: Unquoted local part must not contain consecutive dots.
	if strings.Contains(localPart, "..") {
		return errors.New("unquoted local part cannot contain consecutive dots")
	}

	// Rule 6: Unquoted local part must not contain special characters that are only allowed in quotes.
	if strings.ContainsAny(localPart, "\"(),:;<>@[\\]") {
		return errors.New("unquoted local part cannot contain special characters: \"(),:;<>@[\\]")
	}

	// Rule 7: Ensure all characters in the unquoted local part are printable.
	for i, char := range localPart {
		if !unicode.IsPrint(char) {
			return fmt.Errorf("unquoted local part contains non-printable character '%c' at index %d", char, i)
		}
	}

	return nil
}

// parseEmailDomainIP extracts and parses the IP address from a domain that is enclosed in square brackets,
// as is allowed for email domains. If the domain is in IPv6 format, it can also handle the optional "IPv6:" prefix.
// The function returns a valid net.IP object if parsing is successful or nil if the domain is not a valid IP address.
func parseEmailDomainIP(domain string) net.IP {
	// Remove the square brackets that enclose the domain.
	domain = strings.Trim(domain, "[]")

	// Check for IPv6 prefix and trim it if present.
	domain = strings.TrimPrefix(domain, "IPv6:")

	// Attempt to parse the domain as an IP address.
	return net.ParseIP(domain)
}
