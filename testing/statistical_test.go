package testing

import (
	"fmt"
	"github.com/Pencroff/fluky"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestFrequencyMonobits(t *testing.T) {
	size := int(20 * Size1Gb / 8)
	stat := make(map[string]float64)
	statMonobit := make(map[string]Monobit)
	minName := ""
	minVal := 1.0
	for _, el := range RngTbl {
		total := CountMonobit(size, el.rnd)
		statMonobit[el.name] = total
		prob := float64(total.Ones) / float64(total.Zeros)
		assert.InDelta(t, prob, 1, 1e-05)
		stat[el.name] = math.Abs(1 - prob)
		if stat[el.name] < minVal {
			minName = el.name
			minVal = stat[el.name]
		}
	}

	PrintMonobitStats(size, stat, statMonobit, minName, minVal)

}

func CountMonobit(size int, rng fluky.RandomGenerator) Monobit {
	total := Monobit{}
	for i := 0; i < size; i++ {
		v := rng.Uint64()
		m := MonobitCount(v)
		total.Ones += m.Ones
		total.Zeros += m.Zeros
	}
	return total
}

func PrintMonobitStats(size int, stat map[string]float64, statMonobit map[string]Monobit, minName string, minVal float64) {
	fmt.Println()
	fmt.Printf("TestFrequencyMonobits stats (%d numbers) ~ %d Gb of random bytes\n", size, uint64(size)*8/Size1Gb)
	fmt.Println()
	fmt.Printf("|Name| 1 - (ones / zeros) - should be close to 0| Values |\n")
	fmt.Printf("| --- | --- | --- |\n")
	for key, val := range stat {
		fmt.Printf("| %16s | %e | [ 1: %d ; 0: %d ] |\n", key, val, statMonobit[key].Ones, statMonobit[key].Zeros)
	}
	fmt.Println()
	fmt.Printf("The winner: %s (%e)\n", minName, minVal)
	fmt.Println()
}
