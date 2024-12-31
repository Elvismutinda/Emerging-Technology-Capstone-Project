package user

import (
	"backend/config"
	"backend/handlers/users"
	"backend/utils/middlewares"
	"github.com/gorilla/mux"
	"net/http"
)

type Router struct {
	Router *mux.Router
}

func RegisterRoutes(cfg *config.LocalConfig, r *Router) {
	userRoutes := r.Router.PathPrefix("/user").Subrouter()
	userRoutes.HandleFunc("/signup", users.CreateUserHandler).Methods(http.MethodPost)

	authUserRoutes := r.Router.PathPrefix("/user").Subrouter()
	authUserRoutes.Use(middlewares.AuthenticateUserMiddleware(cfg))
	authUserRoutes.HandleFunc("/{userId}", users.GetUserHandler).Methods(http.MethodGet)
	authUserRoutes.HandleFunc("/{userId}", users.UpdateUserHandler).Methods(http.MethodPatch)
}
