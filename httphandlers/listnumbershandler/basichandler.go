package listnumbershandler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/helloworld-cat/fizzbuzzgo/entities"
	"github.com/helloworld-cat/fizzbuzzgo/interactors/fizzbuzz"
	"github.com/helloworld-cat/fizzbuzzgo/repositories/statsrepository"
)

func NewBasic(ng fizzbuzz.NumbersGenerator, statsRepo statsrepository.StatsRepository) http.Handler {
	return &basicHandler{
		maxLimit:         200,
		numbersGenerator: ng,
		statsRepo:        statsRepo,
	}
}

type (
	basicHandler struct {
		maxLimit         int
		numbersGenerator fizzbuzz.NumbersGenerator
		statsRepo        statsrepository.StatsRepository
	}
)

func (h *basicHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	// Process user request
	userReq, err := entities.NewUserRequest(rw, req)
	if err != nil {
		h.writeJSONResponse(rw, err, http.StatusBadRequest)
		return
	}

	// Validate and prepare inputs
	inputs, errs := userReq.ValidateAndNormalizeInputs(h.maxLimit)
	if len(errs) > 0 {
		h.writeJSONResponse(rw, errs, http.StatusBadRequest)
		return
	}

	if _, err := h.statsRepo.Incr(
		inputs.NumberA(),
		inputs.NumberB(),
		inputs.WordA(),
		inputs.WordB(),
	); err != nil {
		log.Printf("cannot incr. stats: %s", err)
		// pass
		// we do not block the response,
		// stats are not essential for this app.
	}

	// Call interactor
	numbers, err := h.numbersGenerator.Call(
		inputs.Limit(),
		inputs.NumberA(),
		inputs.NumberB(),
		inputs.WordA(),
		inputs.WordB(),
	)

	if err != nil {
		log.Printf("cannot compute fizzbuzz numbers: %s", err)
		h.writeJSONResponse(rw, err, http.StatusBadRequest)
	}

	// Send response
	h.writeJSONResponse(rw, numbers, http.StatusOK)
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
