# Welcome to Fluky documentation

<img
class="logo"
src="images/fluky-v.svg"
alt="Fluky">
</img>

Happening by or depending on chance in Golang.

## Overview

Fluky Library offers a wide range of methods to generate various types of random values.
It relies on the standard `rand.Rand` located in `math/rand` and includes several random number generators such as
[pcg](https://www.pcg-random.org/), [small prng](https://burtleburtle.net/bob/rand/smallprng.html),
[splitmix64](https://prng.di.unimi.it/), [xoshiro256++](https://prng.di.unimi.it/), and [xoshiro256**](https://prng.di.unimi.it/). 
These generators have been evaluated for their quality using the [`dieharder`](https://webhome.phy.duke.edu/~rgb/General/dieharder.php) test suite.

Please check dieharder results in [repo](https://github.com/Pencroff/fluky/tree/main/dieharder-source).

## Install

```bash

go get github.com/Pencroff/fluky -u

```

## Usage example

```go
package main

import (
	"fmt"
	"github.com/Pencroff/fluky"
	src "github.com/Pencroff/fluky/source"
	"math/rand"
)

func main() {
	s := src.NewPcgSource(0)
	r := rand.New(s)
	flk := fluky.NewFluky(r)

	fmt.Println(flk.Uint64())
}

```

