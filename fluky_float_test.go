package fluky

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFluky_Float_NoOptions(t *testing.T) {
	mRng := new(RngMock)
	mRng.On("Float64").Return(0.31415).Times(1)
	mRng.On("Float64").Return(0.11543).Times(1)
	mRng.On("Float64").Return(0.123456789).Times(1)
	// ref - https://play.golang.com/p/JZWbrcw0_fh
	mRng.On("Float64").Return(0.99999999999999988897769753748434595763683319091796875).Times(1)
	mRng.On("Float64").Return(0.00000000000000011102230246251565404236316680908203125).Times(1)

	f := NewFluky(mRng)
	assert.Equal(t, 0.31415, f.Float())
	assert.Equal(t, 0.11543, f.Float())
	assert.Equal(t, 0.123456789, f.Float())
	assert.Equal(t, 0.99999999999999988897769753748434595763683319091796875, f.Float())
	assert.Equal(t, 0.00000000000000011102230246251565404236316680908203125, f.Float())
	mRng.AssertExpectations(t)
}

func TestFluky_Float_WithOptions_Precision(t *testing.T) {
	mRng := new(RngMock)
	mRng.On("Float64").Return(0.31415).Times(1)
	mRng.On("Float64").Return(0.11543).Times(1)
	mRng.On("Float64").Return(0.123456789).Times(1)
	mRng.On("Float64").Return(0.99999999999999988897769753748434595763683319091796875).Times(1)
	mRng.On("Float64").Return(0.00000000000000011102230246251565404236316680908203125)

	f := NewFluky(mRng)
	assert.Equal(t, 0.314, f.Float(WithPrecision(3)))
	assert.Equal(t, 0.1154, f.Float(WithPrecision(4)))
	assert.Equal(t, 0.123457, f.Float(WithPrecision(6)))
	expected := 0.9999999999999998889777
	actual := f.Float(WithPrecision(20))
	assert.Equal(t, expected, actual, "%.53f != %.53f", expected, actual)
	assert.Equal(t, 0.0000000000000001, f.Float(WithPrecision(16)))
	assert.Equal(t, 0.00000000000000011, f.Float(WithPrecision(17)))
	assert.Equal(t, 0.000000000000000111, f.Float(WithPrecision(18)))
	assert.Equal(t, 0.0000000000000001110, f.Float(WithPrecision(19)))
	assert.Equal(t, 0.00000000000000011102, f.Float(WithPrecision(20)))
	assert.Equal(t, 0.00000000000000011102, f.Float(WithPrecision(21+20)))

	mRng.AssertExpectations(t)
}

func TestFluky_Float_WithOptions_MinMax(t *testing.T) {
	mRng := new(RngMock)
	mRng.On("Float64").Return(0.31415).Times(1)
	mRng.On("Float64").Return(0.11543).Times(1)
	mRng.On("Float64").Return(0.123456789).Times(1)
	mRng.On("Float64").Return(0.99999999999999988897769753748434595763683319091796875).Times(1)
	mRng.On("Float64").Return(0.00000000000000011102230246251565404236316680908203125).Times(1)

	f := NewFluky(mRng)
	assert.InDelta(t, -0.7434, f.Float(WithFloatRange(-2, 2)), 1e-15)
	assert.InDelta(t, -1.53828, f.Float(WithFloatRange(-2, 2)), 1e-15)
	assert.InDelta(t, -1.506172844, f.Float(WithFloatRange(-2, 2)), 1e-15)
	assert.InDelta(t, 1.9999999999999996, f.Float(WithFloatRange(-2, 2)), 1e-15)
	assert.InDelta(t, -1.9999999999999996, f.Float(WithFloatRange(-2, 2)), 1e-15)
	mRng.AssertExpectations(t)
}

func TestFluky_Float_WithOptions(t *testing.T) {
	mRng := new(RngMock)
	mRng.On("Float64").Return(0.95).Times(1)
	mRng.On("Float64").Return(0.31415).Times(1)
	mRng.On("Float64").Return(0.11543).Times(1)
	mRng.On("Float64").Return(0.123456789).Times(1)

	f := NewFluky(mRng)
	assert.InDelta(t, 2, f.Float(WithPrecision(0), WithFloatRange(-2, 2)), 1e-15)
	assert.InDelta(t, -0.74, f.Float(WithPrecision(2), WithFloatRange(-2, 2)), 1e-15)
	assert.InDelta(t, -1.5383, f.Float(WithPrecision(4), WithFloatRange(-2, 2)), 1e-15)
	assert.InDelta(t, -1.50617284, f.Float(WithPrecision(8), WithFloatRange(-2, 2)), 1e-15)

	mRng.AssertExpectations(t)
}
