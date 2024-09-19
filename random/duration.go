package random

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func Duration(minDuration, maxDuration time.Duration) (time.Duration, error) {
	if minDuration < 0 {
		return 0, fmt.Errorf("minimum duration %v is negative", minDuration)
	}

	if maxDuration < 0 {
		return 0, fmt.Errorf("maximum duration %v is negative", maxDuration)
	}

	if minDuration > maxDuration {
		return 0, fmt.Errorf("minimum duration %v is less than maximum duration %v", minDuration, maxDuration)
	}

	random, err := rand.Int(rand.Reader, big.NewInt(int64(maxDuration - minDuration) + 1))
	if err != nil {
		return 0, fmt.Errorf("error generating a random number for calculatin the duration: %w", err)
	}

	return minDuration + time.Duration(random.Int64()), nil
}