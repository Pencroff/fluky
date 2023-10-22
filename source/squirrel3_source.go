package source

import "math/rand"

// Based on https://www.youtube.com/watch?v=LWFzPP8ZbdU

type Squirrel3Source struct {
	position uint64
	seed     uint64
}

func (s *Squirrel3Source) Position(pos int64) {
	s.position = uint64(pos)
}

func (s *Squirrel3Source) Seed(seed int64) {
	s.position = 0
	s.seed = uint64(seed)
}

func (s *Squirrel3Source) Int63() int64 {
	return int64(s.Uint64() >> 1)
}

func (s *Squirrel3Source) Uint64() uint64 {
	r := squirrel3(s.position, s.seed)
	s.position += 1
	return r
}

type Squirrel3x2Source struct {
	Squirrel3Source
}

func (s *Squirrel3x2Source) Uint64() uint64 {
	const mask32 = 1<<32 - 1
	x := (squirrel3(s.position, s.seed) >> 16) & mask32 // used high 32 bits
	s.position += 1
	y := (squirrel3(s.position, s.seed) >> 16) & mask32 // used high 32 bits
	s.position += 1
	return x<<32 | y
}

type Squirrel3BigPrimeSource struct {
	Squirrel3Source
}

func (s *Squirrel3BigPrimeSource) Uint64() uint64 {
	r := squirrel3uint64Noise(s.position, s.seed)
	s.position += 1
	return r
}

func squirrel3(pos, seed uint64) uint64 {
	/*
		const unsigned int BIT_NOISE1 = 0x6831DA4; // 0b0110 1000 1110'0011 0001 1101 1010'0100;
		const unsigned int BIT_NOISE2 = 0xB5297A4D; // 0b1011 ' 0101 ' 0010 ' 1001 ' 0111'1010' 0100'1101;
		const unsigned int BIT_NOISE3 - 0x1856C4E9; // 0b0001 1011'0101 0110'1100'0100'1110 1001;
		unsigned int mangledBits = (unsigned int) positionx;
		mangledBits *- BIT_NOISE1;
		mangledBits += seed;
		mangledits ^= (mangledBits » 8);
		mangledBits += BIT_NOISE2;
		mangledits ^= (mangledBits « 8);
		mangledBits *= BIT_NOISE3;
		mangledits ^= (mangledBits » 8);
		return mangledBits;
	*/
	const BIT_NOISE1 = 0x6831DA4  // 0b0110 1000 1110 0011 0001 1101 1010 0100
	const BIT_NOISE2 = 0xB5297A4D // 0b1011 0101 0010 1001 0111 1010 0100 1101
	const BIT_NOISE3 = 0x1856C4E9 // 0b0001 1011'0101 0110'1100'0100'1110 1001
	mangledBits := pos
	mangledBits *= BIT_NOISE1
	mangledBits += seed
	mangledBits ^= mangledBits >> 8
	mangledBits += BIT_NOISE2
	mangledBits ^= mangledBits << 8
	mangledBits *= BIT_NOISE3
	mangledBits ^= mangledBits >> 8
	return mangledBits
}

func squirrel3uint64Noise(pos, seed uint64) uint64 {
	const BIT_NOISE1 = 0x761a4cd0f83414b7
	// 0b 0111 0110 0001 1010 0100 1100 1101 0000 1111 1000 0011 0100 0001 0100 1011 0111
	const BIT_NOISE2 = 0xc731c95ae3d8e029 // 14353474879613362217
	// 0b 1100 0111 0011 0001 1100 1001 0101 1010 1110 0011 1101 1000 1110 0000 0010 1001
	const BIT_NOISE3 = 0x1e6d6e51f71a3bff
	// 0b 0001 1110 0110 1101 0110 1110 0101 0001 1111 0111 0001 1010 0011 1011 1111 1111
	mangledBits := pos
	mangledBits *= BIT_NOISE1
	mangledBits += seed
	mangledBits ^= mangledBits >> 8
	mangledBits += BIT_NOISE2
	mangledBits ^= mangledBits << 8
	mangledBits *= BIT_NOISE3
	mangledBits ^= mangledBits >> 8
	return mangledBits
}

func NewSquirrel3Source(seed int64) rand.Source64 {
	r := &Squirrel3Source{}
	r.Seed(seed)
	return r
}

func NewSquirrel3x2Source(seed int64) rand.Source64 {
	r := &Squirrel3x2Source{}
	r.Seed(seed)
	return r
}

func NewSquirrel3BigPrimeSource(seed int64) rand.Source64 {
	r := &Squirrel3BigPrimeSource{}
	r.Seed(seed)
	return r
}
