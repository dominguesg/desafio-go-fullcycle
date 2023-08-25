package models

type Location struct {
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Route struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Source      Location `json:"source"`
	Destination Location `json:"destination"`
}
