package rng

/*
#include <stdio.h>
#include <stdint.h>
static uint64_t x; // The state can be seeded with any value.

uint64_t next() {
uint64_t z = (x += 0x9e3779b97f4a7c15);
z = (z ^ (z >> 30)) * 0xbf58476d1ce4e5b9;
z = (z ^ (z >> 27)) * 0x94d049bb133111eb;
return z ^ (z >> 31);
}
void seedSplitMix64(uint64_t s) {
  x = s;
}
*/
import "C"

// SplitMix64 is a 64-bit pseudo-random number generator.
// It is a fast, high-quality generator. For generating just floating-point
// numbers, SplitMix64 is even faster than xoroshiro128+.
// The state must be seeded so that it is not everywhere zero. If you have
// a 64-bit seed, we suggest to seed a SplitMix64 generator and use its
// output to fill s.
type SplitMix64 struct {
	state uint64
}

// NewSplitMix64 returns a new SplitMix64 generator to generate
// non-negative pseudo-random 64-bit integers.
func NewSplitMix64(seed int64) *SplitMix64 {
	rng := &SplitMix64{}
	rng.Seed(seed)
	return rng
}

// Uint64 returns a pseudo-random 64-bit integer as a uint64.
func (s *SplitMix64) Uint64() uint64 {
	s.state += 0x9e3779b97f4a7c15
	z := s.state
	z = (z ^ (z >> 30)) * 0xbf58476d1ce4e5b9
	z = (z ^ (z >> 27)) * 0x94d049bb133111eb
	return z ^ (z >> 31)
}

// Int63 returns a non-negative pseudo-random 63-bit integer as an int64.
func (s *SplitMix64) Int63() int64 {
	return int64(s.Uint64() >> 1)
}

// Seed uses the provided seed value to initialize the generator to a
// deterministic state.
func (s *SplitMix64) Seed(seed int64) {
	s.state = uint64(seed)
	C.seedSplitMix64(C.uint64_t(seed))
}

// Uint64C Uint64 returns a pseudo-random 64-bit integer as a uint64.
func (s *SplitMix64) Uint64C() uint64 {
	return uint64(C.next())
}
