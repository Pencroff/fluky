package fluky

type IntegerOptionsFn func(i *IntegerOptions)

type IntegerOptions struct {
	min int
	max int
}

// WithIntRange configure min and max values for integer random
func WithIntRange(min, max int) IntegerOptionsFn {
	return func(i *IntegerOptions) {
		i.min = min
		i.max = max
	}
}

// Integer random integer value from range [min, max]
// from −(2^62) to 2^62 − 1 by default (int64)
// to avoid overflow for int64 positive range
// using values out of this range doesn't warranty
// that result will be in range [min, max]
func (f *Fluky) Integer(opts ...IntegerOptionsFn) int {
	o := &IntegerOptions{min: minInt63, max: maxInt63}
	for _, optFn := range opts {
		optFn(o)
	}
	if o.max == o.min {
		return o.min
	}
	return int(f.rng.Int63())%(o.max-o.min+1) + o.min
}
