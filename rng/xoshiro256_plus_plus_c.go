package rng

/*
#include <stdio.h>
#include <stdint.h>

// This is xoshiro256++ 1.0, one of our all-purpose, rock-solid generators.
// It has excellent (sub-ns) speed, a state (256 bits) that is large
// enough for any parallel application, and it passes all tests we are
// aware of.
//
// For generating just floating-point numbers, xoshiro256+ is even faster.
//
// The state must be seeded so that it is not everywhere zero. If you have
// a 64-bit seed, we suggest to seed a splitmix64 generator and use its
// output to fill s.

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

void print_xoshiro256(char *pref, xoshiro256_state *x ) {
	printf("%s:\n", pref);
	printf("a: %016llx\tb: %016llx\t c: %016llx\t d: %016llx\n", x->s0, x->s1, x->s2, x->s3);
}
void print_xoshiro256_u8(char *pref, u8 v) {
	printf("%s: %016llx\n", pref, v);
}

u8 xoshiro256_next(xoshiro256_state* rng) {
	const u8 result = rotl(rng->s0 + rng->s3, 23) + rng->s0;

	const u8 t = rng->s1 << 17;

	rng->s2 ^= rng->s0;
	rng->s3 ^= rng->s1;
	rng->s1 ^= rng->s2;
	rng->s0 ^= rng->s3;

	rng->s2 ^= t;

	rng->s3 = rotl(rng->s3, 45);

	return result;
}


// This is the jump function for the generator. It is equivalent
// to 2^128 calls to next(); it can be used to generate 2^128
// non-overlapping subsequences for parallel computations.

void xoshiro256_jump(xoshiro256_state* rng) {
	static const u8 JUMP[] = { 0x180ec6d33cfd0aba, 0xd5a61266f0c9392c, 0xa9582618e03fc9aa, 0x39abdc4529b1661c };

	u8 s0 = 0;
	u8 s1 = 0;
	u8 s2 = 0;
	u8 s3 = 0;
	for (int i = 0; i < sizeof JUMP / sizeof *JUMP; i++)
	for (int b = 0; b < 64; b++) {
		if (JUMP[i] & UINT64_C(1) << b) {
			s0 ^= rng->s0;
			s1 ^= rng->s1;
			s2 ^= rng->s2;
			s3 ^= rng->s3;
		}
		xoshiro256_next(rng);
	}

	rng->s0 = s0;
	rng->s1 = s1;
	rng->s2 = s2;
	rng->s3 = s3;
}



// This is the long-jump function for the generator. It is equivalent to
// 2^192 calls to next(); it can be used to generate 2^64 starting points,
// from each of which jump() will generate 2^64 non-overlapping
// subsequences for parallel distributed computations.

void xoshiro256_long_jump(xoshiro256_state* rng) {
	static const u8 LONG_JUMP[] = { 0x76e15d3efefdcbbf, 0xc5004e441c522fb3, 0x77710069854ee241, 0x39109bb02acbe635 };

	u8 s0 = 0;
	u8 s1 = 0;
	u8 s2 = 0;
	u8 s3 = 0;
	for (int i = 0; i < sizeof LONG_JUMP / sizeof *LONG_JUMP; i++)
	for (int b = 0; b < 64; b++) {
		if (LONG_JUMP[i] & UINT64_C(1) << b) {
			s0 ^= rng->s0;
			s1 ^= rng->s1;
			s2 ^= rng->s2;
			s3 ^= rng->s3;
		}
		xoshiro256_next(rng);
	}

	rng->s0 = s0;
	rng->s1 = s1;
	rng->s2 = s2;
	rng->s3 = s3;
}

void xoshiro256_init(xoshiro256_state* rng, u8 seed ) {
	u8 i;
	print_xoshiro256_u8("seed", seed);
	rng->s0 = 0xf1ea5eed, rng->s1 = rng->s2 = rng->s3 = seed;
	for (i=0; i<31; ++i) {
		(void)xoshiro256_next(rng);
	}
	print_xoshiro256("state", rng);
	printf("====================\n\n\n");
}
*/
import "C"

type Xoshiro256ppC struct {
	st *C.xoshiro256_state
}

func (x *Xoshiro256ppC) Int63() int64 {
	return int64(x.Uint64() >> 1)
}

func (x *Xoshiro256ppC) Seed(seed int64) {
	C.xoshiro256_init(x.st, C.u8(seed))
}

func (x *Xoshiro256ppC) Uint64() uint64 {
	return uint64(C.xoshiro256_next(x.st))
}

func NewXoshiro256ppC() *Xoshiro256ppC {
	seed := uint64(11111)
	st := &(C.xoshiro256_state{})
	C.xoshiro256_init(st, C.u8(seed))
	r := &Xoshiro256ppC{
		st,
	}
	return r
}
