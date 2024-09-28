package random

import (
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"math/big"
	"strconv"
	"strings"
)

const (
	minPassphraseWords uint = 2
	maxPassphraseWords uint = 128
)

// Passphrase generates a cryptographically-secure random passphrase with a number of words of between the specified minWords and maxWords from the specified word list.
// You can not use space as separator.
//
// If wordList is nil, AGWordList (1Password) will be used as the default word list https://1password.com/txt/agwordlist.txt.
//
// Parameters:
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
//
// Note:
//   - You can not use space as separator.
//   - If wordList is nil, AGWordList (1Password) will be used as the default word list https://1password.com/txt/agwordlist.txt.
//
// Available Word Lists:
//   - AGWordList (1Password)
//   - EFFLongWordList (Bitwarden)
//
// AGWordList information:
//   - Word List Length: 18176 words
//   - Minimum Character Length: 5 characters
//   - Maximum Character Length: 10 characters
//   - Average Character Length: 8.2 characters
//   - Recommended Word Count: 4 words or more
//   - Safe Maximum Recommended Word Count: 6 words (Not for MariaDB/MySQL)
//
// EFFLongWordList information:
//   - Word List Length: 7776 words
//   - Minimum Word Length: 3 characters
//   - Maximum Word Length: 9 characters
//   - Average Word Length: 7.0 characters
//   - Recommended Word Count: 5 words or more
//   - Safe Maximum Recommended Word Count: 7 words (Not for MariaDB/MySQL)
//
// It would be better to use a random password instead of passphrase for MariaDB/MySQL)
//
// Safe Maximum Passphrase Word Count (AGWordList/EFFLongWordList):
//   - OpenSSH 23/25 words
//   - Linux PAM 11/12 words
//   - Windows 11/12 words
//   - PostgreSQL 9/10 words
//   - GitHub 6/7 words
//   - Facebook 18/20 words
//   - Twitter 9/10 words
//   - Google 9/10 words
func Passphrase(randomness io.Reader, minWords, maxWords uint, separator string, capitalize bool, number bool, wordList []string) (string, error) {
	// Ensure that maxWords is not less than minWords.
	if maxWords < minWords {
		return "", errors.New("maximum number of words allowed in the passphrase can not be less than minimum words allowed")
	}

	// Validate the number of words (must be between 2 and 128).
	if minWords < minPassphraseWords || maxWords > maxPassphraseWords {
		return "", fmt.Errorf("number of words must be between 2 and 128")
	}

	// Determine the actual number of words in the passphrase to be generated.
	random1, err := rand.Int(randomness, big.NewInt(int64(maxWords-minWords+1)))
	if err != nil {
		return "", fmt.Errorf("error generating a random number for number of words in passphrase: %w", err)
	}

	wordNumber := int(random1.Int64()) + int(minWords)

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
	for i := 0; i < wordNumber; i++ {
		// Select a random word from the word list.
		random2, err := rand.Int(randomness, big.NewInt(int64(len(wordList))))
		if err != nil {
			return "", fmt.Errorf("error generating a random index for selecting a word from word list: %w", err)
		}
		word := wordList[random2.Int64()]

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
		random3, err := rand.Int(randomness, big.NewInt(int64(len(passphraseWords))))
		if err != nil {
			return "", fmt.Errorf("error generating a random index for selecting a word from words of passphrase: %w", err)
		}
		randomIndex := random3.Int64()
		// Generate a random digit (0-9) not equal to separator.
		random4, err := rand.Int(randomness, big.NewInt(10))
		if err != nil {
			return "", fmt.Errorf("error generating a random number for: %w", err)
		}
		randomDigit := random4.Int64()

		if strings.ContainsAny(separator, strconv.Itoa(int(randomDigit))) {
			randomDigit = (randomDigit + 1) % 10
		}

		// Append the digit to the selected word.
		passphraseWords[randomIndex] += strconv.FormatInt(int64(randomDigit), 10)
	}

	// Join the words with the specified separator.
	return strings.Join(passphraseWords, separator), nil
}
