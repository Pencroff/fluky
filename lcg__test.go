package fluky

import (
	"fmt"
	"github.com/Pencroff/go-toolkit/bitfield"
	"github.com/fogleman/gg"
	"testing"
)

func Test_randomness_32bit(t *testing.T) {
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
			v := lcg.NextUint64()
			switch el.extract {
			case "32bit":
				x1, y1, r, g, b := extract32(v)
				dc.SetRGB255(int(r), int(g), int(b))
				dc.SetPixel(int(x1), int(y1))
			case "48bit":
				x1, y1, r, g, b := extract48(v)
				dc.SetRGB255(int(r), int(g), int(b))
				dc.SetPixel(int(x1), int(y1))
			case "64bit":
				x1, y1, r, g, b := extract48(v)
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
		v := lcg.NextUint64()
		x1, y1, c := extractZx81(v)
		dc.SetRGB255(int(c), int(c), int(c))
		dc.SetPixel(int(x1), int(y1))
	}
	outPath := string("out/" + Zx81 + "_out.png")
	fmt.Println(outPath)
	err := dc.SavePNG(outPath)
	if err != nil {
		fmt.Println(err)
	}
}

func extract32(v uint64) (x1, y1, r, g, b uint64) {
	x1 = bitfield.Extract(v, 0, 10)
	y1 = bitfield.Extract(v, 10, 10)
	r = bitfield.Extract(v, 20, 4) * 64
	g = bitfield.Extract(v, 24, 4) * 64
	b = bitfield.Extract(v, 28, 4) * 64
	return
}
func extract48(v uint64) (x1, y1, r, g, b uint64) {
	x1 = bitfield.Extract(v, 0, 10)
	y1 = bitfield.Extract(v, 10, 10)
	r = bitfield.Extract(v, 20, 8)
	g = bitfield.Extract(v, 28, 8)
	b = bitfield.Extract(v, 36, 8)
	return
}

func extractZx81(v uint64) (x1, y1, c uint64) {
	x1 = bitfield.Extract(v, 0, 7)
	y1 = bitfield.Extract(v, 7, 7)
	c = bitfield.Extract(v, 14, 2) * 64
	return
}
