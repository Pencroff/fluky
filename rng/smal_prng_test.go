package rng

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSmallPrngC(t *testing.T) {
	rng := NewSmallPrngC()
	rng.Seed(0x2a)
	lst := []uint64{
		0xa5719fd503fff432, 0x6076cbc48ac7a8da, 0x33e07875edf9b45a,
		0xb3c7f3cd329083e1, 0xe99b850931402707, 0x58294a12f5007957,
		0x4f9771e159d7906d, 0x3103ee57275f07b1, 0xd6468481aa17ec74,
		0x72e585b8879f22d2, 0x0494fca496e515ce, 0xedb73cb92f8222bb}
	for _, v := range lst {
		rnd := rng.Uint64()
		//fmt.Printf("0x%x,", rnd)
		assert.Equal(t, v, rnd, "%x != %x", rnd, v)
	}
}

func TestSmallPrng(t *testing.T) {
	rng := NewSmallPrng()
	rng.Seed(0x2a)
	lst := []uint64{
		0xa5719fd503fff432, 0x6076cbc48ac7a8da, 0x33e07875edf9b45a,
		0xb3c7f3cd329083e1, 0xe99b850931402707, 0x58294a12f5007957,
		0x4f9771e159d7906d, 0x3103ee57275f07b1, 0xd6468481aa17ec74,
		0x72e585b8879f22d2, 0x0494fca496e515ce, 0xedb73cb92f8222bb}
	for _, v := range lst {
		rnd := rng.Uint64()
		//fmt.Printf("0x%x,", rnd)
		assert.Equal(t, v, rnd, "%x != %x", rnd, v)
	}
}
