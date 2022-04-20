
package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"emergency_call_rest_api/src/data/api"
)

func main() {
	// Listen through specific port
	var port string = "8080"
	router := mux.NewRouter()

	apiRouter := router.PathPrefix("/api/").Subrouter()
	apiRouter.HandleFunc("/location", api.CreateLocation).Methods("POST")
	apiRouter.HandleFunc("/location/{id}", api.GetLocation).Methods("GET")

	fmt.Printf("Server running at port %s", port)
	http.ListenAndServe(":" + port, router)
}