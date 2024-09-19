package validate

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

const (
	maxPasswordLengthAllowed uint   = 4096
	badPassFilePath   string = "badpass3-5bytehash"
)

var badPass struct {
	pool   map[[5]byte]struct{} // Map to store bad password hashes (5-byte truncated hash)
	loaded bool                 // Flag to indicate if the bad password pool has been loaded
	mu     sync.RWMutex         // Mutex for thread-safe access to the pool
}

var ErrBadPass error = errors.New("password is found in the list of common bad passwords (OWASP Top 1 million)")

// Password validates the password based on length and complexity rules, and checks if it's in the bad password list.
// It returns an error if the password is too short or too long, or lacks required character types.
// minLength and maxLength must be less than 4096.
// The minimum characters allowed for minLength and maxLength equals to the number of boolean requirements (lower, upper, digit, special) that are true. If all are false, the number is one.
// If you also want to check if a password is in the OWASP 1 million bad passwords, use validate.PasswordNotBad function, instead.
func Password(password string, minLength, maxLength uint, hasLower, hasUpper, hasDigit, hasSpecial bool) error {
	var minPasswordLengthAllowed uint

	if hasLower {
		minPasswordLengthAllowed++
	}
	if hasUpper {
		minPasswordLengthAllowed++
	}
	if hasDigit {
		minPasswordLengthAllowed++
	}
	if hasSpecial {
		minPasswordLengthAllowed++
	}

	// If no requirement is true, password will have only ASCII lower case letters.
	if !hasLower && !hasUpper && !hasDigit && !hasSpecial {
		minPasswordLengthAllowed = 1
	}

	err := checkLength(len([]rune(password)), minLength, maxLength, minPasswordLengthAllowed, maxPasswordLengthAllowed, "characters")
	if err != nil {
		return err
	}

	// Flags to track whether the password meets the required complexity rules.
	hasOneLowercase := !hasLower
	hasOneUppercase := !hasUpper
	hasOneDigit := !hasDigit
	hasOneSpecial := !hasSpecial

	// Iterate through the password to validate its complexity.
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
		case char < 32 || char > 126: // Ensure all characters are printable ASCII characters.
			return errors.New("password contains a non-printable ASCII character")
		}
	}

	// Check if all required complexity conditions are met.
	if !hasOneLowercase {
		return errors.New("password must contain at least one lowercase letter")
	}
	if !hasOneUppercase {
		return errors.New("password must contain at least one uppercase letter")
	}
	if !hasOneDigit {
		return errors.New("password must contain at least one digit")
	}
	if !hasOneSpecial {
		return errors.New("password must contain at least one special character")
	}

	// If all checks pass, the password is valid.
	return nil
}

// PasswordNotBad does exactly everything 'validate.Password' does but also returns an error if the password is in the OWASP 1 million bad password list.
func PasswordNotBad(password string, minLength, maxLength uint, hasLower, hasUpper, hasDigit, hasSpecial bool) error {
	err := Password(password, minLength, maxLength, hasLower, hasUpper, hasDigit, hasSpecial)
	if err != nil {
		return err
	}

	length := len([]rune(password))

	if length >= 3 && length < 40 {
		// Load the bad password list if not already loaded.
		if !badPass.loaded {
			if err := loadBadPass(); err != nil {
				return fmt.Errorf("error loading bad password list: %w", err)
			}
		}

		// Check if the password is in the bad password list.
		if IsBadPass(password) {
			return ErrBadPass
		}
	}

	return nil
}

// IsBadPass checks if the provided password's hash exists in the pool of known bad passwords.
// It uses a truncated SHA-256 hash of the password and compares it against the bad password pool.
// It might have false positive (wrong error) specially for simple passwords with 3 or 4 letters similar to common words.
func IsBadPass(password string) bool {
	length := len([]rune(password))
	if length < 3 || length >= 40 {
		return false
	}
	fullHash := sha256.Sum256([]byte(password)) // Generate the SHA-256 hash of the password
	var truncatedHash [5]byte
	copy(truncatedHash[:], fullHash[:5]) // Truncate the hash to 5 bytes

	// Use read-lock to safely access the shared badPass pool
	badPass.mu.RLock()
	defer badPass.mu.RUnlock()

	_, exists := badPass.pool[truncatedHash] // Check if the hash exists in the pool
	return exists
}

// loadBadPass loads truncated password hashes from a binary file into a map for fast lookup.
// Each entry in the file is expected to be a 5-byte long hash of a bad password.
func loadBadPass() error {
	file, err := os.Open(badPassFilePath)
	if err != nil {
		return fmt.Errorf("error opening bad password file: %w", err)
	}
	defer file.Close()

	badPass.pool = make(map[[5]byte]struct{}) // Initialize the map for bad password hashes
	hashBuf := make([]byte, 5)                // Buffer to store the 5-byte hashes

	// Read the file until EOF, loading each hash into the map
	for {
		_, err := io.ReadFull(file, hashBuf)
		if err == io.EOF {
			break
		} else if err != nil {
			return fmt.Errorf("error reading bad password file: %w", err)
		}

		// Copy the truncated hash into a fixed-size array and store it in the pool
		var hash [5]byte
		copy(hash[:], hashBuf[:])
		badPass.pool[hash] = struct{}{}
	}

	// Mark that the bad password pool has been successfully loaded
	badPass.loaded = true
	return nil
}

// PasswordProfile defines the structure for specifying password complexity requirements for a specific service.
type PasswordProfile struct {
	minLength  uint
	maxLength  uint
	hasLower   bool
	hasUpper   bool
	hasDigit   bool
	hasSpecial bool
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

// PasswordFor validates a password based on a predefined password profile.
func PasswordFor(password string, profile PasswordProfile) error {
	err := Password(password, profile.minLength, profile.maxLength, profile.hasLower, profile.hasUpper, profile.hasDigit, profile.hasSpecial)
	if err != nil {
		return err
	}

	return nil
}
