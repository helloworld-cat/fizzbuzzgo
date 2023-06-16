package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/helloworld-cat/fizzbuzzgo/httphandlers/listnumbershandler"
	"github.com/helloworld-cat/fizzbuzzgo/interactors/fizzbuzz"
)

const numbersLimit = 200

func main() {
	// Prepare interactors
	numberBuilder := fizzbuzz.NewBasicNumberBuilder()
	numbersGenerator := fizzbuzz.NewBasicNumberdGenerator(numberBuilder, numbersLimit)

	// Prepare HTTP handlers
	listNumbersHandler := listnumbershandler.NewBasic(numbersGenerator)

	// Prepare routes
	router := mux.NewRouter()
	router.Handle("/numbers", listNumbersHandler).Methods(http.MethodPost)

	// Serve, no need graceful logic: no database
	addr := ":8080"
	log.Printf("Listen %s", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatalf("cannot listen and serve: %s", err)
	}
}
