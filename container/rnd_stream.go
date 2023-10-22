package main

import (
	"flag"
	"fmt"
	rng "github.com/Pencroff/fluky/source"
	"log"
	"math/rand"
	"os"
	"reflect"
	"time"
)

const notePeriodSec = 10

func main() {
	numbPtr := flag.Int64("seed", 0, "seed value")
	flag.Parse()
	butchSize := 8 * 1024
	bff := make([]byte, butchSize)

	src := rng.NewSquirrel3x2Source(*numbPtr)

	os.Stderr.WriteString(fmt.Sprintf("Seed: %d\n", *numbPtr))
	os.Stderr.WriteString(fmt.Sprintf("Source: %s\n", reflect.TypeOf(src).String()))

	start := time.Now()
	tm := start.Add(time.Second * notePeriodSec)

	rnd := rand.New(src)

	sizeCnt := 0

	for true {
		rnd.Read(bff)
		_, err := os.Stdout.Write(bff)
		sizeCnt += butchSize
		if err != nil {
			log.Fatal(err)
		}
		n := time.Now()
		if tm.After(n) {
			continue
		}
		tm = tm.Add(time.Second * notePeriodSec)
		fSize := float64(sizeCnt) / (1024 * 1024 * 1024)
		since := n.Sub(start).Round(time.Second)
		speed := float64(sizeCnt) / (1024 * 1024) / since.Seconds()
		os.Stderr.WriteString(fmt.Sprintf("%7s:%7.2f Gb (%.2f Mb/s)\n", since, fSize, speed))
	}
}
