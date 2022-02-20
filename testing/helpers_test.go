package testing

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
