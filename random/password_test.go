package random_test

import (
	"crypto/rand"
	"errors"
	"testing"

	"github.com/copartner6412/input/random"
	"github.com/copartner6412/input/validate"
)

func FuzzPasswordFunc(f *testing.F) {
	f.Fuzz(func(t *testing.T, min, max uint, lower, upper, digit, special bool) {
		var minPasswordLength uint
		if lower {
			minPasswordLength++
		}
		if upper {
			minPasswordLength++
		}
		if digit {
			minPasswordLength++
		}
		if special {
			minPasswordLength++
		}
		if !lower && !upper && !digit && !special {
			minPasswordLength++
		}
		minLength := (min % (4096 - minPasswordLength + 1)) + minPasswordLength
		maxLength := minLength + max%(4096-minLength+1)
		password, err := random.Password(rand.Reader, minLength, maxLength, lower, upper, digit, special)
		if err != nil {
			t.Fatalf("error generating a random password: %v", err)
		}
		err = validate.Password(password, minLength, maxLength, lower, upper, digit, special)
		if err != nil {
			t.Fatalf("unexpected error for a valid random password %s: %v", password, err)
		}
	})
}

var profileMap = map[string]struct {
	random   random.PasswordProfile
	validate validate.PasswordProfile
}{
	"PasswordProfileTLSCAKey":             {random.PasswordProfileTLSCAKey, validate.PasswordProfileTLSCAKey},
	"PasswordProfileSSHCAKey":             {random.PasswordProfileSSHCAKey, validate.PasswordProfileSSHCAKey},
	"PasswordProfileTLSKey":               {random.PasswordProfileTLSKey, validate.PasswordProfileTLSKey},
	"PasswordProfileSSHKey":               {random.PasswordProfileSSHKey, validate.PasswordProfileSSHKey},
	"PasswordProfileLinuxServerUser":      {random.PasswordProfileLinuxServerUser, validate.PasswordProfileLinuxServerUser},
	"PasswordProfileLinuxWorkstationUser": {random.PasswordProfileLinuxWorkstationUser, validate.PasswordProfileLinuxWorkstationUser},
	"PasswordProfileWindowsServerUser":    {random.PasswordProfileWindowsServerUser, validate.PasswordProfileWindowsServerUser},
	"PasswordProfileWindowsDesktopUser":   {random.PasswordProfileWindowsDesktopUser, validate.PasswordProfileWindowsDesktopUser},
	"PasswordProfileMariaDB":              {random.PasswordProfileMariaDB, validate.PasswordProfileMariaDB},
}

func FuzzPasswordFor(f *testing.F) {
	f.Fuzz(func(t *testing.T, r uint8) {
		var names []string
		randomProfileSlice := make([]random.PasswordProfile, len(profileMap))
		validateProfileSlice := make([]validate.PasswordProfile, len(profileMap))
		for name := range profileMap {
			names = append(names, name)
		}
		for i, name := range names {
			randomProfileSlice[i] = profileMap[name].random
			validateProfileSlice[i] = profileMap[name].validate
		}
		randomProfile := randomProfileSlice[r%uint8(len(randomProfileSlice))]
		validateProfile := validateProfileSlice[r%uint8(len(validateProfileSlice))]
		password, err := random.PasswordFor(rand.Reader, randomProfile)
		if err != nil {
			t.Fatalf("error generating a pseudo-random password: %v", err)
		}
		err = validate.PasswordFor(password, validateProfile)
		if err != nil {
			if !errors.Is(err, validate.ErrBadPass) {
				t.Fatalf("expected no error for valid pseudo-random password %s for password profile %v, but got error: %v", password, validateProfile, err)
			}
		}
	})
}
