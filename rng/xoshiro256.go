package rng

import (
	"github.com/Pencroff/go-toolkit/bits"
)

type Xoshiro256pp struct {
	s0 uint64
	s1 uint64
	s2 uint64
	s3 uint64
}

func NewXoshiro256pp(seed int64) *Xoshiro256pp {
	r := &Xoshiro256pp{}
	r.Seed(seed)
	return r
}

func (x *Xoshiro256pp) Int63() int64 {
	return int64(x.Uint64() >> 1)
}

func (x *Xoshiro256pp) Uint64() uint64 {
	res := bits.RotateL64(x.s0+x.s3, 23) + x.s0
	t := x.s1 << 17
	x.s2 ^= x.s0
	x.s3 ^= x.s1
	x.s1 ^= x.s2
	x.s0 ^= x.s3
	x.s2 ^= t
	x.s3 = bits.RotateL64(x.s3, 45)
	return res
}

func (x *Xoshiro256pp) Seed(seed int64) {
	x.s0 = 0x25038409e9bcae15
	x.s1 = uint64(seed)
	x.s2 = uint64(seed)
	x.s3 = uint64(seed)
	for i := 0; i < 29; i++ {
		_ = x.Uint64()
	}
}
