package fizzbuzz

import "fmt"

// NewBasicNumberBuilder returns fizzbuzz engine.
// See "Call" function for more details.
func NewBasicNumberdGenerator(numberBuilder NumberBuilder, max int) NumbersGenerator {
	return &basicNumbersGenerator{
		numberBuilder: numberBuilder,
		max:           max,
	}
}

type (
	basicNumbersGenerator struct {
		numberBuilder NumberBuilder
		max           int
	}

	InvalidValueError struct {
		Value int
	}

	ErrorLimitTooHigh struct {
		Value int
		Max   int
	}
)

// Call returns list of numbers prepared by NumberBuilder.
// See attached NumberBuilder for more detalis.
func (bng *basicNumbersGenerator) Call(limit, na, nb int, wa, wb string) (Numbers, error) {
	if limit > bng.max {
		return nil, &ErrorLimitTooHigh{Max: bng.max, Value: limit}
	}

	if na == 0 {
		return nil, &InvalidValueError{Value: na}
	}

	if nb == 0 {
		return nil, &InvalidValueError{Value: nb}
	}

	numbers := make(Numbers, 0)
	for i := 1; i <= limit; i++ {
		n := bng.numberBuilder.Call(i, na, nb, wa, wb)
		numbers = append(numbers, n)
	}

	return numbers, nil
}

func (e *InvalidValueError) Error() string {
	return fmt.Sprintf("must be greater than zero, value: %d", e.Value)
}

func (e *ErrorLimitTooHigh) Error() string {
	return fmt.Sprintf("limit too high, max: %d, value: %d", e.Max, e.Value)
}
