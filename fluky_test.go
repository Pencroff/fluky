package fluky

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestFluky_Seed(t *testing.T) {
	mRng := new(RngMock)
	mRng.On("Seed", int64(10))
	mRng.On("Seed", int64(math.MaxInt64-1))
	f := NewFluky(mRng)
	f.Seed(10)
	f.Seed(math.MaxInt64 - 1)
	mRng.AssertExpectations(t)
}

func TestFluky_Uint64(t *testing.T) {
	mRng := new(RngMock)
	mRng.On("Uint64").Return(uint64(0xaf)).Times(2)
	mRng.On("Uint64").Return(uint64((1 << 64) - 1)).Times(1)
	f := NewFluky(mRng)
	assert.Equal(t, uint64(0xaf), f.Uint64())
	assert.Equal(t, uint64(0xaf), f.Uint64())
	assert.Equal(t, uint64((1<<64)-1), f.Uint64())
	mRng.AssertExpectations(t)
}

func TestFluky_Int63(t *testing.T) {
	mRng := new(RngMock)
	mRng.On("Int63").Return(int64(0xaf)).Times(2)
	mRng.On("Int63").Return(int64((1 << 63) - 1)).Times(1)
	f := NewFluky(mRng)
	assert.Equal(t, int64(0xaf), f.Int63())
	assert.Equal(t, int64(0xaf), f.Int63())
	assert.Equal(t, int64((1<<63)-1), f.Int63())
	mRng.AssertExpectations(t)
}

func TestFluky_Float64(t *testing.T) {
	mRng := new(RngMock)
	mRng.On("Float64").Return(0.31415).Times(1)
	mRng.On("Float64").Return(0.11543).Times(1)
	f := NewFluky(mRng)
	assert.Equal(t, 0.31415, f.Float64())
	assert.Equal(t, 0.11543, f.Float64())
	mRng.AssertExpectations(t)
}
