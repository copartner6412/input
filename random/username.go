package random

import (
	"crypto/rand"
	"fmt"
	"io"
	"math/big"
	"strconv"
	"strings"
)

// Username generates a random username. You can capitalize the username or add 4 random digits to the end of username.
//
// If wordList is nil, AGWordList (1Password) will be used as the default word list.
//
// Available Word Lists:
//   - AGWordList (1Password)
//   - EFFLongWordList (Bitwarden)
func Username(randomness io.Reader, capitalize bool, number bool, wordList []string) (string, error) {
	// Validate the input.
	if wordList == nil {
		wordList = AGWordList
	}

	// Choose a random word from the word list.
	randomIndex, err := rand.Int(randomness, big.NewInt(int64(len(wordList))))
	if err != nil {
		return "", fmt.Errorf("error generating random index: %w", err)
	}
	word := wordList[randomIndex.Int64()]

	// Capitalize the word if specified.
	if capitalize {
		word = strings.ToUpper(string(word[0])) + string(word[1:])
	}

	// Add random digits to the end of the word.
	for i := 0; i < 4; i++ {
		// Generate a random digit (0-9)
		randomDigit, err := rand.Int(randomness, big.NewInt(10))
		if err != nil {
			return "", fmt.Errorf("error generating random digit: %w", err)
		}
		word += strconv.Itoa(int(randomDigit.Int64()))
	}

	return word, nil
}
