package fluky_testing

import (
	"fmt"
	"github.com/Pencroff/fluky/rng"
	"github.com/Pencroff/go-toolkit/bitfield"
	"github.com/fogleman/gg"
	"math/rand"
	"os"
	"testing"
)

func Test_randomness_draw_16bit_32bit_64bit(t *testing.T) {
	tbl := []struct {
		name    string
		size    int
		rnd     rng.RandomGenerator
		extract string
	}{
		{
			name:    rng.ANSI_C,
			size:    1024,
			rnd:     rng.NewLcg(rng.ANSI_C),
			extract: "32bit",
		}, {
			name:    rng.Turbo_Pascal,
			size:    1024,
			rnd:     rng.NewLcg(rng.Turbo_Pascal),
			extract: "32bit",
		}, {
			name:    rng.Apple_CarbonLib,
			size:    1024,
			rnd:     rng.NewLcg(rng.Apple_CarbonLib),
			extract: "32bit",
		}, {
			name:    rng.Cplus_11,
			size:    1024,
			rnd:     rng.NewLcg(rng.Cplus_11),
			extract: "32bit",
		}, {
			name:    rng.Posix_rand48,
			size:    1024,
			rnd:     rng.NewLcg(rng.Posix_rand48),
			extract: "48bit",
		}, {
			name:    rng.MMIX,
			size:    1024,
			rnd:     rng.NewLcg(rng.MMIX),
			extract: "64bit",
		}, {
			name:    rng.Musl,
			size:    1024,
			rnd:     rng.NewLcg(rng.Musl),
			extract: "64bit",
		}, {
			name:    string(rng.Zx81),
			size:    128,
			rnd:     rng.NewLcg(rng.Zx81),
			extract: "16bit",
		}, {
			name:    "small_prng",
			size:    1024,
			rnd:     rng.NewSmallPrng(),
			extract: "64bit",
		}, {
			name:    "squares",
			size:    1024,
			rnd:     rng.NewSquares(),
			extract: "64bit",
		}, {
			name:    "pcg64",
			size:    1024,
			rnd:     rng.NewPcgRng(),
			extract: "64bit",
		}, {
			name:    "built-in",
			size:    1024,
			rnd:     rand.New(rand.NewSource(11111)),
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
		for i = 0; i <= n; i += 1 {
			v := el.rnd.Uint64()
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
