package validate_test

import (
	"errors"
	"math/rand/v2"
	"sort"
	"testing"

	"github.com/copartner6412/input/pseudorandom"
	"github.com/copartner6412/input/validate"
)

const (
	maxPasswordLengthAllowed uint = 4096
)

func FuzzPasswordSuccessfulForValidPseudorandomInput(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed1, seed2 uint64, min, max uint, lower, upper, digit, special bool) {
		var minPasswordLengthAllowed uint

		if lower {
			minPasswordLengthAllowed++
		}
		if upper {
			minPasswordLengthAllowed++
		}
		if digit {
			minPasswordLengthAllowed++
		}
		if special {
			minPasswordLengthAllowed++
		}
		if !lower && !upper && !digit && !special {
			minPasswordLengthAllowed++
		}

		r1, r2, minLength, maxLength := randoms(seed1, seed2, min, max, minPasswordLengthAllowed, maxPassphraseWordsAllowed)

		if minLength < minPasswordLengthAllowed {
			minLength = minPasswordLengthAllowed
		}

		password1, err := pseudorandom.Password(r1, minLength, maxLength, lower, upper, digit, special)
		if err != nil {
			t.Fatalf("error generating a pseudo-random password: %v", err)
		}
		
		err = validate.Password(password1, minLength, maxLength, lower, upper, digit, special)
		if err != nil {
			t.Fatalf("expected no error for a valid pseudo-random password %s: but got error: %v", password1, err)
		}

		password2, err := pseudorandom.Password(r2, minLength, maxLength, lower, upper, digit, special)
		if err != nil {
			t.Fatalf("error regenerating the pseudo-random password: %v", err)
		}
		if password1 != password2 {
			t.Fatalf("not deterministic")
		}
	})
}

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

var profileMap = map[string]struct {
	pseudorandom pseudorandom.PasswordProfile
	validate     validate.PasswordProfile
}{
	"PasswordProfileTLSCAKey":             {pseudorandom.PasswordProfileTLSCAKey, validate.PasswordProfileTLSCAKey},
	"PasswordProfileSSHCAKey":             {pseudorandom.PasswordProfileSSHCAKey, validate.PasswordProfileSSHCAKey},
	"PasswordProfileTLSKey":               {pseudorandom.PasswordProfileTLSKey, validate.PasswordProfileTLSKey},
	"PasswordProfileSSHKey":               {pseudorandom.PasswordProfileSSHKey, validate.PasswordProfileSSHKey},
	"PasswordProfileLinuxServerUser":      {pseudorandom.PasswordProfileLinuxServerUser, validate.PasswordProfileLinuxServerUser},
	"PasswordProfileLinuxWorkstationUser": {pseudorandom.PasswordProfileLinuxWorkstationUser, validate.PasswordProfileLinuxWorkstationUser},
	"PasswordProfileWindowsServerUser":    {pseudorandom.PasswordProfileWindowsServerUser, validate.PasswordProfileWindowsServerUser},
	"PasswordProfileWindowsDesktopUser":   {pseudorandom.PasswordProfileWindowsDesktopUser, validate.PasswordProfileWindowsDesktopUser},
	"PasswordProfileMariaDB":              {pseudorandom.PasswordProfileMariaDB, validate.PasswordProfileMariaDB},
}

func FuzzPasswordForSuccessfulForValidPseudorandomPassword(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed1, seed2 uint64, random uint8) {
		r1 := rand.New(rand.NewPCG(seed1, seed2))
		var names []string
		pseudorandomProfileSlice := make([]pseudorandom.PasswordProfile, len(profileMap))
		validateProfileSlice := make([]validate.PasswordProfile, len(profileMap))
		for name := range profileMap {
			names = append(names, name)
		}
		sort.Strings(names)
		for i, name := range names {
			pseudorandomProfileSlice[i] = profileMap[name].pseudorandom
			validateProfileSlice[i] = profileMap[name].validate
		}
		pseudorandomProfile := pseudorandomProfileSlice[random%uint8(len(pseudorandomProfileSlice))]
		validateProfile := validateProfileSlice[random%uint8(len(validateProfileSlice))]
		password1, err := pseudorandom.PasswordFor(r1, pseudorandomProfile)
		if err != nil {
			t.Fatalf("error generating a pseudo-random password: %v", err)
		}
		err = validate.PasswordFor(password1, validateProfile)
		if err != nil {
			if !errors.Is(err, validate.ErrBadPass) {
				t.Fatalf("expected no error for valid pseudo-random password %s for password profile %v, but got error: %v", password1, validateProfile, err)
			}
		}
		r2 := rand.New(rand.NewPCG(seed1, seed2))
		password2, err := pseudorandom.PasswordFor(r2, pseudorandomProfile)
		if err != nil {
			t.Fatalf("error regenerating the pseudo-random password: %v", err)
		}
		if password1 != password2 {
			t.Fatal("not deterministic")
		}

	})
}

func TestPasswordForSuccessfulForValidPassword(t *testing.T) {
	testCases := map[string]struct {
		password string
		profile  validate.PasswordProfile
	}{
		"TLS CA Key": {
			password: "ValidTLSPassword123456789!",
			profile:  validate.PasswordProfileTLSCAKey,
		},
		"SSH CA Key": {
			password: "SSHCaKeyPassword123456789@",
			profile:  validate.PasswordProfileSSHCAKey,
		},
		"TLS Key": {
			password: "TlsKeyValidPass123456789@",
			profile:  validate.PasswordProfileTLSKey,
		},
		"SSH Key": {
			password: "ValidSshPassword123456789", // No special characters required
			profile:  validate.PasswordProfileSSHKey,
		},
		"Linux Server User": {
			password: "LinuxServerPass123456789", // No special characters required
			profile:  validate.PasswordProfileLinuxServerUser,
		},
		"Linux Workstation User": {
			password: "WorkPass123", // Shorter password, no upper case required
			profile:  validate.PasswordProfileLinuxWorkstationUser,
		},
		"Windows Server User": {
			password: "ValidWinServerPass123456789", // No special characters required
			profile:  validate.PasswordProfileWindowsServerUser,
		},
		"Windows Desktop User": {
			password: "DesktopPass123", // No upper case or special characters required
			profile:  validate.PasswordProfileWindowsDesktopUser,
		},
		"MariaDB": {
			password: "MariaDBPass123456789", // No special characters required, short max length
			profile:  validate.PasswordProfileMariaDB,
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			err := validate.PasswordFor(testCase.password, testCase.profile)
			if err != nil {
				if !errors.Is(err, validate.ErrBadPass) {
					t.Errorf("expected no error for valid password %q, but got error: %v", testCase.password, err)
				}
			}
		})
	}
}

func TestPasswordForFailsForNotComplexInput(t *testing.T) {
	testCases := map[string]struct {
		password string
		profile  validate.PasswordProfile
	}{
		"TLS CA Key - Missing special character": {
			password: "ValidPassword12345abcdefghijk", // Missing special character
			profile:  validate.PasswordProfileTLSCAKey,
		},
		"SSH CA Key - Missing digit": {
			password: "ValidPasswordSpecial!abcdefghij", // Missing digit
			profile:  validate.PasswordProfileSSHCAKey,
		},
		"TLS Key - Missing upper case": {
			password: "validpassword123!abcdefghijk", // Missing upper case letter
			profile:  validate.PasswordProfileTLSKey,
		},
		"SSH Key - Missing digit": {
			password: "ValidPasswordNoDigitabcdefghij", // Missing digit
			profile:  validate.PasswordProfileSSHKey,
		},
		"Linux Server User - Missing digit": {
			password: "ValidPasswordNoDigitabcdefghi", // Missing digit
			profile:  validate.PasswordProfileLinuxServerUser,
		},
		"Linux Workstation User - Missing digit": {
			password: "validpasswordabcdef", // Missing digit
			profile:  validate.PasswordProfileLinuxWorkstationUser,
		},
		"Windows Server User - Missing digit": {
			password: "ValidPasswordNoDigit", // Missing digit
			profile:  validate.PasswordProfileWindowsServerUser,
		},
		"Windows Desktop User - Missing digit": {
			password: "validpassword", // Missing digit
			profile:  validate.PasswordProfileWindowsDesktopUser,
		},
		"MariaDB - Missing digit": {
			password: "ValidPasswordNoDigit", // Missing digit
			profile:  validate.PasswordProfileMariaDB,
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			err := validate.PasswordFor(testCase.password, testCase.profile)
			if err == nil || errors.Is(err, validate.ErrBadPass) {
				t.Errorf("expected error for password %q in profile %q, but got no error", testCase.password, name)
			}
		})
	}
}
