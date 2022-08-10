package source

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPcgSource(t *testing.T) {
	src := NewPcgSource(0x2a)
	lst := []uint64{0x86b1da1d72062b68, 0x1304aa46c9853d39, 0xa3670e9e0dd50358,
		0xf9090e529a7dae00, 0xc85b9fd837996f2c, 0x606121f8e3919196,
		0x7ce1c7ff478354ba, 0xcbc4ac70e541310e, 0x74be71999ec37f2c,
		0xb81f9c99a934f1a7, 0x120e9901a900c97f, 0x0f983bad4b19f493}
	for _, v := range lst {
		rnd := src.Uint64()
		assert.Equal(t, v, rnd, "%x != %x", rnd, v)
	}
}
