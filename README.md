# fluky

Happening by or depending on chance in Golang

All tests rng provides uint64 values (1.845e19)

## Stats

TestFrequencyMonobits stats (2,684,354,560 numbers) ~ 20 Gb of random bytes (66.96s)

|Name| 1 - (ones / zeros) - should be close to 0| Values |
| --- | --- | --- |
|             mmix | 1.858732e-06 | [ 1: 85899266088 ; 0: 85899425752 ] |
|             Musl | 8.780273e-07 | [ 1: 85899308209 ; 0: 85899383631 ] |
|         built-in | 5.307109e-06 | [ 1: 85899573858 ; 0: 85899117982 ] |
|       small_prng | 2.628172e-06 | [ 1: 85899458799 ; 0: 85899233041 ] |

The winner: Musl (8.780273e-07)

## References

* [List of random number generators](https://en.wikipedia.org/wiki/List_of_random_number_generators)