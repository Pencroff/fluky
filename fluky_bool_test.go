package fluky

import (
	src "github.com/Pencroff/fluky/source"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func TestFluky_Bool_NoOptions(t *testing.T) {
	mRng := new(RngMock)
	mRng.On("Float64").Return(0.75).Times(1)
	mRng.On("Float64").Return(0.25).Times(1)
	mRng.On("Float64").Return(0.5).Times(1)
	mRng.On("Float64").Return(0.4999999999999).Times(1)

	f := NewFluky(mRng)
	assert.Equal(t, false, f.Bool())
	assert.Equal(t, true, f.Bool())
	assert.Equal(t, false, f.Bool())
	assert.Equal(t, true, f.Bool())

	mRng.AssertExpectations(t)
}

func TestFluky_Bool_WithOptions(t *testing.T) {
	mRng := new(RngMock)
	mRng.On("Float64").Return(0.75).Times(1)
	mRng.On("Float64").Return(0.25).Times(1)
	mRng.On("Float64").Return(0.75).Times(1)
	mRng.On("Float64").Return(0.74999999999999).Times(1)

	f := NewFluky(mRng)
	assert.Equal(t, true, f.Bool(WithLikelihood(0.9)))
	assert.Equal(t, false, f.Bool(WithLikelihood(0.1)))
	assert.Equal(t, false, f.Bool(WithLikelihood(0.75)))
	assert.Equal(t, true, f.Bool(WithLikelihood(0.75)))

	mRng.AssertExpectations(t)
}

func TestFluky_Bool_IncorrectInput(t *testing.T) {
	mRng := new(RngMock)
	mRng.On("Float64").Return(0.5)

	f := NewFluky(mRng)
	assert.Equal(t, false, f.Bool(WithLikelihood(-10)))
	assert.Equal(t, false, f.Bool(WithLikelihood(-1)))
	assert.Equal(t, false, f.Bool(WithLikelihood(0)))
	assert.Equal(t, true, f.Bool(WithLikelihood(1)))
	assert.Equal(t, true, f.Bool(WithLikelihood(10)))

	mRng.AssertExpectations(t)
}

func TestFluky_Bool_WithRnd(t *testing.T) {
	s := src.NewSplitMix64Source(time.Now().UnixNano())
	r := rand.New(s)
	f := NewFluky(r)
	n := 100000
	likelihood := 0.33
	threshold := likelihood * float64(n)
	count := 0
	for i := 0; i < n; i++ {
		if f.Bool(WithLikelihood(likelihood)) {
			count++
		}
	}
	assert.InDelta(t, float64(count), threshold, 0.01*threshold)
}
