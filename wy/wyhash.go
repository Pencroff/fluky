package wy

import (
	"encoding/binary"
	"math/bits"
)

// version 4.2

const (
	wyp0 = 0x2d358dccaa6c78a5
	wyp1 = 0x8bb84b93962eacc9
	wyp2 = 0x4b33a62ed433d4a3
	wyp3 = 0x4d5a2da51de1aa47
)

/*
//a useful 64bit-64bit mix function to produce deterministic pseudo random numbers that can pass BigCrush and PractRand
static inline uint64_t wyhash64(uint64_t A, uint64_t B){ A^=0x2d358dccaa6c78a5ull; B^=0x8bb84b93962eacc9ull; _wymum(&A,&B); return _wymix(A^0x2d358dccaa6c78a5ull,B^0x8bb84b93962eacc9ull);}

//The wyrand PRNG that pass BigCrush and PractRand
static inline uint64_t wyrand(uint64_t *seed){ *seed+=0x2d358dccaa6c78a5ull; return _wymix(*seed,*seed^0x8bb84b93962eacc9ull);}
*/

func WyMixRand(a, b uint64) uint64 {
	a ^= wyp0
	b ^= wyp1
	a, b = wymum(a, b)
	return wymix(a^wyp0, b^wyp1)
}

func WyRand(seed *uint64) uint64 {
	*seed += wyp0
	return wymix(*seed, *seed^wyp1)
}

// Sum64 computes a 64-bit hash for the given key and seed.
func Sum64(seed uint64, key []byte) uint64 {
	return sum64(key, seed)
}

// sum64 implements the Wyhash algorithm based on the C++ implementation.
//
// It starts by mixing the seed with a secret constant; for inputs of 16 bytes or fewer
// it uses two 32-bit (or a 3-byte read for very short inputs) reads combined into 64-bit values.
// For longer inputs it processes 48-byte blocks followed by 16-byte blocks, then uses the final
// 16 bytes to form two 64-bit numbers. Finally, the two numbers and the seed are combined via
// in-place multiplication (wymum_inplace) and a final mix (wymix).
func sum64(key []byte, seed uint64) uint64 {
	// Initial seed mixing: seed ^= wymix(seed^wyp0, wyp1)
	seed ^= wymix(seed^wyp0, wyp1)

	var a, b uint64
	n := len(key)
	u := uint64(n)

	if n == 0 {
		a = wyp1
		b = seed
		a, b = wymum(a, b)
		return wymix(a^wyp0^u, b^wyp1)
	} else if n < 4 {
		// For inputs [1,3] bytes, use three 16-bit words combined into 64-bit values.
		a = wyr3(key, n) ^ wyp1
		b = seed
		a, b = wymum(a, b)
		return wymix(a^wyp0^u, b^wyp1)
	} else if n <= 16 {
		// For inputs [4,16] bytes, use two 32-bit words combined into 64-bit values.
		p1AVal := wyr4(key[0:4])
		idxP2A := (n >> 3) << 2 // if n=4..7, idx=0; if n=8..15, idx=4; if n=16, idx=8
		p2AVal := wyr4(key[idxP2A : idxP2A+4])
		valA := (p1AVal << 32) | p2AVal

		p1BVal := wyr4(key[n-4 : n])
		idxP2B := n - 4 - ((n >> 3) << 2) // if n=4..7, idx=n-4; if n=8..15, idx=n-8; if n=16, idx=4
		p2BVal := wyr4(key[idxP2B : idxP2B+4])
		valB := (p1BVal << 32) | p2BVal

		// Apply the XORs as in the original Go code's pattern for this branch
		a = valA ^ wyp1
		b = valB ^ seed

		// The rest of the original logic for this branch remains the same
		a, b = wymum(a, b)
		return wymix(a^wyp0^u, b^wyp1)

	}
	// For inputs longer than 16 bytes.
	i := n
	p := key
	// Process 48-byte blocks if available.
	if i >= 48 {
		see1, see2 := seed, seed
		for i >= 48 {
			seed = wymix(wyr8(p[0:8])^wyp1, wyr8(p[8:16])^seed)
			see1 = wymix(wyr8(p[16:24])^wyp2, wyr8(p[24:32])^see1)
			see2 = wymix(wyr8(p[32:40])^wyp3, wyr8(p[40:48])^see2)
			p = p[48:]
			i -= 48
		}
		seed ^= see1 ^ see2
	}
	// Process remaining 16-byte blocks.
	for i > 16 {
		seed = wymix(wyr8(p[0:8])^wyp1, wyr8(p[8:16])^seed)
		p = p[16:]
		i -= 16
	}
	// Always extract the final 16 bytes from the original key.
	a = wyr8(key[n-16:n-8]) ^ wyp1
	b = wyr8(key[n-8:n]) ^ seed

	a, b = wymum(a, b)
	return wymix(a^wyp0^u, b^wyp1)
}

// wyr3 reads 3 bytes from p to form a uint64.
func wyr3(p []byte, k int) uint64 {
	return (uint64(p[0]) << 16) | (uint64(p[k>>1]) << 8) | uint64(p[k-1])
}

// wyr4 returns a 32-bit little-endian value from p.
func wyr4(p []byte) uint64 {
	return uint64(binary.LittleEndian.Uint32(p))
}

// wyr8 returns a 64-bit little-endian value from p.
func wyr8(p []byte) uint64 {
	return binary.LittleEndian.Uint64(p)
}

// wymum multiplies a and b and returns the low and high 64-bit parts respectively.
func wymum(a, b uint64) (uint64, uint64) {
	hi, lo := bits.Mul64(a, b)
	return lo, hi
}

// wymix performs mixing by multiplying a and b and then XORing the high and low 64-bit parts.
func wymix(a, b uint64) uint64 {
	hi, lo := bits.Mul64(a, b)
	return hi ^ lo
}
