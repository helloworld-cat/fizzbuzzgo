package fizzbuzz

import (
	"errors"
	"testing"
)

func TestBasicNumbersGenerator(t *testing.T) {
	nb := &fakeNumberBuilder{called: false}
	ng := NewBasicNumberdGenerator(nb, 200)

	t.Run("it should not return any errors", func(t *testing.T) {
		if _, err := ng.Call(100, 7, 9, "Fizz", "Buzz"); err != nil {
			t.Errorf("unexpected error: %s", err)
		}
	})

	var invalidValueError *InvalidValueError
	var errorLimitTooHigh *ErrorLimitTooHigh

	t.Run("when the limit is reached, it returns an error", func(t *testing.T) {
		if _, err := ng.Call(9999, 7, 9, "Fizz", "Buzz"); errors.As(err, &errorLimitTooHigh) == false {
			t.Errorf("unexpected error type: %s", err)
		}
	})

	t.Run("when number a is zero, it returns an error", func(t *testing.T) {
		if _, err := ng.Call(100, 0, 9, "Fizz", "Buzz"); errors.As(err, &invalidValueError) == false {
			t.Errorf("unexpected error type: %s", err)
		}
	})

	t.Run("when number b is zero, it returns an error", func(t *testing.T) {
		if _, err := ng.Call(100, 7, 0, "Fizz", "Buzz"); errors.As(err, &invalidValueError) == false {
			t.Errorf("unexpected error type: %s", err)
		}
	})

	t.Run("it calls NumberBuilder", func(t *testing.T) {
		nb := &fakeNumberBuilder{called: false}
		ng := NewBasicNumberdGenerator(nb, 200)

		ng.Call(10, 1, 2, "foo", "bar")

		if nb.called == false {
			t.Errorf("NumberBuilder never called")
		}
	})

	t.Run("it generates correct numbers", func(t *testing.T) {
		numbers, _ := ng.Call(10, 1, 2, "foo", "bar")
		if l := len(numbers); l != 10 {
			t.Errorf("unexpected length, want %d, got: %d", 10, l)
		}

		if n := numbers[0]; n != "foo" {
			t.Errorf("unexpected number result, want `foo`, got: %s", n)
		}

		if n := numbers[1]; n != 2 {
			t.Errorf("unexpected number result, want `foo`, got: %s", n)
		}
	})

}

type (
	fakeNumberBuilder struct {
		called bool
	}
)

func (it *fakeNumberBuilder) Call(v, na, nb int, wa, wb string) interface{} {
	it.called = true

	if v == 1 {
		return "foo"
	}

	return v
}
