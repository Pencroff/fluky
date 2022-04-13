package rng

import "testing"

func BenchmarkRndGen(b *testing.B) {
	pcgc := NewPcgCRng()
	pcg := NewPcgRng()
	builtIn := NewBuiltIn()
	smallPrng := NewSmallPrng()
	squares := NewSquares()

	b.Run("BuiltIn", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			builtIn.Uint64()
		}
	})
	b.Run("PcgCRng - with CGO", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			pcgc.Uint64()
		}
	})
	b.Run("PcgRng - pure GO", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			pcg.Uint64()
		}
	})
	b.Run("Small Prng", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			smallPrng.Uint64()
		}
	})
	b.Run("Squares", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			squares.Uint64()
		}
	})
}
