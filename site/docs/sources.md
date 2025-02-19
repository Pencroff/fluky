# External Sources

Fluky library provides a set of external sources for random number generation.
They have a good statistical quality and are fast for general purpose cases.
All sources are implement the [`Source`](https://pkg.go.dev/math/rand#Source) and
[`Source64`](https://pkg.go.dev/math/rand#Source64) interfaces from `math/rand` package.

**Import all sources**

```go
import "github.com/Pencroff/fluky/sources"
```

## SplitMix64

!!! info inline end "Reference"

    [splitmix64.c](https://prng.di.unimi.it/splitmix64.c)

SplitMix64 is a pseudorandom number generator (PRNG) algorithm that generates a 64-bit output using a 64-bit seed.
It works by repeatedly applying a basic function to the seed value, but it may not provide enough randomness
for some applications, and its output is not suitable for cryptographic purposes. However, SplitMix64 is known for
its speed and simplicity, and it can be used as a basic building block for more complex PRNGs.

Current implementations of `xoshiro256++` and `xoshiro256**` use `SplitMix64` as a seed generator.

**Usage**

```go
src := sources.NewSplitMix64Source(0) // where 0 is a seed, int64
```

??? example "SplitMix64 implementation"

    `state` has `uint64` type.

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

## Xoroshiro256++ / Xoroshiro256**

!!! info inline end "Reference"

    * [xoshiro generators](https://prng.di.unimi.it/)
    * [xoroshiro256++](https://prng.di.unimi.it/xoshiro256plusplus.c)
    * [xoroshiro256**](https://prng.di.unimi.it/xoshiro256starstar.c)

Xoshiro is a family of fast and efficient pseudorandom number generators (PRNGs) designed by David Blackman
and Sebastiano Vigna. They are known for producing high-quality output with a relatively small state size,
using a combination of bitwise operations and nonlinear transformations.
Fluky provides Xoroshiro256++ and Xoroshiro256** implementations (64 bit output and 2^256^-1 period),
but they are not suitable for cryptographic purposes due to vulnerabilities in their design.
Overall, Xoshiro generators are ideal for simulations, games, and other applications where high-quality
randomness is required.

Fluky implementation of `xoshiro256++` and `xoshiro256**` use `SplitMix64` as a seed generator.

**Usage**

=== "xoshiro256++"

    ```go
    src := sources.NewXoshiro256ppSource(0) // where 0 is a seed, int64
    ```

=== "xoshiro256**"

    ```go
    src := sources.NewXoshiro256ssSource(0) // where 0 is a seed, int64
    ```

Xoroshiro256 implementations provide `Jump` and `LongJump` methods which is return new source with current state,
and move state of existing source to 2^128^ and 2^192^ iterations accordingly.

``` { .go .no-copy }
s := src.(*src.Xoshiro256ppSource) 
// Same cast to Xoshiro256ssSource:
// s := src.(*src.Xoshiro256ssSource)
// Note: state A
s1 := s.Jump() 
// Note: s1 is state A
// s is now state B (equivalent to 2^128 iterations of s.Uint64())
s2 := s.LongJump()
// Note: s2 is state B
// s is now state C (equivalent to 2^192 iterations of s.Uint64())
```

??? example "Xoroshiro256 implementation"

    `s0..s4` have `uint64` type.

    === "xoshiro256++"

        ```go
        func (x *Xoshiro256ppSource) Uint64() uint64 {
            res := bits.RotateLeft64(x.s0+x.s3, 23) + x.s0
            t := x.s1 << 17
            x.s2 ^= x.s0
            x.s3 ^= x.s1
            x.s1 ^= x.s2
            x.s0 ^= x.s3
            x.s2 ^= t
            x.s3 = bits.RotateLeft64(x.s3, 45)
            return res
        }
        ```

    === "xoshiro256**"

        ```go
        func (x *Xoshiro256ssSource) Uint64() uint64 {
            res := bits.RotateLeft64(x.s1*5, 7) * 9
            t := x.s1 << 17
            x.s2 ^= x.s0
            x.s3 ^= x.s1
            x.s1 ^= x.s2
            x.s0 ^= x.s3
            x.s2 ^= t
            x.s3 = bits.RotateLeft64(x.s3, 45)
            return res
        }
        ```

    === "Seed (same for both)"

        ```go
        sm := NewSplitMix64Source(seed)
        x.s0 = sm.Uint64()
        x.s1 = sm.Uint64()
        x.s2 = sm.Uint64()
        x.s3 = sm.Uint64()
        ```

## PCG-XSL-RR

!!! info inline end "Reference"

    * [PCG](https://www.pcg-random.org/)
    * [PCG-XSL-RR - 6.3.3, page 44](https://www.pcg-random.org/pdf/hmc-cs-2014-0905.pdf)
    * [Practical seed-recovery for the PCG...](https://tosc.iacr.org/index.php/ToSC/article/view/8700)

PCG (Permuted Congruential Generator) is a family of pseudorandom number generators (PRNGs) designed by Melissa O'Neill.
They are known for their speed, statistical quality, and flexibility, and can produce high-quality random numbers
with relatively small state sizes. PCG generators work by combining a linear congruential generator (LCG) with a
permutation function that shuffles the output, resulting in improved statistical properties.
Fluky provides PCG-XSL-RR implementation (128 bit state, 64 bit output), and they are suitable for variety of
applications, such as games, simulations (used in the NumPy scientific computing package).
PCG-XSL-RR is not suitable for cryptographic purposes.

**Usage**

```go
src := sources.NewPcgXslRrSource(0) // where 0 is a seed, int64
```

??? example "PCG-XSL-RR implementation"

    `state`, `mul` and `inc` have `uint128` type.

    ```go
    func (s *PcgXslRrSource) Uint64() uint64 {
        s.step()
        return s.stateToValue()
    }
    func (s *PcgXslRrSource) step() {
        s.state = s.state.Mul(s.mul).Add(s.inc)
    }
    func (s *PcgXslRrSource) stateToValue() (v uint64) {
        v = s.state.Lo ^ s.state.Hi
        rot := int(s.state.Rsh(122).Lo&0x3F) * -1
        v = bits.RotateLeft64(v, rot)
        return
    }
    ```

## Small PRNG

!!! info inline end "Reference"

    * [A small noncryptographic PRNG](https://burtleburtle.net/bob/rand/smallprng.html)
    * [Bob Jenkins's Small PRNG...](https://www.pcg-random.org/posts/bob-jenkins-small-prng-passes-practrand.html)

This fast and compact pseudorandom number generator is appropriate for large statistical computations,
but not intended for cryptographic purposes due to its lack of sufficient strength. The cycle length is not
guaranteed to meet a minimum requirement, but on average, it is estimated to produce approximately 2^126^ results.
Fluky provides Small PRNG implementation with 256 bit state and 64 bit output. Upon conducting multiple iterations
of the dieharder test suite, the Small PRNG exhibited just 0-2 weak results (of 114 tests).

**Usage**

```go
src := sources.NewSmallPrngSource(0) // where 0 is a seed, int64
```

??? example "Small PRNG implementation"

    `a`, `b`, `c` and `d` have `uint64` type.

    ```go
        func (s *SmallPrngSource) Uint64() uint64 {
        e := s.a - bits.RotateLeft64(s.b, 7)
        s.a = s.b ^ bits.RotateLeft64(s.c, 13)
        s.b = s.c + bits.RotateLeft64(s.d, 37)
        s.c = s.d + e
        s.d = e + s.a
        return s.d
    }
    ```

## Positional

### Squirrel3

!!! info inline end "Reference"

    * [Squirrel3](https://youtu.be/LWFzPP8ZbdU)
    * [Squirrel3 implementation 1](https://github.com/Descrout/squirrel3-rs)
    * [Squirrel3 implementation 2](https://github.com/archer884/squirrel-rng)

Squirrel3 is a pseudorandom number generator (PRNG) algorithm that generates a 64-bit output using a 64-bit seed and
position (64 bits).
Main idea is to use position and seed as input bits, and them mix them with a simple function.
This approach allow to make a set of different random number generators with the same statistical quality, but different
sequences.

There are 3 implementations of Squirrel3 in Fluky:

* `Squirrel3Source` - original implementation, with 64 bit seed and 64 bit position
* `Squirrel3x2Source` - implementation which use high bits of two random numbers combined to single 64 bits random
  number, iterate 2 positions at once
* `Squirrel3Prime64Source` - implementation with 64 bits prime number for noise bits, prime numbers generated randomly
  with ~ 59% probability of one bits

**Usage**

```go
src := sources.NewSquirrel3Source(0) // init with seed=0, position=0
src.SetPosition(10) // set position to 10
src.Uint64() // get 10th random number
src.Uint64() // get 11th random number
```

or

```go
src := sources.NewSquirrel3x2Source(0) // init with seed=0, position=0
```

or

```go
src := sources.Squirrel3Prime64Source(0) // init with seed=0, position=0
```

### Sxm

`SxmSource` and `SxmMixSource` are implementations of Positional source with inspiration from xxhash and Squirrel3.
`SxmMixSource` on final round used mix64 step from xxhash ([avalanche](https://github.com/Cyan4973/xxHash/blob/dev/doc/xxhash_spec.md#step-6-final-mix-avalanche-1)).

> The final mix ensures that all input bits have a chance to impact any bit in the output digest,
> resulting in an unbiased distribution. This is also called avalanche effect.

**Usage**

```go
src := sources.NewSxmSource(0) // init with seed=0, position=0
src.SetPosition(10) // set position to 10
src.Uint64() // get 10th random number
src.Uint64() // get 11th random number
```

or

```go
src := sources.NewSxmMixSource(0) // init with seed=0, position=0
```

### Wy family

`WySource` and `WyMixSource` are implementations based on constants and mix functions used in [wyhash](https://github.com/wangyi-fudan/wyhash) v4.2.
WyMixSource is positional source.

Comments from author of wyhash:

> The wyrand PRNG that pass BigCrush and PractRand (WySource)
> 
> a useful 64bit-64bit mix function to produce deterministic pseudo random numbers that can pass BigCrush and PractRand (WyMixSource)

**Usage**

```go
src := sources.NewWySource(0) // init with seed=0
src.Uint64() // get random number
```

or

```go
src := sources.NewWyMixSource(0) // init with seed=0, position=0
src.SetPosition(10) // set position to 10
src.Uint64() // get 10th random number
src.Uint64() // get 11th random number
```

## Summary

### Statistical quality (dieharder)

> SEED: 1234567
>
> Test data: 229GiB

Please check `dieharder` test results below.

| Name             |    Time    | PASS | WEAK | FAIL | Total |
|------------------|:----------:|:----:|:----:|:----:|:------|
| Built In         |   35:33    | 112  |  2   |  0   | 114   |
| PCG-XSL-RR       |   36:12    | 112  |  2   |  0   | 114   |
| Small Prng       |   37:49    | 114  |  0   |  0   | 114   | 
| Xoshiro256++     |   39:35    | 112  |  2   |  0   | 114   |                                  
| Xoshiro256**     |   37:01    | 108  |  6   |  0   | 114   |
| SplitMix64       |   38:06    | 113  |  1   |  0   | 114   |
| Squirrel3        | 19:50 (M2) | 108  |  6   |  0   | 114   |
| Squirrel3x2      | 19:50 (M2) | 111  |  3   |  0   | 114   |
| Squirrel3Prime64 | 19:10 (M2) | 113  |  1   |  0   | 114   |
| Sxm              | 19:10 (M2) | 113  |  1   |  0   | 114   |
| SxmMix           | 19:10 (M2) | 113  |  1   |  0   | 114   |
| Wy               | 21:00 (M2) | 112  |  2   |  0   | 114   |
| WyMix            | 21:30 (M2) | 111  |  3   |  0   | 114   |

!!! info "Reference"

    * Detailed results available in the repo's [`dieharder-source`](https://github.com/Pencroff/fluky/tree/main/dieharder-source) folder.
    * Built In source is a default Go source - [ref](https://pkg.go.dev/math/rand)

### Statistical quality (draw test)

??? abstract "Expand to check results (each picture is 2 Mb)"

    === "Built In"

        <figure markdown>
            ![Built In](//raw.githubusercontent.com/Pencroff/fluky/main/out/built-in_source_out.png){ loading=lazy width="768" }
        </figure>

    === "SplitMix64"

        <figure markdown>
            ![SplitMix64](//raw.githubusercontent.com/Pencroff/fluky/main/out/splitmix64_source_out.png){ loading=lazy width="768" }
        </figure>

    === "Xoshiro256++"

        <figure markdown>
            ![Xoshiro256++](//raw.githubusercontent.com/Pencroff/fluky/main/out/xoshiro256pp_source_out.png){ loading=lazy width="768" }
        </figure>

    === "Xoshiro256**"

        <figure markdown>
            ![Xoshiro256**](//raw.githubusercontent.com/Pencroff/fluky/main/out/xoshiro256ss_source_out.png){ loading=lazy width="768" }
        </figure>

    === "PCG-XSL-RR"

        <figure markdown>
            ![PCG-XSL-RR](//raw.githubusercontent.com/Pencroff/fluky/main/out/pcg-xsl-rr_source_out.png){ loading=lazy width="768" }
        </figure>

    === "Small Prng"

        <figure markdown>
            ![Small Prng](//raw.githubusercontent.com/Pencroff/fluky/main/out/small-prng_source_out.png){ loading=lazy width="768" }
        </figure>

    === "Squirrel3"
        
        <figure markdown>
            ![Squirrel3](//raw.githubusercontent.com/Pencroff/fluky/main/out/squirrel3_source_out.png){ loading=lazy width="768" }
        </figure>

    === "Squirrel3x2"

          <figure markdown>
              ![Squirrel3x2](//raw.githubusercontent.com/Pencroff/fluky/main/out/squirrel3x2_source_out.png){ loading=lazy width="768" }
          </figure>

    === "Squirrel3Prime64"

          <figure markdown>
              ![Squirrel3Prime64](//raw.githubusercontent.com/Pencroff/fluky/main/out/squirrel3_prime64_source_out.png){ loading=lazy width="768" }
          </figure>

    === "Sxm"

          <figure markdown>
              ![Sxm](//raw.githubusercontent.com/Pencroff/fluky/main/out/sxm_source_out.png){ loading=lazy width="768" }
          </figure>

    === "SxmMix"  
    
          <figure markdown>
              ![SxmMix](//raw.githubusercontent.com/Pencroff/fluky/main/out/sxmmix_source_out.png){ loading=lazy width="768" }
          </figure>

    === "Wy"

          <figure markdown>
              ![SxmMix](//raw.githubusercontent.com/Pencroff/fluky/main/out/wy_source_out.png){ loading=lazy width="768" }
          </figure>

    === "WyMix"  
    
          <figure markdown>
              ![SxmMix](//raw.githubusercontent.com/Pencroff/fluky/main/out/wymix_source_out.png){ loading=lazy width="768" }
          </figure>

### Performance

**Manual run**

```shell
go test -bench=. -timeout 30m ./fluky_testing
```

```shell
Ξ go test -bench=. -timeout 30m ./fluky_testing
goos: darwin
goarch: arm64
pkg: github.com/Pencroff/fluky/fluky_testing
cpu: Apple M2
BenchmarkSource/BuiltIn_Source-8                512898049                2.182 ns/op
BenchmarkSource/Pcg_Source-8                    294336325                4.054 ns/op
BenchmarkSource/SmallPrng_Source-8              440967740                2.743 ns/op
BenchmarkSource/xoshiro256++_Source-8           417001135                2.866 ns/op
BenchmarkSource/xoshiro256**_Source-8           413702812                2.784 ns/op
BenchmarkSource/SplitMix64-8                    590126202                2.019 ns/op
BenchmarkSource/Squirrel3-8                     582244454                2.075 ns/op
BenchmarkSource/Squirrel3x2-8                   510214744                2.356 ns/op
BenchmarkSource/Squirrel3Prime64-8              592099540                2.012 ns/op
BenchmarkSource/SxmSource-8                     535098463                2.277 ns/op
BenchmarkSource/SxmMixSource-8                  354509968                3.411 ns/op
BenchmarkSource/WySource-8                      597069507                2.004 ns/op
BenchmarkSource/WyMixSource-8                   595322499                1.999 ns/op
PASS
ok      github.com/Pencroff/fluky/fluky_testing 19.137s
```

{==
TBW - requires tests on Intel, M2, with latest Go version
==}

