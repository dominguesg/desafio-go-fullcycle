package entity

type RouteRepository interface {
	Create(route *Route) error
	FindByID(id string) (*Route, error)
	FindAll() ([]*Route, error)
}

type Route struct {
	ID          string
	Name        string
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
