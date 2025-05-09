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

wyr8At
Ξ go test -bench=. -timeout 30m ./wy
goos: darwin
goarch: arm64
pkg: github.com/Pencroff/fluky/wy
cpu: Apple M2
BenchmarkHash3Bytes-8           465490204                2.471 ns/op    1213.86 MB/s
BenchmarkHash8Bytes-8           421994221                2.842 ns/op    2814.47 MB/s
BenchmarkHash320Bytes-8         49417201                24.25 ns/op     13194.58 MB/s
BenchmarkHash1K-8               16127685                74.51 ns/op     13743.80 MB/s
BenchmarkHash8K-8                2091560               572.3 ns/op      14315.12 MB/s

go:inline
Ξ go test -bench=. -timeout 30m ./wy
goos: darwin
goarch: arm64
pkg: github.com/Pencroff/fluky/wy
cpu: Apple M2
BenchmarkHash3Bytes-8           451815946                2.483 ns/op    1208.40 MB/s
BenchmarkHash8Bytes-8           422200411                2.854 ns/op    2802.73 MB/s
BenchmarkHash320Bytes-8         68553800                16.99 ns/op     18836.32 MB/s
BenchmarkHash1K-8               24986118                47.53 ns/op     21544.83 MB/s
BenchmarkHash8K-8                3283581               364.9 ns/op      22452.06 MB/s

// *(*uint64)(unsafe.Pointer(&p[0]))
Ξ go test -bench=. -timeout 30m ./wy
goos: darwin
goarch: arm64
pkg: github.com/Pencroff/fluky/wy
cpu: Apple M2
BenchmarkHash3Bytes-8           459399382                2.480 ns/op    1209.43 MB/s
BenchmarkHash8Bytes-8           426460575                2.813 ns/op    2843.82 MB/s
BenchmarkHash320Bytes-8         74800530                16.04 ns/op     19949.05 MB/s
BenchmarkHash1K-8               26140420                45.87 ns/op     22324.28 MB/s
BenchmarkHash8K-8                3446850               348.1 ns/op      23534.71 MB/s


