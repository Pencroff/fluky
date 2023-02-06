package source

import (
	"testing"
)

func TestXoshiro256ppSource_Seed(t *testing.T) {
	tbl := []struct {
		name string
		seed int64
		res  []uint64
	}{
		{"test-1", 0,
			[]uint64{
				0x53175d61490b23df,
				0x61da6f3dc380d507,
				0x5c0fdf91ec9a7bfc,
				0x2eebf8c3bbe5e1a,
				0x7eca04ebaf4a5eea,
			}},
		{"test-2", 1,
			[]uint64{
				0xcfc5d07f6f03c29b,
				0xbf424132963fe08d,
				0x19a37d5757aaf520,
				0xbf08119f05cd56d6,
				0x2f47184b86186fa4,
			}},
		{"test-3", 2,
			[]uint64{
				0xc3e67584b5c4fc2a,
				0x89837ec39e40f2c8,
				0xa6bb0b2987ac94cd,
				0x4b31e5fbdd210a72,
				0x7dad9d2b709a04f7,
			}},
		{"test-4", 10,
			[]uint64{
				0x38f1349ff3c8329c,
				0xcefca9b16c1d4aea,
				0x16ec5f8677445a68,
				0xc804054bf768e31a,
				0xc4e2294cfb3eddc0,
			}},
		{"test-5", 1234567,
			[]uint64{
				0x610e053dd55ab68,
				0x70c979e26e27fbac,
				0xfb95f99f9f6bb2de,
				0x3890aaecd9fa80a,
				0x536acabd892d9406,
			}},
	}
	for _, tt := range tbl {
		t.Run(tt.name, func(t *testing.T) {
			g := NewXoshiro256ppSource(tt.seed)
			//fmt.Printf("seed: %d\n", tt.seed)
			for _, r := range tt.res {
				rng := g.Uint64()
				//fmt.Printf("0x%x,\n", rng)
				if rng != r {
					t.Errorf("rngC: 0x%x, extected: 0x%x", rng, r)
				}
			}
		})
	}
}

func TestXoshiro256pp_Jump(t *testing.T) {
	tbl := []struct {
		name string
		seed int64
		resA []uint64
		resB []uint64
	}{
		{"test-1", 0,
			[]uint64{
				0x53175d61490b23df,
				0x61da6f3dc380d507,
				0x5c0fdf91ec9a7bfc,
				0x2eebf8c3bbe5e1a,
				0x7eca04ebaf4a5eea,
			},
			[]uint64{
				0x2107d23f5380538b,
				0x860c46fba09246f0,
				0xe824e1ac3bb3b014,
				0x5fcec05a1c2523c9,
				0x92790ab81295cbdb,
			},
		}, {
			"test-2", 1,
			[]uint64{
				0xcfc5d07f6f03c29b,
				0xbf424132963fe08d,
				0x19a37d5757aaf520,
				0xbf08119f05cd56d6,
				0x2f47184b86186fa4,
			},
			[]uint64{
				0xdafd92f1adffc5b9,
				0x89d5ed6828f5becf,
				0xc81a7b85673e9dac,
				0xe3ed98a07ef5a746,
				0xe294a7e13e75c33c,
			},
		},
	}
	for _, tt := range tbl {
		t.Run(tt.name, func(t *testing.T) {
			s := NewXoshiro256ppSource(tt.seed)
			g := s.(*Xoshiro256ppSource)
			a := g.Jump()
			//fmt.Printf("seed: %d\n", tt.seed)
			for i, rA := range tt.resA {
				rB := tt.resB[i]
				rngA := a.Uint64()
				rngB := g.Uint64()
				//fmt.Printf("0x%x,\n", rng)
				if rngA != rA {
					t.Errorf("rngA: 0x%x, extected: 0x%x", rngA, rA)
				}
				if rngB != tt.resB[i] {
					t.Errorf("rngB: 0x%x, extected: 0x%x", rngB, rB)
				}
			}
		})
	}
}

func TestXoshiro256pp_LongJump(t *testing.T) {
	tbl := []struct {
		name string
		seed int64
		resA []uint64
		resB []uint64
	}{
		{"test-1", 0,
			[]uint64{
				0x53175d61490b23df,
				0x61da6f3dc380d507,
				0x5c0fdf91ec9a7bfc,
				0x2eebf8c3bbe5e1a,
				0x7eca04ebaf4a5eea,
			},
			[]uint64{
				0x708919b147f78af3,
				0xf391447947dcccec,
				0x8619b00c868c7e42,
				0xcb148b88c2929741,
				0x2952ff7548b916f6,
			},
		}, {
			"test-2", 1,
			[]uint64{
				0xcfc5d07f6f03c29b,
				0xbf424132963fe08d,
				0x19a37d5757aaf520,
				0xbf08119f05cd56d6,
				0x2f47184b86186fa4,
			},
			[]uint64{
				0xc6e0f3d2b09d8eec,
				0x55ad95eef7a40e42,
				0x8cc0e5594cb97ab0,
				0x708019a0cb2b42e8,
				0x62c8bf2965d869ba,
			},
		},
	}
	for _, tt := range tbl {
		t.Run(tt.name, func(t *testing.T) {
			s := NewXoshiro256ppSource(tt.seed)
			g := s.(*Xoshiro256ppSource)
			a := g.LongJump()
			//fmt.Printf("seed: %d\n", tt.seed)
			for i, rA := range tt.resA {
				rB := tt.resB[i]
				rngA := a.Uint64()
				rngB := g.Uint64()
				//fmt.Printf("0x%x,\n", rngB)
				if rngA != rA {
					t.Errorf("rngA: 0x%x, extected: 0x%x", rngA, rA)
				}
				if rngB != rB {
					t.Errorf("rngB: 0x%x, extected: 0x%x", rngB, rB)
				}
			}
		})
	}
}
