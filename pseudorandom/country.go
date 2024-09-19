package pseudorandom

import (
	"math/rand/v2"
	"sort"
)

// CountryName returns a deterministic pseudo-random valid country name.
func CountryName(r *rand.Rand) string {
	// Pre-allocate the slice with the exact size.
	countriesSlice := make([]string, 0, len(Countries))
	for country := range Countries {
		countriesSlice = append(countriesSlice, country)
	}

	// Sort the slice to make returned value deterministic.
	sort.Strings(countriesSlice)

	// Return the pseudo-randomly selected country.
	return countriesSlice[r.IntN(len(countriesSlice))]
}

// CountryCode2 returns a deterministic pseudo-random valid 2-letter country code.
func CountryCode2(r *rand.Rand) string {
	// Pre-allocate the slice with the exact size.
	code2Slice := make([]string, 0, len(Countries))
	for _, country := range Countries {
		code2Slice = append(code2Slice, country.Code2)
	}

	// Sort the slice to make returned value deterministic.
	sort.Strings(code2Slice)

	// Return the pseudo-randomly selected country code 2.
	return code2Slice[r.IntN(len(code2Slice))]
}

// CountryCode3 returns a deterministic pseudo-random valid 3-letter country code.
func CountryCode3(r *rand.Rand) string {
	// Pre-allocate the slice with the exact size.
	code3Slice := make([]string, 0, len(Countries))
	for _, country := range Countries {
		code3Slice = append(code3Slice, country.Code3)
	}

	// Sort the slice to make returned value deterministic.
	sort.Strings(code3Slice)

	// Return the pseudo-randomly selected country code 3.
	return code3Slice[r.IntN(len(code3Slice))]
}