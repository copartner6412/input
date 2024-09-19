package validate_test

import (
	"bytes"
	"testing"

	"github.com/copartner6412/input/pseudorandom"
	"github.com/copartner6412/input/validate"
)

const (
	minByteSliceLengthAllowed uint = 1
	maxByteSliceLengthAllowed uint = 8192
)

func FuzzByteSuccessfulForValidPseudorandomInput(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed1, seed2 uint64, min, max uint) {
		r1, r2, minLength, maxLength := randoms(seed1, seed2, min, max, minByteSliceLengthAllowed, maxByteSliceLengthAllowed)
		
		bytes1, err := pseudorandom.Bytes(r1, minLength, maxLength)
		if err != nil {
			t.Fatalf("error generating a pseudo-random byte slice: %v", err)
		}

		err = validate.Bytes(bytes1, minLength, maxLength)
		if err != nil {
			t.Fatalf("expected no error for valid pseudo-random byte slice, but got error: %v", err)
		}
		
		bytes2, err := pseudorandom.Bytes(r2, minLength, maxLength)
		if err != nil {
			t.Fatalf("error regenerating the pseudo-random byte slice: %v", err)
		}

		if !bytes.Equal(bytes1, bytes2) {
			t.Fatal("not deterministic")
		}
	})
}

func TestByteSuccessfulForValidByteSlice(t *testing.T) {
	testCases := map[string]struct {
		byteSlice []byte
		minLength uint
		maxLength uint
	}{
		"Valid byte slice of length equal to minLength": {
			byteSlice: []byte{0x01, 0x02, 0x03},
			minLength: 3,
			maxLength: 5,
		},
		"Valid byte slice of length equal to maxLength": {
			byteSlice: []byte{0x01, 0x02, 0x03, 0x04, 0x05},
			minLength: 3,
			maxLength: 5,
		},
		"Valid byte slice of length between minLength and maxLength": {
			byteSlice: []byte{0x01, 0x02, 0x03, 0x04},
			minLength: 3,
			maxLength: 5,
		},
		"Valid byte slice at the boundary of max length": {
			byteSlice: []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08},
			minLength: 3,
			maxLength: 8,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			err := validate.Bytes(tc.byteSlice, tc.minLength, tc.maxLength)
			if err != nil {
				t.Errorf("expected no error for byte slice %v, but got error: %v", tc.byteSlice, err)
			}
		})
	}
}

func TestByteFailsForInvalidByteSlice(t *testing.T) {
	testCases := map[string]struct {
		byteSlice []byte
		minLength uint
		maxLength uint
	}{
		"Valid empty byte slice with minLength 0": {
			byteSlice: []byte{},
			minLength: 0,
			maxLength: 5,
		},
		"maxLength less than minLength": {
			byteSlice: []byte{0x01, 0x02, 0x03},
			minLength: 4,
			maxLength: 3,
		},
		"minLength below system minimum (minByteSliceLength)": {
			byteSlice: []byte{0x01},
			minLength: 0, // Assuming minByteSliceLength is greater than 0
			maxLength: 5,
		},
		"maxLength exceeds system maximum (maxByteSliceLength)": {
			byteSlice: []byte{0x01, 0x02},
			minLength: 1,
			maxLength: 8193, // Assuming maxByteSliceLength is less than 1000
		},
		"byte slice shorter than minLength": {
			byteSlice: []byte{0x01, 0x02},
			minLength: 3,
			maxLength: 5,
		},
		"byte slice longer than maxLength": {
			byteSlice: []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06},
			minLength: 3,
			maxLength: 5,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			err := validate.Bytes(tc.byteSlice, tc.minLength, tc.maxLength)
			if err == nil {
				t.Errorf("expected error for byte slice %v, but got no error", tc.byteSlice)
			}
		})
	}
}


