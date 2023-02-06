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

func (x *Xoshiro256pp) Jump() *Xoshiro256pp {
	s0, s1, s2, s3 := x.s0, x.s1, x.s2, x.s3
	r := &Xoshiro256pp{s0, s1, s2, s3}
	jmp := []uint64{0x180ec6d33cfd0aba, 0xd5a61266f0c9392c, 0xa9582618e03fc9aa, 0x39abdc4529b1661c}
	s0, s1, s2, s3 = 0, 0, 0, 0
	for _, jump := range jmp {
		var b uint
		for b = 0; b < 64; b++ {
			if jump&(1<<b) != 0 {
				s0 ^= x.s0
				s1 ^= x.s1
				s2 ^= x.s2
				s3 ^= x.s3
			}
			x.Uint64()
		}
	}
	x.s0 = s0
	x.s1 = s1
	x.s2 = s2
	x.s3 = s3
	return r
}

func (x *Xoshiro256pp) LongJump() *Xoshiro256pp {
	s0, s1, s2, s3 := x.s0, x.s1, x.s2, x.s3
	r := &Xoshiro256pp{s0, s1, s2, s3}
	jmp := []uint64{0x76e15d3efefdcbbf, 0xc5004e441c522fb3, 0x77710069854ee241, 0x39109bb02acbe635}
	s0, s1, s2, s3 = 0, 0, 0, 0
	for _, jump := range jmp {
		var b uint
		for b = 0; b < 64; b++ {
			if jump&(1<<b) != 0 {
				s0 ^= x.s0
				s1 ^= x.s1
				s2 ^= x.s2
				s3 ^= x.s3
			}
			x.Uint64()
		}
	}
	x.s0 = s0
	x.s1 = s1
	x.s2 = s2
	x.s3 = s3
	return r
}
