package validate_test

import (
	"testing"

	"github.com/copartner6412/input/validate"
)

const (
	minStringLengthAllowed = 1    // Minimum acceptable string length
	maxStringLengthAllowed = 8192 // Maximum acceptable string length
)

func TestStringSuccessfulForValidString(t *testing.T) {
	testCases := map[string]struct {
		str       string
		minLength uint
		maxLength uint
		justASCII bool
	}{
		"valid ASCII string within length range": {
			str:       "HelloWorld",
			minLength: 5,
			maxLength: 15,
			justASCII: true,
		},
		"valid ASCII string exactly minLength": {
			str:       "Hi",
			minLength: 2,
			maxLength: 10,
			justASCII: true,
		},
		"valid ASCII string exactly maxLength": {
			str:       "GoLangTest",
			minLength: 5,
			maxLength: 10,
			justASCII: true,
		},
		"valid Unicode string within length range": {
			str:       "こんにちは",
			minLength: 3,
			maxLength: 10,
			justASCII: false,
		},
		"valid mixed Unicode and ASCII string": {
			str:       "Hello世界",
			minLength: 5,
			maxLength: 15,
			justASCII: false,
		},
		"valid ASCII string with spaces allowed": {
			str:       "Go Lang Test",
			minLength: 5,
			maxLength: 20,
			justASCII: true,
		},
		"valid string with special ASCII characters": {
			str:       "Valid!@#$%",
			minLength: 5,
			maxLength: 15,
			justASCII: true,
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			err := validate.String(testCase.str, testCase.minLength, testCase.maxLength, testCase.justASCII)
			if err != nil {
				t.Errorf("expected no error for valid string:\n%s\nbut got error: %v", testCase.str, err)
			}
		})
	}
}

func TestStringFailsForInvalidString(t *testing.T) {
	testCases := map[string]struct {
		str       string
		minLength uint
		maxLength uint
		justASCII bool
	}{
		"max less than min length": {
			str:       "abc",
			minLength: 5,
			maxLength: 3,
			justASCII: true,
		},
		"max more than maxStringLength": {
			str:       "hi",
			minLength: 3,
			maxLength: 8193,
			justASCII: true,
		},
		"min less than minStringLength": {
			str:       "hi",
			minLength: 0,
			maxLength: 8192,
			justASCII: true,
		},
		"too short string": {
			str:       "hi",
			minLength: 3,
			maxLength: 10,
			justASCII: true,
		},
		"too long string": {
			str:       "This is a very long string",
			minLength: 5,
			maxLength: 10,
			justASCII: true,
		},
		"empty string": {
			str:       "",
			minLength: 1,
			maxLength: 5,
			justASCII: true,
		},
		"string with non-ASCII character when justASCII is true": {
			str:       "hello€",
			minLength: 1,
			maxLength: 10,
			justASCII: true,
		},
		"string with non-printable Unicode character": {
			str:       "hello\x00world",
			minLength: 1,
			maxLength: 15,
			justASCII: false,
		},
		"string within valid length but contains non-ASCII": {
			str:       "hello世界",
			minLength: 5,
			maxLength: 10,
			justASCII: true,
		},
		"too long string with non-printable character": {
			str:       "abcd\x00efghij",
			minLength: 1,
			maxLength: 5,
			justASCII: false,
		},
		"valid length but non-printable Unicode character when justASCII is false": {
			str:       "hi\x07there",
			minLength: 1,
			maxLength: 10,
			justASCII: false,
		},
		"non-printable Unicode character with valid length": {
			str:       "hello\x0bworld",
			minLength: 5,
			maxLength: 15,
			justASCII: false,
		},
		"string at min length boundary but contains non-ASCII": {
			str:       "abc€",
			minLength: 4,
			maxLength: 6,
			justASCII: true,
		},
		"string exceeding max length boundary": {
			str:       "abcdefghij",
			minLength: 1,
			maxLength: 9,
			justASCII: true,
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			err := validate.String(testCase.str, testCase.minLength, testCase.maxLength, testCase.justASCII)
			if err == nil {
				t.Errorf("expected error for invalid string \"%s\", but got no error", testCase.str)
			}
		})
	}
}
