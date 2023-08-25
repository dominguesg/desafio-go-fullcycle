// database/routes_db.go
package database

import (
	"database/sql"

	"github.com/dominguesg/desafio-go-fulcycle"
	"github.com/dominguesg/desafio-go-fullcycle/api/models"
)

var db *sql.DB

func InitDB(databaseURL string) error {
	var err error
	db, err = sql.Open("mysql", databaseURL)
	if err != nil {
		return err
	}
	return nil
}

func InsertRoute(route models.Route) (int, error) {
	query := "INSERT INTO routes (name, source, destination) VALUES (?, ?, ?)"
	result, err := db.Exec(query, route.Name, route.Source, route.Destination)
	if err != nil {
		return 0, err
	}

	insertID, _ := result.LastInsertId()
	return int(insertID), nil
}

func FetchRoutes() ([]models.Route, error) {
	rows, err := db.Query("SELECT id, name, source, destination FROM routes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var routes []models.Route
	for rows.Next() {
		var route models.Route
		err := rows.Scan(&route.ID, &route.Name, &route.Source, &route.Destination)
		if err != nil {
			return nil, err
		}
		routes = append(routes, route)
	}

	return routes, nil
}
