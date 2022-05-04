package helper

import (
	"emergency_call_rest_api/src/data/models"
	"fmt"
	"net/http"
	"strconv"
)

func DecodeBody(req *http.Request) (models.Location, bool) {
	var latitude, _ = strconv.ParseFloat(req.FormValue("latitude"), 8)
	var longitude, _ = strconv.ParseFloat(req.FormValue("longitude"), 8)
	var location = models.Location{0, latitude, longitude}

	//err := json.NewDecoder(req.Body).Decode(&location)

	/*if location == nil {
		return models.Location{}, false
	}*/

	fmt.Printf("location decoded: %s", location)

	return location, true

}

func IsValidLocation(location models.Location) bool {
	fmt.Println(location)
	return true
}
