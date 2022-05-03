package rng

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"unsafe"
)

type PcgRngData struct {
	state []byte
	inc   []byte
}

type PcgRngMock struct {
	v uint64
}

func (r *PcgRngMock) Seed(v int64) {
	r.v = uint64(v)
}

func (r *PcgRngMock) Uint64() uint64 {
	return r.v
}

func (r *PcgRngMock) Int63() int64 {
	return int64(r.Uint64() >> 1)
}

func (r *PcgRngMock) Float64() float64 {
	rnd := r.Int63()
	var res float64
	if rnd < 0x7ffffffffffffbff {
		res = float64(rnd) / (1 << 63)
	} else {
		res = float64(rnd-1024) / (1 << 63)
	}
	return res
}

func TestPcgRng(t *testing.T) {
	rng := NewPcgRng()
	rng.Seed(0x2a)
	lst := []uint64{0x86b1da1d72062b68, 0x1304aa46c9853d39, 0xa3670e9e0dd50358,
		0xf9090e529a7dae00, 0xc85b9fd837996f2c, 0x606121f8e3919196,
		0x7ce1c7ff478354ba, 0xcbc4ac70e541310e, 0x74be71999ec37f2c,
		0xb81f9c99a934f1a7, 0x120e9901a900c97f, 0x0f983bad4b19f493}
	for _, v := range lst {
		rnd := rng.Uint64()
		assert.Equal(t, v, rnd, "%x != %x", rnd, v)
	}
}

func TestPcgStateDataEquality(t *testing.T) {
	debugState := false
	pcgc := NewPcgCRng()
	cData := ExtractPcgCRng(pcgc, debugState)

	pcg := NewPcgRng()
	goData := ExtractPcgRng(pcg, debugState)
	assert.Equal(t, cData, goData)

	pcgcRnd := pcgc.Uint64()
	pcgRnd := pcg.Uint64()
	cDataRnd := ExtractPcgCRng(pcgc, debugState)
	goDataRnd := ExtractPcgRng(pcg, debugState)
	assert.Equal(t, cDataRnd, goDataRnd)
	assert.Equal(t, pcgcRnd, pcgRnd)

}

func TestPcgCRng(t *testing.T) {
	rng := NewPcgCRng()
	rng.Seed(0x2a)
	lst := []uint64{0x86b1da1d72062b68, 0x1304aa46c9853d39, 0xa3670e9e0dd50358,
		0xf9090e529a7dae00, 0xc85b9fd837996f2c, 0x606121f8e3919196,
		0x7ce1c7ff478354ba, 0xcbc4ac70e541310e, 0x74be71999ec37f2c,
		0xb81f9c99a934f1a7, 0x120e9901a900c97f, 0x0f983bad4b19f493}
	for _, v := range lst {
		rnd := rng.Uint64()
		assert.Equal(t, v, rnd, "%x != %x", rnd, v)
	}
}

func TestInt63Logic(t *testing.T) {
	rng := PcgRngMock{32}
	assert.Equal(t, int64(16), rng.Int63())
	rng = PcgRngMock{1<<64 - 1}
	assert.Equal(t, int64(1<<63-1), rng.Int63())
}

func TestFloat64NotReachOneLogic(t *testing.T) {
	rng := PcgRngMock{1<<64 - 1}
	assert.Equal(t, rng.Uint64(), uint64(0xFFFFFFFFFFFFFFFF))
	assert.Equal(t, rng.Int63(), int64(0x7FFFFFFFFFFFFFFF))
	v := rng.Float64()
	assert.NotEqual(t, 1, v)
	assert.Less(t, v, float64(1))
	assert.InDelta(t, v, float64(1), 1e-15)
}

func ExtractPcgCRng(r *PcgCRng, debug bool) *PcgRngData {
	if debug {
		fmt.Println("PcgCRng")
	}
	vRng := reflect.ValueOf(r).Elem().FieldByName("rng")
	vState := vRng.FieldByName("state")
	dataState := CUint128ToBytes(vState)
	if debug {
		fmt.Printf("state: %x\n", dataState)
	}
	//fmt.Printf("state: %v\n", dataState)
	vInc := vRng.FieldByName("inc")
	dataInc := CUint128ToBytes(vInc)
	if debug {
		fmt.Printf("inc:   %x\n", dataInc)
	}

	return &PcgRngData{
		state: dataState,
		inc:   dataInc,
	}
}

func ExtractPcgRng(r *PcgRng, debug bool) *PcgRngData {
	if debug {
		fmt.Println("PcgRng")
	}
	vRng := reflect.ValueOf(r).Elem()
	vState := vRng.FieldByName("state")
	vInc := vRng.FieldByName("inc")
	dataState := Uint128ToBytes(vState)
	if debug {
		fmt.Printf("state: %x\n", dataState)
	}
	dataInc := Uint128ToBytes(vInc)
	if debug {
		fmt.Printf("inc:   %x\n", dataInc)
	}
	return &PcgRngData{
		state: dataState,
		inc:   dataInc,
	}
}

func Uint128ToBytes(v reflect.Value) []byte {
	res := make([]byte, 16)
	if v.Type().Name() == "Uint128" {
		nv := reflect.NewAt(v.Type(), unsafe.Pointer(v.UnsafeAddr())).Elem()
		vptr := nv.Addr()
		method := vptr.MethodByName("ToBytes")
		bt := method.Call([]reflect.Value{})
		res = bt[0].Interface().([]byte)
	}
	return res
}

func CUint128ToBytes(v reflect.Value) []byte {
	res := make([]byte, 16)
	if v.Type().Name() == "_Ctype___int128unsigned" {
		bytePtr := (*[16]byte)(unsafe.Pointer(v.UnsafeAddr()))
		for i := 0; i < 16; i++ {
			res[i] = bytePtr[i]
		}
	}
	return res
}
