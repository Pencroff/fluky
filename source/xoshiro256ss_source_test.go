package source

import (
	"testing"
)

func TestXoshiro256ssSource_Seed(t *testing.T) {
	tbl := []struct {
		name string
		seed int64
		res  []uint64
	}{
		{"test-1", 0,
			[]uint64{
				0x99ec5f36cb75f2b4,
				0xbf6e1f784956452a,
				0x1a5f849d4933e6e0,
				0x6aa594f1262d2d2c,
				0xbba5ad4a1f842e59,
			}},
		{"test-2", 1,
			[]uint64{
				0xb3f2af6d0fc710c5,
				0x853b559647364cea,
				0x92f89756082a4514,
				0x642e1c7bc266a3a7,
				0xb27a48e29a233673,
			}},
	}
	for _, tt := range tbl {
		t.Run(tt.name, func(t *testing.T) {
			g := NewXoshiro256ssSource(tt.seed)
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
