package main

import (
	"github.com/Pencroff/fluky/rng"
	"log"
	"math/rand"
	"os"
)

func main() {
	butchSize := 8 * 1024
	bff := make([]byte, butchSize)
	src := rng.NewXoshiro256pp(11111)
	rnd := rand.New(src)
	for true {
		rnd.Read(bff)
		_, err := os.Stdout.Write(bff)
		if err != nil {
			log.Fatal(err)
		}
	}
}
