package pseudorandom

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"strconv"
	"strings"
)

const (
	minPassphraseWordsAllowed uint = 2
	maxPassphraseWordsAllowed uint = 128
)

// Passphrase generates a deterministic pseudo-random passphrase with a number of words of between the specified minWords and maxWords from the specified word list.
// You can not use space as separator.
//
// If wordList is nil, AGWordList (1Password) will be used as the default word list https://1password.com/txt/agwordlist.txt.
//
// Parameters:
//   - r: Randomness source.
//   - minWords: The minimum number of words in the passphrase (between 2 and 128).
//   - maxWords: The maximum number of words in the passphrase (between 2 and 128).
//   - separator: The string used to separate words (cannot be a space).
//   - capitalize: If true, capitalize the first letter of each word.
//   - number: If true, append a random digit to one of the words.
//   - wordList: A custom word list to use for generating the passphrase. If nil, AGWordList is used.
//
// Returns:
//   - A string containing the generated passphrase.
//   - An error if invalid parameters are passed.
func Passphrase(r *rand.Rand, minWords, maxWords uint, separator string, capitalize bool, number bool, wordList []string) (string, error) {
	if minWords == 0 && maxWords == 0 {
		minWords = minPassphraseWordsAllowed
		maxWords = maxPassphraseWordsAllowed
	} else {
		if maxWords < minWords {
			return "", errors.New("maximum number of words allowed in the passphrase can not be less than minimum words allowed")
		}

		if minWords < minPassphraseWordsAllowed || maxWords > maxPassphraseWordsAllowed {
			return "", fmt.Errorf("number of words must be between 2 and 128")
		}
	}

	// Determine the actual number of words in the passphrase to be generated.
	wordNumber := minWords + r.UintN(maxWords-minWords+1)

	// Validate that the separator does not have a space.
	if strings.Contains(separator, " ") {
		return "", fmt.Errorf("separator cannot be a space")
	}

	// Use the default word list (AGWordList) if none is provided.
	if wordList == nil {
		wordList = AGWordList
	}

	// Generate the random passphrase.
	passphraseWords := []string{} // Preallocate the slice for efficiency.
	for i := 0; i < int(wordNumber); i++ {
		// Select a random word from the word list.
		word := wordList[r.IntN(len(wordList))]

		// Capitalize the first letter of the word if requested.
		if capitalize {
			word = (strings.ToUpper(string([]rune(word)[0])) + string([]rune(word)[1:]))
		}

		// Append the word to the passphrase.
		passphraseWords = append(passphraseWords, word)
	}

	// Optionally, append a random digit to one of the words.
	if number {
		// Randomly select a word to append the digit to.
		randomIndex := r.IntN(len(passphraseWords))
		// Generate a random digit (0-9) not equal to separator.
		randomDigit := r.IntN(10)

		if strings.ContainsAny(separator, strconv.Itoa(randomDigit)) {
			randomDigit = (randomDigit + 1) % 10
		}

		// Append the digit to the selected word.
		passphraseWords[randomIndex] += strconv.FormatInt(int64(randomDigit), 10)
	}

	// Join the words with the specified separator.
	return strings.Join(passphraseWords, separator), nil
}
