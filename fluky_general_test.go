package fluky

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestFlukyGeneral_Seed(t *testing.T) {
	mRng := new(RngMock)
	mRng.On("Seed", int64(10))
	mRng.On("Seed", int64(math.MaxInt64-1))
	f := NewFluky(mRng)
	f.Seed(10)
	f.Seed(math.MaxInt64 - 1)
	mRng.AssertExpectations(t)
}

func TestFlukyGeneral_Weighted(t *testing.T) {
	lst := []struct {
		Weight  float64
		Values  []string
		Weights []uint
		Result  string
	}{
		{Weight: 0.0, Values: []string{"a", "b", "c"}, Weights: []uint{1, 2, 5}, Result: "a"},
		{Weight: 0.1, Values: []string{"a", "b", "c"}, Weights: []uint{1, 2, 5}, Result: "a"},
		{Weight: 0.125, Values: []string{"a", "b", "c"}, Weights: []uint{1, 2, 5}, Result: "a"},
		{Weight: 0.2, Values: []string{"a", "b", "c"}, Weights: []uint{1, 2, 5}, Result: "b"},
		{Weight: 0.3, Values: []string{"a", "b", "c"}, Weights: []uint{1, 2, 5}, Result: "b"},
		{Weight: 0.375, Values: []string{"a", "b", "c"}, Weights: []uint{1, 2, 5}, Result: "b"},
		{Weight: 0.4, Values: []string{"a", "b", "c"}, Weights: []uint{1, 2, 5}, Result: "c"},
		{Weight: 0.5, Values: []string{"a", "b", "c"}, Weights: []uint{1, 2, 5}, Result: "c"},
		{Weight: 0.6, Values: []string{"a", "b", "c"}, Weights: []uint{1, 2, 5}, Result: "c"},
		{Weight: 0.7, Values: []string{"a", "b", "c"}, Weights: []uint{1, 2, 5}, Result: "c"},
		{Weight: 0.8, Values: []string{"a", "b", "c"}, Weights: []uint{1, 2, 5}, Result: "c"},
		{Weight: 0.9, Values: []string{"a", "b", "c"}, Weights: []uint{1, 2, 5}, Result: "c"},
		// Skip unmatched weights/values elements
		{Weight: 0.25, Values: []string{"a", "b", "c"}, Weights: []uint{1, 2}, Result: "a"},
		{Weight: 0.33, Values: []string{"a", "b"}, Weights: []uint{1, 2, 5}, Result: "a"},
		{Weight: 0.75, Values: []string{"a", "b"}, Weights: []uint{1, 2, 5}, Result: "b"},
	}

	for idx, el := range lst {
		mRng := new(RngMock)
		mRng.On("Float64").Return(el.Weight)
		f := NewFluky(mRng)
		result := f.Weighted(el.Values, el.Weights)
		assert.Equal(t, el.Result, result, "[%d] Weighted test: %v - %v: %s!=%s", idx, el.Values, el.Weights, el.Result, result)
		mRng.AssertExpectations(t)
	}
}
