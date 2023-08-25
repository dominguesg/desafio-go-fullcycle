package repository

import "database/sql"

type RouteRepositoryMySQL struct {
	db *sql.DB
}

func NewRouteRepositoryMySQL(db *sql.DB) *RouteRepositoryMySQL {
	return &RouteRepositoryMySQL{
		db: db,
	}
}

func 
