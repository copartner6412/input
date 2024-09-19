package validate

import (
	"fmt"
	"time"
)

func Duration(duration time.Duration, minDuration, maxDuration time.Duration) error {
	if duration < 0 {
		return fmt.Errorf("duration %v is negative", duration)
	}

	if minDuration < 0 {
		return fmt.Errorf("minimum duration %v is negative", minDuration)
	}

	if maxDuration < 0 {
		return fmt.Errorf("maximum duration %v is negative", maxDuration)
	}

	if minDuration > maxDuration {
		return fmt.Errorf("minimum duration %v is greater than maximum duration %v", minDuration, maxDuration)
	}

	if duration < minDuration {
		return fmt.Errorf("duration %v is less than minimum duration %v", duration, minDuration)
	}

	if duration > maxDuration {
		return fmt.Errorf("duration %v is more than maximum duration %v", duration, maxDuration)
	}

	return nil
}