
# Fluky API

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
	flk := f.NewFluky(r)                    // create fluky instance
	// ready to use
	fmt.Println(flk.Bool()) // true
	fmt.Println(flk.Integer()) // 3203168211198807973
	// with parameters
	fmt.Println(flk.Bool(f.WithLikelihood(0.25))) // 25% chance to return true
	fmt.Println(flk.Integer(f.WithIntRange(0, 100))) // 0..99, 100 is excluded
}
```

All Fluky methods accept set of optional parameters. They are used for result customization of the method.

## Sections

* [Basic](basic.md)
* [Helpers](helpers.md)
