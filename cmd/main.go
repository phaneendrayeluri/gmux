package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/phaneendrayeluri/gmux/decorators"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/healthcheck", decorators.ResponseWriterJSON(
		http.StatusOK,
		struct{ Status string }{"OK - GMUX"},
	)).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":8080", r))

}
