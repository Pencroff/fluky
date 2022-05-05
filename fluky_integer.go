package fluky

type IntegerOptionsFn func(i *IntegerOptions)

type IntegerOptions struct {
	min int64
	max int64
}

// WithIntRange configure min and max values for integer random
func WithIntRange(min, max int64) IntegerOptionsFn {
	return func(i *IntegerOptions) {
		i.min = min
		i.max = max
	}
}

// Integer random integer value from range [min, max]
// from −(2^63) to 2^63 − 1 by default
func (f Fluky) Integer(opts ...IntegerOptionsFn) int64 {
	//var i IntegerOptions
	//for _, o := range opts {
	//	o(&i)
	//}
	//return f.rng.Int63(i.max - i.min) + i.min
	return f.rng.Int63()
}
