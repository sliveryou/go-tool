package randx

import "crypto/rand"

var (
	// stdSource standard source string.
	stdSource = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	// stdNumberSource standard number source string.
	stdNumberSource = "0123456789"
)

// StdSource returns standard source string.
func StdSource() string {
	return stdSource
}

// StdNumberSource returns standard number source string.
func StdNumberSource() string {
	return stdNumberSource
}

// NewString returns a new random string of the provided length, consisting
// of the standard source string.
// It panics if source length is wrong (<1 or >256) or rand.Read occurs an error.
func NewString(length int) string {
	return NewWithSource(length, StdSource())
}

// NewNumber returns a new random string of the provided length, consisting
// of the standard number source string.
// It panics if source length is wrong (<1 or >256) or rand.Read occurs an error.
func NewNumber(length int) string {
	return NewWithSource(length, StdNumberSource())
}

// NewWithSource returns a new random string of the provided length, consisting
// of the provided source string.
// It panics if source length is wrong (<1 or >256) or rand.Read occurs an error.
func NewWithSource(length int, source string) string {
	if length == 0 {
		return ""
	}

	sl := len(source)
	if sl < 1 || sl > 256 {
		panic("randx: wrong source length")
	}

	rbMax := 255 - (256 % sl)
	b := make([]byte, length)
	r := make([]byte, length+length/2) // storage for random bytes
	i := 0

	for {
		if _, err := rand.Read(r); err != nil {
			panic(err)
		}

		for _, rb := range r {
			v := int(rb)
			if v > rbMax { // skip to avoid modulo bias
				continue
			}

			b[i] = source[v%sl]
			i++

			if i == length {
				return string(b)
			}
		}
	}
}
