package user

import (
	"backend/config"
	"backend/handlers"
	"backend/utils/middlewares"
	"github.com/gorilla/mux"
	"net/http"
)

type Router struct {
	Router *mux.Router
}

func RegisterRoutes(cfg *config.LocalConfig, r *Router) {
	userRoutes := r.Router.PathPrefix("/user").Subrouter()
	userRoutes.HandleFunc("/signup", handlers.CreateUserHandler).Methods(http.MethodPost)

	authUserRoutes := r.Router.PathPrefix("/user").Subrouter()
	authUserRoutes.Use(middlewares.AuthenticateUserMiddleware(cfg))
	authUserRoutes.HandleFunc("/{userId}", handlers.GetUserHandler).Methods(http.MethodGet)

}
