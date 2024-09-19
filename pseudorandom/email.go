package pseudorandom

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

const (
	minEmailLengthAllowed           uint = minEmailLocalPartLengthAllowed + 1 + minEmailDomainPartLengthAllowed // 1 for @
	maxEmailLengthAllowed           uint = maxEmailLocalPartLengthAllowed + 1 + maxEmailDomainPartLengthAllowed
	minEmailLocalPartLengthAllowed  uint = 1
	maxEmailLocalPartLengthAllowed  uint = 64
	minEmailDomainPartLengthAllowed uint = minDomainLengthAllowed
	maxEmailDomainPartLengthAllowed uint = maxDomainLengthAllowed
	minEmailDomainPartIPLength uint = 4 + 3 + 2 // 4 for numbers, 3 for dots, 2 for brackets
	maxEmailDomainPartIPLength uint = 32 + 7 + 7 // 32 for hexadecimal numbers, 7 for brackets, 7 for [IPv6:]
)

// Email generates a deterministic pseudo-random valid email address using the provided random source.
// The local part of the email can be either quoted or unquoted, and the domain
// part can be a regular domain name, an IPv4 address, or an IPv6 address.
// The email format follows the pattern: localPart@domain.
func Email(r *rand.Rand, minLength, maxLength uint, quotedLocalPart, ipDomainPart bool) (email string, err error) {
	length, err := checkLength(r, minLength, maxLength, minEmailLengthAllowed, maxEmailLengthAllowed)
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

		switch random := r.IntN(2); random {
		case 0:
			ipv4, err := IPv4(r, "")
			if err != nil {
				return "", fmt.Errorf("error generating a pseudo-random IPv4 as E-mail domain part: %w", err)
			}
			domain = fmt.Sprintf("[%s]", ipv4.String())
		case 1:
			ipv6, err := IPv4(r, "")
			if err != nil {
				return "", fmt.Errorf("error generating a pseudo-random IPv6 as E-mail domain part: %w", err)
			}
			domain = fmt.Sprintf("[IPv6:%s]", ipv6.String())
		}

		domainPartLength = uint(len(domain))

		localPartLength = length - domainPartLength - 1 // 1 for @
	} else {
		localPartLength = minEmailLocalPartLengthAllowed + r.UintN(maxLocalPartLength - minEmailLocalPartLengthAllowed+1)

		domainPartLength = length - localPartLength - 1 // 1 for @

		if domainPartLength > maxDomainLengthAllowed {
			domainPartLength = maxDomainLengthAllowed
			localPartLength = length - domainPartLength - 1 // 1 for @
		}

		domain, err = Domain(r, domainPartLength, domainPartLength)
		if err != nil {
			return "", fmt.Errorf("error generating a pseudo-random domain as E-mail domain part: %w", err)
		}
	}

	if quotedLocalPart {
		localPart = pseudorandomEmailLocalPartQuoted(r, localPartLength)
	} else {
		localPart = pseudorandomEmailLocalPartUnquoted(r, localPartLength)
	}

	parts := []string{localPart, domain}
	return strings.Join(parts, "@"), nil
}

// pseudorandomEmailLocalPartUnquoted generates a deterministic pseudo-random unquoted local part for an email address.
// The local part will consist of alphanumeric and printable characters, with a length between 1 and 64.
func pseudorandomEmailLocalPartUnquoted(r *rand.Rand, length uint) string {
	// Create a slice of runes to store the local part characters.
	localPart := make([]rune, int(length))

	allowedCharactersForUnquotedWithoutDot := append(alphanumericalRunes, []rune("!#$%&'*+-/=?^_`{|}~")...)
	allowedCharactersForUnquoted := append(alphanumericalRunes, []rune("!#$%&'*+-./=?^_`{|}~")...)

	// The first character must be alphanumeric.
	localPart[0] = allowedCharactersForUnquotedWithoutDot[r.IntN(len(allowedCharactersForUnquotedWithoutDot))]

	// Label to retry generation if we encounter invalid patterns like "..".
generate:
	for i := 1; i < int(length)-1; i++ {
		// Randomly select printable characters for the middle characters of the local part.
		localPart[i] = allowedCharactersForUnquoted[r.IntN(len(allowedCharactersForUnquoted))]
	}

	// The last character must be alphanumeric.
	localPart[length-1] = allowedCharactersForUnquotedWithoutDot[r.IntN(len(allowedCharactersForUnquotedWithoutDot))]

	// If the generated local part contains "..", regenerate the local part.
	if strings.Contains(string(localPart), "..") {
		goto generate // Regenerate if ".." is found.
	}

	// Convert the rune slice back to a string and return the local part.
	return string(localPart)
}

// pseudorandomEmailLocalPartQuoted generates a deterministic pseudo-random quoted local part for an email address.
func pseudorandomEmailLocalPartQuoted(r *rand.Rand, length uint) string {
	// Create a slice of runes to store the local part characters.
	localPart := make([]rune, length)

	allowedCharacters := append(printableRunes, ' ')

	// The first and last character must be double quote.
	localPart[0] = '"'
	localPart[len(localPart)-1] = '"'

	for i := 1; i < int(length)-1; i++ {
		// Randomly select printable characters for the middle characters of the local part.
		localPart[i] = allowedCharacters[r.IntN(len(allowedCharacters))]
	}

	// Convert the rune slice back to a string and return the local part.
	return string(localPart)
}
