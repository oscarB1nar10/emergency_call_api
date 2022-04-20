

package helper

import (
	"fmt"
	"encoding/json"
	"emergency_call_rest_api/src/data/models"
	"net/http"
)

func DecodeBody(req *http.Request) (models.Location, bool) {
	var location models.Location
	err := json.NewDecoder(req.Body).Decode(&location)

	if err != nil {
		return models.Location{}, false
	}

	return location, true

}

func IsValidLocation(location models.Location) bool {
	fmt.Println(location)
	return true
}