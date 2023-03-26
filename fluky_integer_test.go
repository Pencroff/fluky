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
	mRng.On("Int63").Return(int64(1<<63 - 1)).Times(1)
	mRng.On("Int63").Return(int64(1 << 62)).Times(1)
	mRng.On("Int63").Return(int64(0)).Times(1)
	mRng.On("Int63").Return(int64(1)).Times(1)

	f := NewFluky(mRng)
	assert.Equal(t, 4611686018427387903, f.Integer())
	assert.Equal(t, 0, f.Integer())
	assert.Equal(t, -4611686018427387904, f.Integer())
	assert.Equal(t, -4611686018427387903, f.Integer())
	mRng.AssertExpectations(t)
}

func TestFluky_Integer_WithOptions(t *testing.T) {
	mRng := new(RngMock)
	mRng.On("Int63").Return(int64(1<<63 - 1)).Times(1)
	mRng.On("Int63").Return(int64(1 << 62)).Times(1)
	mRng.On("Int63").Return(int64(0)).Times(1)
	mRng.On("Int63").Return(int64(1)).Times(1)
	mRng.On("Int63").Return(int64(7351)).Times(1)

	f := NewFluky(mRng)
	assert.Equal(t, 75, f.Integer(WithIntRange(-100, 100)))
	assert.Equal(t, -12, f.Integer(WithIntRange(-100, 100)))
	assert.Equal(t, -100, f.Integer(WithIntRange(-100, 100)))
	assert.Equal(t, -99, f.Integer(WithIntRange(-100, 100)))
	assert.Equal(t, 15, f.Integer(WithIntRange(-100, 100)))
	assert.Equal(t, 5, f.Integer(WithIntRange(5, 5)))
	mRng.AssertExpectations(t)
}

func TestFluky_Integer_IncorrectInput(t *testing.T) {
	mRng := new(RngMock)
	mRng.On("Int63").Return(int64(1))

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
		assert.LessOrEqual(t, v, max)
	}
}
