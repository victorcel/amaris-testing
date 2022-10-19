package routers

import (
	"github.com/gorilla/mux"
	"github.com/victorcel/amaris-testing/handlers"
	"github.com/victorcel/amaris-testing/server"
)

func BindRoutes(s server.Server, r *mux.Router) {

	routeApi := r.PathPrefix("/api/v1/").Subrouter()

	routeApi.HandleFunc("/orderText", handlers.OrderText()).Methods("POST")
	routeApi.HandleFunc("/getPokemon/{id}", handlers.SearchPokemon()).Methods("GET")

}
