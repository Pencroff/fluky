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

## Summary

### Statistical quality (dieharder)

> SEED: 1234567
> 
> Test data: 229GiB

| Name         | Time  | PASS | WEAK | FAIL | Total |
|--------------|:-----:|:----:|:----:|:----:|:------|
| Built In     | 35:33 | 112  |  2   |  0   | 114   |
| PCG-XSL-RR   | 36:12 | 112  |  2   |  0   | 114   |
| Small Prng   | 37:49 | 114  |  0   |  0   | 114   | 
| Xoshiro256++ | 39:35 | 112  |  2   |  0   | 114   |                                  
| Xoshiro256** | 37:01 | 108  |  6   |  0   | 114   |
| SplitMix64   | 38:06 | 113  |  1   |  0   | 114   |


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

### Performance

**Manual run**

```shell
go test -bench=. -timeout 30m ./fluky_testing
```

{==
TBW - requires tests on Intel, M2, with latest Go version
==}

