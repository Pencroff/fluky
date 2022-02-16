package fluky

//inline static uint64_t squares64(uint64_t ctr, uint64_t key) {
//
//t, x, y, z ui
//
//y = x = ctr * key; z = y + key;
//
//x = x*x + y; x = (x>>31) | (x<<32);       /* round 1 */
//
//x = x*x + z; x = (x>>31) | (x<<32);       /* round 2 */
//
//x = x*x + y; x = (x>>31) | (x<<32);       /* round 3 */
//
//t = x = x*x + z; x = (x>>31) | (x<<32);   /* round 4 */
//
//return t ^ ((x*x + y) >> 32);             /* round 5 */
//
//}

type Cbms struct {
	modulus      uint64
	modulusFloat float64
	multiplier   uint64
	increment    uint64
	seed         uint64
}
