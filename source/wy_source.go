package source

import (
	"github.com/Pencroff/fluky/wy"
)

type WySource struct {
	seed uint64
}

func (s *WySource) Seed(seed int64) {
	s.seed = uint64(seed)
}

func (s *WySource) Int63() int64 {
	return int64(wy.WyRand(&s.seed) >> 1)
}

func (s *WySource) Uint64() uint64 {
	return wy.WyRand(&s.seed)
}

func NewWySource(seed int64) *WySource {
	return &WySource{seed: uint64(seed)}
}

type WyMixSource struct {
	seed     uint64
	position uint64
}

func (s *WyMixSource) Seed(seed int64) {
	s.seed = uint64(seed)
	s.position = 0
}

func (s *WyMixSource) Uint64() uint64 {
	r := wy.WyMixRand(s.position, s.seed)
	s.position += 1
	return r
}

func (s *WyMixSource) Int63() int64 {
	return int64(s.Uint64() >> 1)
}

func NewWyMixSource(seed int64) *WyMixSource {
	r := &WyMixSource{}
	r.Seed(seed)
	return r
}
