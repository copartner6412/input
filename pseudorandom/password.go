package pseudorandom

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

const (
	maxPasswordLengthAllowed uint = 4096
)

// Behavior:
//  - The function first validates the input parameters to ensure minLength and maxLength are within an acceptable range.
//  - It then builds a character set based on the boolean flags (lower, upper, digit, special) and ensures that at least
//    one flag is true, ensuring a valid character set.
//  - The function generates a password of random length (between minLength and maxLength) using the selected character set.
//  - It ensures that the generated password meets all the specified complexity requirements (e.g., at least one lowercase letter
//    if lower is true, etc.).
//  - The function repeats password generation until all specified complexity conditions are met.

// Password generates a deterministic pseudo-random password of a length between the specified minLength and maxLength and complexity requirements from printable ASCII characters.
//
// Parameters:
//   - r: Randomness source.
//   - minLength: The minimum length of the password (up to 4096 characters).
//   - maxLength: The maximum length of the password (up to 4096 characters).
//   - lower: Whether to include lowercase letters.
//   - upper: Whether to include uppercase letters.
//   - digit: Whether to include digits.
//   - special: Whether to include special characters.
//
// Returns:
//   - A string containing the generated password.
//   - An error if the parameters are invalid.
//
// The minimum characters allowed for minLength and maxLength equls to the number of boolean requirements (lower, upper, digit, special) that are true. If all are false, the number is one.
func Password(r *rand.Rand, minLength, maxLength uint, lower bool, upper bool, digit bool, special bool) (string, error) {
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

	// If no requirement is true, password will have only ASCII lower case letters and the minimum length allowed for password is one.
	if !lower && !upper && !digit && !special {
		lower = true
		minPasswordLengthAllowed = 1
	}

	length, err := checkLength(r, minLength, maxLength, minPasswordLengthAllowed, maxPasswordLengthAllowed)
	if err != nil {
		return "", err
	}

	// Build the character set based on selected options.
	var allowedCharacters []rune
	if lower {
		allowedCharacters = append(allowedCharacters, lowerCaseRunes...)
	}
	if upper {
		allowedCharacters = append(allowedCharacters, upperCaseRunes...)
	}
	if digit {
		allowedCharacters = append(allowedCharacters, digitRunes...)
	}
	if special {
		allowedCharacters = append(allowedCharacters, specialRunes...)
	}

	// Create a rune slice to hold the generated password.
	password := make([]rune, length)

	// Loop until a valid password that meets the complexity requirements is generated.
	for {
		// Randomly generate the password using the allowed characters.
		for i := 0; i < int(length); i++ {
			password[i] = allowedCharacters[r.IntN(len(allowedCharacters))]
		}

		// Flags to track whether the password meets the required complexity rules.
		hasOneLowercase := !lower
		hasOneUppercase := !upper
		hasOneDigit := !digit
		hasOneSpecial := !special

		// Validate the password's complexity by checking if it contains the required types of characters.
		for _, char := range password {
			switch {
			case strings.ContainsRune(string(lowerCaseRunes), char):
				hasOneLowercase = true
			case strings.ContainsRune(string(upperCaseRunes), char):
				hasOneUppercase = true
			case strings.ContainsRune(string(digitRunes), char):
				hasOneDigit = true
			case strings.ContainsRune(string(specialRunes), char):
				hasOneSpecial = true
			}
		}

		// If the password satisfies all complexity requirements, break out of the loop.
		if hasOneLowercase && hasOneUppercase && hasOneDigit && hasOneSpecial {
			break
		}

		// Otherwise, generate a new password in the next iteration.
	}

	return string(password), nil
}

// PasswordProfile defines the structure for specifying password complexity requirements for using in PasswordFor function.
type PasswordProfile struct {
	minLength  uint // Minimum number of characters required in the password.
	maxLength  uint // Maximum allowed number of characters in the password.
	hasLower   bool // Whether the password must include lowercase letters.
	hasUpper   bool // Whether the password must include uppercase letters.
	hasDigit   bool // Whether the password must include digits.
	hasSpecial bool // Whether the password must include special characters.
}

var (
	// Password profile for TLS CA key:
	//  - minLength: 20
	//  - maxLength: 255
	//  - hasLower: true
	//  - hasUpper: true
	//  - hasDigit: true
	//  - hasSpecial: true
	PasswordProfileTLSCAKey = PasswordProfile{minLength: 20, maxLength: 255, hasLower: true, hasUpper: true, hasDigit: true, hasSpecial: true}
	// Password Profile for SSH CA key:
	//  - minLength: 20
	//  - maxLength: 255
	//  - hasLower: true
	//  - hasUpper: true
	//  - hasDigit: true
	//  - hasSpecial: true
	PasswordProfileSSHCAKey = PasswordProfile{minLength: 20, maxLength: 255, hasLower: true, hasUpper: true, hasDigit: true, hasSpecial: true}
	// Password profile for TLS key:
	//  - minLength: 20
	//  - maxLength: 127
	//  - hasLower: true
	//  - hasUpper: true
	//  - hasDigit: true
	//  - hasSpecial: true
	PasswordProfileTLSKey = PasswordProfile{minLength: 20, maxLength: 127, hasLower: true, hasUpper: true, hasDigit: true, hasSpecial: true}
	// Password profile for SSH key:
	//  - minLength: 20
	//  - maxLength: 127
	//  - hasLower: true
	//  - hasUpper: true
	//  - hasDigit: true
	//  - hasSpecial: false
	PasswordProfileSSHKey = PasswordProfile{minLength: 20, maxLength: 127, hasLower: true, hasUpper: true, hasDigit: true, hasSpecial: false}
	// Password profile for Linux server user:
	//  - minLength: 20
	//  - maxLength: 63
	//  - hasLower: true
	//  - hasUpper: true
	//  - hasDigit: true
	//  - hasSpecial: false
	PasswordProfileLinuxServerUser = PasswordProfile{minLength: 20, maxLength: 63, hasLower: true, hasUpper: true, hasDigit: true, hasSpecial: false}
	// Password profile for Linux workstation user:
	//  - minLength: 10
	//  - maxLength: 20
	//  - hasLower: true
	//  - hasUpper: false
	//  - hasDigit: true
	//  - hasSpecial: false
	PasswordProfileLinuxWorkstationUser = PasswordProfile{minLength: 10, maxLength: 20, hasLower: true, hasUpper: false, hasDigit: true, hasSpecial: false}
	// Password profile for Windows server user:
	//  - minLength: 20
	//  - maxLength: 63
	//  - hasLower: true
	//  - hasUpper: true
	//  - hasDigit: true
	//  - hasSpecial: false
	PasswordProfileWindowsServerUser = PasswordProfile{minLength: 20, maxLength: 63, hasLower: true, hasUpper: true, hasDigit: true, hasSpecial: false}
	// Password profile for Windows desktop user:
	//  - minLength: 10
	//  - maxLength: 20
	//  - hasLower: true
	//  - hasUpper: false
	//  - hasDigit: true
	//  - hasSpecial: false
	PasswordProfileWindowsDesktopUser = PasswordProfile{minLength: 10, maxLength: 20, hasLower: true, hasUpper: false, hasDigit: true, hasSpecial: false}
	// 	Password profile for MariaDB
	//  - minLength: 20
	//  - maxLength: 31
	//  - hasLower: true
	//  - hasUpper: true
	//  - hasDigit: true
	//  - hasSpecial: false
	PasswordProfileMariaDB = PasswordProfile{minLength: 20, maxLength: 31, hasLower: true, hasUpper: true, hasDigit: true, hasSpecial: false}
)

// PasswordFor generates a deterministic pseudo-random password based on a predefined profile using the provided random source.
// Parameters:
//   - r: Randomness source.
//   - profile: A PasswordProfile struct containing the configuration for the password.
//
// Returns:
//   - A string containing the generated password.
//   - An error if something goes wrong during password generation.
func PasswordFor(r *rand.Rand, profile PasswordProfile) (string, error) {
	password, err := Password(r, profile.minLength, profile.maxLength, profile.hasLower, profile.hasUpper, profile.hasDigit, profile.hasSpecial)
	if err != nil {
		return "", fmt.Errorf("error generating a pseudo-random password for profile: %w", err)
	}

	return password, nil
}
