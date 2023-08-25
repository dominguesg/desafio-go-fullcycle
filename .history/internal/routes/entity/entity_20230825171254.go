package entity

type RouteRepository interface {
	Create(route *Route) error
	FindByID(id string) (*Route, error)
	FindAll() ([]*Route, error)
}

type Route struct {
	ID          string
	Name        string
	Source      string
	Destination string
}

func NewRoute(id, name, source, destination string) *Route {
	return &Route{
		ID:          id,
		Name:        name,
		source:      source,
		destination: destination,
	}
}
