package listnumbershandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type (
	request struct {
		Limit   int    `json:"limit"`
		NumberA int    `json:"number_a"`
		NumberB int    `json:"number_b"`
		WordA   string `json:"word_a"`
		WordB   string `json:"word_b"`
	}

	inputs struct {
		limit   int
		numberA int
		numberB int
		wordA   string
		wordB   string
	}

	requestError struct {
		Name    string      `json:"name"`
		Value   interface{} `json:"value"`
		Message string      `json:"message"`
	}
)

func NewUserRequest(rw http.ResponseWriter, req *http.Request) (*request, error) {
	body := http.MaxBytesReader(rw, req.Body, 1048576)

	dec := json.NewDecoder(body)
	dec.DisallowUnknownFields()

	userReq := &request{}
	if err := dec.Decode(userReq); err != nil {
		return nil, err
	}

	return userReq, nil
}

func (r *request) ValidateAndNormalizeInputs(maxLimit int) (*inputs, []error) {
	inputs := &inputs{
		numberA: r.NumberA,
		numberB: r.NumberB,
		wordA:   r.WordA,
		wordB:   r.WordB,
		limit:   r.Limit,
	}

	if inputs.limit > maxLimit || inputs.limit == 0 {
		inputs.limit = maxLimit
	}

	errs := make([]error, 0)
	if inputs.limit < 0 {
		errs = append(
			errs,
			&requestError{
				Message: "limit parameter must be greater than zero",
				Name:    "limit",
				Value:   r.Limit,
			},
		)
	}

	if inputs.numberA == 0 {
		errs = append(
			errs,
			&requestError{
				Name:    "number_a",
				Message: "cannot be zero",
				Value:   r.NumberA,
			},
		)
	}

	if inputs.numberB == 0 {
		errs = append(
			errs,
			&requestError{
				Name:    "number_b",
				Message: "cannot be zero",
				Value:   r.NumberA,
			},
		)
	}

	if len(errs) > 0 {
		return nil, errs
	}

	inputs.wordA = r.normalizeWord(inputs.wordA, "Fizz")
	inputs.wordB = r.normalizeWord(inputs.wordB, "Buzz")

	return inputs, nil
}

func (r *request) normalizeWord(s, defaultValue string) string {
	s = strings.TrimSpace(s)
	if strings.Compare(s, "") == 0 {
		return defaultValue
	}
	return s
}

func (e *requestError) Error() string {
	return fmt.Sprintf("Invalid parameter `%s` (name: `%s`, value: `%s`)", e.Message, e.Name, e.Value)
}

func (i *inputs) NumberA() int  { return i.numberA }
func (i *inputs) NumberB() int  { return i.numberB }
func (i *inputs) WordA() string { return i.wordA }
func (i *inputs) WordB() string { return i.wordB }
func (i *inputs) Limit() int    { return i.limit }
