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
	_, err := r.db.Query(sql, route.ID, route.Name, route.Source, route.Destination)
	if err != nil {
		return err
	}
	return nil
}

func (r *RouteRepositoryMySQL) FindByID(id string) (*entity.Route, error) {
	var route entity.Route
	sql := "SELECT id, name, source, destination FROM routes WHERE id = ?"
	err := r.db.QueryRow(sql, id).Scan(&route.ID, &route.Name, &route.Source, &route.Destination)
	if err != nil {
		return nil, err
	}
	return &route, nil
}

func (r *RouteRepositoryMySQL) FindAll() ([]*entity.Route, error) {
	var routes []*entity.Route
	sql := "SELECT id, name, source, destination FROM routes"
	rows, err := r.db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var route entity.Route
		err := rows.Scan(&route.ID, &route.Name, &route.Source, &route.Destination)
		if err != nil {
			return nil, err
		}
		routes = append(routes, &route)
	}
	return routes, nil
}


