package auth

import (
	"backend/config"
	"backend/handlers/auth"
	"github.com/gorilla/mux"
)

type Router struct {
	Router *mux.Router
}

func RegisterRoutes(cfg *config.LocalConfig, r *Router) {
	authRoutes := r.Router.PathPrefix("/auth").Subrouter()
	authRoutes.HandleFunc("/login", auth.AuthLoginHandler).Methods("POST")
}
