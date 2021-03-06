package fluky

import (
	. "github.com/Pencroff/go-toolkit/general"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlukyGeneral_Weighted(t *testing.T) {
	lst := []struct {
		Weight  float64
		Values  []interface{}
		Weights []uint
		Result  string
	}{
		{Weight: 0.0, Values: InterfaceSlice([]string{"a", "b", "c"}), Weights: []uint{1, 2, 5}, Result: "a"},
		{Weight: 0.1, Values: InterfaceSlice([]string{"a", "b", "c"}), Weights: []uint{1, 2, 5}, Result: "a"},
		{Weight: 0.125, Values: InterfaceSlice([]string{"a", "b", "c"}), Weights: []uint{1, 2, 5}, Result: "a"},
		{Weight: 0.2, Values: InterfaceSlice([]string{"a", "b", "c"}), Weights: []uint{1, 2, 5}, Result: "b"},
		{Weight: 0.3, Values: InterfaceSlice([]string{"a", "b", "c"}), Weights: []uint{1, 2, 5}, Result: "b"},
		{Weight: 0.375, Values: InterfaceSlice([]string{"a", "b", "c"}), Weights: []uint{1, 2, 5}, Result: "b"},
		{Weight: 0.4, Values: InterfaceSlice([]string{"a", "b", "c"}), Weights: []uint{1, 2, 5}, Result: "c"},
		{Weight: 0.5, Values: InterfaceSlice([]string{"a", "b", "c"}), Weights: []uint{1, 2, 5}, Result: "c"},
		{Weight: 0.6, Values: InterfaceSlice([]string{"a", "b", "c"}), Weights: []uint{1, 2, 5}, Result: "c"},
		{Weight: 0.7, Values: InterfaceSlice([]string{"a", "b", "c"}), Weights: []uint{1, 2, 5}, Result: "c"},
		{Weight: 0.8, Values: InterfaceSlice([]string{"a", "b", "c"}), Weights: []uint{1, 2, 5}, Result: "c"},
		{Weight: 0.9, Values: InterfaceSlice([]string{"a", "b", "c"}), Weights: []uint{1, 2, 5}, Result: "c"},
		// Skip unmatched weights/values elements
		{Weight: 0.25, Values: InterfaceSlice([]string{"a", "b", "c"}), Weights: []uint{1, 2}, Result: "a"},
		{Weight: 0.33, Values: InterfaceSlice([]string{"a", "b"}), Weights: []uint{1, 2, 5}, Result: "a"},
		{Weight: 0.75, Values: InterfaceSlice([]string{"a", "b"}), Weights: []uint{1, 2, 5}, Result: "b"},
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

func TestFluky_PickOne(t *testing.T) {
	lst := []struct {
		Rnd    uint64
		Values []interface{}
		Idx    int
		Result string
	}{
		{Rnd: 0, Values: InterfaceSlice([]string{"a", "b", "c"}), Idx: 0, Result: "a"},
		{Rnd: 1, Values: InterfaceSlice([]string{"a", "b", "c"}), Idx: 1, Result: "b"},
		{Rnd: 2, Values: InterfaceSlice([]string{"a", "b", "c"}), Idx: 2, Result: "c"},
		{Rnd: 3, Values: InterfaceSlice([]string{"a", "b", "c"}), Idx: 0, Result: "a"},
	}

	for idx, el := range lst {
		mRng := new(RngMock)
		mRng.On("Uint64").Return(el.Rnd)
		f := NewFluky(mRng)
		rndIdx, result := f.PickOne(el.Values)
		assert.Equal(t, rndIdx, el.Idx)
		assert.Equal(t, el.Result, result, "[%d] PickOne test: %v - %v: %s!=%s", idx, el.Values, el.Rnd, el.Result, result)
		mRng.AssertExpectations(t)
	}
}

func TestFluky_PickOne_Empty(t *testing.T) {
	lst := []struct {
		Values []interface{}
		Result interface{}
	}{
		{Values: InterfaceSlice([]string{}), Result: interface{}(nil)},
		{Values: []interface{}(nil), Result: interface{}(nil)},
		{Values: InterfaceSlice([]string{}), Result: nil},
		{Values: []interface{}(nil), Result: nil},
	}

	for idx, el := range lst {
		mRng := new(RngMock)
		mRng.On("Uint64").Maybe()
		f := NewFluky(mRng)
		rndIdx, result := f.PickOne(el.Values)
		assert.Equal(t, -1, rndIdx)
		assert.Equal(t, el.Result, result, "[%d] PickOne empty test: %s!=%s", idx, el.Result, result)
		mRng.AssertExpectations(t)
	}

}
