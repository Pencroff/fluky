package source

import (
	"math/bits"
	"math/rand"
)

const (
	BIT_NOISE_A uint64 = 0xa55d924b5499ede9 // 0b1010010101011101100100100100101101010100100110011110110111101001
	BIT_NOISE_B uint64 = 0x531e8dd1c779c9d1 // 0b0101001100011110100011011101000111000111011110011100100111010001
	BIT_NOISE_C uint64 = 0xf6e7845905e176cb // 0b1111011011100111100001000101100100000101111000010111011011001011
)

type SxmSource struct {
	position uint64
	seed     uint64
}

func (s *SxmSource) Position() int64 {
	return int64(s.position)
}

func (s *SxmSource) SetPosition(pos int64) {
	s.position = uint64(pos)
}

func (s *SxmSource) Seed(seed int64) {
	s.position = 0
	s.seed = uint64(seed)
}

func (s *SxmSource) Int63() int64 {
	return int64(s.Uint64() >> 1)
}

func (s *SxmSource) Uint64() uint64 {
	r := sxmFunc(s.position, s.seed)
	s.position += 1
	return r
}

type SxmMixSource struct {
	SxmSource
}

func (s *SxmMixSource) Int63() int64 {
	return int64(s.Uint64() >> 1)
}

func (s *SxmMixSource) Uint64() uint64 {
	return mix64(s.SxmSource.Uint64())
}

func NewSxmSource(seed int64) rand.Source64 {
	r := &SxmSource{}
	r.Seed(seed)
	return r
}

func NewSxmMixSource(seed int64) rand.Source64 {
	r := &SxmMixSource{}
	r.Seed(seed)
	return r
}

func sxmFunc(pos, seed uint64) uint64 {
	//const BIT_NOISE1 = 0x761a4cd0f83414b7
	//// 0b 0111 0110 0001 1010 0100 1100 1101 0000 1111 1000 0011 0100 0001 0100 1011 0111
	//const BIT_NOISE2 = 0xc731c95ae3d8e029 // 14353474879613362217
	//// 0b 1100 0111 0011 0001 1100 1001 0101 1010 1110 0011 1101 1000 1110 0000 0010 1001
	//const BIT_NOISE3 = 0x1e6d6e51f71a3bff
	//// 0b 0001 1110 0110 1101 0110 1110 0101 0001 1111 0111 0001 1010 0011 1011 1111 1111
	//mangledBits := pos
	//mangledBits *= BIT_NOISE1
	//mangledBits += seed
	//mangledBits ^= mangledBits >> 8
	//mangledBits += BIT_NOISE2
	//mangledBits ^= mangledBits << 8
	//mangledBits *= BIT_NOISE3
	//mangledBits ^= mangledBits >> 8
	//return mangledBits

	//s.state += 0x9e3779b97f4a7c15
	//z := s.state
	//z ^= z >> 30
	//z *= 0xbf58476d1ce4e5b9
	//z ^= z >> 27
	//z *= 0x94d049bb133111eb
	//z ^= z >> 31
	//return z

	//e := s.a - bits.RotateLeft64(s.b, 7)
	//s.a = s.b ^ bits.RotateLeft64(s.c, 13)
	//s.b = s.c + bits.RotateLeft64(s.d, 37)
	//s.c = s.d + e
	//s.d = e + s.a
	//return s.d

	// v1 and v2 - 3 WEAK in test
	//const BIT_NOISE_A = 0xa84c44880b088313
	//const BIT_NOISE_B = 0x9e1921f1204915db
	//const BIT_NOISE_C = 0x3b2bb426c5ce355d

	//r := pos
	//r *= BIT_NOISE_A
	//r += seed
	//r -= bits.RotateLeft64(r, 13)
	////r *= BIT_NOISE_B // v1
	//r += BIT_NOISE_B // v2
	//r ^= bits.RotateLeft64(r, 29)
	//r *= BIT_NOISE_C
	//r += bits.RotateLeft64(r, 37)
	//return r

	//const BIT_NOISE_A = 0xa84c44880b088313
	//const BIT_NOISE_B = 0x9e1921f1204915db
	//const BIT_NOISE_C = 0x3b2bb426c5ce355d

	/* v3 + v4
	const BIT_NOISE_A = 5913243947771178181  // 0x5210100a71212cc5 - 0b0101001000010000000100000000101001110001001000010010110011000101
	const BIT_NOISE_B = 1305701637361584433  // 0x121ec8b8a5954931 - 0b0001001000011110110010001011100010100101100101010100100100110001
	const BIT_NOISE_C = 2390192657074336187  // 0x212bab8ecda9b1bb - 0b0010000100101011101010111000111011001101101010011011000110111011
	const BIT_NOISE_D = 13483123982118542761 // 0xbb1dadd338b9dda9 - 0b1011101100011101101011011101001100111000101110011101110110101001
	const BIT_NOISE_E = 10264185916200023893 // 0x8e71b6b573ab7755 - 0b1000111001110001101101101011010101110011101010110111011101010101
	const PRIME64_2 = 0xC2B2AE3D27D4EB4F     // 0b1100001010110010101011100011110100100111110101001110101101001111
	const PRIME64_3 = 0x165667B19E3779F9     // 0b0001011001010110011001111011000110011110001101110111100111111001

	in1 := (pos + BIT_NOISE_A) * BIT_NOISE_C
	in2 := (seed + BIT_NOISE_B) * BIT_NOISE_D

	r := in1 * in2
	r ^= bits.RotateLeft64(r, 13)
	r *= BIT_NOISE_E
	r ^= bits.RotateLeft64(r, 29)
	r *= BIT_NOISE_B
	r ^= bits.RotateLeft64(r, 37)
	r *= BIT_NOISE_A
	// v4
	//acc = acc xor (acc >> 33);
	//acc = acc * PRIME64_2;
	//acc = acc xor (acc >> 29);
	//acc = acc * PRIME64_3;
	//acc = acc xor (acc >> 32);
	// from xxhash
	r ^= r >> 33
	r *= PRIME64_2
	r ^= r >> 29
	r *= PRIME64_3
	r ^= r >> 32
	return r
	*/

	// 1 week in test
	//const BIT_NOISE_A = 5913243947771178181  // 0x5210100a71212cc5 - 0b0101001000010000000100000000101001110001001000010010110011000101
	//const BIT_NOISE_B = 1305701637361584433  // 0x121ec8b8a5954931 - 0b0001001000011110110010001011100010100101100101010100100100110001
	//const BIT_NOISE_C = 2390192657074336187  // 0x212bab8ecda9b1bb - 0b0010000100101011101010111000111011001101101010011011000110111011
	//const BIT_NOISE_D = 13483123982118542761 // 0xbb1dadd338b9dda9 - 0b1011101100011101101011011101001100111000101110011101110110101001
	//const BIT_NOISE_E = 10264185916200023893 // 0x8e71b6b573ab7755 - 0b1000111001110001101101101011010101110011101010110111011101010101
	//
	//in1 := (pos + BIT_NOISE_A) * BIT_NOISE_C
	//in2 := (seed + BIT_NOISE_B) * BIT_NOISE_D
	//
	//// v7.1
	//r := in1 * in2
	//r ^= bits.RotateLeft64(r, 13)
	//r *= BIT_NOISE_E
	//r ^= bits.RotateLeft64(r, 29)
	//r *= BIT_NOISE_A
	//r ^= bits.RotateLeft64(r, 37)
	//r *= BIT_NOISE_B
	//r ^= r >> 33
	//r *= BIT_NOISE_C
	//r ^= r >> 29
	//r *= BIT_NOISE_D
	//r ^= r >> 32
	//return r

	//const BIT_NOISE_A = 5913243947771178181  // 0x5210100a71212cc5 - 0b0101001000010000000100000000101001110001001000010010110011000101
	//const BIT_NOISE_B = 1305701637361584433  // 0x121ec8b8a5954931 - 0b0001001000011110110010001011100010100101100101010100100100110001
	//const BIT_NOISE_C = 2390192657074336187  // 0x212bab8ecda9b1bb - 0b0010000100101011101010111000111011001101101010011011000110111011
	//const BIT_NOISE_D = 13483123982118542761 // 0xbb1dadd338b9dda9 - 0b1011101100011101101011011101001100111000101110011101110110101001
	//const BIT_NOISE_E = 10264185916200023893 // 0x8e71b6b573ab7755 - 0b1000111001110001101101101011010101110011101010110111011101010101

	/*
		const (
			BIT_NOISE_A uint64 = 11400714785074694791
			BIT_NOISE_B uint64 = 14029467366897019727
			BIT_NOISE_C uint64 = 1609587929392839161
			BIT_NOISE_D uint64 = 9650029242287828579
			BIT_NOISE_E uint64 = 2870177450012600261
		)

		// v7.2
		r := pos * BIT_NOISE_E
		r += seed + BIT_NOISE_E
		r ^= bits.RotateLeft64(r, 37)
		r += BIT_NOISE_A
		r ^= bits.RotateLeft64(r, 31)
		r *= BIT_NOISE_B
		r ^= bits.RotateLeft64(r, 29)
		r += BIT_NOISE_C
		r ^= bits.RotateLeft64(r, 23)
		r *= BIT_NOISE_D
		r ^= bits.RotateLeft64(r, 19)
		r += BIT_NOISE_E
		r ^= bits.RotateLeft64(r, 17)
		r *= BIT_NOISE_A
		r ^= bits.RotateLeft64(r, 13)
		r *= BIT_NOISE_B
		r ^= bits.RotateLeft64(r, 11)
		r *= BIT_NOISE_C
		r ^= bits.RotateLeft64(r, 7)
		r *= BIT_NOISE_D
		r ^= bits.RotateLeft64(r, 5)
		r *= BIT_NOISE_E
		return r
	*/

	// aka xxhash
	//v1 := pos + BIT_NOISE_A
	//v2 := seed + BIT_NOISE_A
	//v3 := pos - BIT_NOISE_A
	//
	//r := bits.RotateLeft64(v1, 37) + bits.RotateLeft64(v2, 29) + bits.RotateLeft64(v3, 19)
	//r = mergeRound64(r, v1)
	//r = mergeRound64(r, v2)
	//r = mergeRound64(r, v3)
	//
	//r = mix64(r)
	//return r
	// v8.1
	//v1 := (pos + BIT_NOISE_B) * BIT_NOISE_A
	//v2 := (seed + BIT_NOISE_C) * BIT_NOISE_D
	//r := v1 * v2
	//r ^= bits.RotateLeft64(r, 13)
	//r *= BIT_NOISE_A
	//r ^= bits.RotateLeft64(r, 29)
	//r *= BIT_NOISE_D
	//r ^= bits.RotateLeft64(r, 37)
	// v8.2
	v1 := pos + BIT_NOISE_A
	v2 := seed + BIT_NOISE_B
	v3 := pos - BIT_NOISE_C
	v4 := seed - BIT_NOISE_C
	r := bits.RotateLeft64(v1, 1) + bits.RotateLeft64(v2, 7) + bits.RotateLeft64(v3, 11) + bits.RotateLeft64(v4, 13)
	r *= BIT_NOISE_A
	r ^= bits.RotateLeft64(r, 37)
	r *= BIT_NOISE_B
	r ^= bits.RotateLeft64(r, 29)
	r *= BIT_NOISE_C
	r ^= bits.RotateLeft64(r, 19)
	return r
}

func sxmMixFunc(pos, seed uint64) uint64 {
	return mix64(sxmFunc(pos, seed))
}

func mix64(h uint64) uint64 {
	h ^= h >> 33
	h *= BIT_NOISE_B
	h ^= h >> 29
	h *= BIT_NOISE_C
	h ^= h >> 32
	return h
}
