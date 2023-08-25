package repository

import (
	"database/sql"

	"github.com/dominguesg/desafio-go-fullcycle/internal/routes/entity"
)

type RouteRepositoryMySQL struct {
	db *sql.DB
}

func NewRouteRepositoryMySQL(db *sql.DB) *RouteRepositoryMySQL {
	return &RouteRepositoryMySQL{
		db: db,
	}
}

func (r *RouteRepositoryMySQL) Create(route *entity.Route) error {
	
	_, err := r.db.Exec("INSERT INTO routes (id, name, source, destination) VALUES (?, ?, ?, ?)", route.ID, route.name, route.source, route.destination)
	return err
}
