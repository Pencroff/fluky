package experiments

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"math/rand"
	"testing"
)

type MinMax struct {
	Min float64
	Max float64
}

func TestMinMax(t *testing.T) {
	size := int(1 * Size1Gb / 8)
	m := make(map[string]MinMax)
	for _, el := range RngTbl {
		measure := MinMax{1, 0}
		r := rand.New(el.rnd)
		for i := 0; i < size; i++ {
			f := r.Float64()
			measure.Min = math.Min(measure.Min, f)
			measure.Max = math.Max(measure.Max, f)
		}
		assert.InDelta(t, measure.Min, 0, 1e-06, el.name)
		assert.InDelta(t, measure.Max, 1, 1e-06, el.name)
		m[el.name] = measure
	}
	err := PrettyPrint(m)
	if err != nil {
		fmt.Println(err)
	}
}

func TestBuckets(t *testing.T) {
	size := 1e7
	buckets := 100.0
	perBucket := size / buckets
	m := make(map[string]MinMax)

	for _, el := range RngTbl {
		bucketMap := make(map[int]float64)
		r := rand.New(el.rnd)
		for i := 0; i < int(size); i++ {
			f := r.Float64()
			key := int(math.Floor(f * buckets))
			bucketMap[key] += 1
		}
		measure := MinMax{0, 0}
		for idx, val := range bucketMap {
			bucketMap[idx] = math.Round((val-perBucket)/perBucket*10000) / 100
			measure.Min = math.Min(measure.Min, bucketMap[idx])
			measure.Max = math.Max(measure.Max, bucketMap[idx])
		}
		m[el.name] = measure
	}
	err := PrettyPrint(m)
	if err != nil {
		fmt.Println(err)
	}
}

func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}
