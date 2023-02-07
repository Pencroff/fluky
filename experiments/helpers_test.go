package experiments

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMonobitCount(t *testing.T) {
	var tbl = []struct {
		value uint64
		cnt0  uint64
		cnt1  uint64
	}{
		{
			value: 0b0000000000000000000000000000000000000000000000000000000000000001,
			cnt0:  63,
			cnt1:  1,
		},
		{
			value: 0b0000111111111111111111111111111111111111111111111111111111111111,
			cnt0:  4,
			cnt1:  60,
		},
		{
			value: 0b1010101010101010101010101010101010101010101010101010101010101010,
			cnt0:  32,
			cnt1:  32,
		},
	}
	for _, el := range tbl {
		m := MonobitCount(el.value)
		assert.Equal(t, el.cnt0, m.Zeros)
		assert.Equal(t, el.cnt1, m.Ones)
	}
}

func TestSizeToRange(t *testing.T) {
	var tbl = []struct {
		size   uint64
		groups uint64
		out    [][]uint64
	}{
		{
			size:   100,
			groups: 2,
			out:    [][]uint64{{0, 49}, {50, 99}},
		},
		{
			size:   100,
			groups: 3,
			out:    [][]uint64{{0, 32}, {33, 65}, {66, 99}},
		},
		{
			size:   100,
			groups: 4,
			out:    [][]uint64{{0, 24}, {25, 49}, {50, 74}, {75, 99}},
		},
		{
			size:   1e6,
			groups: 2,
			out:    [][]uint64{{0, 499_999}, {500_000, 999_999}},
		},
		{
			size:   100,
			groups: 12,
			out: [][]uint64{
				{0, 7}, {8, 15}, {16, 23}, {24, 31}, {32, 39},
				{40, 47}, {48, 55}, {56, 63}, {64, 71}, {72, 79},
				{80, 87}, {88, 99},
			},
		},
	}
	for _, el := range tbl {
		v := SizeToRange(el.size, el.groups)
		assert.Equal(t, el.out, v)
	}
}
