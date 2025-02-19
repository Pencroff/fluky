package external

import (
	crand "crypto/rand"
	"encoding/binary"
	"github.com/Pencroff/fluky/source"
	"io"
	"math/rand"
)

const (
	SIZE_8k  = 1 << 13 // 8 kB
	SIZE_16k = 1 << 14 // 16 kB
	SIZE_32k = 1 << 15 // 32 kB
	SIZE_64k = 1 << 16 // 64 kB
)

// ShieldedRandReader is a random bytes reader that use a cryptographically secure pseudorandom number generator for
// seeding quick random number generator such as xorshift++ and let it generate limited number of random numbers.
// After reaching the limit, it will reseed the quick random number generator with a new seed from the cryptographically
// secure pseudorandom number generator.
// by default, it will use the Xoshiro256pp source and generate 16k batches of random bytes.
type ShieldedRandReader struct {
	counter uint64
	allowed uint64
	src     rand.Source64
	pos     uint8
	val     uint64
}

func (r *ShieldedRandReader) Read(b []byte) (n int, err error) {
	n = len(b)
	for i := 0; i < n; i++ {
		if r.counter >= r.allowed {
			seed, err := readSeed()
			if err != nil {
				return 0, err
			}
			r.src.Seed(seed)
			r.counter = 0
		}
		if r.pos == 0 {
			// Using top 7 bytes of 8 bytes random number, lower bytes might be less random
			r.val = r.src.Uint64() >> 8
			r.pos = 7
		}
		b[i] = byte(r.val)
		r.val >>= 8
		r.pos -= 1
		r.counter += 1
	}
	return n, nil
}

func NewShieldedRandReader() (io.Reader, error) {
	seed, err := readSeed()
	if err != nil {
		return nil, err
	}
	size := uint64(SIZE_16k)
	return NewShieldedRandReaderWithSizeAndSource(size, source.NewXoshiro256ppSource(seed))
}

func NewShieldedRandReaderWithSize(size uint64) (io.Reader, error) {
	seed, err := readSeed()
	if err != nil {
		return nil, err
	}
	return NewShieldedRandReaderWithSizeAndSource(size, source.NewXoshiro256ppSource(seed))
}

func NewShieldedRandReaderWithSizeAndSource(size uint64, source64 rand.Source64) (io.Reader, error) {
	return &ShieldedRandReader{
		counter: 0,
		allowed: size,
		src:     source64,
	}, nil
}

func readSeed() (int64, error) {
	b := make([]byte, 8)
	_, err := crand.Read(b)
	if err != nil {
		return -1, err
	}
	u := binary.BigEndian.Uint64(b)
	seed := int64(u >> 1)
	return seed, nil
}
