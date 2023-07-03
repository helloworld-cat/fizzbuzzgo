package fetchstatshandler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestServeHTTP(t *testing.T) {

	value := 123
	fakeStatsRepo := &fakeStatsRepo{value: value}

	payload := `
	  {"number_a": 7, "number_b": 9, "word_a": "foo", "word_b": "bar"}
        `

	req, err := http.NewRequest(http.MethodPost, "/stats", strings.NewReader(payload))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	NewBasicHandler(fakeStatsRepo).ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf(
			"unexpected status code: want %+v, got: %+v",
			http.StatusOK,
			status,
		)
	}

	data := struct {
		NumberA int    `json:"number_a"`
		NumberB int    `json:"number_b"`
		WordA   string `json:"word_a"`
		WordB   string `json:"word_b"`
		Stats   int    `json:"stats"`
	}{}

	if err := json.Unmarshal(rr.Body.Bytes(), &data); err != nil {
		t.Errorf(
			"unexpected json error: %s",
			err,
		)
	}

	if data.Stats != value {
		t.Errorf("unexpected value. Want %d, but got: %d", value, data.Stats)
	}

	// TODO: check other parameter: worda, wordb, etc.
}

type (
	fakeStatsRepo struct {
		value int
	}
)

func (r *fakeStatsRepo) Incr(nb, nc int, wa, wb string) (int, error) {
	return 0, nil
}

// TODO: check parameter (worda, wordb, etc.)
func (r *fakeStatsRepo) Fetch(nb, nc int, wa, wb string) (int, error) {
	return r.value, nil
}
