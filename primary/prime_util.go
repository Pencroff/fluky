package primary

import (
	"math/big"
	"math/rand"
	"strings"
)

func Int64Str(r *rand.Rand, size int, oneProb float64) string {
	var b strings.Builder
	for i := 0; i < size; i += 1 {
		if r.Float64() < oneProb {
			b.WriteString("1")
		} else {
			b.WriteString("0")
		}
		//if i > 0 && i%4 == 3 {
		//	b.WriteString("_")
		//}
	}
	bStr := b.String()
	b.Reset()
	b.WriteString("0b")
	l := len(bStr)
	if l < 64 {
		for i := 0; i < 64-l; i++ {
			b.WriteByte('0')
		}
	}
	for i := l - 1; i >= 0; i-- {
		b.WriteByte(bStr[i])
	}
	return b.String()
}

func CalcOneProbability(s string) (float64, int) {
	oneCnt := 0
	for _, c := range s {
		if c == '1' {
			oneCnt++
		}
	}
	l := len(s)
	if strings.Index(s, "0b") == 0 {
		l -= 2
	}
	return float64(oneCnt) / float64(l), oneCnt
}

func NextPrimeLess(n uint64) uint64 {
	var bi big.Int
	bi.SetUint64(n)

	for {
		if bi.ProbablyPrime(64) {
			break
		}
		bi.Sub(&bi, big.NewInt(1))
	}
	return bi.Uint64()
}
