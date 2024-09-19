package random_test

func randoms(min, max, minLengthAllowed, maxLengthAllowed uint) (minLength, maxLength uint) {
	minLength = minLengthAllowed + min%(maxLengthAllowed - minLengthAllowed + 1)
	maxLength = minLength + max%(maxLengthAllowed - minLength + 1)
	return minLength, maxLength
}