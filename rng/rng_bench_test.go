package rng

import "testing"

var res uint64

/*
λ go test -bench=BenchmarkRndGen -timeout 30m
goos: windows
goarch: amd64
pkg: github.com/Pencroff/fluky/rng
cpu: Intel(R) Core(TM) i9-8950HK CPU @ 2.90GHz
BenchmarkRndGen/BuiltIn-12               		228414409                4.987 ns/op
BenchmarkRndGen/PcgCRng_-_with_CGO-12           14487975                77.26 ns/op
BenchmarkRndGen/PcgRng_-_pure_GO-12             192425984                5.896 ns/op
BenchmarkRndGen/Small_Prng_-_with_CGO-12        15584030                77.81 ns/op
BenchmarkRndGen/Small_Prng-12                   405450333                2.685 ns/op
PASS
ok      github.com/Pencroff/fluky/rng   7.610s
*/

func BenchmarkRndGen(b *testing.B) {
	pcgc := NewPcgCRng()
	pcg := NewPcgRng()
	builtIn := NewBuiltIn()
	smallPrng := NewSmallPrng()
	smallPrngC := NewSmallPrngC()
	//squares := NewSquares()

	b.Run("BuiltIn", func(b *testing.B) {
		v := uint64(0)
		for i := 0; i < b.N; i++ {
			v = builtIn.Uint64()
		}
		res = v
	})
	b.Run("PcgCRng - with CGO", func(b *testing.B) {
		v := uint64(0)
		for i := 0; i < b.N; i++ {
			v = pcgc.Uint64()
		}
		res = v
	})
	b.Run("PcgRng - pure GO", func(b *testing.B) {
		v := uint64(0)
		for i := 0; i < b.N; i++ {
			v = pcg.Uint64()
		}
		res = v
	})
	b.Run("Small Prng - with CGO", func(b *testing.B) {
		v := uint64(0)
		for i := 0; i < b.N; i++ {
			v = smallPrngC.Uint64()
		}
		res = v
	})
	b.Run("Small Prng", func(b *testing.B) {
		v := uint64(0)
		for i := 0; i < b.N; i++ {
			v = smallPrng.Uint64()
		}
		res = v
	})
	//b.Run("Squares", func(b *testing.B) {
	//	for i := 0; i < b.N; i++ {
	//		squares.Uint64()
	//	}
	//})
}
