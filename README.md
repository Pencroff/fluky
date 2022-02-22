# fluky

Happening by or depending on chance in Golang

All tests rng provides uint64 values (1.845e19)

## Docker compose

`docker-compose --file=./container/docker-compose.yml up -d dieharder`

## Nist tests

TestFrequencyMonobits stats (1342177280 numbers) ~ 10 Gb of random bytes each

|Name| P-value | Values | Status |
| --- | --- | --- |  --- |
|             mmix | 5.940685e-01 | [ 1: 42949594860 ; 0: 42949751060 ] | PASS |
|             Musl | 9.624453e-01 | [ 1: 42949679860 ; 0: 42949666060 ] | PASS |
|         built_in | 1.700484e-01 | [ 1: 42949874024 ; 0: 42949471896 ] | PASS |
|       small_prng | 4.221443e-01 | [ 1: 42949790591 ; 0: 42949555329 ] | PASS |

Max value: Musl (9.624453e-01)

--- PASS: TestFrequencyMonobits (40.95s)

## References

* [List of random number generators](https://en.wikipedia.org/wiki/List_of_random_number_generators)
* [Dieharder test](https://webhome.phy.duke.edu/~rgb/General/dieharder.php)
* [NIST Statistical Test Suite](https://csrc.nist.gov/Projects/Random-Bit-Generation/Documentation-and-Software)
* [NIST Statistical Test Suite implementation (Python)](https://github.com/GINARTeam/NIST-statistical-test)