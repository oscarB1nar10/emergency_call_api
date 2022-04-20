
package api

import (
	"emergency_call_rest_api/src/data/models"
	"emergency_call_rest_api/src/data/helper"
	"github.com/gorilla/mux"
	"encoding/json"
	"net/http"
)

type Data struct {
	Success bool `json: "success"`
	Data models.Location `json: "data"`
	Errors []string `json:"errors"`
}

func CreateLocation(w http.ResponseWriter, req *http.Request) {
	location, success := helper.DecodeBody(req)
	if success != true {
		http.Error(w, "Could not decode body", http.StatusBadRequest)
		return 
	}

	var data Data = Data{Errors: make([]string, 0)}
	
	if !helper.IsValidLocation(location) {
		data.Success = false
		data.Errors = append(data.Errors, "Invalid location")

		json, _ := json.Marshal(data)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(json)
		return
	}

	location, success := models.Insert(location.latitude, location.longitude)
	
	if success != true {
		data.Errors = append(data.Errors, "Could not create location")

	} 


	data.Success = true
	data.Data = append(data.Data, location)

	json, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func GetLocation(w http.ResponseWriter, req *http.Request){
	vars := mux.Vars(req)
	id := vars["id"]

	var data Data

	var location models.Location
	var location, success = models.Get(id)

	if success != true {
		data.Success = false
		data.Errors = append(data.Errors, "location not found")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(json)
		return
	}

	data.Success = true
	data.Data = append(data.Data, location)

	json, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}