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

func TestPrimary_nextPrimeLess(t *testing.T) {
	testCases := []struct {
		name string
		n    uint64
		want uint64
	}{
		{"2", 2, 2},
		{"3", 3, 3},
		{"4", 4, 3},
		{"5", 5, 5},
		{"11", 11, 11},
		{"20", 20, 19},
		{"1<<64-1-59", 1<<64 - 1, 1<<64 - 59},
		{"1<<64-1-59", 1<<64 - 59, 1<<64 - 59},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := NextPrimeLess(tc.n)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestPrimary_findPrime(t *testing.T) {
	prob := 0.15
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

	n1 := NextPrimeLess(0x1050021020000047) // Remember 0x1052930420443f9d
	sq := n1 * n1
	n2 := NextPrimeLess(sq)
	sq2 := sq * sq
	n3 := NextPrimeLess(sq2)

	fmt.Println()

	fmt.Printf("N1: %#064b\n", n1)
	fmt.Printf("Dec: %d\t\t\tHex: %#016x\n", n1, n1)
	fmt.Printf("N2: %#064b\n", n2)
	fmt.Printf("Dec: %d\t\t\tHex: %#016x\n", n2, n2)
	fmt.Printf("N3: %#064b\n", n3)
	fmt.Printf("Dec: %d\t\t\tHex: %#016x\n", n3, n3)

	fmt.Println()

	fmt.Printf("Pr: %#064b\n", n1)
	fmt.Printf("Dec: %d\t\t\tHex: %#016x\n", n1, n1)
	//sq := n1 * n1
	fmt.Printf("Sq: %#064b\n", sq)
	fmt.Printf("Dec: %d\t\t\tHex: %#016x\n", sq, sq)
	di := n1 / n2
	fmt.Printf("Di: %#064b\n", di)
	fmt.Printf("Dec: %d\t\t\tHex: %#016x\n", di, di)
	pl := n1 + n1
	fmt.Printf("Pl: %#064b\n", pl)
	fmt.Printf("Dec: %d\t\t\tHex: %#016x\n", pl, pl)
}
