package validate

import (
	"errors"
	"fmt"
)

func checkLength(length int, minLength, maxLength, minLengthAllowed, maxLengthAllowed uint, units string) error {
	if minLength == 0 && maxLength == 0 {
		minLength = minLengthAllowed
		maxLength = maxLengthAllowed
	} else {
		if maxLength < minLength {
			return fmt.Errorf("maximum length can not be less than minimum length")
		}

		var errs []error

		if minLength < minLengthAllowed {
			errs = append(errs, fmt.Errorf("minimum length must not be less than %d", minLengthAllowed))
		}

		if maxLength > maxLengthAllowed {
			errs = append(errs, fmt.Errorf("maximum length must not exceed %d", maxLengthAllowed))
		}

		if len(errs) > 0 {
			return errors.Join(errs...)
		}
	}

	if uint(length) < minLength {
		return fmt.Errorf("length of %d is less than minimum length of %d %s", length, minLength, units)
	}

	if uint(length) > maxLength {
		return fmt.Errorf("length of %d exceeds maximum length of %d %s", length, maxLength, units)
	}

	return nil
}
