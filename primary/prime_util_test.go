package primary

import (
	"fmt"
	rng "github.com/Pencroff/fluky/source"
	"github.com/stretchr/testify/assert"
	"math/big"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestPrimary_largeIntStr(t *testing.T) {
	seed := time.Now().UnixNano()
	source := rng.NewSmallPrngSource(seed)
	rnd := rand.New(source)
	size := 63
	oneProb := 0.5
	s := Int64Str(rnd, size, oneProb)
	prob, _ := CalcOneProbability(s)
	delta := oneProb / 4
	//fmt.Printf("Res: %s, oneCnt: %d, oneProb: %f (%f)\n", s, oneCnt, prob, delta)
	assert.InDelta(t, prob, oneProb, delta)
}

func TestPrimary_largeIntStrToInt(t *testing.T) {
	seed := time.Now().UnixNano()
	source := rng.NewSmallPrngSource(seed)
	rnd := rand.New(source)
	size := 60
	oneProb := 0.5
	s := Int64Str(rnd, size, oneProb)
	//fmt.Printf("Result: %s\n", s)
	if i, err := strconv.ParseUint(s, 0, 64); err == nil {
		//fmt.Printf("Parsed: %#064b\n", i)
		//fmt.Printf("%T, %#016x\n", i, i)
		assert.Equal(t, fmt.Sprintf("%#064b", i), s)
	}
}

func TestPrimary_calcOneProbability(t *testing.T) {
	s := "0b0011"
	oneProb, oneCnt := CalcOneProbability(s)
	assert.Equal(t, oneProb, 0.5)
	assert.Equal(t, oneCnt, 2)
}

func TestPrimary_isPrime(t *testing.T) {
	testCases := []struct {
		name string
		n    uint64
		k    int
		want bool
	}{
		{"2", 2, 5, true},
		{"3", 3, 5, true},
		{"4", 4, 5, false},
		{"5", 5, 5, true},
		{"11", 11, 5, true},
		{"20", 20, 5, false},
		{"1<<64-1-59", 1<<64 - 59, 20, true},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var bi big.Int
			bi.SetUint64(tc.n)
			got := bi.ProbablyPrime(tc.k)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestPrimary_findPrime(t *testing.T) {
	prob := 0.33
	delta := prob / 4
	minProb := prob - delta
	maxProb := prob + delta
	seed := time.Now().UnixNano()
	source := rng.NewSmallPrngSource(seed)
	rnd := rand.New(source)
	size := 64
	var s string
	for {
		s = Int64Str(rnd, size, prob)
		realProb, oneCnt := CalcOneProbability(s)
		fmt.Printf("Result: %s, oneCnt: %d, oneProb: %f (%f)\n", s, oneCnt, realProb, delta)
		if realProb > minProb && realProb < maxProb {
			break
		}
	}
	var i uint64
	i, _ = strconv.ParseUint(s, 0, 64)
	fmt.Printf("Parsed: %#064b\n", i)

	var bi big.Int
	bi.SetUint64(i)

	for {
		if bi.ProbablyPrime(100) {
			break
		}
		bi.Sub(&bi, big.NewInt(1))
	}

	fmt.Printf("Prime:  %#064b\n", bi.Uint64())
	fmt.Printf("Dec: %d\t\t\tHex: %#016x\n", bi.Uint64(), bi.Uint64())
}
