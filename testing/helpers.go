package testing

import "math/bits"

func MonobitCount(v uint64) (r Monobit) {
	cnt := uint64(bits.OnesCount64(v))
	r.Ones = cnt
	r.Zeros = 64 - cnt
	return
}
