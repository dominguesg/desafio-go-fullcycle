package usecase

type CreateRouteInput struct {
	ID		  string
	Name	  string
	Source	  string
	Destination string
}

type CreateRouteOutput struct {
	ID string
	Name string
	