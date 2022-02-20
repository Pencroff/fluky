package testing

import (
	"github.com/Pencroff/fluky"
	"math/rand"
)

var RngTbl = []struct {
	name string
	rnd  fluky.RandomGenerator
	size int
}{
	{
		name: fluky.MMIX,
		rnd:  fluky.NewLcg(fluky.MMIX),
		size: 64,
	}, {
		name: fluky.Musl,
		rnd:  fluky.NewLcg(fluky.Musl),
		size: 64,
	}, {
		name: "built-in",
		rnd:  rand.New(rand.NewSource(11111)),
		size: 64,
	}, {
		name: "small_prng",
		rnd:  fluky.NewSmallPrng(),
		size: 64,
	},
}
