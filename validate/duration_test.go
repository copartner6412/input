package validate_test

import (
	"testing"
	"time"

	"github.com/copartner6412/input/validate"
)

func TestDurationSuccessfulForValidInput(t *testing.T) {
	testCases := map[string]struct{
		duration, minDuration, maxDuration time.Duration
	}{
		"minEqualMax": {
			duration:    5 * time.Second,
			minDuration: 5 * time.Second,
			maxDuration: 5 * time.Second,
		},
		"durationEqualsMin": {
			duration:    2 * time.Second,
			minDuration: 2 * time.Second,
			maxDuration: 10 * time.Second,
		},
		"durationEqualsMax": {
			duration:    10 * time.Second,
			minDuration: 2 * time.Second,
			maxDuration: 10 * time.Second,
		},
		"durationBetweenMinMax": {
			duration:    5 * time.Second,
			minDuration: 2 * time.Second,
			maxDuration: 10 * time.Second,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			err := validate.Duration(tc.duration, tc.minDuration, tc.maxDuration)
			if err != nil {
				t.Errorf("expected no error for valid input, but got error: %v", err)
			}
		})
	}
}


func TestDurationFailsForInvalidInput(t *testing.T) {
	testCases := map[string]struct{
		duration, minDuration, maxDuration time.Duration
	}{
		"minGreaterThanMax": {
			duration:    5 * time.Second,
			minDuration: 10 * time.Second,
			maxDuration: 2 * time.Second,
		},
		"durationLessThanMin": {
			duration:    1 * time.Second,
			minDuration: 2 * time.Second,
			maxDuration: 10 * time.Second,
		},
		"durationGreaterThanMax": {
			duration:    15 * time.Second,
			minDuration: 2 * time.Second,
			maxDuration: 10 * time.Second,
		},
		"negativeDuration": {
			duration:    -5 * time.Second,
			minDuration: 2 * time.Second,
			maxDuration: 10 * time.Second,
		},
		"negativeMinDuration": {
			duration:    5 * time.Second,
			minDuration: -1 * time.Second,
			maxDuration: 10 * time.Second,
		},
		"negativeMaxDuration": {
			duration:    5 * time.Second,
			minDuration: 2 * time.Second,
			maxDuration: -10 * time.Second,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			err := validate.Duration(tc.duration, tc.minDuration, tc.maxDuration)
			if err == nil {
				t.Error("expected error for invalid input, but got no error")
			}
		})
	}
}
