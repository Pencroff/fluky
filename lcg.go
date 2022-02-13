package fluky

type Lcg struct {
	modulus      uint64
	modulusFloat float64
	multiplier   uint64
	increment    uint64
	seed         uint64
}

type LcgMode string

const (
	Zx81            LcgMode = "zx81"
	ANSI_C                  = "ansi_c"
	Turbo_Pascal            = "turbo_pascal"
	Apple_CarbonLib         = "minstd_rand0"
	Cplus_11                = "minstd_rand"
	Posix_rand48            = "posix_rand48"
	MMIX                    = "mmix"
	musl                    = "musl"
)

func NewLcg(mode LcgMode) *Lcg {
	var gen Lcg
	gen.seed = 11111
	switch mode {
	case Zx81:
		gen.modulus = 1<<16 + 1
		gen.multiplier = 75
		gen.increment = 74
	case ANSI_C:
		gen.modulus = 1 << 31
		gen.multiplier = 1103515245
		gen.increment = 12345
	case Turbo_Pascal:
		gen.modulus = 1 << 32
		gen.multiplier = 134775813
		gen.increment = 1
	case Apple_CarbonLib:
		gen.modulus = 1<<31 - 1
		gen.multiplier = 16807
		gen.increment = 0
	case Cplus_11:
		gen.modulus = 1<<31 - 1
		gen.multiplier = 48271
		gen.increment = 0
	case Posix_rand48:
		gen.modulus = 1 << 48
		gen.multiplier = 25214903917
		gen.increment = 11
	case MMIX:
		gen.modulus = 1<<64 - 1
		gen.multiplier = 6364136223846793005
		gen.increment = 1442695040888963407
	case musl:
		gen.modulus = 1<<64 - 1
		gen.multiplier = 6364136223846793005
		gen.increment = 1
	}
	gen.modulusFloat = float64(gen.modulus)
	return &gen
}

func (l *Lcg) Seed(v uint64) {
	l.seed = v
}

func (l *Lcg) NextUint64() uint64 {
	l.seed = l.multiplier*l.seed + l.increment
	return l.seed % l.modulus
}

func (l *Lcg) NextFloat64() float64 {
	rnd := l.NextUint64()
	return float64(rnd) / l.modulusFloat
}
