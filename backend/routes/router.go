package routes

import (
	"backend/config"
	"backend/handlers"
	"github.com/gorilla/mux"
)

type Router struct {
	Router *mux.Router
}

func (r *Router) RegisterRoutes(cfg *config.LocalConfig) {
	userRoutes := r.Router.PathPrefix("/user").Subrouter()
	userRoutes.HandleFunc("/create", handlers.CreateUserHandler).Methods("POST")
}
