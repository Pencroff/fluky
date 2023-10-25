package primary

import (
	"fmt"
	"math/big"
	"math/rand"
	"strconv"
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
	if strings.Index(s, "0b") == 0 {
		s = s[2:]
	}
	for _, c := range s {
		if c == '1' {
			oneCnt++
		}
	}
	l := len(s)
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

func GetPrimeListWithParams(r *rand.Rand, n int, oneProb float64, delta float64) []uint64 {
	var res []uint64
	minProb := oneProb - delta
	maxProb := oneProb + delta

	fuse := n * 256

	for n > 0 {
		var s string
		for {
			s = Int64Str(r, 64, oneProb)
			p, _ := CalcOneProbability(s)
			if p >= minProb && p <= maxProb {
				break
			}
		}
		v, _ := strconv.ParseUint(s, 0, 64)
		prime := NextPrimeLess(v)
		sPrime := fmt.Sprintf("%#064b", prime)
		p, _ := CalcOneProbability(sPrime)
		fuse--
		if fuse < 0 {
			res = append(res, prime)
			fmt.Println("FUSE")
			fmt.Println(len(res))
			break
		}
		if p >= minProb && p <= maxProb {
			if strings.Index(sPrime, "11111") > -1 {
				continue
			}
			res = append(res, prime)
			n--
		}
	}
	return res
}
