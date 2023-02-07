package fluky_testing

import (
	"github.com/Pencroff/fluky/source"
	"math/rand"
	"testing"
)

var res uint64

func BenchmarkSource(b *testing.B) {
	var seed int64 = 1234567
	tbl := []struct {
		name string
		src  rand.Source
	}{
		{"BuiltIn Source", rand.NewSource(seed)},
		{"Pcg Source", source.NewPcgSource(seed)},
		{"SmallPrng Source", source.NewSmallPrngSource(seed)},
		{"xoshiro256++ Source", source.NewXoshiro256ppSource(seed)},
		{"xoshiro256** Source", source.NewXoshiro256ssSource(seed)},
		{"SplitMix64", source.NewSplitMix64Source(seed)},
	}
	for _, t := range tbl {
		r := t.src.(rand.Source64)
		b.Run(t.name, func(b *testing.B) {
			v := uint64(0)
			for i := 0; i < b.N; i++ {
				v = r.Uint64()
			}
			res = v
		})
		res += 1
	}
}