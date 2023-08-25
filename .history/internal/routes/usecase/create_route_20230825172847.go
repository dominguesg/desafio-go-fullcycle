package usecase

import "github.com/dominguesg/desafio-go-fullcycle/internal/routes/entity"

type CreateRouteInput struct {
	ID          string
	Name        string
	Source      string
	Destination string
}

type CreateRouteOutput struct {
	ID          string
	Name        string
	Source      string
	Destination string
}

type CreateRouteUseCase struct {
	Repository entity.RouteRepository
	Route      entity.Route
}

func NewCreateRouteUseCase(repository entity.RouteRepository, route entity.Route) *CreateRouteUseCase {
	return &CreateRouteUseCase{
		Repository: repository,
		Route:      route,
	}
}

func (u *CreateRouteInput) Execute() (input CreateRouteOutput, error) {
	route := entity.NewRoute(u.ID, u.Name, u.Source, u.Destination)
	err := u.Repository.Create(route)
	if err != nil {
		return nil, err
	}
	return &CreateRouteOutput{
		ID:          route.ID,
		Name:        route.Name,
		Source:      route.Source,
		Destination: route.Destination,
	}, nil
}
