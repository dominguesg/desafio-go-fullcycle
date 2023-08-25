package models

type Route struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	Source      JSONString `json:"source"`
	Destination JSONString `json:"destination"`
}
