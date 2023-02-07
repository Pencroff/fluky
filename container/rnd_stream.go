package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"reflect"
)

func main() {
	numbPtr := flag.Int64("seed", 0, "seed value")
	flag.Parse()
	butchSize := 8 * 1024
	bff := make([]byte, butchSize)

	src := rand.NewSource(*numbPtr)

	os.Stderr.WriteString(fmt.Sprintf("Seed: %d\n", *numbPtr))
	os.Stderr.WriteString(fmt.Sprintf("Source: %s\n", reflect.TypeOf(src).String()))

	rnd := rand.New(src)
	for true {
		rnd.Read(bff)
		_, err := os.Stdout.Write(bff)
		if err != nil {
			log.Fatal(err)
		}
	}
}
