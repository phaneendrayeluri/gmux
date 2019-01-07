package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/phaneendrayeluri/gmux/decorators"
)

func main() {

	routingPath := flag.String("routing.path", "/gmux/v1", "routing path configured in apigateway, default pattern /NameOfProject/LatestVersion")

	r := mux.NewRouter()

	r.HandleFunc("/healthcheck", decorators.ResponsdAsJSON(
		http.StatusOK,
		struct{ Status string }{"OK - GMUX"},
	)).Methods(http.MethodGet)

	s := r.PathPrefix(*routingPath).Subrouter()

	s.HandleFunc("/whoareyou", decorators.ResponsdAsJSON(
		http.StatusOK,
		struct{ Name string }{"I am a microservice"},
	)).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", r))

}
