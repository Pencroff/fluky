package fluky_testing

import (
	"fmt"
	"github.com/Pencroff/fluky/rng"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestFrequencyMonobits(t *testing.T) {
	size := int(10 * Size1Gb / 8)
	stat := make(map[string]float64)
	statMonobit := make(map[string]Monobit)
	maxName := ""
	maxVal := 0.0
	for _, el := range RngTbl {
		total := CountMonobit(0, uint64(size), el.rnd)
		statMonobit[el.name] = total
		s := math.Abs(float64(total.Ones) - float64(total.Zeros))
		n := float64(total.Ones + total.Zeros)
		sOb := s / math.Sqrt(n)
		pVal := math.Erfc(sOb / math.Sqrt(2.0))
		assert.True(t, pVal >= 0.01,
			"%s: P-value less then 0.01 - %e, [s=%e, n=%e, s_ob=%e]",
			el.name, pVal, s, n, sOb)
		stat[el.name] = pVal
		if stat[el.name] > maxVal {
			maxName = el.name
			maxVal = stat[el.name]
		}
	}
	PrintMonobitStats("TestFrequencyMonobits", size, stat, statMonobit, maxName, maxVal)
}

func CountMonobit(from, to uint64, rng rng.RandomGenerator) Monobit {
	total := Monobit{}
	for i := from; i < to; i++ {
		v := rng.Uint64()
		m := MonobitCount(v)
		total.Ones += m.Ones
		total.Zeros += m.Zeros
	}
	return total
}

func PrintMonobitStats(name string, size int, stat map[string]float64, statMonobit map[string]Monobit, minName string, minVal float64) {
	fmt.Println()
	fmt.Printf("%s stats (%d numbers) ~ %d Gb of random bytes each\n", name, size, uint64(size)*8/Size1Gb)
	fmt.Println()
	fmt.Printf("|Name| P-value | Values | Status |\n")
	fmt.Printf("| --- | --- | --- |  --- |\n")
	for key, val := range stat {
		status := "PASS"
		if val < 0.01 {
			status = "FAIL"
		}
		fmt.Printf("| %16s | %e | [ 1: %d ; 0: %d ] | %s |\n", key, val, statMonobit[key].Ones, statMonobit[key].Zeros, status)
	}
	fmt.Println()
	fmt.Printf("Max value: %s (%e)\n", minName, minVal)
	fmt.Println()
}
