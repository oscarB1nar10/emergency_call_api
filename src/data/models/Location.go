package models

import (
	"emergency_call_rest_api/src/data"
)

type Location struct {
	Id        int     `json:"id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func Insert(latitude float64, longitude float64) (Location, bool) {
	db := data.GetConnection()

	var location_id int
	db.QueryRow("INSERT INTO location(latitude, longitude) VALUES($1, $2) RETURNING id", latitude, longitude).Scan(&location_id)

	if location_id == 0 {
		return Location{}, false
	}

	return Location{location_id, latitude, longitude}, true
}

func Get(id string) (Location, bool) {
	db := data.GetConnection()
	row := db.QueryRow("SELECT * FROM location where id = $1", id)

	var id_get int
	var latitude float64
	var longitude float64

	err := row.Scan(&id, &latitude, &longitude)

	if err != nil {
		return Location{}, false
	}

	return Location{id_get, latitude, longitude}, true
}
