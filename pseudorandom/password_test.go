package pseudorandom_test

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

func FuzzPasswordFunc(f *testing.F) {
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

func FuzzPasswordFor(f *testing.F) {
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
