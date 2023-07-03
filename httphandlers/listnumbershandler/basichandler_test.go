package listnumbershandler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/helloworld-cat/fizzbuzzgo/interactors/fizzbuzz"
)

func TestBasicHandler(t *testing.T) {

	t.Run("Allow user to prints the numbers from 1 to 200 with 7 and 9 multiples", func(t *testing.T) {
		payload := `
			{"number_a": 7, "number_b": 9}
		`

		expected := map[int]interface{}{
			6:  "Fizz",
			8:  "Buzz",
			62: "FizzBuzz",
		}

		runTest(t, payload, expected, 200)
	})

	t.Run("Allow user to define the limit of numbers that he wants to display", func(t *testing.T) {
		payload := `
			{"number_a": 7, "number_b": 9, "limit": 100}
		`
		expected := map[int]interface{}{
			6:  "Fizz",
			8:  "Buzz",
			62: "FizzBuzz",
		}
		runTest(t, payload, expected, 100)

	})

	t.Run("Allow user to define the two multiples numbers he wants to override (ex: 3 and 5)", func(t *testing.T) {
		payload := `
			{"number_a": 3, "number_b": 5}
		`
		expected := map[int]interface{}{
			2:  "Fizz",
			4:  "Buzz",
			14: "FizzBuzz",
		}

		runTest(t, payload, expected, 200)
	})

	t.Run("Allow user to define the two words he wants to display", func(t *testing.T) {
		payload := `
			{"number_a": 7, "number_b": 9, "word_a": "foo", "word_b": "bar"}
		`

		expected := map[int]interface{}{
			6:  "foo",
			8:  "bar",
			62: "foobar",
		}

		runTest(t, payload, expected, 200)
	})

}

func runTest(t *testing.T, payload string, expected map[int]interface{}, expectedSize int) {
	fakeStatsRepo := &fakeStatsRepo{called: false}

	basicNumberBuilder := fizzbuzz.NewBasicNumberBuilder()
	basicNumbersGenerator := fizzbuzz.NewBasicNumberdGenerator(basicNumberBuilder, 200)

	req, err := http.NewRequest(http.MethodPost, "/numbers", strings.NewReader(payload))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	NewBasic(basicNumbersGenerator, fakeStatsRepo).ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf(
			"unexpected status code: want %+v, got: %+v",
			http.StatusOK,
			status,
		)
	}

	numbers := make([]interface{}, 0)
	if err := json.Unmarshal(rr.Body.Bytes(), &numbers); err != nil {
		t.Errorf(
			"unexpected json error: %s",
			err,
		)
	}

	if len(numbers) != expectedSize {
		t.Errorf("unexpected length of `numbers`")
	}

	for k, v := range expected {
		got := numbers[k]
		if got != v {
			t.Errorf("unexpected values, %d, want %+v, got %+v", k, v, got)
		}
	}

	if fakeStatsRepo.called == false {
		t.Errorf("fakeStatsRepo.Incr never called")
	}

}

type (
	fakeStatsRepo struct {
		called bool
	}
)

func (r *fakeStatsRepo) Incr(nb, nc int, wa, wb string) (int, error) {
	r.called = true
	return 1, nil
}

func (r *fakeStatsRepo) Fetch(nb, nc int, wa, wb string) (int, error) {
	return 0, nil
}
