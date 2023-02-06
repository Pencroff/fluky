package experiments

import (
	"encoding/binary"
	"io"
	"log"
	"math"
	"math/rand"
	"os"
	"path"
	"path/filepath"
	"testing"
)

func TestDataExport(t *testing.T) {
	size := int64(1 * Size1Kb)
	folderPath, err := filepath.Abs("../container/rnd_data")
	checkErr(err)
	err = os.Chdir(folderPath)
	if err != nil {
		err = os.Mkdir(folderPath, 0755)
		checkErr(err)
	}

	for _, el := range RngTbl {
		if el.exportData {
			dataPath := path.Join(folderPath, el.name+"_data.bin")
			if stat, err := os.Stat(dataPath); err == nil {
				if stat.Size() == size {
					continue
				}
			}
			WriteToFile(dataPath, size, el.rnd)
		}
	}
}

func WriteToFile(filePath string, size int64, rnd rand.Source64) {
	out, err := os.Create(filePath)
	checkErr(err)
	defer out.Close()
	reader := NewRngReader(rnd, binary.LittleEndian, int64(32*Size1Kb))

	n, err := io.CopyN(out, reader, size)
	checkErr(err)
	log.Printf("Fill %d random bytes to %s file\n", n, filePath)
}

type RngReader struct {
	rng       rand.Source64
	order     binary.ByteOrder
	batchSize int64
}

func NewRngReader(rng rand.Source64, order binary.ByteOrder, butchSize int64) *RngReader {
	return &RngReader{rng: rng, order: order, batchSize: butchSize}
}

func (r *RngReader) Read(p []byte) (n int, err error) {
	b := make([]byte, 8)
	s := math.Min(float64(cap(p)), float64(r.batchSize))
	itterNum := int(s / 8)
	for i := 0; i < itterNum; i++ {
		v := r.rng.Uint64()
		binary.LittleEndian.PutUint64(b, v)
		for j := 0; j < 8; j++ {
			idx := 8*i + j
			p[idx] = b[j]
		}
	}

	return int(s), nil
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
