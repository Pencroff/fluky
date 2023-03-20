# External Sources

All sources implement the [`Source`](https://pkg.go.dev/math/rand#Source) and
[`Source64`](https://pkg.go.dev/math/rand#Source64) interface from `math/rand` package.

## SplitMix64 ([ref](https://prng.di.unimi.it/splitmix64.c))

The SplitMix64 generator is a simple, fast generator for 64-bit numbers. It is a very fast generator passing
BigCrush without systematic failures, but due to the relatively short period it is acceptable only for
applications with a mild amount of parallelism; otherwise, use a xorshift or pcg generator.

Current implementations of `xoshiro256++` and `xoshiro256**` use `SplitMix64` as a seed generator.

The type of `state` is `uint64`.

??? example "SplitMix64 implementation"
    ```go
    func (s *SplitMix64Source) Uint64() uint64 {
        s.state += 0x9e3779b97f4a7c15
        z := s.state
        z ^= z >> 30
        z *= 0xbf58476d1ce4e5b9
        z ^= z >> 27
        z *= 0x94d049bb133111eb
        z ^= z >> 31
        return z
    }
    ```

