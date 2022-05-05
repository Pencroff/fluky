package fluky

const maxPrecision = 53

type FloatOptionsFn func(f *FloatOptions)

type FloatOptions struct {
	min       float64
	max       float64
	precision uint8
}

// WithFloatRange configure min and max values for float random
func WithFloatRange(min, max float64) FloatOptionsFn {
	return func(f *FloatOptions) {
		f.min = min
		f.max = max
	}
}

// WithPrecision configure max precision for float random
func WithPrecision(precision uint8) FloatOptionsFn {
	return func(f *FloatOptions) {
		f.precision = precision % maxPrecision
	}
}

// Float random float value from range [min, max]
// by default min = 0, max = 1
func (f Fluky) Float(opts ...FloatOptionsFn) float64 {
	//var i FloatOptions
	//for _, o := range opts {
	//	o(&i)
	//}
	//return f.rng.Float64(i.max-i.min) + i.min
	return f.rng.Float64()
}
