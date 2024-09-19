package random

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

const (
	minEmailLocalPartLengthAllowed  uint = 1
	maxEmailLocalPartLengthAllowed  uint = 64
	minEmailDomainPartLengthAllowed uint = minDomainLengthAllowed
	maxEmailDomainPartLengthAllowed uint = maxDomainLengthAllowed
	minEmailLengthAllowed           uint = minEmailLocalPartLengthAllowed + 1 + minEmailDomainPartLengthAllowed // 1 for @
	maxEmailLengthAllowed           uint = maxEmailLocalPartLengthAllowed + 1 + maxEmailDomainPartLengthAllowed
	minEmailDomainPartIPLength uint = 4 + 3 + 2 // 4 for numbers, 3 for dots, 2 for brackets
	maxEmailDomainPartIPLength uint = 32 + 7 + 7 // 32 for hexadecimal numbers, 7 for brackets, 7 for [IPv6:]
)

// Email generates a deterministic pseudo-random valid email address using the provided random source.
// The local part of the email can be either quoted or unquoted, and the domain
// part can be a regular domain name, an IPv4 address, or an IPv6 address.
// The email format follows the pattern: localPart@domain.
func Email(minLength, maxLength uint, quotedLocalPart, ipDomainPart bool) (email string, err error) {
	length, err := checkLength(minLength, maxLength, minEmailLengthAllowed, maxEmailLengthAllowed)
	if err != nil {
		return "", err
	}

	var localPartLength, domainPartLength uint
	var localPart, domain string

	maxLocalPartLength := length - minEmailDomainPartLengthAllowed - 1 // 1 for @
	if maxLocalPartLength > maxEmailLocalPartLengthAllowed {
		maxLocalPartLength = maxEmailLocalPartLengthAllowed
	}

	if ipDomainPart {
		maxLengthAllowed := maxEmailLocalPartLengthAllowed + minEmailDomainPartIPLength + 1
		minLengthAllowed := minEmailLocalPartLengthAllowed + maxEmailDomainPartIPLength + 1
		if maxLength > maxLengthAllowed {
			return "", fmt.Errorf("can not create an E-mail with IP as domain part when maximum length is more than %d", maxLengthAllowed)
		}
		if minLength < minLengthAllowed {
			return "", fmt.Errorf("can not create an E-mail with IP as domain part when minimum length is less than %d", minLengthAllowed)
		}

		random1, err := rand.Int(rand.Reader, big.NewInt(int64(2)))
		if err != nil {
			return "", fmt.Errorf("error generating a random number for chance of domain part to be IPv4 or IPv6: %w", err)
		}
		switch random1.Int64() {
		case 0:
			ipv4, err := IPv4("")
			if err != nil {
				return "", fmt.Errorf("error generating a random IPv4 as E-mail domain part: %w", err)
			}
			domain = fmt.Sprintf("[%s]", ipv4.String())
		case 1:
			ipv6, err := IPv4("")
			if err != nil {
				return "", fmt.Errorf("error generating a random IPv6 as E-mail domain part: %w", err)
			}
			domain = fmt.Sprintf("[IPv6:%s]", ipv6.String())
		}

		domainPartLength = uint(len(domain))

		localPartLength = length - domainPartLength - 1 // 1 for @
	} else {
		random2, err := rand.Int(rand.Reader, big.NewInt(int64(maxLocalPartLength - minEmailLocalPartLengthAllowed+1)))
		if err != nil {
			return "", fmt.Errorf("error generating a random number for calculating local part length: %w", err)
		}

		localPartLength = minEmailLocalPartLengthAllowed + uint(random2.Int64())

		domainPartLength = length - localPartLength - 1 // 1 for @

		if domainPartLength > maxDomainLengthAllowed {
			domainPartLength = maxDomainLengthAllowed
			localPartLength = length - domainPartLength - 1 // 1 for @
		}

		domain, err = Domain(domainPartLength, domainPartLength)
		if err != nil {
			return "", fmt.Errorf("error generating a pseudo-random domain as E-mail domain part: %w", err)
		}
	}

	if quotedLocalPart {
		localPart, err = randomEmailLocalPartQuoted(localPartLength)
		if err != nil {
			return "", fmt.Errorf("error generating a random quoted local part: %w", err)
		}
	} else {
		localPart, err = randomEmailLocalPartUnquoted(localPartLength)
		if err != nil {
			return "", fmt.Errorf("error generating a random unquoted local part: %w", err)
		}
	}

	parts := []string{localPart, domain}
	return strings.Join(parts, "@"), nil
}

// pseudorandomEmailLocalPartUnquoted generates a deterministic pseudo-random unquoted local part for an email address.
// The local part will consist of alphanumeric and printable characters, with a length between 1 and 64.
func randomEmailLocalPartUnquoted(length uint) (string, error) {
	// Create a slice of runes to store the local part characters.
	localPart := make([]rune, int(length))

	allowedCharactersForUnquotedWithoutDot := append(alphanumericalRunes, []rune("!#$%&'*+-/=?^_`{|}~")...)
	allowedCharactersForUnquoted := append(alphanumericalRunes, []rune("!#$%&'*+-./=?^_`{|}~")...)

	// The first character must be alphanumeric.
	random1, err := rand.Int(rand.Reader, big.NewInt(int64(len(allowedCharactersForUnquotedWithoutDot))))
	if err != nil {
		return "", fmt.Errorf("error generating a random index for the first character of E-mail unquoted local part: %w", err)
	}
	localPart[0] = allowedCharactersForUnquotedWithoutDot[random1.Int64()]

	// Label to retry generation if we encounter invalid patterns like "..".
generate:
	for i := 1; i < int(length)-1; i++ {
		random2, err := rand.Int(rand.Reader, big.NewInt(int64(len(allowedCharactersForUnquoted))))
		if err != nil {
			return "", fmt.Errorf("error generating a random index for middle character of E-mail unquoted local part: %w", err)
		}
		// Randomly select printable characters for the middle characters of the local part.
		localPart[i] = allowedCharactersForUnquoted[random2.Int64()]
	}

	random3, err := rand.Int(rand.Reader, big.NewInt(int64(len(allowedCharactersForUnquotedWithoutDot))))
		if err != nil {
			return "", fmt.Errorf("error generating a random index for the last character of E-mail unquoted local part: %w", err)
		}
	// The last character must be alphanumeric.
	localPart[length-1] = allowedCharactersForUnquotedWithoutDot[random3.Int64()]

	// If the generated local part contains "..", regenerate the local part.
	if strings.Contains(string(localPart), "..") {
		goto generate // Regenerate if ".." is found.
	}

	// Convert the rune slice back to a string and return the local part.
	return string(localPart), nil
}

// pseudorandomEmailLocalPartQuoted generates a deterministic pseudo-random quoted local part for an email address.
func randomEmailLocalPartQuoted(length uint) (string, error) {
	// Create a slice of runes to store the local part characters.
	localPart := make([]rune, length)

	allowedCharacters := append(printableRunes, ' ')

	// The first and last character must be double quote.
	localPart[0] = '"'
	localPart[len(localPart)-1] = '"'

	for i := 1; i < int(length)-1; i++ {
		random1, err := rand.Int(rand.Reader, big.NewInt(int64(len(allowedCharacters))))
		if err != nil {
			return "", fmt.Errorf("error generating a random index for a character of E-mail quoted local part: %w", err)
		}
		// Randomly select printable characters for the middle characters of the local part.
		localPart[i] = allowedCharacters[random1.Int64()]
	}

	// Convert the rune slice back to a string and return the local part.
	return string(localPart), nil
}
