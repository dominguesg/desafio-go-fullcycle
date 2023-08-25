package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dominguesg/desafio-go-fullcycle/api/database"
	"github.com/dominguesg/desafio-go-fullcycle/api/models"
	"github.com/go-chi/chi"
)

func main() {
	err := database.InitDB("root:rootpassword@tcp(mysql:3306)/mydb")
	if err != nil {
		log.Fatal(err)
	}
	defer database.CloseDB()

	r := chi.NewRouter()
	r.Post("/api/routes", createRoute)
	r.Get("/api/routes", listRoutes)

	http.ListenAndServe(":8080", r)
}

func createRoute(w http.ResponseWriter, r *http.Request) {
	var route models.Route
	err := json.NewDecoder(r.Body).Decode(&route)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	insertID, err := database.InsertRoute(route)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	route.ID = insertID

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(route)
}

func listRoutes(w http.ResponseWriter, r *http.Request) {
	routes, err := database.FetchRoutes()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(routes)
}
