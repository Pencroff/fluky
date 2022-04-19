package fluky

import (
	"github.com/Pencroff/fluky/rng"
	"github.com/Pencroff/go-toolkit/math"
)

func NewFluky(r rng.RandomGenerator) *Fluky {
	return &Fluky{r}
}

type Fluky struct {
	rng rng.RandomGenerator
}

// Seed internal RNG, reset seed value
func (f *Fluky) Seed(v int64) {
	f.rng.Seed(v)
}

func (f *Fluky) Weighted(values []string, weights []uint) string {
	sum := uint(0)
	mxIdx := math.Min(len(values), len(weights)) - 1
	for idx, weight := range weights {
		if idx > mxIdx {
			break
		}
		sum += weight
	}
	rnd := f.rng.Float64() * float64(sum)
	for idx, weight := range weights {
		rnd -= float64(weight)
		if rnd <= 0 {
			return values[idx]
		}
	}
	return ""

}
