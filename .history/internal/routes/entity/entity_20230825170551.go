package entity

type RouteRepository interface {
	Create

type Route struct {
	ID          string
	name        string
	source      string
	destination string
}

func NewRoute(id, name, source, destination string) *Route {
	return &Route{
		ID:          id,
		name:        name,
		source:      source,
		destination: destination,
	}
}

