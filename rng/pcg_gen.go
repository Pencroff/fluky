package rng

//pcg64_srandom_r(&rng, 42u, 54u);
//
//pcg64_random_r(&rng)
//
//#define pcg64_srandom_r                 pcg_setseq_128_srandom_r
//
//struct pcg_state_128 {
//    pcg128_t state;
//};
//
//inline void pcg_setseq_128_srandom_r(struct pcg_state_setseq_128* rng,
//                                     pcg128_t initstate, pcg128_t initseq)
//{
//    rng->state = 0U;
//    rng->inc = (initseq << 1u) | 1u;
//    pcg_setseq_128_step_r(rng);
//    rng->state += initstate;
//    pcg_setseq_128_step_r(rng);
//}
//#endif
//
//#if PCG_HAS_128BIT_OPS
//inline uint64_t
//pcg_setseq_128_xsl_rr_64_random_r(struct pcg_state_setseq_128* rng)
//{
//    pcg_setseq_128_step_r(rng);
//    return pcg_output_xsl_rr_128_64(rng->state);
//}
//#endif
// inline void pcg_setseq_128_step_r(struct pcg_state_setseq_128* rng)
//{
//    rng->state = rng->state * PCG_DEFAULT_MULTIPLIER_128 + rng->inc;
//}
