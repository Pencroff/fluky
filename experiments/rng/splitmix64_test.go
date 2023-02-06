package rng

import (
	"testing"
)

func TestSplitMix64_Uint64(t *testing.T) {
	tbl := []struct {
		seed int64
		res  []uint64
	}{
		{
			seed: 0,
			res: []uint64{
				0xe220a8397b1dcdaf,
				0x6e789e6aa1b965f4,
				0x6c45d188009454f,
				0xf88bb8a8724c81ec,
				0x1b39896a51a8749b,
			},
		},
		{
			seed: 1,
			res: []uint64{
				0x910a2dec89025cc1,
				0xbeeb8da1658eec67,
				0xf893a2eefb32555e,
				0x71c18690ee42c90b,
				0x71bb54d8d101b5b9,
			},
		},
		{
			seed: 2,
			res: []uint64{
				0x975835de1c9756ce,
				0xbfc846100bfc1e42,
				0x987bbcbfdd7e532f,
				0xc3f2827affe7f664,
				0x4fc446b53f17fb29,
			},
		},
		{
			seed: 10,
			res: []uint64{
				0x88712be8a582fca,
				0xbbff7c596e26ce46,
				0x21876e7a2aec4a3d,
				0xd77e91a249eb9308,
				0xdb3658ea52921cc8,
			},
		},
		{
			seed: 1234567,
			res: []uint64{
				0x599ed017fb08fc85,
				0x2c73f08458540fa5,
				0x883ebce5a3f27c77,
				0x3fbef740e9177b3f,
				0xe3b8346708cb5ecd,
			},
		},
	}

	for _, tt := range tbl {
		s := NewSplitMix64(tt.seed)
		for _, r := range tt.res {
			rng := s.Uint64()
			rngC := s.Uint64C()
			if rng != r {
				t.Errorf("Go => Seed: %d, got: 0x%x, expected: 0x%x", tt.seed, rng, r)
			}
			if rngC != r {
				t.Errorf("CGo => Seed: %d, got: 0x%x, expected: 0x%x", tt.seed, rngC, r)
			}
		}
	}

}
