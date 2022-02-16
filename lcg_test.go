package fluky

import (
	"fmt"
	"github.com/Pencroff/go-toolkit/bitfield"
	"github.com/fogleman/gg"
	"math/rand"
	"testing"
)

func Test_randomness_32bit_64bit(t *testing.T) {
	tbl := []struct {
		size    int
		mode    LcgMode
		extract string
	}{
		{
			size:    1024,
			mode:    ANSI_C,
			extract: "32bit",
		}, {
			size:    1024,
			mode:    Turbo_Pascal,
			extract: "32bit",
		}, {
			size:    1024,
			mode:    Apple_CarbonLib,
			extract: "32bit",
		}, {
			size:    1024,
			mode:    Cplus_11,
			extract: "32bit",
		}, {
			size:    1024,
			mode:    Posix_rand48,
			extract: "48bit",
		}, {
			size:    1024,
			mode:    MMIX,
			extract: "64bit",
		}, {
			size:    1024,
			mode:    musl,
			extract: "64bit",
		},
	}
	for _, el := range tbl {
		size := el.size
		n := uint64(size*size) * 8
		lcg := NewLcg(el.mode)
		dc := gg.NewContext(size, size)
		var i uint64
		for i = 0; i <= n; i += 1 {
			v := lcg.Uint64()
			switch el.extract {
			case "32bit":
				x1, y1, c := extract32G(v)
				dc.SetRGB255(int(c), int(c), int(c))
				dc.SetPixel(int(x1), int(y1))
			case "48bit":
				x1, y1, r, g, b := extract48G(v)
				dc.SetRGB255(int(r), int(g), int(b))
				dc.SetPixel(int(x1), int(y1))
			case "64bit":
				x1, y1, r, g, b := extract64G(v)
				dc.SetRGB255(int(r), int(g), int(b))
				dc.SetPixel(int(x1), int(y1))
			}
		}
		outPath := string("out/" + el.mode + "_out.png")
		err := dc.SavePNG(outPath)
		if err != nil {
			fmt.Println(err)
		}
	}

}
func Test_randomness_zx81(t *testing.T) {
	size := 2 << 6
	n := uint64(size * size * size)
	lcg := NewLcg(Zx81)
	dc := gg.NewContext(size, size)
	var i uint64
	for i = 0; i <= n; i += 1 {
		v := lcg.Uint64()
		x1, y1, c := extractZx81G(v)
		dc.SetRGB255(int(c), int(c), int(c))
		dc.SetPixel(int(x1), int(y1))
	}
	outPath := string("out/" + Zx81 + "_out.png")
	err := dc.SavePNG(outPath)
	if err != nil {
		fmt.Println(err)
	}
}
func Test_randomness_built_in(t *testing.T) {
	size := 1024
	n := uint64(size*size) * 8
	rand.Seed(11111)
	dc := gg.NewContext(size, size)
	var i uint64
	for i = 0; i <= n; i += 1 {
		v := rand.Uint64()
		x1, y1, r, g, b := extract64G(v)
		dc.SetRGB255(int(r), int(g), int(b))
		dc.SetPixel(int(x1), int(y1))
	}
	outPath := "out/built-in_out.png"
	err := dc.SavePNG(outPath)
	if err != nil {
		fmt.Println(err)
	}
}

//func extract32(v uint64) (x1, y1, r, g, b uint64) {
//	x1 = bitfield.Extract(v, 0, 10)
//	y1 = bitfield.Extract(v, 10, 10)
//	r = bitfield.Extract(v, 20, 4) * 64
//	g = bitfield.Extract(v, 24, 4) * 64
//	b = bitfield.Extract(v, 28, 4) * 64
//	return
//}
func extract32G(v uint64) (x1, y1, c uint64) {
	x1 = bitfield.Extract(v, 21, 10)
	y1 = bitfield.Extract(v, 11, 10)
	c = bitfield.Extract(v, 3, 8)
	return
}

//func extract48(v uint64) (x1, y1, r, g, b uint64) {
//	x1 = bitfield.Extract(v, 0, 10)
//	y1 = bitfield.Extract(v, 10, 10)
//	r = bitfield.Extract(v, 20, 8)
//	g = bitfield.Extract(v, 28, 8)
//	b = bitfield.Extract(v, 36, 8)
//	return
//}
func extract48G(v uint64) (x1, y1, r, g, b uint64) {
	x1 = bitfield.Extract(v, 38, 10)
	y1 = bitfield.Extract(v, 28, 10)
	r = bitfield.Extract(v, 22, 6) * 4
	g = bitfield.Extract(v, 16, 6) * 4
	b = bitfield.Extract(v, 10, 6) * 4
	return
}

func extract64G(v uint64) (x1, y1, r, g, b uint64) {
	x1 = bitfield.Extract(v, 54, 10)
	y1 = bitfield.Extract(v, 44, 10)
	r = bitfield.Extract(v, 36, 8)
	g = bitfield.Extract(v, 28, 8)
	b = bitfield.Extract(v, 20, 8)
	return
}

//func extractZx81(v uint64) (x1, y1, c uint64) {
//	x1 = bitfield.Extract(v, 0, 7)
//	y1 = bitfield.Extract(v, 7, 7)
//	c = bitfield.Extract(v, 14, 2) * 64
//	return
//}

func extractZx81G(v uint64) (x1, y1, c uint64) {
	x1 = bitfield.Extract(v, 9, 7)
	y1 = bitfield.Extract(v, 2, 7)
	c = bitfield.Extract(v, 0, 2) * 64
	return
}
