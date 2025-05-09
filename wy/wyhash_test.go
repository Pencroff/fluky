package wy

import (
	"encoding/binary"
	"testing"
)

/*
test_vector v4.2 wyhash_cpp

```bash
Ξ ./wyhash_cpp/test_vector
# wyhash         message,  seed
93228a4de0eec5a2 "",  0
c5bac3db178713c4 "a",  1
a97f2f7b1d9b3314 "abc",  2
d3fa000fdb422cf "1234",  3
160efc6277fcec80 "12345",  4
b2205dc629f3ef52 "123456",  5
fcbd904b20d5022c "1234567",  6
b4d64b1fc5efc3f8 "12345678",  7
84c12a0bfe6d670b "123456789",  8
63f35c1f12d4c922 "123456789A",  9
4a7b5d200a5860b2 "123456789AB",  a
313b0ac37c8cc1ce "123456789ABC",  b
73a21fc7838e0aec "123456789ABCD",  c
4fd0b1e18c59f85d "123456789ABCDE",  d
b3d88614f07f1708 "123456789ABCDEF",  e
d0f623b2f2c1015f "123456789ABCDEFG",  f
e580b70f3cb71bf0 "message digest",  10
425f943510e04fc3 "abcdefghijklmnopqrstuvwxyz",  11
70c9a2236889e825 "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789",  12
b650386c8c24729e "12345678901234567890123456789012345678901234567890123456789012345678901234567890",  13
```
*/

func TestSum64(t *testing.T) {
	tests := []struct {
		s        string
		seed     uint64
		expected uint64
	}{
		{"", 0, 0x93228a4de0eec5a2},
		{"a", 1, 0xc5bac3db178713c4},
		{"abc", 2, 0xa97f2f7b1d9b3314},
		{"1234", 3, 0xd3fa000fdb422cf},
		{"12345", 4, 0x160efc6277fcec80},
		{"123456", 5, 0xb2205dc629f3ef52},
		{"1234567", 6, 0xfcbd904b20d5022c},
		{"12345678", 7, 0xb4d64b1fc5efc3f8},
		{"123456789", 8, 0x84c12a0bfe6d670b},
		{"123456789A", 9, 0x63f35c1f12d4c922},
		{"123456789AB", 0xa, 0x4a7b5d200a5860b2},
		{"123456789ABC", 0xb, 0x313b0ac37c8cc1ce},
		{"123456789ABCD", 0xc, 0x73a21fc7838e0aec},
		{"123456789ABCDE", 0xd, 0x4fd0b1e18c59f85d},
		{"123456789ABCDEF", 0xe, 0xb3d88614f07f1708},
		{"123456789ABCDEFG", 0xf, 0xd0f623b2f2c1015f},
		{"message digest", 0x10, 0xe580b70f3cb71bf0},
		{"Studico", 0x28032025, 0x263839fe116b8681},
		{"abcdefghijklmnopqrstuvwxyz", 0x11, 0x425f943510e04fc3},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789", 0x12, 0x70c9a2236889e825},
		{"12345678901234567890123456789012345678901234567890123456789012345678901234567890", 0x13, 0xb650386c8c24729e},
	}
	for idx, test := range tests {
		if got := Sum64(test.seed, []byte(test.s)); got != test.expected {
			t.Errorf("%x: unexpected digest.\nexpected: %x, but got: %x", idx, test.expected, got)
		}
	}
}

var buf = make([]byte, 8192)

func benchmarkSize(b *testing.B, size int) {
	b.SetBytes(int64(size))
	sum := make([]byte, 8)
	for i := 0; i < b.N; i++ {
		binary.LittleEndian.PutUint64(sum, Sum64(0, buf[:size]))
	}
}

// 410189038			2.829 ns/op			1060.34 MB/s - v4
// 455456869			2.493 ns/op			1203.49 MB/s - v4.2
func BenchmarkHash3Bytes(b *testing.B) {
	benchmarkSize(b, 3)
}

// 403622229			2.931 ns/op			2729.36 MB/s - v4
// 381856410			3.139 ns/op			2548.89 MB/s - v4.2
func BenchmarkHash8Bytes(b *testing.B) {
	benchmarkSize(b, 8)
}

// 63395768			16.68 ns/op			19179.43 MB/s - v4
// 68088165			16.87 ns/op			18967.92 MB/s - v4.2
func BenchmarkHash320Bytes(b *testing.B) {
	benchmarkSize(b, 320)
}

// 23834745			48.61 ns/op			21065.33 MB/s - v4
// 25285848			47.63 ns/op			21497.81 MB/s - v4.2
func BenchmarkHash1K(b *testing.B) {
	benchmarkSize(b, 1024)
}

// 3132324			366.3 ns/op			22361.95 MB/s - v4
// 3283228			366.5 ns/op			22350.06 MB/s - v4.2
func BenchmarkHash8K(b *testing.B) {
	benchmarkSize(b, 8192)
}
