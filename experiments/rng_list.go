package experiments

import (
	"github.com/Pencroff/fluky/experiments/rng"
	"math/rand"
)

var RngTbl = []struct {
	name       string
	rnd        rand.Source64
	size       int
	exportData bool
	measure    bool
}{
	{
		name:       rng.MMIX,
		rnd:        rng.NewLcg(rng.MMIX),
		size:       64,
		exportData: false,
		measure:    false,
	}, {
		name:       rng.Musl,
		rnd:        rng.NewLcg(rng.Musl),
		size:       64,
		exportData: false,
		measure:    false,
	}, {
		name:       "built_in",
		rnd:        rng.NewBuiltIn(),
		size:       64,
		exportData: false,
		measure:    false,
	}, {
		name:       "small_prng",
		rnd:        rng.NewSmallPrng(),
		size:       64,
		exportData: false,
		measure:    false,
	}, {
		name:       "squares",
		rnd:        rng.NewSquares(),
		size:       64,
		exportData: false,
		measure:    false,
	}, {
		name:       "pcg64",
		rnd:        rng.NewPcgRng(),
		size:       64,
		exportData: false,
		measure:    false,
	}, {
		name:       "xoshiro256pp",
		rnd:        rng.NewXoshiro256pp(11111),
		exportData: false,
		measure:    true,
	}, {
		name:       "SplitMix64",
		rnd:        rng.NewSplitMix64(11111),
		exportData: false,
		measure:    true,
	},
}
