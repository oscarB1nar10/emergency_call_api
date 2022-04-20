
import ("encoding/json")

type Location struct {
	id int `json:"id"`
    latitude   string `json:"latitude"`
    longitude string `json:"longitude"`
}

type UserPhone struct {
    id int `json:"id"`
    name     string `json:"name"`
    location Location `json:"location"`
}