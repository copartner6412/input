package validate_test

import (
	"errors"
	"testing"

	"github.com/copartner6412/input/validate"
)

func TestPasswordSuccessfulForValidInput(t *testing.T) {
	testCases := map[string]struct {
		password string
		length   uint
		lower    bool
		upper    bool
		digit    bool
		special  bool
	}{
		"LowerOnly":                 {"abcdefgh", 8, true, false, false, false},
		"UpperOnly":                 {"ABCDEFGH", 8, false, true, false, false},
		"DigitOnly":                 {"123456789", 9, false, false, true, false},
		"SpecialOnly":               {"!@#$%^&*()_-+=", 14, false, false, false, true},
		"LowerAndUpper":             {"AbCdEfGhIjKlMn", 14, true, true, false, false},
		"LowerAndDigit":             {"abc123def456", 12, true, false, true, false},
		"LowerAndSpecial":           {"abc!@#$%^&", 10, true, false, false, true},
		"UpperAndDigit":             {"ABC123EFG987", 12, false, true, true, false},
		"UpperAndSpecial":           {"ABC!@#$%^&*()_+-=:;?/><.,", 25, false, true, false, true},
		"DigitAndSpecial":           {"123!@~!|()", 10, false, false, true, true},
		"LowerUpperAndDigit":        {"AbC123MnB765", 12, true, true, true, false},
		"LowerUpperAndSpecial":      {"AbC!@$%^bGt*&^SdfQ", 18, true, true, false, true},
		"LowerDigitAndSpecial":      {"abc123!@", 8, true, false, true, true},
		"UpperDigitAndSpecial":      {"ABC123!@", 8, false, true, true, true},
		"LowerUpperDigitAndSpecial": {"AbC123!@", 8, true, true, true, true},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			err := validate.Password(testCase.password, testCase.length, testCase.length, testCase.lower, testCase.upper, testCase.digit, testCase.special)
			if err != nil {
				t.Errorf("expected no error for valid password %q, but got error: %v", testCase.password, err)
			}
		})
	}
}

func TestPasswordFailsForTooShortOrTooLongPassword(t *testing.T) {
	testCases := map[string]struct {
		password  string
		length    uint
		minLength uint
		maxLength uint
	}{
		// minLength less than the allowed minimum (minLength < 3)
		"MinLengthTooShort": {
			password:  "abcefg",
			length:    6,
			minLength: 3, // invalid minLength
			maxLength: 10,
		},

		// maxLength greater than the allowed maximum (maxLength > 1024)
		"MaxLengthTooLong": {
			password:  "abcde",
			length:    5,
			minLength: 4,
			maxLength: 4097, // invalid maxLength
		},

		// maxLength less than minLength (maxLength < minLength)
		"MaxLengthLessThanMinLength": {
			password:  "abcdef",
			length:    6,
			minLength: 7, // invalid minLength
			maxLength: 5, // invalid maxLength (max < min)
		},

		// Empty password
		"EmptyPassword": {
			password:  "",
			length:    0,
			minLength: 4,
			maxLength: 5,
		},

		// Password length less than minLength
		"PasswordLengthLessThanMinLength": {
			password:  "abc",
			length:    3, // password length less than minLength
			minLength: 4,
			maxLength: 5,
		},

		// Password length greater than maxLength
		"PasswordLengthGreaterThanMaxLength": {
			password:  "abcdef",
			length:    6, // password length greater than maxLength
			minLength: 4,
			maxLength: 5, // maxLength is exceeded
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			err := validate.Password(testCase.password, testCase.minLength, testCase.maxLength, true, true, true, true)
			if err == nil || errors.Is(err, validate.ErrBadPass) {
				t.Errorf("expected error for invalid password length in case %s, but got none", name)
			}
		})
	}
}

func TestPasswordFailsForInvalidComplexityRequirementInput(t *testing.T) {
	testCases := map[string]struct {
		password string
		length   uint
		lower    bool
		upper    bool
		digit    bool
		special  bool
	}{
		// Expected lower, but only upper provided
		"LowerRequiredButMissing": {"ABCD", 4, true, false, false, false},

		// Expected upper, but only lower provided
		"UpperRequiredButMissing": {"abcd", 4, false, true, false, false},

		// Expected digit, but none provided
		"DigitRequiredButMissing": {"AbcD", 4, false, false, true, false},

		// Expected special, but none provided
		"SpecialRequiredButMissing": {"Abc123", 6, false, false, false, true},

		// Lower and upper required, but only digits
		"LowerUpperRequiredButOnlyDigits": {"123456", 6, true, true, false, false},

		// Lower and digit required, but only upper and special provided
		"LowerDigitRequiredButOnlyUpperSpecial": {"ABC!@", 5, true, false, true, false},

		// Upper and special required, but only lower and digits provided
		"UpperSpecialRequiredButOnlyLowerDigits": {"abc123", 6, false, true, false, true},

		// Lower, upper, and digit required, but only special characters
		"LowerUpperDigitRequiredButOnlySpecial": {"!@#$%", 5, true, true, true, false},

		// Lower and upper required, but only digits and special characters
		"LowerUpperRequiredButOnlyDigitsSpecial": {"123!@", 5, true, true, false, true},

		// Lower, upper, and digit required, but only lower provided
		"LowerUpperDigitRequiredButOnlyLower": {"abcdef", 6, true, true, true, false},

		// Lower, upper, and special required, but only digits provided
		"LowerUpperSpecialRequiredButOnlyDigits": {"123456", 6, true, true, false, true},

		// Upper, digit, and special required, but only lower provided
		"UpperDigitSpecialRequiredButOnlyLower": {"abcdef", 6, false, true, true, true},

		// Lower, digit, and special required, but only upper provided
		"LowerDigitSpecialRequiredButOnlyUpper": {"ABCDEF", 6, true, false, true, true},

		// Upper, digit, and special required, but only lower and digits provided
		"UpperDigitSpecialRequiredButOnlyLowerDigits": {"abc123", 6, false, true, true, true},

		// Special required, but only lower, upper, and digits provided
		"LowerUpperDigitSpecial": {"AbC123", 6, true, true, true, true},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			err := validate.Password(testCase.password, testCase.length, testCase.length, testCase.lower, testCase.upper, testCase.digit, testCase.special)
			if err == nil {
				t.Errorf("expected error for invalid password %q, but got none", testCase.password)
			}
		})
	}
}

func TestPasswordNotBadSuccessfulForNotBadPassword(t *testing.T) {
	testCases := []string{
		"find2",
		"comr",
		"badee",
		"yett",
		"bye3",
		"ByE",
		"odD",
		"whyY",
		"Howw",
		"suprising",
		"sun9",
		"Ioutil",
		"golang",
		"pseudorandom",
		"ab",
		"Iwanttocheckapasswordwithmorethan40characters",
	}
	for _, testCase := range testCases {
		t.Run(testCase, func(t *testing.T) {
			err := validate.PasswordNotBad(testCase, 2, 1024, false, false, false, false)
			if errors.Is(err, validate.ErrBadPass) {
				t.Errorf("expected no error for bad password %s, but got: %v", testCase, err)
			}
		})
	}
}

func TestPasswordNotBadFailsForBadPasswords(t *testing.T) {
	testCases := []string{
		"!y8AjEWuveNeqa",
		"#y@AMe@uHuTy2u",
		"22092000m",
		"DeadDroi",
		"WU55667N",
		"agronomis",
		"otju1otju",
		"reloaded216465",
		"rellorts",
		"supralingua",
		"war7yesout",
		"xzxz1234",
		"acerbic",
		"Remorse",
		"INSTANT",
	}
	for _, testCase := range testCases {
		t.Run(testCase, func(t *testing.T) {
			err := validate.PasswordNotBad(testCase, 1, 1024, false, false, false, false)
			if !errors.Is(err, validate.ErrBadPass) {
				t.Errorf("expected error for bad password %s, but got: %v", testCase, err)
			}
		})
	}
}

func TestIsBadPassTrueForNotBadPassword(t *testing.T) {
	testCases := []string{
		"find2",
		"comr",
		"badee",
		"yett",
		"bye3",
		"ByE",
		"odD",
		"whyY",
		"Howw",
		"suprising",
		"sun9",
		"Ioutil",
		"golang",
		"pseudorandom",
		"ab",
		"Iwanttocheckapasswordwithmorethan40characters",
	}
	for _, testCase := range testCases {
		t.Run(testCase, func(t *testing.T) {
			if validate.IsBadPass(testCase) {
				t.Errorf("expected false for bad password %s, but got true", testCase)
			}
		})
	}
}

func TestIsBadPassTrueForBadPasswords(t *testing.T) {
	testCases := []string{
		"!y8AjEWuveNeqa",
		"#y@AMe@uHuTy2u",
		"22092000m",
		"DeadDroi",
		"WU55667N",
		"agronomis",
		"otju1otju",
		"reloaded216465",
		"rellorts",
		"supralingua",
		"war7yesout",
		"xzxz1234",
		"acerbic",
		"Remorse",
		"INSTANT",
	}
	for _, testCase := range testCases {
		t.Run(testCase, func(t *testing.T) {
			if !validate.IsBadPass(testCase) {
				t.Errorf("expected ture for bad password %s, but got false", testCase)
			}
		})
	}
}
