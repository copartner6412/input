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
