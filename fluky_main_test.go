package fluky

import (
	"github.com/Pencroff/fluky/source"
	"github.com/stretchr/testify/assert"
	"math"
	"math/rand"
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

func TestFluky_PcgSource(t *testing.T) {
	src := source.NewPcgXslRrSource(0x2a)
	rng := rand.New(src)
	f := NewFluky(rng)
	lst := []uint64{
		0x86b1da1d72062b68, 0x1304aa46c9853d39, 0xa3670e9e0dd50358,
		0xf9090e529a7dae00, 0xc85b9fd837996f2c, 0x606121f8e3919196,
		0x7ce1c7ff478354ba, 0xcbc4ac70e541310e, 0x74be71999ec37f2c,
		0xb81f9c99a934f1a7, 0x120e9901a900c97f, 0x0f983bad4b19f493}
	for _, v := range lst {
		rnd := f.Uint64()
		assert.Equal(t, v, rnd, "%x != %x", rnd, v)
	}
}

func TestFluky_SmallPrngSource(t *testing.T) {
	src := source.NewSmallPrngSource(0x2a)
	rng := rand.New(src)
	f := NewFluky(rng)
	lst := []uint64{
		0xa5719fd503fff432, 0x6076cbc48ac7a8da, 0x33e07875edf9b45a,
		0xb3c7f3cd329083e1, 0xe99b850931402707, 0x58294a12f5007957,
		0x4f9771e159d7906d, 0x3103ee57275f07b1, 0xd6468481aa17ec74,
		0x72e585b8879f22d2, 0x0494fca496e515ce, 0xedb73cb92f8222bb}
	for _, v := range lst {
		rnd := f.Uint64()
		assert.Equal(t, v, rnd, "%x != %x", rnd, v)
	}
}
