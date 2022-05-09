package fluky

import (
	"github.com/stretchr/testify/assert"
	"testing"
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
