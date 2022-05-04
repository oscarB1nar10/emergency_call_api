package api

import (
	"emergency_call_rest_api/src/data/helper"
	"emergency_call_rest_api/src/data/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Data struct {
	Success bool            `json:"success"`
	Data    models.Location `json:"data"`
	Errors  []string        `json:"errors"`
}

func CreateLocation(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("Server running at port")
	location, success := helper.DecodeBody(req)
	if success != true {
		http.Error(w, "Could not decode body", http.StatusBadRequest)
		return
	}

	var data = Data{Errors: make([]string, 0)}

	if !helper.IsValidLocation(location) {
		data.Success = false
		data.Errors = append(data.Errors, "Invalid location")

		json, _ := json.Marshal(data)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(json)
		return
	}

	location, success = models.Insert(location.Latitude, location.Longitude)

	if success != true {
		data.Errors = append(data.Errors, "Could not create location")
	}

	data.Success = true

	json, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func GetLocation(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	var data Data

	var location, success = models.Get(id)

	if success != true {
		data.Success = false
		data.Errors = append(data.Errors, "location not found")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		log.Fatal(json.NewEncoder(w).Encode(data))
		return
	}

	data.Success = true
	data.Data = location

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	var response = json.NewEncoder(w).Encode(data)
	log.Printf("result: %s", response)
}
