package validate_test

import (
	"strings"
	"testing"

	"github.com/copartner6412/input/validate"
)

const (
	minEmailLocalPartLengthAllowed  uint = 1
	maxEmailLocalPartLengthAllowed  uint = 64
	minEmailDomainPartLengthAllowed uint = minDomainLengthAllowed
	maxEmailDomainPartLengthAllowed uint = maxDomainLengthAllowed
	minEmailLengthAllowed           uint = minEmailLocalPartLengthAllowed + 1 + minEmailDomainPartLengthAllowed // 1 for @
	maxEmailLengthAllowed           uint = maxEmailLocalPartLengthAllowed + 1 + maxEmailDomainPartLengthAllowed
)

func TestEmailSuccessfulForValidInput(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		email           string
		minLength       uint
		maxLength       uint
		quotedLocalPart bool
		ipDomainPart    bool
	}{
		// Valid: regular email with standard local and domain parts
		"regularEmail": {
			email:           "user@example.com",
			minLength:       minEmailLengthAllowed,
			maxLength:       maxEmailLengthAllowed,
			quotedLocalPart: false,
			ipDomainPart:    false,
		},

		// Valid: quoted local part with '@' symbol
		"quotedLocalPartWithAtSymbol": {
			email:           `"user@quoted"@example.com`,
			minLength:       minEmailLengthAllowed,
			maxLength:       maxEmailLengthAllowed,
			quotedLocalPart: true,
			ipDomainPart:    false,
		},

		// Valid: domain part as IP address enclosed in square brackets
		"ipAddressDomain": {
			email:           "user@[192.168.0.1]",
			minLength:       minEmailLengthAllowed,
			maxLength:       maxEmailLengthAllowed,
			quotedLocalPart: false,
			ipDomainPart:    true,
		},

		// Valid: local part with the maximum allowed length (64 characters)
		"localPartWithMaxLength": {
			email:           strings.Repeat("a", 64) + "@example.com",
			minLength:       minEmailLengthAllowed,
			maxLength:       maxEmailLengthAllowed,
			quotedLocalPart: false,
			ipDomainPart:    false,
		},

		// Valid: minimal length email with shortest valid local and domain part
		"minimalLengthEmail": {
			email:           "a@b.co",
			minLength:       minEmailLengthAllowed,
			maxLength:       maxEmailLengthAllowed,
			quotedLocalPart: false,
			ipDomainPart:    false,
		},

		// Valid: email with subdomains in the domain part
		"emailWithSubdomains": {
			email:           "user@mail.sub.example.com",
			minLength:       minEmailLengthAllowed,
			maxLength:       maxEmailLengthAllowed,
			quotedLocalPart: false,
			ipDomainPart:    false,
		},

		// Valid: email with special characters in the local part
		"localPartWithSpecialCharacters": {
			email:           "user.name+alias@example.com",
			minLength:       minEmailLengthAllowed,
			maxLength:       maxEmailLengthAllowed,
			quotedLocalPart: false,
			ipDomainPart:    false,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			err := validate.Email(tc.email, tc.minLength, tc.maxLength, tc.quotedLocalPart, tc.ipDomainPart)
			if err != nil {
				t.Errorf("expected no error for valid input %q, but got error: %v", tc.email, err)
			}
		})
	}
}


func TestEmailFailsForInvalidInput(t *testing.T) {
    t.Parallel()

    testCases := map[string]struct{
        email          string
        minLength      uint
        maxLength      uint
        quotedLocalPart bool
        ipDomainPart   bool
    }{
        // Empty email string
        "emptyEmailString": {
            email: "",
            minLength: minEmailLengthAllowed,  // min length requires at least "x@y"
            maxLength: maxEmailLengthAllowed,
            quotedLocalPart: false,
            ipDomainPart: false,
        },

        // Email shorter than minLength
        "emailShorterThanMinLength": {
            email: "a@b.com",  // 7 characters
            minLength: 8,
            maxLength: maxEmailLengthAllowed,
            quotedLocalPart: false,
            ipDomainPart: false,
        },

        // Email exceeding maxLength
        "emailExceedsMaxLength": {
            email: strings.Repeat("a", 65) + "@domain.com",  // exceeds max local part length of 64
            minLength: minEmailLengthAllowed,
            maxLength: maxEmailLengthAllowed,
            quotedLocalPart: false,
            ipDomainPart: false,
        },

        // maxLength is less than minLength
        "maxLengthLessThanMinLength": {
            email: "user@domain.com",
            minLength: 10,
            maxLength: 5,
            quotedLocalPart: false,
            ipDomainPart: false,
        },


        "onePart": {
            email: "abc",
            minLength: minEmailLengthAllowed,
            maxLength: maxEmailLengthAllowed,
            quotedLocalPart: false,
            ipDomainPart: false,
        },

        // Invalid format: multiple "@" symbols without quoted local part
        "multipleAtSymbolsWithoutQuotedLocalPart": {
            email: "user@domain@com",
            minLength: minEmailLengthAllowed,
            maxLength: maxEmailLengthAllowed,
            quotedLocalPart: false,
            ipDomainPart: false,
        },

        // Invalid characters in unquoted local part
        "unquotedLocalPartWithInvalidCharacters": {
            email: "user(name)@domain.com",
            minLength: minEmailLengthAllowed,
            maxLength: maxEmailLengthAllowed,
            quotedLocalPart: false,
            ipDomainPart: false,
        },

        // Quoted local part containing non-printable characters
        "quotedLocalPartWithNonPrintableCharacters": {
            email: "\"user\x00name\"@domain.com",
            minLength: minEmailLengthAllowed,
            maxLength: maxEmailLengthAllowed,
            quotedLocalPart: true,
            ipDomainPart: false,
        },

        // Unquoted local part starts with a dot
        "unquotedLocalPartStartsWithDot": {
            email: ".username@domain.com",
            minLength: minEmailLengthAllowed,
            maxLength: maxEmailLengthAllowed,
            quotedLocalPart: false,
            ipDomainPart: false,
        },

        // Unquoted local part ends with a dot
        "unquotedLocalPartEndsWithDot": {
            email: "username.@domain.com",
            minLength: minEmailLengthAllowed,
            maxLength: maxEmailLengthAllowed,
            quotedLocalPart: false,
            ipDomainPart: false,
        },

        // Unquoted local part contains consecutive dots
        "unquotedLocalPartWithConsecutiveDots": {
            email: "user..name@domain.com",
            minLength: minEmailLengthAllowed,
            maxLength: maxEmailLengthAllowed,
            quotedLocalPart: false,
            ipDomainPart: false,
        },

        // Domain part is an invalid IP address (for ipDomainPart)
        "domainPartWithInvalidIP": {
            email: "user@[999.999.999.999]",  // invalid IP address
            minLength: minEmailLengthAllowed,
            maxLength: maxEmailLengthAllowed,
            quotedLocalPart: false,
            ipDomainPart: true,
        },

        // Domain part contains invalid domain name characters
        "domainPartWithInvalidCharacters": {
            email: "user@domain[dot]com",  // invalid characters in domain
            minLength: minEmailLengthAllowed,
            maxLength: maxEmailLengthAllowed,
            quotedLocalPart: false,
            ipDomainPart: false,
        },

        // Domain part is missing (only local part and "@")
        "missingDomainPart": {
            email: "user@",
            minLength: minEmailLengthAllowed,
            maxLength: maxEmailLengthAllowed,
            quotedLocalPart: false,
            ipDomainPart: false,
        },

        // Local part contains spaces (invalid for unquoted local part)
        "unquotedLocalPartWithSpaces": {
            email: "user name@domain.com",
            minLength: minEmailLengthAllowed,
            maxLength: maxEmailLengthAllowed,
            quotedLocalPart: false,
            ipDomainPart: false,
        },
    }

    for name, tc := range testCases {
        t.Run(name, func(t *testing.T) {
            err := validate.Email(tc.email, tc.minLength, tc.maxLength, tc.quotedLocalPart, tc.ipDomainPart)
            if err == nil {
                t.Errorf("expected error for invalid input %q, but got none", tc.email)
            }
        })
    }
}


