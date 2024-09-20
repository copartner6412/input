package pseudorandom

import (
	"math/rand/v2"
)

type Reader struct {
	src *rand.Rand
}

// New creates a new Reader with the given seed.
func New(r *rand.Rand) Reader {
	return Reader{
		src: r,
	}
}

// Read implements the io.Reader interface.
func (r Reader) Read(p []byte) (n int, err error) {
	for n < len(p) {
		p[n] = byte(r.src.UintN(256))
		n += 1
	}

	return n, nil
}
