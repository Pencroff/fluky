package rng

/*
#include <stdio.h>

typedef unsigned long long u8;
typedef struct ranctx { u8 a; u8 b; u8 c; u8 d; } ranctx;

#define rot(x,k) (((x)<<(k))|((x)>>(64-(k))))
void print_ranctx(char *pref, ranctx *x ) {
	printf("%s:\n", pref);
	printf("a: %016llx\tb: %016llx\t c: %016llx\t d: %016llx\n", x->a, x->b, x->c, x->d);
}
void print_u8(char *pref, u8 v) {
	printf("%s: %016llx\n", pref, v);
}

u8 ranval( ranctx *x ) {
	u8 e = x->a - rot(x->b, 7);
	//print_u8("e", e);
	x->a = x->b ^ rot(x->c, 13);
	//print_ranctx("step 1", x);
	x->b = x->c + rot(x->d, 37);
	//print_ranctx("step 2", x);
	x->c = x->d + e;
	//print_ranctx("step 3", x);
	x->d = e + x->a;
	//print_ranctx("step 4", x);
	//printf("-------\n");
	//print_u8("res", x->d);
	//printf("\n");
	return x->d;
}

void raninit( ranctx *x, u8 seed ) {
	u8 i;
	//print_u8("seed", seed);
	x->a = 0xf1ea5eed, x->b = x->c = x->d = seed;
	for (i=0; i<20; ++i) {
		(void)ranval(x);
	}
	//print_ranctx("state", x);
	//printf("====================\n\n\n");
}

*/
import "C"

type SmallPrngC struct {
	v *C.ranctx
}

func (g *SmallPrngC) Seed(seed int64) {
	C.raninit(g.v, C.u8(seed))
}

func (g *SmallPrngC) Uint64() uint64 {
	return uint64(C.ranval(g.v))
}

func (g *SmallPrngC) Int63() int64 {
	return int64(g.Uint64() >> 1)
}

func (g *SmallPrngC) Float64() float64 {
	rnd := g.Uint64() >> (uint64Bits - precisionBits)
	var res float64
	if rnd == maxDoublePrecision {
		rnd -= 1
	}
	res = float64(rnd) / maxDoublePrecision
	return res
}

func NewSmallPrngC() *SmallPrngC {
	seed := uint64(11111)
	n := &(C.ranctx{})
	C.raninit(n, C.u8(seed))
	r := &SmallPrngC{
		v: n,
	}
	return r
}
