package validate

import "math/big"

const (
	minBitSizeAllowed uint = 1
	maxBitSizeAllowed uint = 4096
)

func BigInteger(serialNumber *big.Int, minBitSize, maxBitSize uint) error {
	err := checkLength(serialNumber.BitLen(), minBitSize, maxBitSize, minBitSizeAllowed, maxBitSizeAllowed, "bits")
	if err != nil {
		return err
	}

	return nil
}