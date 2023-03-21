package source

import (
	et "github.com/Pencroff/go-toolkit/extended_type"
	"math/bits"
	"math/rand"
)

type PcgSource struct {
	state             et.Uint128
	inc               et.Uint128
	defaultMultiplier et.Uint128
}

func (s *PcgSource) Seed(seed int64) {
	initseq := uint64(54)
	s.state = et.ZeroUint128
	s.inc = et.From64(initseq).Lsh(1).Or64(1)
	s.step()
	s.state = s.state.Add64(uint64(seed))
	s.step()
}

func (s *PcgSource) Uint64() uint64 {
	s.step()
	return s.stateToValue()
}

func (s *PcgSource) Int63() int64 {
	return int64(s.Uint64() >> 1)
}

func (s *PcgSource) step() {
	s.state = s.state.Mul(s.defaultMultiplier).Add(s.inc)
}

func (s *PcgSource) stateToValue() (v uint64) {
	v = s.state.Lo ^ s.state.Hi
	rot := int(s.state.Rsh(122).Lo&0x3F) * -1
	v = bits.RotateLeft64(v, rot)
	return
}

func NewPcgSource(seed int64) rand.Source64 {
	r := &PcgSource{
		state:             et.Uint128{},
		inc:               et.Uint128{},
		defaultMultiplier: et.New(0x4385df649fccf645, 0x2360ed051fc65da4),
	}
	r.Seed(seed)
	return r
}
