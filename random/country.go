package random

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// CountryName returns a random valid country name.
func CountryName() (string, error) {
	// Pre-allocate the slice with the exact size.
	countriesSlice := make([]string, 0, len(Countries))
	for country := range Countries {
		countriesSlice = append(countriesSlice, country)
	}

	// Generate a random index within the range of available countries.
	index, err := rand.Int(rand.Reader, big.NewInt(int64(len(countriesSlice))))
	if err != nil {
		return "", fmt.Errorf("error generating random index: %w", err)
	}

	// Return the randomly selected country.
	return countriesSlice[index.Int64()], nil
}

// CountryCode2 returns a random valid 2-letter country code.
func CountryCode2() (string, error) {
	// Pre-allocate the slice with the exact size.
	code2Slice := make([]string, 0, len(Countries))
	for _, country := range Countries {
		code2Slice = append(code2Slice, country.Code2)
	}

	// Generate a random index within the range of available countries.
	index, err := rand.Int(rand.Reader, big.NewInt(int64(len(code2Slice))))
	if err != nil {
		return "", fmt.Errorf("error generating random index: %w", err)
	}

	// Return the randomly selected country code 2.
	return code2Slice[index.Int64()], nil
}

// CountryCode3 returns a random valid 3-letter country code.
func CountryCode3() (string, error) {
	// Pre-allocate the slice with the exact size.
	code3Slice := make([]string, 0, len(Countries))
	for _, country := range Countries {
		code3Slice = append(code3Slice, country.Code3)
	}

	// Generate a random index within the range of available countries.
	index, err := rand.Int(rand.Reader, big.NewInt(int64(len(code3Slice))))
	if err != nil {
		return "", fmt.Errorf("error generating random index: %w", err)
	}

	// Return the randomly selected country code 3.
	return code3Slice[index.Int64()], nil
}