package fizzbuzz

import "fmt"

type (
	basicNumberBuilder struct{}
)

// NewBasicNumberBuilder prepare NumberBuilder with basic implementation.
func NewBasicNumberBuilder() NumberBuilder {
	return &basicNumberBuilder{}
}

// Call returns computed value like this:
// if value is multiple of n1, then it returns word1
// if value is multiple of n2, then it returns word2
// if value is multiple of n1 and n2, then it returns word1+word2"
// otherwise Call returns value.
func (bnb *basicNumberBuilder) Call(v, na, nb int, wa, wb string) interface{} {
	if v%na == 0 && v%nb == 0 {
		return fmt.Sprintf("%s%s", wa, wb)
	}
	if v%na == 0 {
		return wa
	}
	if v%nb == 0 {
		return wb
	}

	return v
}
