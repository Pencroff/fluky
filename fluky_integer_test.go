package fluky

import (
	src "github.com/Pencroff/fluky/source"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func TestFluky_Integer_NoOptions(t *testing.T) {
	mRng := new(RngMock)
	mRng.On("Uint64").Return(uint64(1<<64 - 1)).Times(1)
	mRng.On("Uint64").Return(uint64(1<<63 - 1)).Times(1)
	mRng.On("Uint64").Return(uint64(1 << 63)).Times(1)
	mRng.On("Uint64").Return(uint64(1)).Times(1)

	f := NewFluky(mRng)
	assert.Equal(t, -1, f.Integer())
	assert.Equal(t, 0x7fffffffffffffff, f.Integer())
	assert.Equal(t, -0x8000000000000000, f.Integer())
	assert.Equal(t, 1, f.Integer())
	mRng.AssertExpectations(t)
}

func TestFluky_Integer_WithOptions(t *testing.T) {
	mRng := new(RngMock)
	mRng.On("Uint64").Return(uint64(1<<64 - 1)).Times(1)
	mRng.On("Uint64").Return(uint64(1<<63 - 1)).Times(1)
	mRng.On("Uint64").Return(uint64(1 << 63)).Times(1)
	mRng.On("Uint64").Return(uint64(1)).Times(1)
	mRng.On("Uint64").Return(uint64(7351)).Times(1)

	f := NewFluky(mRng)
	assert.Equal(t, -101, f.Integer(WithIntRange(-100, 100)))
	assert.Equal(t, -93, f.Integer(WithIntRange(-100, 100)))
	assert.Equal(t, -108, f.Integer(WithIntRange(-100, 100)))
	assert.Equal(t, -99, f.Integer(WithIntRange(-100, 100)))
	assert.Equal(t, 51, f.Integer(WithIntRange(-100, 100)))
	assert.Equal(t, 5, f.Integer(WithIntRange(5, 5)))
	mRng.AssertExpectations(t)
}

func TestFluky_Integer_IncorrectInput(t *testing.T) {
	mRng := new(RngMock)
	mRng.On("Uint64").Return(uint64(1))

	f := NewFluky(mRng)
	assert.Equal(t, 11, f.Integer(WithIntRange(10, -10)))
	assert.Equal(t, 1, f.Integer(WithIntRange(0, -10)))

	mRng.AssertExpectations(t)
}

func TestFluky_Integer_WithRand(t *testing.T) {
	s := src.NewSplitMix64Source(time.Now().UnixNano())
	r := rand.New(s)
	f := NewFluky(r)
	n := 10000
	min := 256
	max := 512
	for i := 0; i < n; i++ {
		v := f.Integer(WithIntRange(min, max))
		assert.GreaterOrEqual(t, v, min)
		assert.Less(t, v, max)
	}
}
