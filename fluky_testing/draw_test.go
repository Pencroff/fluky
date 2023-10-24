package fluky_testing

import (
	"fmt"
	"github.com/Pencroff/fluky/source"
	"github.com/Pencroff/go-toolkit/bitfield"
	"github.com/fogleman/gg"
	"math/rand"
	"os"
	"testing"
)

func Test_randomness_draw_16bit_32bit_64bit(t *testing.T) {
	seed := int64(1234567)
	tbl := []struct {
		name    string
		size    int
		src     rand.Source
		extract string
	}{
		{
			name:    "built-in_source",
			size:    768,
			src:     rand.NewSource(seed),
			extract: "64bit",
		}, {
			name:    "small-prng_source",
			size:    768,
			src:     source.NewSmallPrngSource(seed),
			extract: "64bit",
		}, {
			name:    "pcg-xsl-rr_source",
			size:    768,
			src:     source.NewPcgXslRrSource(seed),
			extract: "64bit",
		}, {
			name:    "xoshiro256pp_source",
			size:    768,
			src:     source.NewXoshiro256ppSource(seed),
			extract: "64bit",
		}, {
			name:    "xoshiro256ss_source",
			size:    768,
			src:     source.NewXoshiro256ssSource(seed),
			extract: "64bit",
		}, {
			name:    "splitmix64_source",
			size:    768,
			src:     source.NewSplitMix64Source(seed),
			extract: "64bit",
		}, {
			name:    "squirrel3_source",
			size:    768,
			src:     source.NewSquirrel3Source(seed),
			extract: "64bit",
		}, {
			name:    "squirrel3x2_source",
			size:    768,
			src:     source.NewSquirrel3x2Source(seed),
			extract: "64bit",
		}, {
			name:    "squirrel3_big_prime_source",
			size:    768,
			src:     source.NewSquirrel3BigPrimeSource(seed),
			extract: "64bit",
		}, {
			name:    "noise_source",
			size:    768,
			src:     source.NewNoiseSource(seed),
			extract: "64bit",
		},
	}
	for _, el := range tbl {
		outPath := "../out/" + el.name + "_out.png"
		if _, err := os.Stat(outPath); err == nil {
			continue
		}
		size := el.size
		n := uint64(size*size) * 8
		dc := gg.NewContext(size, size)
		var i uint64
		r := el.src.(rand.Source64)
		for i = 0; i <= n; i += 1 {
			v := r.Uint64()
			switch el.extract {
			case "16bit":
				x1, y1, c := extractZx81(v)
				dc.SetRGB255(int(c), int(c), int(c))
				dc.SetPixel(int(x1), int(y1))
			case "32bit":
				x1, y1, c := extract32(v)
				dc.SetRGB255(int(c), int(c), int(c))
				dc.SetPixel(int(x1), int(y1))
			case "48bit":
				x1, y1, r, g, b := extract48(v)
				dc.SetRGB255(int(r), int(g), int(b))
				dc.SetPixel(int(x1), int(y1))
			case "64bit":
				x1, y1, r, g, b := extract64(v)
				dc.SetRGB255(int(r), int(g), int(b))
				dc.SetPixel(int(x1), int(y1))
			}
		}

		err := dc.SavePNG(outPath)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func extractZx81(v uint64) (x1, y1, c uint64) {
	x1 = bitfield.Extract(v, 9, 7)
	y1 = bitfield.Extract(v, 2, 7)
	c = bitfield.Extract(v, 0, 2) * 64
	return
}

func extract32(v uint64) (x1, y1, c uint64) {
	x1 = bitfield.Extract(v, 21, 10)
	y1 = bitfield.Extract(v, 11, 10)
	c = bitfield.Extract(v, 3, 8)
	return
}

func extract48(v uint64) (x1, y1, r, g, b uint64) {
	x1 = bitfield.Extract(v, 38, 10)
	y1 = bitfield.Extract(v, 28, 10)
	r = bitfield.Extract(v, 22, 6) * 4
	g = bitfield.Extract(v, 16, 6) * 4
	b = bitfield.Extract(v, 10, 6) * 4
	return
}

func extract64(v uint64) (x1, y1, r, g, b uint64) {
	x1 = bitfield.Extract(v, 54, 10)
	y1 = bitfield.Extract(v, 44, 10)
	r = bitfield.Extract(v, 36, 8)
	g = bitfield.Extract(v, 28, 8)
	b = bitfield.Extract(v, 20, 8)
	return
}
