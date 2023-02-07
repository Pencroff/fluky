package experiments

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
