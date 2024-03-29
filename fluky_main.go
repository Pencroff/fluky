package fluky

const (
	maxInt63 = 1<<62 - 1
	minInt63 = -1 << 62
)

// NewFluky create new Fluky instance and return pointer to it
func NewFluky(r RandomGenerator) *Fluky {
	return &Fluky{r}
}

type Fluky struct {
	rng RandomGenerator
}

// Seed internal RNG, reset seed value
func (f *Fluky) Seed(v int64) {
	f.rng.Seed(v)
}

// Uint64 returns random uint64 value
func (f *Fluky) Uint64() uint64 {
	return f.rng.Uint64()
}

// Int63 returns random int64 value in range [0, 2^63)
func (f *Fluky) Int63() int64 {
	return f.rng.Int63()
}

// Float64 returns random float64 value in range [0, 1)
func (f *Fluky) Float64() float64 {
	return f.rng.Float64()
}
