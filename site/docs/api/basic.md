# Basic

## New Fluky

```go
package main

import (
	"fmt"
	f "github.com/Pencroff/fluky"
	src "github.com/Pencroff/fluky/source"
	"math/rand"
)

func main() {
	s := src.NewSplitMix64Source(1234567)   // create source
	r := rand.New(s)                        // create rand.Rand instance
	flk := f.NewFluky(r)                // create fluky instance
	// ready to use
	fmt.Println(flk.Bool(f.WithLikelihood(0.25)))
	fmt.Println(flk.Bool())
}
```

All Fluky methods accept set of optional parameters. They are used for result customization of the method.

## Bool

By default, `Bool` returns a random bool value with equal probability.

```go
flk.Bool()
```

### Parameters

#### Likelihood

`WithLikelihood` parameter allows to set the likelihood of the `True` result.

```go
flk.Bool(f.WithLikelihood(0.25)) // 25% chance to return true
```


