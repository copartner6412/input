package pseudorandom_test

import "math/rand/v2"

func randoms(seed1, seed2 uint64, min, max, minLengthAllowed, maxLengthAllowed uint) (r1, r2 *rand.Rand, minLength, maxLength uint) {
	r1 = rand.New(rand.NewPCG(seed1, seed2))
	r2 = rand.New(rand.NewPCG(seed1, seed2))
	minLength = min%(maxLengthAllowed - minLengthAllowed + 1) + minLengthAllowed
	maxLength = minLength + max%(maxLengthAllowed - minLength + 1)
	return r1, r2, minLength, maxLength
}