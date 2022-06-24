package models

type Image struct {
	Pk     int    `json:"pk"`
	UserId string `json:"user_id"`
	Url    string `json:"url"`
}
