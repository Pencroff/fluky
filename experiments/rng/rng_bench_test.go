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
---
λ go test -bench=BenchmarkRndGen -timeout 30m
goos: windows
goarch: amd64
pkg: github.com/Pencroff/fluky/rng
cpu: Intel(R) Core(TM) i9-8950HK CPU @ 2.90GHz
BenchmarkRndGen/BuiltIn-12              213480511                5.354 ns/op
BenchmarkRndGen/PcgCRng_CGO-12          14445613                90.44 ns/op
BenchmarkRndGen/PcgRng-12               174679074                6.538 ns/op
BenchmarkRndGen/Small_Prng_CGO-12       14638108                80.58 ns/op
BenchmarkRndGen/Small_Prng-12           389652643                2.980 ns/op
BenchmarkRndGen/xoshiro256++-12         349818225                3.496 ns/op
BenchmarkRndGen/xoshiro256++_CGO-12     13426557                81.26 ns/op
PASS
ok      github.com/Pencroff/fluky/rng   10.817s
--λ go test -bench=BenchmarkRndGen -timeout 30m
goos: windows
goarch: amd64
pkg: github.com/Pencroff/fluky/rng
cpu: Intel(R) Core(TM) i9-8950HK CPU @ 2.90GHz
BenchmarkRndGen/BuiltIn-12              216459781                5.512 ns/op
BenchmarkRndGen/PcgCRng_CGO-12          15000430                86.05 ns/op
BenchmarkRndGen/PcgRng-12               191395281                6.234 ns/op
BenchmarkRndGen/Small_Prng_CGO-12       15421210                79.06 ns/op
BenchmarkRndGen/Small_Prng-12           387534187                2.821 ns/op
BenchmarkRndGen/xoshiro256++-12         360043672                3.388 ns/op
BenchmarkRndGen/xoshiro256++_CGO-12             15192853                78.41 ns/op
BenchmarkRndGen/SplitMix64-12                   657538569                1.839 ns/op
BenchmarkRndGen/SplitMix64C-12                  15702715                67.15 ns/op
PASS
ok      github.com/Pencroff/fluky/rng   13.374s
*/

func BenchmarkRndGen(b *testing.B) {
	pcgc := NewPcgCRng()
	pcg := NewPcgRng()
	builtIn := NewBuiltIn()
	smallPrng := NewSmallPrng()
	smallPrngC := NewSmallPrngC()
	xoshiro256pp := NewXoshiro256pp(11111)
	xoshiro256ppC := NewXoshiro256ppC(11111)
	splitMix64Rng := NewSplitMix64(11111)
	//squares := NewSquares()

	b.Run("BuiltIn", func(b *testing.B) {
		v := uint64(0)
		for i := 0; i < b.N; i++ {
			v = builtIn.Uint64()
		}
		res = v
	})
	b.Run("PcgCRng CGO", func(b *testing.B) {
		v := uint64(0)
		for i := 0; i < b.N; i++ {
			v = pcgc.Uint64()
		}
		res = v
	})
	b.Run("PcgRng", func(b *testing.B) {
		v := uint64(0)
		for i := 0; i < b.N; i++ {
			v = pcg.Uint64()
		}
		res = v
	})
	b.Run("Small Prng CGO", func(b *testing.B) {
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
	b.Run("xoshiro256++", func(b *testing.B) {
		v := uint64(0)
		for i := 0; i < b.N; i++ {
			v = xoshiro256pp.Uint64()
		}
		res = v
	})
	b.Run("xoshiro256++ CGO", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			xoshiro256ppC.Uint64()
		}
	})
	b.Run("SplitMix64", func(b *testing.B) {
		v := uint64(0)
		for i := 0; i < b.N; i++ {
			v = splitMix64Rng.Uint64()
		}
		res = v
	})
	b.Run("SplitMix64 CGO", func(b *testing.B) {
		v := uint64(0)
		for i := 0; i < b.N; i++ {
			v = splitMix64Rng.Uint64C()
		}
		res = v
	})
}
