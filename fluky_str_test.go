package fluky

import (
	"fmt"
	src "github.com/Pencroff/fluky/source"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func TestFluky_Str_StringOptionsFn(t *testing.T) {
	lst := []struct {
		fn       StringOptionsFn
		actual   StringOptions
		expected StringOptions
	}{
		{fn: WithLength(10),
			actual:   StringOptions{},
			expected: StringOptions{minLen: 10, maxLen: 10}},
		{fn: WithLengthRange(5, 7),
			actual:   StringOptions{},
			expected: StringOptions{minLen: 5, maxLen: 7}},
		{fn: WithAlphabet("abc"),
			actual:   StringOptions{},
			expected: StringOptions{alphabet: "abc"}},
		{fn: AndAlphabet("abc"),
			actual:   StringOptions{alphabet: "012"},
			expected: StringOptions{alphabet: "012abc"}},
		{fn: WithUrlSafeAlphabet(),
			actual:   StringOptions{},
			expected: StringOptions{alphabet: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_-"}},
		{fn: WithHexAlphabet(),
			actual:   StringOptions{},
			expected: StringOptions{alphabet: "0123456789abcdef"}},
		{fn: WithNumericAlphabet(),
			actual:   StringOptions{},
			expected: StringOptions{alphabet: "0123456789"}},
		{fn: AndNumericAlphabet(),
			actual:   StringOptions{alphabet: "abcdef"},
			expected: StringOptions{alphabet: "abcdef0123456789"}},
		{fn: WithLowerAlphabet(),
			actual:   StringOptions{alphabet: "1234567890"},
			expected: StringOptions{alphabet: "abcdefghijklmnopqrstuvwxyz"}},
		{fn: AndLowerAlphabet(),
			actual:   StringOptions{alphabet: "1234567890"},
			expected: StringOptions{alphabet: "1234567890abcdefghijklmnopqrstuvwxyz"}},
		{fn: WithUpperAlphabet(),
			actual:   StringOptions{alphabet: "1234567890"},
			expected: StringOptions{alphabet: "ABCDEFGHIJKLMNOPQRSTUVWXYZ"}},
		{fn: AndUpperAlphabet(),
			actual:   StringOptions{alphabet: "1234567890"},
			expected: StringOptions{alphabet: "1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ"}},
		{fn: WithSymbolsAlphabet(),
			actual:   StringOptions{alphabet: "1234567890"},
			expected: StringOptions{alphabet: "!@#$%^&*+=_-"}},
		{fn: AndSymbolsAlphabet(),
			actual:   StringOptions{alphabet: "1234567890"},
			expected: StringOptions{alphabet: "1234567890!@#$%^&*+=_-"}},
		{fn: WithSymbolsUrlSafeAlphabet(),
			actual:   StringOptions{alphabet: "1234567890"},
			expected: StringOptions{alphabet: "_-"}},
		{fn: AndSymbolsUrlSafeAlphabet(),
			actual:   StringOptions{alphabet: "1234567890"},
			expected: StringOptions{alphabet: "1234567890_-"}},
	}
	for _, el := range lst {
		el.fn(&el.actual)
		assert.Equal(t, el.expected, el.actual)
	}
}

func TestFluky_Str_NoOptions(t *testing.T) {
	mRng := new(RngMock)
	mRng.On("Int63").Return(int64(10)).Times(1)
	mRng.On("Int63").Return(int64(27)).Times(5)
	mRng.On("Int63").Return(int64(57)).Times(5)
	mRng.On("Int63").Return(int64(64)).Times(5)
	// default options
	// min 5
	// max 20
	// alphabet "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*+=_-"
	f := NewFluky(mRng)
	str := f.String()
	assert.Equal(t, 15, len(str)) // 10%(20 - 5) + 5
	assert.Equal(t, "BBBBB55555#####", str)
}

func TestFluky_Str_UnicodeAlphabet_One(t *testing.T) {
	mRng := new(RngMock)
	mRng.On("Int63").Return(int64(1))
	// default options
	// min 5
	// max 20
	// alphabet "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*+=_-"
	f := NewFluky(mRng)
	actual := f.String(WithAlphabet("❤"), WithLength(5))
	expected := "❤❤❤❤❤"
	fmt.Printf("actual: %s != %s\n", actual, expected)
	assert.Equal(t, expected, actual, "%s != %s", expected, actual)
}

func TestFluky_Str_UnicodeAlphabet(t *testing.T) {
	mRng := new(RngMock)
	mRng.On("Int63").Return(int64(1)).Times(2)
	mRng.On("Int63").Return(int64(2)).Times(1)
	mRng.On("Int63").Return(int64(3)).Times(2)
	// default options
	// min 5
	// max 20
	// alphabet "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*+=_-"
	f := NewFluky(mRng)
	actual := f.String(WithAlphabet("❤🎈🍕😉"), WithLength(5))
	expected := "🎈🎈🍕❤❤"
	assert.Equal(t, expected, actual, "%s != %s", expected, actual)
}

func TestFluky_Str_WithRnd(t *testing.T) {
	// default options
	// min 5
	// max 20
	// alphabet "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*+=_-"
	s := src.NewSplitMix64Source(time.Now().UnixNano())
	r := rand.New(s)
	f := NewFluky(r)
	n := 1000
	alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*+=_-"
	for i := 0; i < n; i++ {
		str := f.String()
		assert.GreaterOrEqual(t, len(str), 5)
		assert.LessOrEqual(t, len(str), 20)
		for _, c := range str {
			assert.Contains(t, alphabet, string(c))
		}
	}
}
