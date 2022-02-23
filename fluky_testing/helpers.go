package fluky_testing

import "math/bits"

func SizeToRange(size uint64, groups uint64) (r [][]uint64) {
	oneGroupCnt := size / groups
	lastGroup := int(groups - 1)
	firstIdx := uint64(0)
	lastIdx := size - 1
	for i := 0; i < int(groups); i++ {
		if i == lastGroup {
			r = append(r, []uint64{firstIdx, lastIdx})
		} else {
			r = append(r, []uint64{firstIdx, firstIdx + oneGroupCnt - 1})
			firstIdx += oneGroupCnt
		}
	}
	return
}

func MonobitCount(v uint64) (r Monobit) {
	cnt := uint64(bits.OnesCount64(v))
	r.Ones = cnt
	r.Zeros = 64 - cnt
	return
}
