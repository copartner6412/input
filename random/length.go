package random

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
)

func checkLength(minLength, maxLength, minLengthAllowed, maxLengthAllowed uint) (uint, error) {
	if minLength == 0 && maxLength == 0 {
		minLength = minLengthAllowed
		maxLength = maxLengthAllowed
	} else {
		if maxLength < minLength {
			return 0, fmt.Errorf("maximum length can not be less than minimum length")
		}

		var errs []error

		if minLength < minLengthAllowed {
			errs = append(errs, fmt.Errorf("minimum length must not be less than %d", minLengthAllowed))
		}

		if maxLength > maxLengthAllowed {
			errs = append(errs, fmt.Errorf("maximum length must not exceed %d", maxLengthAllowed))
		}

		if len(errs) > 0 {
			return 0, errors.Join(errs...)
		}
	}

	random, err := rand.Int(rand.Reader, big.NewInt(int64(maxLength-minLength)+1))
	if err != nil {
		return 0, fmt.Errorf("error generating a random number for length: %w", err)
	}

	length := minLength + uint(random.Int64())

	return length, nil
}
