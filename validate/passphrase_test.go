package validate_test

import (
	"testing"

	"github.com/copartner6412/input/pseudorandom"
	"github.com/copartner6412/input/validate"
)

const (
	minPassphraseWordsAllowed uint = 2
	maxPassphraseWordsAllowed uint = 128
)

func FuzzPassphraseSuccessfulForValidPseudorandomInput(f *testing.F) {
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

func TestPassphraseSuccessfulForValidPassphrase(t *testing.T) {
	testCases := map[string]struct {
		passphrase string
		length     uint
		separator  string
		capitalize bool
		number     bool
		wordList   []string
	}{

		"3 words from AGWordList with dash as separator": {
			passphrase: "abalone-balcony-gadfly", // Basic valid passphrase
			length:     3,
			separator:  "-",
			capitalize: false,
			number:     false,
			wordList:   validate.AGWordList,
		},
		"3 words capitalized from AGWordList with 2-dash as separator": {
			passphrase: "Telepath--Council--Rental", // Capitalized words
			length:     3,
			separator:  "--",
			capitalize: true,
			number:     false,
			wordList:   validate.AGWordList,
		},
		"3 words with number from AGWordList with underscore as separator": {
			passphrase: "cleanup4_telepath00_rental3", // Passphrase with number
			length:     3,
			separator:  "_",
			capitalize: false,
			number:     true,
			wordList:   validate.AGWordList,
		},
		"3 words with number from AGWordList with plus as separator": {
			passphrase: "9balcony751+cleanup+gadfly", // Valid passphrase with plus separator
			length:     3,
			separator:  "+",
			capitalize: false,
			number:     false,
			wordList:   validate.AGWordList,
		},
		"4 words from EEFLongWordList with '-1%' as separator": {
			passphrase: "ablaze-1%basket-1%caravan-1%uncommon", // Basic valid passphrase
			length:     4,
			separator:  "-1%",
			capitalize: false,
			number:     false,
			wordList:   validate.EEFLongWordList,
		},
		"3 words with number from EFFLongWordList with '89' as separator": {
			passphrase: "Reliable89Charter8765432989Nineteen", // Capitalized words with number suffix
			length:     3,
			separator:  "89",
			capitalize: true,
			number:     true,
			wordList:   validate.EEFLongWordList,
		},
		"3 words with number from EFFLongWordList with underscore as separator": {
			passphrase: "eatery_affidavit_bunkhouse5", // Passphrase with number
			length:     3,
			separator:  "_",
			capitalize: false,
			number:     true,
			wordList:   validate.EEFLongWordList,
		},
		"3 words capitalized from EFFLongWordList with dash as separtor": {
			passphrase: "Hassle+!Errand+!Catatonic", // Capitalized words
			length:     3,
			separator:  "+!",
			capitalize: true,
			number:     false,
			wordList:   validate.EEFLongWordList,
		},
		"4 random words with '0abC00' as separator": {
			passphrase: "random0abC00passphrase0abC00without0abC00list", // Passphrase without word list
			length:     4,
			separator:  "0abC00",
			capitalize: false,
			number:     false,
			wordList:   nil, // No word list provided
		},
		"3 random words capitalized with number and '-1zG' as separator": {
			passphrase: "Some-1zGRandom-1zGPhrase4", // Capitalized and number, no word list
			length:     3,
			separator:  "-1zG",
			capitalize: true,
			number:     true,
			wordList:   nil,
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			err := validate.Passphrase(testCase.passphrase, testCase.length, testCase.length, testCase.separator, testCase.capitalize, testCase.number, testCase.wordList)
			if err != nil {
				t.Errorf("expected no error for valid pseudo-random passphrase \"%s\", but got error: %v", testCase.passphrase, err)
			}
		})
	}
}

func TestPassphraseFailsForInvalidPassphrase(t *testing.T) {
	testCases := map[string]struct {
		passphrase string
		minWords   uint
		maxWords   uint
		separator  string
		capitalize bool
		number     bool
		wordList   []string
	}{
		"maxWords less than minWords": {
			passphrase: "abalone-balcony", // 1 word, less than 2
			minWords:   3,
			maxWords:   2,
			separator:  "-",
			capitalize: false,
			number:     false,
			wordList:   validate.AGWordList,
		},
		"minWords less than minPassphraseWords": {
			passphrase: "abalone-balcony", // 1 word, less than 2
			minWords:   1,
			maxWords:   2,
			separator:  "-",
			capitalize: false,
			number:     false,
			wordList:   validate.AGWordList,
		},
		"maxWords more than maxPassphraseWords": {
			passphrase: "abalone-balcony", // 1 word, less than 2
			minWords:   2,
			maxWords:   129,
			separator:  "-",
			capitalize: false,
			number:     false,
			wordList:   validate.AGWordList,
		},
		"passphrase has less words than minWords": {
			passphrase: "abalone", // 1 word, less than 2
			minWords:   2,
			maxWords:   3,
			separator:  "-",
			capitalize: false,
			number:     false,
			wordList:   validate.AGWordList,
		},
		"passphrase has more words than maxWords": {
			passphrase: "abalone-balcony-gadfly-telepath", // 4 words, more than 3
			minWords:   2,
			maxWords:   3,
			separator:  "-",
			capitalize: false,
			number:     false,
			wordList:   validate.AGWordList,
		},
		"separator is empty": {
			passphrase: "AbaloneBalconyGadfly", // 'gadfly' is not capitalized
			minWords:   3,
			maxWords:   3,
			separator:  "",
			capitalize: false,
			number:     false,
			wordList:   validate.AGWordList,
		},
		"separator is space": {
			passphrase: "Abalone Balcony gadfly", // 'gadfly' is not capitalized
			minWords:   3,
			maxWords:   3,
			separator:  " ",
			capitalize: false,
			number:     false,
			wordList:   validate.AGWordList,
		},
		"separator only consists of letters": {
			passphrase: "AbaloneXyBalconyXyGadfly", // 'gadfly' is not capitalized
			minWords:   3,
			maxWords:   3,
			separator:  "Xy",
			capitalize: false,
			number:     false,
			wordList:   validate.AGWordList,
		},
		"separator doesn't match and causes empty word": {
			passphrase: "Abalone--Balcony--adfly", // 'gadfly' is not capitalized
			minWords:   3,
			maxWords:   3,
			separator:  "-",
			capitalize: false,
			number:     false,
			wordList:   validate.AGWordList,
		},
		"separator doesn't match": {
			passphrase: "Abalone-%Balcony-%Gadfly", // 'gadfly' is not capitalized
			minWords:   3,
			maxWords:   3,
			separator:  "-",
			capitalize: false,
			number:     false,
			wordList:   validate.AGWordList,
		},
		"capitalize is true but passphrase has a non-capitalized word": {
			passphrase: "Abalone-Balcony-gadfly", // 'gadfly' is not capitalized
			minWords:   3,
			maxWords:   3,
			separator:  "-",
			capitalize: true,
			number:     false,
			wordList:   validate.AGWordList,
		},
		"number is true but no number at the end of any word": {
			passphrase: "abalone-balcony-gadfly", // No numbers
			minWords:   3,
			maxWords:   3,
			separator:  "-",
			capitalize: false,
			number:     true,
			wordList:   validate.AGWordList,
		},
		"wordlist specified but passphrase has word not in wordlist": {
			passphrase: "abalone-balcony-nonexistent", // 'nonexistent' not in AGWordList
			minWords:   3,
			maxWords:   3,
			separator:  "-",
			capitalize: false,
			number:     false,
			wordList:   validate.AGWordList,
		},
		"less words than minWords and no number": {
			passphrase: "balcony", // Only 1 word and no number
			minWords:   2,
			maxWords:   3,
			separator:  "-",
			capitalize: false,
			number:     true,
			wordList:   validate.AGWordList,
		},
		"more words than maxWords and no capitalization": {
			passphrase: "abalone-balcony-gadfly-telepath", // 4 words, should be max 3
			minWords:   2,
			maxWords:   3,
			separator:  "-",
			capitalize: true, // No words are capitalized
			number:     false,
			wordList:   validate.AGWordList,
		},
		"non-capitalized word and no number at the end": {
			passphrase: "abalone-balcony-gadfly", // 'gadfly' not capitalized, no number
			minWords:   3,
			maxWords:   3,
			separator:  "-",
			capitalize: true,
			number:     true,
			wordList:   validate.AGWordList,
		},
		"non-wordlist word and no number": {
			passphrase: "abalone-balcony-fakeword", // 'fakeword' not in wordList, no number
			minWords:   3,
			maxWords:   3,
			separator:  "-",
			capitalize: false,
			number:     true,
			wordList:   validate.AGWordList,
		},
		"non-capitalized word, no number, and word not in wordlist": {
			passphrase: "abalone-balcony-incorrect", // 'incorrect' not in wordlist, no capitalization, no number
			minWords:   3,
			maxWords:   3,
			separator:  "-",
			capitalize: true,
			number:     true,
			wordList:   validate.AGWordList,
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			err := validate.Passphrase(testCase.passphrase, testCase.minWords, testCase.maxWords, testCase.separator, testCase.capitalize, testCase.number, testCase.wordList)
			if err == nil {
				t.Errorf("expected error for invalid passphrase \"%s\", but got no error", testCase.passphrase)
			}
		})
	}
}
