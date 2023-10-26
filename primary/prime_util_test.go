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

func TestGetPrimeListWithParams(t *testing.T) {
	seed := time.Now().UnixNano()
	source := rng.NewSmallPrngSource(seed)
	rnd := rand.New(source)
	// 11 13 17 19 23 29 31 37 41 43 47 53 59 61 67 71 73 79 83 89 97
	oneProb := 0.53
	delta := 0.01 // oneProb / 32
	minProb := oneProb - delta
	maxProb := oneProb + delta
	n := 128
	res := GetPrimeListWithParams(rnd, n, oneProb, delta)
	assert.Equal(t, n, len(res))

	for _, v := range res {
		s64 := fmt.Sprintf("%#064b", v)
		fmt.Println("---")
		fmt.Printf("Dec: %d\t\t\tHex: %#016x\n", v, v)
		fmt.Printf("%s\n", s64)
		p, c := CalcOneProbability(s64)
		fmt.Printf("oneCnt: %d, %.4f < %.4f < %.4f\n", c, maxProb, p, minProb)
	}
	for _, v := range res {
		prob, _ := CalcOneProbability(fmt.Sprintf("%#064b", v))
		assert.InDelta(t, oneProb, prob, delta)
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

func TestPrimary_MeasureOneProb(t *testing.T) {
	//var lst = []uint64{0x6831DA4, 0xB5297A4D, 0x1856C4E9} // 37.5% 53.1% 43.7%
	//
	// https://github.com/Cyan4973/xxHash/blob/dev/doc/xxhash_spec.md
	lst := []uint64{
		0b1001111000110111011110011011000110000101111010111100101010000111,
		0b1100001010110010101011100011110100100111110101001110101101001111,
		0b0001011001010110011001111011000110011110001101110111100111111001,
		0b1000010111101011110010100111011111000010101100101010111001100011,
		0b0010011111010100111010110010111100010110010101100110011111000101,
	}
	/*
		0b1001111000110111011110011011000110000101111010111100101010000111
		oneCnt: 36, oneProb: 0.562500
		Dec: 11400714785074694791			Hex: 0x9e3779b185ebca87
		---
		0b1100001010110010101011100011110100100111110101001110101101001111
		oneCnt: 36, oneProb: 0.562500
		Dec: 14029467366897019727			Hex: 0xc2b2ae3d27d4eb4f
		---
		0b1011001010110011001111011000110011110001101110111100111111001
		oneCnt: 37, oneProb: 0.606557
		Dec: 1609587929392839161			Hex: 0x165667b19e3779f9
		---
		0b1000010111101011110010100111011111000010101100101010111001100011
		oneCnt: 35, oneProb: 0.546875
		Dec: 9650029242287828579			Hex: 0x85ebca77c2b2ae63
		---
		0b10011111010100111010110010111100010110010101100110011111000101
		oneCnt: 35, oneProb: 0.564516
		Dec: 2870177450012600261			Hex: 0x27d4eb2f165667c5
	*/
	for _, v := range lst {
		fmt.Printf("0x%08x\n", v)
		s := fmt.Sprintf("%#032b", v)
		realProb, oneCnt := CalcOneProbability(s)
		fmt.Println(s)
		fmt.Printf("oneCnt: %d, oneProb: %f\n", oneCnt, realProb)
		fmt.Printf("Dec: %d\t\t\tHex: %#016x\n", v, v)
	}
}

func TestPrimary_BIT_NOISE_A(t *testing.T) {
	var BIT_NOISE_A uint64 = 0x5210100a71212cc5
	cnt := 0
	for i := 1; i < 1024; i++ {
		r := BIT_NOISE_A * uint64(i)
		b := r / uint64(i)
		if b != BIT_NOISE_A {
			cnt++
			fmt.Println("==========")

			fmt.Printf("i: %d\n", i)
			fmt.Printf("n: %#064b\n", BIT_NOISE_A)
			fmt.Printf("b: %#064b\n", b)
			fmt.Printf("r: %#064b\n", r)
		}
	}
	fmt.Printf("cnt: %d\n", cnt)
}
