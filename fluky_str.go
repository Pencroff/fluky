package fluky

import (
	"strings"
)

const (
	latinLower     = "abcdefghijklmnopqrstuvwxyz"
	latinUpper     = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers        = "0123456789"
	symbols        = "!@#$%^&*+="
	linkSymbols    = "_-"
	ukrainianLower = "абвгґдеєжзиіїйклмнопрстуфхцчшщьюя"
	ukrainianUpper = "АБВГҐДЕЄЖЗИІЇЙКЛМНОПРСТУФХЦЧШЩЬЮЯ"
	greekLower     = "αβγδεζηθικλμνξοπρστυφχψω"
	greekUpper     = "ΑΒΓΔΕΖΗΘΙΚΛΜΝΞΟΠΡΣΤΥΦΧΨΩ"
)

type StringOptionsFn func(s *StringOptions)

type StringOptions struct {
	minLen   uint8
	maxLen   uint8
	alphabet string
}

// WithLength makes closure with passed length for options object
func WithLength(v uint8) StringOptionsFn {
	return func(s *StringOptions) {
		s.minLen = v
		s.maxLen = v
	}
}

// WithLengthRange makes closure with passed min and max length for options object
func WithLengthRange(min, max uint8) StringOptionsFn {
	return func(s *StringOptions) {
		s.minLen = min
		s.maxLen = max
	}
}

// WithAlphabet makes closure with passed alphabet for options object
func WithAlphabet(alphabet string) StringOptionsFn {
	return func(s *StringOptions) {
		s.alphabet = alphabet
	}
}

// AndAlphabet extend configured alphabet for options object
func AndAlphabet(alphabet string) StringOptionsFn {
	return func(s *StringOptions) {
		s.alphabet += alphabet
	}
}

// WithNumericAlphabet configure numeric alphabet
func WithNumericAlphabet() StringOptionsFn {
	return func(s *StringOptions) {
		s.alphabet = numbers
	}
}

// AndNumericAlphabet extend configured alphabet with numbers
func AndNumericAlphabet() StringOptionsFn {
	return func(s *StringOptions) {
		s.alphabet += numbers
	}
}

// WithLatinLowerAlphabet configure latin lower alphabet
func WithLatinLowerAlphabet() StringOptionsFn {
	return func(s *StringOptions) {
		s.alphabet = latinLower
	}
}

// AndLatinLowerAlphabet extend configured alphabet with latin lower characters
func AndLatinLowerAlphabet() StringOptionsFn {
	return func(s *StringOptions) {
		s.alphabet += latinLower
	}
}

// WithLatinUpperAlphabet configure latin upper alphabet
func WithLatinUpperAlphabet() StringOptionsFn {
	return func(s *StringOptions) {
		s.alphabet = latinUpper
	}
}

// AndLatinUpperAlphabet extend configured alphabet with latin upper characters
func AndLatinUpperAlphabet() StringOptionsFn {
	return func(s *StringOptions) {
		s.alphabet += latinUpper
	}
}

// WithSymbolsAlphabet configure symbols alphabet
func WithSymbolsAlphabet() StringOptionsFn {
	return func(s *StringOptions) {
		s.alphabet = symbols + linkSymbols
	}
}

// AndSymbolsAlphabet extend configured alphabet with symbols
func AndSymbolsAlphabet() StringOptionsFn {
	return func(s *StringOptions) {
		s.alphabet += symbols + linkSymbols
	}
}

// WithSymbolsUrlSafeAlphabet configure safe for url usage alphabet symbols
func WithSymbolsUrlSafeAlphabet() StringOptionsFn {
	return func(s *StringOptions) {
		s.alphabet = linkSymbols
	}
}

// AndSymbolsUrlSafeAlphabet extend configured alphabet with symbols
func AndSymbolsUrlSafeAlphabet() StringOptionsFn {
	return func(s *StringOptions) {
		s.alphabet += linkSymbols
	}
}

// WithUrlSafeAlphabet configure safe for url usage alphabet
func WithUrlSafeAlphabet() StringOptionsFn {
	return func(s *StringOptions) {
		s.alphabet = latinLower + latinUpper + numbers + linkSymbols
	}
}

// WithHexAlphabet configure hex alphabet
func WithHexAlphabet() StringOptionsFn {
	return func(s *StringOptions) {
		s.alphabet = numbers + "abcdef"
	}
}

// WithUkrainianLowerAlphabet configure ukrainian lower alphabet
func WithUkrainianLowerAlphabet() StringOptionsFn {
	return func(s *StringOptions) {
		s.alphabet = ukrainianLower
	}
}

// AndUkrainianLowerAlphabet extend configured alphabet with ukrainian lower characters
func AndUkrainianLowerAlphabet() StringOptionsFn {
	return func(s *StringOptions) {
		s.alphabet += ukrainianLower
	}
}

// WithUkrainianUpperAlphabet configure ukrainian upper alphabet
func WithUkrainianUpperAlphabet() StringOptionsFn {
	return func(s *StringOptions) {
		s.alphabet = ukrainianUpper
	}
}

// AndUkrainianUpperAlphabet extend configured alphabet with ukrainian upper characters
func AndUkrainianUpperAlphabet() StringOptionsFn {
	return func(s *StringOptions) {
		s.alphabet += ukrainianUpper
	}
}

// WithGreekLowerAlphabet configure greek lower alphabet
func WithGreekLowerAlphabet() StringOptionsFn {
	return func(s *StringOptions) {
		s.alphabet = greekLower
	}
}

// AndGreekLowerAlphabet extend configured alphabet with greek lower characters
func AndGreekLowerAlphabet() StringOptionsFn {
	return func(s *StringOptions) {
		s.alphabet += greekLower
	}
}

// WithGreekUpperAlphabet configure greek upper alphabet
func WithGreekUpperAlphabet() StringOptionsFn {
	return func(s *StringOptions) {
		s.alphabet = greekUpper
	}
}

// AndGreekUpperAlphabet extend configured alphabet with greek upper characters
func AndGreekUpperAlphabet() StringOptionsFn {
	return func(s *StringOptions) {
		s.alphabet += greekUpper
	}
}

// String returns random string configured by default options
func (f *Fluky) String(opts ...StringOptionsFn) string {
	b := &StringOptions{minLen: 5, maxLen: 20, alphabet: latinLower + latinUpper + numbers + symbols + linkSymbols}
	for _, opt := range opts {
		opt(b)
	}
	l := f.Integer(WithIntRange(int(b.minLen), int(b.maxLen)))
	alphabetRunes := []rune(b.alphabet)
	maxIdx := int64(len(alphabetRunes) - 1)
	builder := strings.Builder{}
	for i := 0; i < l; i++ {
		idx := maxIdx
		if maxIdx != 0 {
			r := f.Int63()
			idx = r % maxIdx
		}
		builder.WriteRune(alphabetRunes[idx])
	}
	return builder.String()
}
