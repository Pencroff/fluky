package fluky

import (
	"github.com/stretchr/testify/assert"
	"testing"
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
