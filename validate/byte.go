package validate

const (
	minByteSliceLengthAllowed uint = 1
	maxByteSliceLengthAllowed uint = 8192
)

// Bytes validates a byte slice to ensure that its length falls between the specified minLength and maxLength.
//
// Parameters:
// - bytes: The byte slice to validate.
// - minLength: The minimum allowed length for the byte slice. Must be between 1 and 8192 bytes.
// - maxLength: The maximum allowed length for the byte slice. Must be between 1 and 8192 bytes.
//
// Returns:
// - An error if the byte slice does not meet the length requirements or if any input validation fails.
func Bytes(bytes []byte, minLength, maxLength uint) error {
	err := checkLength(len(bytes), minLength, maxLength, minByteSliceLengthAllowed, maxByteSliceLengthAllowed, "bytes")
	if err != nil {
		return err
	}

	return nil
}
