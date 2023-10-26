# fluky

![fluky](./assets/fluky-color.min.svg)

Happening by or depending on chance in Golang

[![Go Reference](https://pkg.go.dev/badge/github.com/Pencroff/fluky.svg)](https://pkg.go.dev/github.com/Pencroff/fluky)
[![Docs Reference](https://img.shields.io/badge/-reference-007D9C?style=flat-square&logo=readthedocs&logoWidth=18&labelColor=5C5C5C&logoColor=FAFAFA)](https://pencroff.github.io/fluky/)
[![Go Report Card](https://goreportcard.com/badge/github.com/Pencroff/fluky)](https://goreportcard.com/report/github.com/Pencroff/fluky)
[![License](https://img.shields.io/github/license/Pencroff/fluky)](https://github.com/Pencroff/fluky/blob/main/LICENSE)



## Dieharder source summary

    SEED: 1234567
    Test data: 229GiB

| Name         | Performance<br/>MIN / AVG / MAX<br/>ns/op<br/>_smaller is better_ | Time  | PASS | WEAK | FAIL | Total | Draw test                              | References                                                                                                                                          |
|--------------|:-----------------------------------------------------------------:|:-----:|:----:|:----:|:----:|:------|----------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------|
| Built In     |                       3.183 / 3.245 / 3.307                       | 35:33 | 112  |  2   |  0   | 114   | [img](_out/built-in_source_out.png)     | [Ref](https://pkg.go.dev/math/rand)                                                                                                                 |
| PCG-XSL-RR   |                       5.337 / 5.861 / 6.279                       | 36:12 | 112  |  2   |  0   | 114   | [img](_out/pcg-xsl-rr_source_out.png)   | [Ref](https://www.pcg-random.org/)                                                                                                                  |
| Small Prng   |                     2.494 / **2.607** / 2.934                     | 37:49 | 114  |  0   |  0   | 114   | [img](_out/small-prng_source_out.png)   | [Ref1](https://burtleburtle.net/bob/rand/smallprng.html),<br/>[Ref2](https://www.pcg-random.org/posts/bob-jenkins-small-prng-passes-practrand.html) | 
| Xoshiro256++ |                       2.838 / 3.102 / 3.360                       | 39:35 | 112  |  2   |  0   | 114   | [img](_out/xoshiro256pp_source_out.png) | [Ref](https://prng.di.unimi.it/)                                                                                                                    |                                  
| Xoshiro256** |                       2.678 / 2.935 / 3.186                       | 37:01 | 108  |  6   |  0   | 114   | [img](_out/xoshiro256ss_source_out.png) | [Ref](https://prng.di.unimi.it/)                                                                                                                    |
| SplitMix64   |                       2.282 / 2.32 / 2.411                        | 38:06 | 113  |  1   |  0   | 114   | [img](_out/splitmix64_source_out.png)   | [Ref](https://prng.di.unimi.it/)                                                                                                                    |

Detailed results please check in the [`dieharder-source`](dieharder-source) folder.

Few more RNG and test results can be found in [rngset repo](https://github.com/TyeolRik/rngset).

## Benchmark

    go test -bench=. -timeout 30m ./fluky_testing

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
