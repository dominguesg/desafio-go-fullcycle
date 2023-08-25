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
	Route      *entity.Route
}

func NewCreateRouteUseCase(repository entity.RouteRepository) *CreateRouteUseCase {
	return &CreateRouteUseCase{
		Repository: repository,
		
	}
}