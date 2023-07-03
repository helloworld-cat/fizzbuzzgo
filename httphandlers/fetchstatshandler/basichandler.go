package fetchstatshandler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/helloworld-cat/fizzbuzzgo/entities"
	"github.com/helloworld-cat/fizzbuzzgo/repositories/statsrepository"
)

type (
	basicHandler struct {
		statsRepo statsrepository.StatsRepository
	}
)

func NewBasicHandler(statsRepo statsrepository.StatsRepository) http.Handler {
	return &basicHandler{
		statsRepo: statsRepo,
	}
}

func (h *basicHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	userReq, err := entities.NewUserRequest(rw, req)
	if err != nil {
		h.writeJSONResponse(rw, err, http.StatusBadRequest)
		return
	}

	// Validate and prepare inputs
	inputs, errs := userReq.ValidateAndNormalizeInputs(0)
	if len(errs) > 0 {
		h.writeJSONResponse(rw, errs, http.StatusBadRequest)
		return
	}

	statsValue, err := h.statsRepo.Fetch(
		inputs.NumberA(),
		inputs.NumberB(),
		inputs.WordA(),
		inputs.WordB(),
	)

	if err != nil {
		log.Printf("cannot fetch stats: %s", err)
		h.writeJSONResponse(rw, err, http.StatusBadRequest)
	}

	data := struct {
		NumberA int    `json:"number_a"`
		NumberB int    `json:"number_b"`
		WordA   string `json:"word_a"`
		WordB   string `json:"word_b"`
		Stats   int    `json:"stats"`
	}{
		Stats:   statsValue,
		NumberA: inputs.NumberA(),
		NumberB: inputs.NumberB(),
		WordA:   inputs.WordA(),
		WordB:   inputs.WordB(),
	}

	h.writeJSONResponse(rw, data, http.StatusOK)
}

func (h *basicHandler) writeJSONResponse(rw http.ResponseWriter, data any, httpCode int) {
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.WriteHeader(httpCode)

	blob, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Printf("cannot marshal json response: %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	if _, err := rw.Write(blob); err != nil {
		log.Printf("cannot writer blob: %s", err)
		return
	}
}
