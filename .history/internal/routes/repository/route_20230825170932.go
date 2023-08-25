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
	sql := "INSERT INTO routes (id, name, source, destination) VALUES (?, ?, ?, ?)"
	rows, err := r.db.Query(sql, route.ID, route.name, route.Source, route.Destination)
	return err
}
