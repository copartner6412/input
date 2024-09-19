package pseudorandom

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func Duration(r *rand.Rand, minDuration, maxDuration time.Duration) (time.Duration, error) {
	if minDuration < 0 {
		return 0, fmt.Errorf("minimum duration %v is negative", minDuration)
	}

	if maxDuration < 0 {
		return 0, fmt.Errorf("maximum duration %v is negative", maxDuration)
	}

	if minDuration > maxDuration {
		return 0, fmt.Errorf("minimum duration %v is less than maximum duration %v", minDuration, maxDuration)
	}

	return minDuration + time.Duration(r.Uint64N(uint64(maxDuration - minDuration) + 1)), nil
}