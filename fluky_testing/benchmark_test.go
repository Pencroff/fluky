package fluky_testing

import (
	"github.com/Pencroff/fluky/source"
	"math/rand"
	"testing"
)

/*
λ go test -bench=. -timeout 30m .\fluky_testing
goos: windows
goarch: amd64
pkg: github.com/Pencroff/fluky/fluky_testing
cpu: Intel(R) Core(TM) i9-8950HK CPU @ 2.90GHz
BenchmarkSource/BuiltIn_Source-12               375685155                3.242 ns/op
BenchmarkSource/Pcg_Source-12                   197836590                6.279 ns/op
BenchmarkSource/SmallPrng_Source-12             481054144                2.934 ns/op
BenchmarkSource/xoshiro256++_Source-12          366972813                3.040 ns/op
BenchmarkSource/xoshiro256**_Source-12          426120020                2.725 ns/op
BenchmarkSource/SplitMix64-12                   502013072                2.299 ns/op
PASS
ok      github.com/Pencroff/fluky/fluky_testing 10.582s
------------------------------------------------------------------------------------
λ go test -bench=. -timeout 30m .\fluky_testing
goos: windows
goarch: amd64
pkg: github.com/Pencroff/fluky/fluky_testing
cpu: Intel(R) Core(TM) i9-8950HK CPU @ 2.90GHz
BenchmarkSource/BuiltIn_Source-12               355009314                3.247 ns/op
BenchmarkSource/Pcg_Source-12                   187504272                5.337 ns/op
BenchmarkSource/SmallPrng_Source-12             468746520                2.494 ns/op
BenchmarkSource/xoshiro256++_Source-12          405408144                2.838 ns/op
BenchmarkSource/xoshiro256**_Source-12          423457124                3.152 ns/op
BenchmarkSource/SplitMix64-12                   506380605                2.282 ns/op
PASS
ok      github.com/Pencroff/fluky/fluky_testing 9.223s
------------------------------------------------------------------------------------
λ go test -bench=. -timeout 30m .\fluky_testing
goos: windows
goarch: amd64
pkg: github.com/Pencroff/fluky/fluky_testing
cpu: Intel(R) Core(TM) i9-8950HK CPU @ 2.90GHz
BenchmarkSource/BuiltIn_Source-12               342288259                3.183 ns/op
BenchmarkSource/Pcg_Source-12                   183376828                6.106 ns/op
BenchmarkSource/SmallPrng_Source-12             446078412                2.501 ns/op
BenchmarkSource/xoshiro256++_Source-12          354640320                3.360 ns/op
BenchmarkSource/xoshiro256**_Source-12          434709998                3.186 ns/op
BenchmarkSource/SplitMix64-12                   563374726                2.411 ns/op
PASS
ok      github.com/Pencroff/fluky/fluky_testing 9.643s
------------------------------------------------------------------------------------
λ go test -bench=. -timeout 30m .\fluky_testing
goos: windows
goarch: amd64
pkg: github.com/Pencroff/fluky/fluky_testing
cpu: Intel(R) Core(TM) i9-8950HK CPU @ 2.90GHz
BenchmarkSource/BuiltIn_Source-12               355807413                3.307 ns/op
BenchmarkSource/Pcg_Source-12                   204869928                5.723 ns/op
BenchmarkSource/SmallPrng_Source-12             472371390                2.500 ns/op
BenchmarkSource/xoshiro256++_Source-12          371474130                3.171 ns/op
BenchmarkSource/xoshiro256**_Source-12          443611275                2.678 ns/op
BenchmarkSource/SplitMix64-12                   496955937                2.286 ns/op
PASS
ok      github.com/Pencroff/fluky/fluky_testing 9.324s
*/

var res uint64

func BenchmarkSource(b *testing.B) {
	var seed int64 = 1234567
	tbl := []struct {
		name string
		src  rand.Source
	}{
		{"BuiltIn Source", rand.NewSource(seed)},
		{"Pcg Source", source.NewPcgSource(seed)},
		{"SmallPrng Source", source.NewSmallPrngSource(seed)},
		{"xoshiro256++ Source", source.NewXoshiro256ppSource(seed)},
		{"xoshiro256** Source", source.NewXoshiro256ssSource(seed)},
		{"SplitMix64", source.NewSplitMix64Source(seed)},
	}
	for _, t := range tbl {
		r := t.src.(rand.Source64)
		b.Run(t.name, func(b *testing.B) {
			v := uint64(0)
			for i := 0; i < b.N; i++ {
				v = r.Uint64()
			}
			res = v
		})
		res += 1
	}
}
