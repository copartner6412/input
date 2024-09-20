package pseudorandom

import (
	"encoding/binary"
	"math/rand/v2"
)

type Reader struct {
	src *rand.Rand
}

// New creates a new Reader with the given seed.
func New(seed1, seed2 uint64) *Reader {
	return &Reader{
		src: rand.New(rand.NewPCG(seed1, seed2)),
	}
}

// Read implements the io.Reader interface.
func (r *Reader) Read(p []byte) (n int, err error) {
	for n+8 <= len(p) {
		val := r.src.Uint64()
		binary.LittleEndian.PutUint64(p[n:], uint64(val))
		n += 8
	}
	if n < len(p) {
		val := r.src.Uint64()
		for i := 0; i < len(p)-n; i++ {
			p[n+i] = byte(val)
			val >>= 8
		}
		n = len(p)
	}
	return n, nil
}
