package pseudorandom

import (
	"math/rand/v2"
	"strconv"
	"strings"
)

// Username generates a pseudo-random username. You can capitalize the username or add 4 random digits to the end of username.
//
// If wordList is nil, AGWordList (1Password) will be used as the default word list.
//
// Available Word Lists:
//   - AGWordList (1Password)
//   - EFFLongWordList (Bitwarden)
func Username(r *rand.Rand, capitalize bool, number bool, wordList []string) string {
	// Validate the input.
	if wordList == nil {
		wordList = AGWordList
	}

	// Choose a random word from the word list.
	word := wordList[r.IntN(len(wordList))]

	// Capitalize the word if specified.
	if capitalize {
		word = strings.ToUpper(string(word[0])) + string(word[1:])
	}

	// Add random digits to the end of the word.
	for i := 0; i < 4; i++ {
		// Generate a random digit (0-9)
		word += strconv.Itoa(r.IntN(10))
	}

	return word
}
