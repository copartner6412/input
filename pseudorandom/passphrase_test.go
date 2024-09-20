package pseudorandom_test

import (
	"testing"

	"github.com/copartner6412/input/pseudorandom"
	"github.com/copartner6412/input/validate"
)

const (
	minPassphraseWordsAllowed uint = 2
	maxPassphraseWordsAllowed uint = 128
)

func FuzzPassphrase(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed1, seed2 uint64, min, max uint, lowerSep, upperSep, digitSep, specialSep bool, capitalize, number bool, list uint) {
		if !digitSep && !specialSep {
			t.Skip()
		}

		r1, r2, minWords, maxWords := randoms(seed1, seed2, min, max, minPassphraseWordsAllowed, maxPassphraseWordsAllowed)

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

		separator1, err := pseudorandom.Password(r1, minSeparatorLength, 20, lowerSep, upperSep, digitSep, specialSep)
		if err != nil {
			t.Errorf("error generating a random separator: %v", err)
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

		passphrase1, err := pseudorandom.Passphrase(r1, minWords, maxWords, separator1, capitalize, number, wordList)
		if err != nil {
			t.Fatalf("error generating a random passphrase: %v", err)
		}

		err = validate.Passphrase(passphrase1, minWords, maxWords, separator1, capitalize, number, nil)
		if err != nil {
			t.Fatalf("expected no error for valid pseudo-random passphrase %s, but got error: %v", passphrase1, err)
		}

		separator2, err := pseudorandom.Password(r2, minSeparatorLength, 20, lowerSep, upperSep, digitSep, specialSep)
		if err != nil {
			t.Fatalf("error generating a random separator: %v", err)
		}

		passphrase2, err := pseudorandom.Passphrase(r2, minWords, maxWords, separator2, capitalize, number, wordList)
		if err != nil {
			t.Fatalf("error regenerating the random passphrase: %v", err)
		}

		if passphrase1 != passphrase2 {
			t.Fatal("not deterministic")
		}
	})
}
