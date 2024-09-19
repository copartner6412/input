package pseudorandom

import (
	"fmt"
	"math/rand/v2"
	"strings"

)

const (
	minDomainLengthAllowed   uint = 1
	maxDomainLengthAllowed   uint = 253
	minDomainWithValidTLDLengthAllowed uint = minTLDLengthAllowed + 2
	mindomainWithValidccTLDLength uint = ccTLDLength + 2
	maxSubdomainCount uint = 127 // (253 + 1) / 2
)

// Domain generates a deterministic pseudo-random Internet domain name using the provided random source.
// The domain consists of a random number of subdomains, with a total length of up to 253 characters.
// It ensures at least one subdomain and avoids subdomains with invalid patterns.
func Domain(r *rand.Rand, minLength, maxLength uint) (domain string, err error) {
	length, err := checkLength(r, minLength, maxLength, minDomainLengthAllowed, maxDomainLengthAllowed)
	if err != nil {
		return "", err
	}

	subdomains := []string{} // Slice to hold generated subdomains.
	var trackLength uint     // Keep track of the total domain length as we add subdomains.
	var subdomainLength uint

	// Generate subdomains while ensuring the total length stays within the limit.
	for {
		var subdomain string
		// Add a subdomain which doesn't contain "wwww".
		for {
			subdomain, err = Subdomain(r, minSubdomainLengthAllowed, maxSubdomainLengthAllowed)
			if err != nil {
				return "", err
			}

			// Avoid using "www" for any subdomain other than the first.
			if len(subdomains) != 0 && strings.Contains(subdomain, "www") {
				continue // Regenerate subdomain if it is "www" for non-initial positions.
			} else {
				break
			}
		}

		subdomains = append(subdomains, subdomain) // Add the subdomain to the list.

		subdomainLength = uint(len([]rune(subdomain)))

		trackLength += subdomainLength // Update the total length with the new subdomain.

		if trackLength == length - 1 {
			if subdomainLength == maxSubdomainLengthAllowed {
				subdomains[len(subdomains) - 1] = subdomain[:maxSubdomainLengthAllowed - 1]
				
				if strings.HasSuffix(subdomains[len(subdomains) - 1], "-") {
					subdomains[len(subdomains) - 1] = subdomain[:maxSubdomainLengthAllowed - 2] + string(lowerAlphanumericalRunes[r.IntN(len(lowerAlphanumericalRunes))])
				}

				lastSubdomain, err := Subdomain(r, 1, 1)
				if err != nil {
					return "", fmt.Errorf("error generating the last pseudo-random subdomain: %w", err)
				}

				subdomains = append(subdomains, lastSubdomain)

				break
			} else {
				lastSubdomain, err := Subdomain(r, uint(len(subdomain)) + 1, uint(len(subdomain)) + 1)
				if err != nil {
					return "", fmt.Errorf("error generating the last pseudo-random subdomain: %w", err)
				}

				subdomains[len(subdomains) - 1] = lastSubdomain

				break
			}
		} else if trackLength >= length {

			lastSubdomain := subdomain[:subdomainLength - (trackLength - length)]

			if strings.HasSuffix(lastSubdomain, "-") {
				lastSubdomain = lastSubdomain[:len(lastSubdomain)-1] + string(lowerAlphanumericalRunes[r.IntN(len(lowerAlphanumericalRunes))])
			}

			subdomains[len(subdomains) - 1] = lastSubdomain

			break // Stop adding subdomains if adding more would exceed the total allowed length.
		} else {
			trackLength += 1 // Account for the dot between subdomains.
		}
	}

	// Join subdomains into a full domain string, separated by dots.
	domain = strings.Join(subdomains, ".")

	return domain, nil
}


// DomainWithValidTLD generates a deterministic pseudo-random domain name and replaces the last part with a valid top-level domain (TLD).
func DomainWithValidTLD(r *rand.Rand, minLength, maxLength uint) (domain string, err error) {
	length, err := checkLength(r, minLength, maxLength, minDomainWithValidTLDLengthAllowed, maxDomainLengthAllowed)
	if err != nil {
		return "", err
	}

	maxTLDLength := length - minDomainLengthAllowed - 1 // 1 for dot
	if maxTLDLength > maxTLDLengthAllowed {
		maxTLDLength = maxTLDLengthAllowed
	}

	tldLength := minTLDLengthAllowed + r.UintN(maxTLDLength - minTLDLengthAllowed + 1)

	tld, err := TLD(r, tldLength, tldLength)
	if err != nil {
		return "", fmt.Errorf("error generating a pseudo-random gTLD: %w", err)
	}

	firstPartLength := length - tldLength - 1 // 1 for dot

	firstPart, err := Domain(r, firstPartLength, firstPartLength)
	if err != nil {
		return "", fmt.Errorf("error generating a pseudo-random domain part: %w", err)
	}

	parts := []string{firstPart, tld}
	domain = strings.Join(parts, ".")


	return domain, nil
}

// DomainWithValidCCTLD generates a deterministic pseudo-random domain name and replaces the last part with a valid country-code TLD (ccTLD).
func DomainWithValidCCTLD(r *rand.Rand, minLength, maxLength uint) (domain string, err error) {
	length, err := checkLength(r, minLength, maxLength, mindomainWithValidccTLDLength, maxDomainLengthAllowed)
	if err != nil {
		return "", err
	}

	firstPartLength := length - ccTLDLength - 1

	ccTLD := CCTLD(r)

	firstPart, err := Domain(r, firstPartLength, firstPartLength)
	if err != nil {
		return "", fmt.Errorf("error generating a pseudo-random domain part: %w", err)
	}
	parts := []string{firstPart, ccTLD}
	domain = strings.Join(parts, ".")


	return domain, nil
}