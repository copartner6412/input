package validate_test

import (
	"strings"
	"testing"

	"github.com/copartner6412/input/validate"
)

const (
	minDomainLengthAllowed   uint = 1
	maxDomainLengthAllowed   uint = 253
	minTLDLengthAllowed uint = 2
	maxTLDLengthAllowed uint = 16
	ccTLDLength uint = 2
	minDomainWithValidTLDLengthAllowed uint = minTLDLengthAllowed + 2
	minDomainWithValidCCTLDLengthAllowed uint = ccTLDLength + 2
)

func TestDomainSuccessfulForValidInput(t *testing.T) {
    t.Parallel()

    testCases := map[string]struct{
        domain    string
        minLength uint
        maxLength uint
    }{
        // 1. Test case: domain at minimum allowed length
        "domainAtMinLength": {
            domain: "a.com",   // domain of length 5 (minimum domain with valid ccTLD)
            minLength: minDomainLengthAllowed,
            maxLength: maxDomainLengthAllowed,
        },
        
        // 2. Test case: domain at maximum allowed length
        "domainAtMaxLength": {
            domain: "aa" + strings.Repeat("a.", 124) + "com",  // domain of length 253
            minLength: minDomainLengthAllowed,
            maxLength: maxDomainLengthAllowed,
        },

        // 3. Test case: domain with multiple valid subdomains
        "multipleValidSubdomains": {
            domain: "valid.sub.example.com",
            minLength: minDomainLengthAllowed,
            maxLength: maxDomainLengthAllowed,
        },

        // 4. Test case: valid domain with hyphen in subdomain (but not at the start or end)
        "subdomainWithHyphen": {
            domain: "sub-domain.example.com",   // valid subdomain with hyphen in the middle
            minLength: minDomainLengthAllowed,
            maxLength: maxDomainLengthAllowed,
        },

        // 5. Test case: domain with numeric subdomain
        "numericSubdomain": {
            domain: "123.example.com",   // subdomain containing numbers
            minLength: minDomainLengthAllowed,
            maxLength: maxDomainLengthAllowed,
        },

        // 6. Test case: valid domain with a long subdomain (at the maximum allowed length)
        "maxLengthSubdomain": {
            domain: strings.Repeat("a", 63) + ".example.com",   // subdomain of length 63 (max allowed)
            minLength: minDomainLengthAllowed,
            maxLength: maxDomainLengthAllowed,
        },

        // 7. Test case: valid two-letter country code TLD
        "validCcTLD": {
            domain: "example.ir",   // valid domain with a country code TLD
            minLength: minDomainLengthAllowed,
            maxLength: maxDomainLengthAllowed,
        },

        // 8. Test case: domain with subdomain at minimum length
        "subdomainAtMinLength": {
            domain: "a.example.com",   // subdomain with 1 character (min allowed)
            minLength: minDomainLengthAllowed,
            maxLength: maxDomainLengthAllowed,
        },

        // 9. Test case: valid domain with UTF-8 characters (IDN)
        "validIDNDomain": {
            domain: "xn--fsq.com",  // Punycode for valid internationalized domain (like föß.com)
            minLength: minDomainLengthAllowed,
            maxLength: maxDomainLengthAllowed,
        },
    }

    for name, tc := range testCases {
        t.Run(name, func(t *testing.T) {
            err := validate.Domain(tc.domain, tc.minLength, tc.maxLength)
            if err != nil {
                t.Errorf("expected no error for valid input %q, but got error %v", tc.domain, err)
            }
        })
    }
}


func TestDomainFailsForInvalidInput(t *testing.T) {
    t.Parallel()

    testCases := map[string]struct{
        domain    string
        minLength uint
        maxLength uint
    }{
        // 1. Test case: domain length shorter than minLength
        "domainTooShort": {
            domain: "a",
            minLength: 5,   // minLength > domain length
            maxLength: 10,
        },
        
        // 2. Test case: domain length longer than maxLength
        "domainTooLong": {
            domain: "thisisaverylongdomainnamethatexceedslimit.com",
            minLength: 5,
            maxLength: 10,  // maxLength < domain length
        },

        // 3. Test case: minLength is less than minDomainLengthAllowed
        "minLengthTooSmall": {
            domain: "example.com",
            minLength: 0,    // less than minDomainLengthAllowed (1)
            maxLength: 50,
        },

        // 4. Test case: maxLength exceeds maxDomainLengthAllowed
        "maxLengthTooLarge": {
            domain: "example.com",
            minLength: minDomainLengthAllowed,
            maxLength: 500,  // exceeds maxDomainLengthAllowed (253)
        },

        // 5. Test case: maxLength is less than minLength
        "maxLengthLessThanMinLength": {
            domain: "example.com",
            minLength: 10,   // maxLength < minLength
            maxLength: 5,
        },

        // 6. Test case: domain contains invalid subdomain 'www' not in the first position
        "invalidSubdomainWww": {
            domain: "example.www.com",  // 'www' is not the first subdomain
            minLength: minDomainLengthAllowed,
            maxLength: maxDomainLengthAllowed,
        },

        // 7. Test case: empty domain input
        "emptyDomain": {
            domain: "",
            minLength: minDomainLengthAllowed,
            maxLength: maxDomainLengthAllowed,
        },

        // 8. Test case: domain with subdomain length less than minSubdomainLengthAllowed
        "subdomainTooShort": {
            domain: "e..com",  // first subdomain is too short (length 1)
            minLength: minDomainLengthAllowed,
            maxLength: maxDomainLengthAllowed,
        },

        // 9. Test case: domain with subdomain length greater than maxSubdomainLengthAllowed
        "subdomainTooLong": {
            domain: "thisisaverylongsubdomainthatexceedsallowedlengthwecontinuetomakeitlongerthan63words.com",
            minLength: minDomainLengthAllowed,
            maxLength: maxDomainLengthAllowed,
        },
    }

    for name, tc := range testCases {
        t.Run(name, func(t *testing.T) {
            err := validate.Domain(tc.domain, tc.minLength, tc.maxLength)
            if err == nil {
                t.Errorf("expected error for invalid input: %q, but got none", tc.domain)
            }
        })
    }
}

func TestDomainWithValidTLDSuccessfulForValidInput(t *testing.T) {
    testCases := map[string]struct{
        domain    string
        minLength uint
        maxLength uint
    }{
        // 1. Test case: valid domain with generic TLD
        "validDomainWithGTLD": {
            domain: "example.com",
            minLength: minDomainLengthAllowed,
            maxLength: maxDomainLengthAllowed,
        },

        // 2. Test case: valid domain with a country code TLD
        "validDomainWithCcTLD": {
            domain: "example.ir",   // valid .ir ccTLD, relevant to the project context
            minLength: minDomainLengthAllowed,
            maxLength: maxDomainLengthAllowed,
        },

        // 3. Test case: valid domain at minimum length
        "domainAtMinLength": {
            domain: "a.co",   // 4 characters (minimum domain with valid ccTLD)
            minLength: minDomainLengthAllowed,
            maxLength: maxDomainLengthAllowed,
        },

        // 4. Test case: valid domain with multiple subdomains
        "domainWithMultipleSubdomains": {
            domain: "sub.sub2.example.com",
            minLength: minDomainLengthAllowed,
            maxLength: maxDomainLengthAllowed,
        },

        // 5. Test case: valid domain with hyphen in subdomain
        "domainWithHyphenSubdomain": {
            domain: "sub-domain.example.com",   // valid hyphen usage in the subdomain
            minLength: minDomainLengthAllowed,
            maxLength: maxDomainLengthAllowed,
        },

        // 6. Test case: valid internationalized domain name (IDN)
        "validIDNDomain": {
            domain: "xn--fsq.com",   // Punycode for a valid IDN
            minLength: minDomainLengthAllowed,
            maxLength: maxDomainLengthAllowed,
        },

        // 7. Test case: valid domain with numeric subdomain
        "domainWithNumericSubdomain": {
            domain: "123.example.com",
            minLength: minDomainLengthAllowed,
            maxLength: maxDomainLengthAllowed,
        },
        
        // 8. Test case: valid domain with unusual but valid TLD
        "validDomainWithUnusualTLD": {
            domain: "example.xyz",   // valid, though less common TLD
            minLength: minDomainLengthAllowed,
            maxLength: maxDomainLengthAllowed,
        },
    }

    for name, tc := range testCases {
        t.Run(name, func(t *testing.T) {
            err := validate.DomainWithValidTLD(tc.domain, tc.minLength, tc.maxLength)
            if err != nil {
                t.Errorf("expected no error for a valid domain input \"%s\", but got error: %v", tc.domain, err)
            }
        })
    }
}


func TestDomainWithValidTLDFailsForInvalidInput(t *testing.T) {
    testCases := map[string]struct{
        domain    string
        minLength uint
        maxLength uint
    }{
        // 1. Test case: empty domain string (invalid length)
        "emptyDomainString": {
            domain: "",
            minLength: minDomainLengthAllowed,
            maxLength: maxDomainLengthAllowed,
        },

        // 2. Test case: domain shorter than minimum allowed length
        "domainShorterThanMinLength": {
            domain: "x.com",  // 5 characters, testing with higher minLength
            minLength: 6,
            maxLength: maxDomainLengthAllowed,
        },

        // 3. Test case: domain exceeding maximum allowed length
        "domainExceedsMaxLength": {
            domain: strings.Repeat("a", 254) + ".com",  // Exceeds 253 characters
            minLength: minDomainLengthAllowed,
            maxLength: maxDomainLengthAllowed,
        },

        // 4. Test case: maxLength is less than minLength
        "maxLengthLessThanMinLength": {
            domain: "example.com",
            minLength: 10,
            maxLength: 5,
        },

        // 5. Test case: invalid TLD (non-existent TLD)
        "invalidTLD": {
            domain: "example.invalidtld",
            minLength: minDomainLengthAllowed,
            maxLength: maxDomainLengthAllowed,
        },

        // 6. Test case: domain with invalid characters
        "domainWithInvalidCharacters": {
            domain: "example@domain.com",
            minLength: minDomainLengthAllowed,
            maxLength: maxDomainLengthAllowed,
        },

        // 7. Test case: domain without TLD (no dots)
        "domainWithoutTLD": {
            domain: "exampledomain",
            minLength: minDomainLengthAllowed,
            maxLength: maxDomainLengthAllowed,
        },

        // 8. Test case: TLD with numeric characters (invalid TLD format)
        "domainWithNumericTLD": {
            domain: "example.123",
            minLength: minDomainLengthAllowed,
            maxLength: maxDomainLengthAllowed,
        },

        // 9. Test case: domain with multiple dots but invalid TLD
        "domainWithMultipleDotsInvalidTLD": {
            domain: "sub.domain.notld",
            minLength: minDomainLengthAllowed,
            maxLength: maxDomainLengthAllowed,
        },

        // 10. Test case: domain without a subdomain, only TLD
        "domainWithoutSubdomain": {
            domain: ".com",
            minLength: minDomainLengthAllowed,
            maxLength: maxDomainLengthAllowed,
        },
    }

    for name, tc := range testCases {
        t.Run(name, func(t *testing.T) {
            err := validate.DomainWithValidTLD(tc.domain, tc.minLength, tc.maxLength)
            if err == nil {
                t.Errorf("expected error for domain with invalid input \"%s\", but got no error", tc.domain)
            }
        })
    }
}

func TestDomainWithValidCCTLDSuccessfulForValidInput(t *testing.T) {
    testCases := map[string]struct{
        domain    string
        minLength uint
        maxLength uint
    }{
        // 1. Test case: valid domain with .ir ccTLD
        "validDomainWithIRCcTLD": {
            domain: "example.ir",
            minLength: minDomainLengthAllowed,
            maxLength: maxDomainLengthAllowed,
        },

        // 2. Test case: valid domain with .us ccTLD
        "validDomainWithUSCcTLD": {
            domain: "example.us",
            minLength: minDomainLengthAllowed,
            maxLength: maxDomainLengthAllowed,
        },

        // 3. Test case: valid domain with .fr ccTLD (France)
        "validDomainWithFRCcTLD": {
            domain: "example.fr",
            minLength: minDomainLengthAllowed,
            maxLength: maxDomainLengthAllowed,
        },

        // 4. Test case: valid domain with minimum length and .uk ccTLD
        "domainAtMinLengthWithUKCcTLD": {
            domain: "a.co.uk",   // 6 characters with ccTLD .uk
            minLength: minDomainLengthAllowed,
            maxLength: maxDomainLengthAllowed,
        },

        // 5. Test case: valid domain at maximum length with .de ccTLD (Germany)
        "domainAtMaxLengthWithDECcTLD": {
            domain: "ab" + strings.Repeat("ab.", 83) + "de",   // 253 characters in total
            minLength: minDomainLengthAllowed,
            maxLength: maxDomainLengthAllowed,
        },

        // 6. Test case: valid domain with multiple subdomains and .jp ccTLD
        "domainWithMultipleSubdomainsWithJPCcTLD": {
            domain: "sub.example.co.jp",
            minLength: minDomainLengthAllowed,
            maxLength: maxDomainLengthAllowed,
        },

        // 7. Test case: valid internationalized domain name with .cn ccTLD (China)
        "validIDNDomainWithCNCcTLD": {
            domain: "xn--fsq.cn",   // Punycode for a valid IDN with .cn
            minLength: minDomainLengthAllowed,
            maxLength: maxDomainLengthAllowed,
        },

        // 8. Test case: valid domain with numeric subdomain and .ru ccTLD (Russia)
        "domainWithNumericSubdomainWithRUCcTLD": {
            domain: "123.example.ru",
            minLength: minDomainLengthAllowed,
            maxLength: maxDomainLengthAllowed,
        },

        // 9. Test case: valid domain with hyphenated subdomain and .au ccTLD (Australia)
        "domainWithHyphenSubdomainWithAUCcTLD": {
            domain: "sub-domain.example.au",
            minLength: minDomainLengthAllowed,
            maxLength: maxDomainLengthAllowed,
        },

        // 10. Test case: valid domain with unusual but valid .tv ccTLD (Tuvalu)
        "validDomainWithTVCcTLD": {
            domain: "example.tv",
            minLength: minDomainLengthAllowed,
            maxLength: maxDomainLengthAllowed,
        },
    }

    for name, tc := range testCases {
        t.Run(name, func(t *testing.T) {
            err := validate.DomainWithValidCCTLD(tc.domain, tc.minLength, tc.maxLength)
            if err != nil {
                t.Errorf("expected no error for a valid domain input \"%s\", but got error: %v", tc.domain, err)
            }
        })
    }
}


func TestDomainWithValidCCTLDFailsForInvalidInput(t *testing.T) {
    testCases := map[string]struct{
        domain    string
        minLength uint
        maxLength uint
    }{
        // 1. Test case: empty domain string (invalid length)
        "emptyDomainString": {
            domain: "",
            minLength: minDomainLengthAllowed,
            maxLength: maxDomainLengthAllowed,
        },

        // 2. Test case: domain shorter than minimum allowed length
        "domainShorterThanMinLength": {
            domain: "a.ir",  // 4 characters, but testing with higher minLength
            minLength: 5,
            maxLength: maxDomainLengthAllowed,
        },

        // 3. Test case: domain exceeding maximum allowed length
        "domainExceedsMaxLength": {
            domain: strings.Repeat("a", 254) + ".com",  // Exceeds 253 characters
            minLength: minDomainLengthAllowed,
            maxLength: maxDomainLengthAllowed,
        },

        // 4. Test case: maxLength is less than minLength
        "maxLengthLessThanMinLength": {
            domain: "example.com",
            minLength: 10,
            maxLength: 5,
        },

        // 5. Test case: invalid ccTLD (non-existent TLD)
        "invalidCCTLD": {
            domain: "example.invalidtld",
            minLength: minDomainLengthAllowed,
            maxLength: maxDomainLengthAllowed,
        },

        // 6. Test case: invalid characters in domain
        "domainWithInvalidCharacters": {
            domain: "example@domain.com",
            minLength: minDomainLengthAllowed,
            maxLength: maxDomainLengthAllowed,
        },

        // 7. Test case: domain without TLD (no dots)
        "domainWithoutTLD": {
            domain: "exampledomain",
            minLength: minDomainLengthAllowed,
            maxLength: maxDomainLengthAllowed,
        },

        // 8. Test case: domain with numbers in TLD (invalid TLD format)
        "domainWithNumericTLD": {
            domain: "example.123",
            minLength: minDomainLengthAllowed,
            maxLength: maxDomainLengthAllowed,
        },

        // 9. Test case: domain with only TLD, no subdomain
        "domainWithOnlyTLD": {
            domain: ".com",
            minLength: minDomainLengthAllowed,
            maxLength: maxDomainLengthAllowed,
        },

        // 10. Test case: domain with subdomain exceeding maximum length
        "subdomainExceedsMaxLength": {
            domain: strings.Repeat("a", 250) + ".com",  // Subdomain itself is too long
            minLength: minDomainLengthAllowed,
            maxLength: maxDomainLengthAllowed,
        },
    }

    for name, tc := range testCases {
        t.Run(name, func(t *testing.T) {
            err := validate.DomainWithValidCCTLD(tc.domain, tc.minLength, tc.maxLength)
            if err == nil {
                t.Errorf("expected error for domain with invalid input \"%s\", but got no error", tc.domain)
            }
        })
    }
}
