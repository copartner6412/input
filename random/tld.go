package random

import (
	"crypto/rand"
	"fmt"
	"io"
	"math/big"
	"sort"
)

func TLD(randomness io.Reader, minLength, maxLength uint) (string, error) {
	length, err := checkLength(randomness, minLength, maxLength, minTLDLengthAllowed, maxTLDLengthAllowed)
	if err != nil {
		return "", err
	}

	var tlds = map[int][]string{}
	for tld := range TLDs {
		tlds[len(tld)] = append(tlds[len(tld)], tld)
	}

	for key := range tlds {
		sort.Strings(tlds[key])
	}

	set := tlds[int(length)]

	random1, err := rand.Int(randomness, big.NewInt(int64(len(set))))
	if err != nil {
		return "", fmt.Errorf("error generating a random index for selecting a TLD: %w", err)
	}

	return set[random1.Int64()], nil
}

func CCTLD(randomness io.Reader) (string, error) {
	var tlds []string
	for _, country := range Countries {
		tlds = append(tlds, country.CCTLD)
	}

	index, err := rand.Int(randomness, big.NewInt(int64(len(tlds))))
	if err != nil {
		return "", fmt.Errorf("error generating a random index for selecting a ccTLD: %w", err)
	}

	return tlds[index.Int64()], nil
}
