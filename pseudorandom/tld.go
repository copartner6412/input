package pseudorandom

import (
	"math/rand/v2"
	"sort"
)

// TLD generates and returns a deterministic pseudo-random valid top-level domain (TLD)
// from the list of TLDs recognized by IANA using the provided random source.
func TLD(r *rand.Rand, minLength, maxLength uint) (string, error) {
	length, err := checkLength(r, minLength, maxLength, minTLDLengthAllowed, maxTLDLengthAllowed)
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

	return set[r.IntN(len(set))], nil
}

// CCTLD generates and returns a deterministic pseudo-random valid country code top-level domain (ccTLD)
// based on the list of countries using the provided random source.
func CCTLD(r *rand.Rand) string {
	var tlds []string
	for _, country := range Countries {
		tlds = append(tlds, country.CCTLD)
	}
	sort.Strings(tlds)
	return tlds[r.IntN(len(tlds))]
}
