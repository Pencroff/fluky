# fluky

![fluky](./assets/fluky-logo.min.png)


Happening by or depending on chance in Golang

All rng implements rand.Source64 interface
All passed dieharder tests located in source folder

## Dieharder source summary

    SEED: 1234567
    Test data: 229GiB

| Name         | Time  | PASS | WEAK | FAIL | Total | Draw test                              | References                                                                                                                                          |
|--------------|:-----:|:----:|:----:|:----:|:------|----------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------|
| Built In     | 35:44 | 112  |  2   |  0   | 114   | [img](out/built-in_source_out.png)     | [Ref](https://pkg.go.dev/math/rand)                                                                                                                 |
| Pcg64        | 34:34 | 113  |  1   |  0   | 114   | [img](out/pcg_source_out.png)          | [Ref](https://www.pcg-random.org/)                                                                                                                  |
| Small Prng   | 37:49 | 113  |  1   |  0   | 114   | [img](out/small-prng_source_out.png)   | [Ref1](https://burtleburtle.net/bob/rand/smallprng.html),<br/>[Ref2](https://www.pcg-random.org/posts/bob-jenkins-small-prng-passes-practrand.html) | 
| Xoshiro256++ | 39:35 | 112  |  2   |  0   | 114   | [img](out/xoshiro256pp_source_out.png) | [Ref](https://prng.di.unimi.it/)                                                                                                                    |                                  
| Xoshiro256** | 37:01 | 108  |  6   |  0   | 114   | [img](out/xoshiro256ss_source_out.png) | [Ref](https://prng.di.unimi.it/)                                                                                                                    |
| SplitMix64   | 38:06 | 113  |  1   |  0   | 114   | [img](out/splitmix64_source_out.png)   | [Ref](https://prng.di.unimi.it/)                                                                                                                    |

Detailed results please check in the [`dieharder-result`](experiments/dieharder-result) directory.

Few more RNG and test results can be found in [rngset repo](https://github.com/TyeolRik/rngset).

## Benchmark

    goos: windows
    goarch: amd64
    pkg: github.com/Pencroff/fluky/rng
    cpu: Intel(R) Core(TM) i9-8950HK CPU @ 2.90GHz
    BenchmarkRndGen 

| Name             |    Ops    | Execution time |
|------------------|:---------:|:--------------:|
| BuiltIn          | 243794713 |  4.721 ns/op   |
| PcgCRng with CGO | 18878521  |  63.00 ns/op   |
| PcgRng pure GO   | 205031472 |  5.755 ns/op   |
| Small Prng       | 518096238 |  2.327 ns/op   |
| Squares          | 146871903 |  8.507 ns/op   |


## Execute dieharder tests

Check the [**README.md**](./container/README.md).

### How to increase version

* commit all required changes
* git tag \<version - v0.0.2>
* git push origin --tags
* done - check docs on [pkg.go.dev](https://pkg.go.dev/github.com/Pencroff/fluky)
* install by `go get -u github.com/Pencroff/fluky`

## References

* [List of random number generators](https://en.wikipedia.org/wiki/List_of_random_number_generators)
* [Dieharder test](https://webhome.phy.duke.edu/~rgb/General/dieharder.php)
* [NIST Statistical Test Suite](https://csrc.nist.gov/Projects/Random-Bit-Generation/Documentation-and-Software)
* [NIST Statistical Test Suite implementation (Python)](https://github.com/GINARTeam/NIST-statistical-test)
* Similar projects - [ref1](https://github.com/skeeto/rng-go), [ref2](https://github.com/TyeolRik/rngset)
* [Optional Parameters in Go](https://petomalina.medium.com/dealing-with-optional-parameters-in-go-9780f9bfbd1d)
