package fluky

type RandomGenerator interface {
	Seed(v uint64)
	NextUint64() uint64
	NextFloat64() float64
}
