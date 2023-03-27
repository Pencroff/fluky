# Introduction

<figure markdown>
![Fluky](images/fluky-color.min.png){ width=256px }
</figure>


Happening by or depending on chance in Golang.

[![Go Reference](https://pkg.go.dev/badge/github.com/Pencroff/fluky.svg)](https://pkg.go.dev/github.com/Pencroff/fluky)
[![Go Report Card](https://goreportcard.com/badge/github.com/Pencroff/fluky)](https://goreportcard.com/report/github.com/Pencroff/fluky)
[![License](https://img.shields.io/github/license/Pencroff/fluky)](https://github.com/Pencroff/fluky/blob/main/LICENSE)

## Overview

Fluky Library offers a wide range of methods to generate various types of random values.
It relies on the standard `rand.Rand` located in `math/rand` and includes several random number generators such as
[pcg](https://www.pcg-random.org/), [small prng](https://burtleburtle.net/bob/rand/smallprng.html),
[splitmix64](https://prng.di.unimi.it/), [xoshiro256++](https://prng.di.unimi.it/), and [xoshiro256**](https://prng.di.unimi.it/). 
These generators have been evaluated for their quality using the [`dieharder`](https://webhome.phy.duke.edu/~rgb/General/dieharder.php) test suite.

Please check dieharder results in [repo](https://github.com/Pencroff/fluky/tree/main/dieharder-source).

## Install

```shell
go get -u github.com/Pencroff/fluky
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