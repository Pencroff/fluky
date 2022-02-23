package fluky_testing

import (
	"github.com/Pencroff/fluky"
)

var RngTbl = []struct {
	name       string
	rnd        fluky.RandomGenerator
	size       int
	exportData bool
}{
	{
		name:       fluky.MMIX,
		rnd:        fluky.NewLcg(fluky.MMIX),
		size:       64,
		exportData: false,
	}, {
		name:       fluky.Musl,
		rnd:        fluky.NewLcg(fluky.Musl),
		size:       64,
		exportData: false,
	}, {
		name:       "built_in",
		rnd:        fluky.NewBuiltIn(),
		size:       64,
		exportData: false,
	}, {
		name:       "small_prng",
		rnd:        fluky.NewSmallPrng(),
		size:       64,
		exportData: true,
	},
}
