package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/phaneendrayeluri/gmux/decorators"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/healthcheck", decorators.TypeDecoratorJSON(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(struct{ Status string }{"OK"})
	})).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":8080", r))
}
