package rng

import (
	"fmt"
	"testing"
)

func TestXoshiro256PP_Seed(t *testing.T) {
	type args struct {
		seed int64
	}
	tbl := []struct {
		name string
		args args
	}{
		{"test-1", args{seed: 10}},
		{"test-2", args{seed: 0}},
		{"test-3", args{seed: 11111}},
	}
	for _, tt := range tbl {
		t.Run(tt.name, func(t *testing.T) {
			g := NewXoshiro256ppC()
			g.Seed(tt.args.seed)
			fmt.Printf("seed: %d\n", tt.args.seed)
			fmt.Printf("s0: 0x%x, s1: 0x%x, s2: 0x%x, s3: 0x%x\n", g.st.s0, g.st.s1, g.st.s2, g.st.s3)
		})
	}
}
