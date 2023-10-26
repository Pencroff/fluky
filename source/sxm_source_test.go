package source

import (
	"testing"
)

func TestSxm_compare_with_mix(t *testing.T) {
	testCases := []struct {
		seed int64
		v1   []uint64
		v2   []uint64
	}{

		{
			1,
			[]uint64{0x4fd5840664c064e8, 0xfd8acbd95e3e3596},
			[]uint64{0x9b34821ecf715f22, 0x65277eadfc9669fb},
		},
		{
			1234567,
			[]uint64{0xce63ed58b351e354, 0xc934ba58a1917aa3},
			[]uint64{0x87c54d12581e2bd7, 0x23272668a1ba95e4},
		},
	}
	for _, tc := range testCases {
		sxm := NewSxmSource(tc.seed)
		mix := NewSxmMixSource(tc.seed)
		for i := 0; i < len(tc.v1); i++ {
			v1 := sxm.Uint64()
			v2 := mix.Uint64()
			if v1 != tc.v1[i] {
				t.Errorf("Sxm value %016x is not equal to expected %016x", v1, tc.v1[i])
			}
			if v2 != tc.v2[i] {
				t.Errorf("Mix value %016x is not equal to expected %016x", v2, tc.v2[i])
			}
		}
	}
}
