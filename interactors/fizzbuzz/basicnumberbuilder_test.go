package fizzbuzz

import (
	"testing"
)

func TestNewBasicNumberBuilder(t *testing.T) {
	it := NewBasicNumberBuilder()

	na := 7
	nb := 9
	wa := "Fizz"
	wb := "Buzz"
	wab := "FizzBuzz"

	tests := []struct {
		input    int
		expected interface{}
	}{
		{
			input:    -7,
			expected: wa,
		},
		{
			input:    -1,
			expected: -1,
		},
		{
			input:    0, // "0 / 7" && "0 / 9" -> 0
			expected: wab,
		},
		{
			input:    1,
			expected: 1,
		},
		{
			input:    2,
			expected: 2,
		},
		{
			input:    7,
			expected: wa,
		},
		{
			input:    9,
			expected: wb,
		},
		{
			input:    63,
			expected: wab,
		},
		{
			input:    189,
			expected: wab,
		},
	}

	for _, ct := range tests {
		t.Logf("** input: %d", ct.input)

		v := it.Call(ct.input, na, nb, wa, wb)

		switch expectedValue := ct.expected.(type) {
		case int:
			value, ok := v.(int)
			if !ok {
				t.Errorf("unexpected value type. Want int, but got: %+v", v)
				continue
			}
			if value != expectedValue {
				t.Errorf("unexpected result. Want `%d`, but got: `%d`", expectedValue, ct.input)
				continue
			}
		case string:
			value, ok := v.(string)
			if !ok {
				t.Errorf("unexpected value type. Want string, but got: %+v", v)
				continue
			}
			if value != expectedValue {
				t.Errorf("unexpected result. Want `%s`, but got: `%d`", expectedValue, ct.input)
				continue
			}
		default:
			t.Errorf("cannot convert value `%+v`", v)
		}
	}
}
