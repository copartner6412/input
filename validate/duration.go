package validate

import (
	"errors"
	"fmt"
	"time"
)

func Duration(duration time.Duration, minDuration, maxDuration time.Duration) error {
	if duration < 0 {
		return errors.New("duration can not be negative")
	}

	if minDuration > maxDuration {
		return fmt.Errorf("minimum duration %v is greater than maximum duration %v", minDuration, maxDuration)
	}

	if minDuration < 0 {
		return errors.New("minimum duration can not be negative")
	}

	if duration < minDuration {
		return fmt.Errorf("duration %v is less than minimum duration %v", duration, minDuration)
	}

	if duration > maxDuration {
		return fmt.Errorf("duration %v is more than maximum duration %v", duration, maxDuration)
	}

	return nil
}