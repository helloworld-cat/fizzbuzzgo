package fizzbuzz

type (
	NumberBuilder interface {
		Call(value, numberA, numberB int, wordA, wordB string) interface{}
	}

	NumbersGenerator interface {
		Call(limit, numberA, numberB int, wordA, wordB string) (Numbers, error)
	}

	Numbers []interface{}
)
