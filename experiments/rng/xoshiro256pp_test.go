package rng

import (
	"fmt"
	"testing"
)

func TestXoshiro256PP_Seed(t *testing.T) {
	tbl := []struct {
		name string
		seed int64
		res  []uint64
	}{
		{"test-1", 0,
			[]uint64{
				0xe9a74e1577b5678b,
				0xaf3662eb44a08fcf,
				0x5aaec5f0ecc7bfd5,
				0x745fb861c070ff4a,
				0x1333e932b3f7832d,
				0xd35acc9a143bf01a,
				0xc5fce182c7992cc3,
				0xfa22b5ff5c2bb011,
				0xe974ec941daecd52,
				0x75d9f781e329f119,
			}},
		{"test-2", 1,
			[]uint64{
				0x689d693084f90390,
				0xe7b847261e8aff51,
				0x516961827a899525,
				0xa3f422e983b07d2f,
				0x2a225042ce1215b0,
				0xee463b96321391a0,
				0xe2e46d91f2cb20ff,
				0xc3524d27b7dbddcd,
				0x7f3cd7be17d87880,
				0x2617515e4b920a38,
			}},
		{"test-3", 2,
			[]uint64{
				0x300c23cd60cf140,
				0x87f25ecdcaf80c8a,
				0xd876fae39144cc9d,
				0xb8388fa83ebd4496,
				0x7000e0b690a41c7a,
				0xe2f7ab77e7e311c2,
				0xaf6d7ec680aa86e6,
				0xcd15090a07e26e9e,
				0xbf6ae190af1586bc,
				0xabacd4a2fbd4af68,
			}},
		{"test-4", 10,
			[]uint64{
				0x61155b34dff54664,
				0x2ad7b6938c2e5ae5,
				0x9a1adaef9cb4a995,
				0xb3093cc0d2bd28b5,
				0x4b8b80ce5d278b7d,
				0xc44b44af7726240,
				0xe454ba340c167770,
				0x9ff85c9955f6fc74,
				0x19bf102a000664b,
				0xb24b9df26e055774,
			}},
		{"test-5", 11111,
			[]uint64{
				0x93959632eb345c6d,
				0x8d7f9f99f05aa7ef,
				0x3e5b4eae5fb84aab,
				0xf46d92732422af67,
				0xe76d9095a2d7c87a,
				0xb5f89fb9e8b73e68,
				0x1bfd35d69885b6a0,
				0x57cb34cd4a6f7384,
				0x5ceb399a53cdee7b,
				0x6e366785d5ede17,
			}},
	}
	for _, tt := range tbl {
		t.Run(tt.name, func(t *testing.T) {
			gc := NewXoshiro256ppC(tt.seed)
			g := NewXoshiro256pp(tt.seed)
			//fmt.Printf("seed: %d\n", tt.seed)
			//fmt.Printf("s0: 0x%x, s1: 0x%x, s2: 0x%x, s3: 0x%x\n", gc.st.s0, gc.st.s1, gc.st.s2, gc.st.s3)
			for i := 0; i < 10; i++ {
				rngC := gc.Uint64()
				rng := g.Uint64()
				//fmt.Printf("0x%x,\n", rngC)
				if rngC != tt.res[i] {
					t.Errorf("rngC: 0x%x, extected: 0x%x", rngC, tt.res[i])
				}
				if rng != tt.res[i] {
					t.Errorf("rng: 0x%x, extected: 0x%x", rng, tt.res[i])
				}
			}
		})
	}
}

func TestXoshiro256pp_Jump(t *testing.T) {
	tbl := []struct {
		name string
		seed int64
		res  []uint64
	}{
		{"test-1", 0,
			[]uint64{
				0x8ff9a13b029e8e5f,
				0x3d37a7a8dcbb4a2e,
				0x3796611974a6594b,
				0x1298ce3f3bef596a,
				0x96d0ccc23638e3e4,
			},
		}, {
			"test-2", 1,
			[]uint64{
				0x7c421212ccec7a5,
				0xa56d101945a85981,
				0x3484b69fbd46dfb,
				0x8e88cb479467098b,
				0x6314c51f1f728d9b,
			},
		},
	}
	for _, tt := range tbl {
		t.Run(tt.name, func(t *testing.T) {
			gc := NewXoshiro256ppC(tt.seed)
			g := NewXoshiro256pp(tt.seed)
			gc.Jump()
			g.Jump()
			//fmt.Printf("seed: %d\n", tt.seed)
			for i := 0; i < 5; i++ {
				rngC := gc.Uint64()
				rng := g.Uint64()
				//fmt.Printf("0x%x,\n", rngC)
				if rngC != tt.res[i] {
					t.Errorf("rngC: 0x%x, extected: 0x%x", rngC, tt.res[i])
				}
				if rng != tt.res[i] {
					t.Errorf("rng: 0x%x, extected: 0x%x", rng, tt.res[i])
				}
			}

		})
	}
}

func TestXoshiro256pp_LongJump(t *testing.T) {
	tbl := []struct {
		name string
		seed int64
		res  []uint64
	}{
		{"test-1", 0,
			[]uint64{
				0x2e4279bfb50300bd,
				0xa2369893868bcaa2,
				0xef7625229f28ac28,
				0xf6fbeb051b93cb78,
				0xce3f76fb02cff137,
			},
		}, {
			"test-2", 1,
			[]uint64{
				0xa4d26800129c600e,
				0x30208c9633b786fd,
				0x54ecd04a912bdf59,
				0x7a3ee891958395b8,
				0x15308c1ab41f8c53,
			},
		},
	}
	for _, tt := range tbl {
		t.Run(tt.name, func(t *testing.T) {
			gc := NewXoshiro256ppC(tt.seed)
			g := NewXoshiro256pp(tt.seed)
			gc.LongJump()
			g.LongJump()
			//fmt.Printf("seed: %d\n", tt.seed)
			for i := 0; i < 5; i++ {
				rngC := gc.Uint64()
				rng := g.Uint64()
				//fmt.Printf("0x%x,\n", rngC)
				if rngC != tt.res[i] {
					t.Errorf("rngC: 0x%x, extected: 0x%x", rngC, tt.res[i])
				}
				if rng != tt.res[i] {
					t.Errorf("rng: 0x%x, extected: 0x%x", rng, tt.res[i])
				}
			}

		})
	}
}

func TestXoshiro256ss_Seed(t *testing.T) {
	tbl := []struct {
		name string
		seed int64
		res  []uint64
	}{
		{"test-1", 0,
			[]uint64{
				0xfbddae7499cf0a04,
				0x373768bbec1d527b,
				0xb568706cabe9f460,
				0x8f09a4b107754efa,
				0xe52ccb64ef9f1daa,
			},
		}, {
			"test-2", 1,
			[]uint64{
				0xa7331f834908f3a,
				0x3bade10b38c2910c,
				0x14d4244a9c6c4b9c,
				0x9c1cf7ddfa2d010c,
				0xacb52e7925f3c4f8,
			},
		},
	}
	for _, tt := range tbl {
		t.Run(tt.name, func(t *testing.T) {
			g := NewXoshiro256ss(tt.seed)
			fmt.Printf("seed: %d\n", tt.seed)
			//fmt.Printf("s0: 0x%x, s1: 0x%x, s2: 0x%x, s3: 0x%x\n", gc.st.s0, gc.st.s1, gc.st.s2, gc.st.s3)
			for _, r := range tt.res {
				rngC := g.Uint64C()
				rng := g.Uint64()
				//fmt.Printf("0x%x,\n", rngC)
				if rngC != r {
					t.Errorf("rngC: 0x%x, extected: 0x%x", rngC, r)
				}
				if rng != r {
					t.Errorf("rng: 0x%x, extected: 0x%x", rng, r)
				}
			}
		})

	}
}
