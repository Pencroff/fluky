package main

import (
	"encoding/binary"
	"github.com/Pencroff/fluky/fluky_testing"
	"github.com/Pencroff/fluky/rng"
	"log"
	"os"
)

func main() {
	butchSize := int(8 * fluky_testing.Size1Kb)
	r := rng.NewPcgRng()
	smB := make([]byte, 8)
	b := make([]byte, butchSize)
	itterNum := butchSize / 8
	for true {
		for i := 0; i < itterNum; i++ {
			v := r.Uint64()
			binary.LittleEndian.PutUint64(smB, v)
			for j := 0; j < 8; j++ {
				idx := 8*i + j
				b[idx] = smB[j]
			}
		}
		_, err := os.Stdout.Write(b)
		if err != nil {
			log.Fatal(err)
		}
	}
}
