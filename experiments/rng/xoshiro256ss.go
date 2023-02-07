package rng

/*
#include <stdint.h>

typedef uint64_t u8;

static inline u8 rotl(const u8 x, int k) {
	return (x << k) | (x >> (64 - k));
}

typedef struct xoshiro256_state {
	u8 s0;
	u8 s1;
	u8 s2;
	u8 s3;
} xoshiro256_state;

u8 xoshiro256_next_ss(xoshiro256_state* rng) {
	const u8 result = rotl(rng->s1 *5, 7) * 9;

	const u8 t = rng->s1 << 17;

	rng->s2 ^= rng->s0;
	rng->s3 ^= rng->s1;
	rng->s1 ^= rng->s2;
	rng->s0 ^= rng->s3;

	rng->s2 ^= t;

	rng->s3 = rotl(rng->s3, 45);

	return result;
}

void xoshiro256_init_ss(xoshiro256_state* rng, u8 seed ) {
	u8 i;
	//print_xoshiro256_u8("seed", seed);
	rng->s0 = 0x25038409e9bcae15;
	rng->s1 = seed;
	rng->s2 = seed;
	rng->s3 = seed;
	for (i=0; i<29; ++i) {
		(void)xoshiro256_next_ss(rng);
	}
	//printf("====================\n\n\n");
}
*/
import "C"
import "github.com/Pencroff/go-toolkit/bits"

type Xoshiro256ss struct {
	st *C.xoshiro256_state
	s0 uint64
	s1 uint64
	s2 uint64
	s3 uint64
}

func NewXoshiro256ss(seed int64) *Xoshiro256ss {
	st := &(C.xoshiro256_state{})
	r := &Xoshiro256ss{st: st}
	r.Seed(seed)
	return r
}

func (x *Xoshiro256ss) Int63() int64 {
	return int64(x.Uint64() >> 1)
}

func (x *Xoshiro256ss) Uint64() uint64 {
	res := bits.RotateL64(x.s1*5, 7) * 9
	t := x.s1 << 17
	x.s2 ^= x.s0
	x.s3 ^= x.s1
	x.s1 ^= x.s2
	x.s0 ^= x.s3
	x.s2 ^= t
	x.s3 = bits.RotateL64(x.s3, 45)
	return res
}

func (x *Xoshiro256ss) Seed(seed int64) {
	x.s0 = 0x25038409e9bcae15
	x.s1 = uint64(seed)
	x.s2 = uint64(seed)
	x.s3 = uint64(seed)
	for i := 0; i < 29; i++ {
		_ = x.Uint64()
	}
	C.xoshiro256_init_ss(x.st, C.u8(seed))
}

func (x *Xoshiro256ss) Jump() *Xoshiro256pp {
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

func (x *Xoshiro256ss) LongJump() *Xoshiro256pp {
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

func (x *Xoshiro256ss) Uint64C() uint64 {
	return uint64(C.xoshiro256_next_ss(x.st))
}
