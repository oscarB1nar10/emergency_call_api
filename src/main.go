package main

import (
	"emergency_call_rest_api/src/data/api"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	// Listen through specific port
	var port = "8080"
	router := mux.NewRouter()

	apiRouter := router.PathPrefix("/api/").Subrouter()
	apiRouter.HandleFunc("/location", api.CreateLocation).Methods("POST")
	apiRouter.HandleFunc("/location/{id}", api.GetLocation).Methods("GET")

	fmt.Printf("Server running at port %s", port)

	// Wraps in Log to handle error if happen
	log.Fatal(http.ListenAndServe(":"+port, router))
}
