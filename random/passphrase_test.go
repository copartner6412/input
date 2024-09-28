package random_test

import (
	"crypto/rand"
	"testing"

	"github.com/copartner6412/input/random"
	"github.com/copartner6412/input/validate"
)

func FuzzPassphrase(f *testing.F) {
	f.Fuzz(func(t *testing.T, min, max uint, lowerSep, upperSep, digitSep, specialSep bool, capitalize, number bool, list uint) {
		if !digitSep && !specialSep {
			t.Skip()
		}
		minWords := (min%127 + 2)
		maxWords := minWords + max%(128-minWords+1)
		var minSeparatorLength uint
		if lowerSep {
			minSeparatorLength++
		}
		if upperSep {
			minSeparatorLength++
		}
		if digitSep {
			minSeparatorLength++
		}
		if specialSep {
			minSeparatorLength++
		}
		separator, err := random.Password(rand.Reader, minSeparatorLength, 20, lowerSep, upperSep, digitSep, specialSep)
		if err != nil {
			t.Fatalf("error generating a random separator: %v", err)
		}

		var wordList []string

		switch list % 3 {
		case 1:
			wordList = validate.AGWordList
		case 2:
			wordList = validate.EEFLongWordList
		default:
			wordList = nil
		}
		passphrase, err := random.Passphrase(rand.Reader, minWords, maxWords, separator, capitalize, number, wordList)
		if err != nil {
			t.Fatalf("error generating a random passphrase: %v", err)
		}
		err = validate.Passphrase(passphrase, minWords, maxWords, separator, capitalize, number, nil)
		if err != nil {
			t.Fatalf("expected no error for valid pseudo-random passphrase \"%s\", but got error: %v", passphrase, err)
		}
	})
}
