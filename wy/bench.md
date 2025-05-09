Original

Ξ go test -bench=. -timeout 30m ./wy
goos: darwin
goarch: arm64
pkg: github.com/Pencroff/fluky/wy
cpu: Apple M2
BenchmarkHash3Bytes-8           413977370                2.708 ns/op    1107.76 MB/s
BenchmarkHash8Bytes-8           379983823                3.165 ns/op    2527.55 MB/s
BenchmarkHash320Bytes-8         69030314                17.39 ns/op     18398.17 MB/s
BenchmarkHash1K-8               24083791                48.46 ns/op     21130.27 MB/s
BenchmarkHash8K-8                3250099               368.5 ns/op      22231.74 MB/s

mull no pointers
Ξ go test -bench=. -timeout 30m ./wy
goos: darwin
goarch: arm64
pkg: github.com/Pencroff/fluky/wy
cpu: Apple M2
BenchmarkHash3Bytes-8           450967466                2.480 ns/op    1209.68 MB/s
BenchmarkHash8Bytes-8           419240212                2.883 ns/op    2774.83 MB/s
BenchmarkHash320Bytes-8         69352635                17.36 ns/op     18430.93 MB/s
BenchmarkHash1K-8               24265914                48.23 ns/op     21233.65 MB/s
BenchmarkHash8K-8                3234781               374.0 ns/op      21903.06 MB/s

case 12
Ξ go test -bench=. -timeout 30m ./wy
goos: darwin
goarch: arm64
pkg: github.com/Pencroff/fluky/wy
cpu: Apple M2
BenchmarkHash3Bytes-8           456105758                2.458 ns/op    1220.29 MB/s
BenchmarkHash8Bytes-8           424674283                2.838 ns/op    2819.33 MB/s
BenchmarkHash320Bytes-8         68888808                16.61 ns/op     19264.54 MB/s
BenchmarkHash1K-8               25690086                46.66 ns/op     21945.59 MB/s
BenchmarkHash8K-8                3341960               360.0 ns/op      22753.84 MB/s
