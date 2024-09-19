package pseudorandom_test

import (
	"testing"

	"github.com/copartner6412/input/pseudorandom"
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
	f.Fuzz(func(t *testing.T, seed1, seed2 uint64, min, max uint, quotedLocalPart, ipDomainPart bool) {
		r1, r2, minLength, maxLength := randoms(seed1, seed2, min, max, minEmailLengthAllowed, maxEmailLengthAllowed)
		
		if (maxLength > 74 || minLength < 48) && ipDomainPart {
			t.Skip()
		}

		email1, err := pseudorandom.Email(r1, minLength, maxLength, quotedLocalPart, ipDomainPart)
		if err != nil {
			t.Fatalf("error generating a pseudo-random E-mail: %v", err)
		}

		err = validate.Email(email1, minLength, maxLength, quotedLocalPart, ipDomainPart)
		if err != nil {
			t.Fatalf("expected no error for valid pseudo-random E-mail \"%s\": %v", email1, err)
		}

		email2, err := pseudorandom.Email(r2, minLength, maxLength, quotedLocalPart, ipDomainPart)
		if err != nil {
			t.Fatalf("error regenerating the pseudo-random E-mail: %v", err)
		}

		if email1 != email2 {
			t.Fatal("not deterministic")
		}
	})
}
