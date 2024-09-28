package random_test

import (
	"crypto/rand"
	"testing"

	"github.com/copartner6412/input/random"
	"github.com/copartner6412/input/validate"
)

const (
	minEmailLocalPartLengthAllowed  uint = 1
	maxEmailLocalPartLengthAllowed  uint = 64
	minEmailDomainPartLengthAllowed uint = minDomainLengthAllowed
	maxEmailDomainPartLengthAllowed uint = maxDomainLengthAllowed
	minEmailLengthAllowed           uint = minEmailLocalPartLengthAllowed + 1 + minEmailDomainPartLengthAllowed // 1 for @
	maxEmailLengthAllowed           uint = maxEmailLocalPartLengthAllowed + 1 + maxEmailDomainPartLengthAllowed
)

func FuzzEmail(f *testing.F) {
	f.Fuzz(func(t *testing.T, min, max uint, quotedLocalPart, ipDomainPart bool) {
		minLength, maxLength := randoms(min, max, minEmailLengthAllowed, maxEmailLengthAllowed)

		if (maxLength > 74 || minLength < 48) && ipDomainPart {
			t.Skip()
		}

		email1, err := random.Email(rand.Reader, minLength, maxLength, quotedLocalPart, ipDomainPart)
		if err != nil {
			t.Fatalf("error generating a random E-mail: %v", err)
		}

		err = validate.Email(email1, minLength, maxLength, quotedLocalPart, ipDomainPart)
		if err != nil {
			t.Fatalf("expected no error for valid random E-mail \"%s\": %v", email1, err)
		}
	})
}
