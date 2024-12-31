package routes

import (
	"backend/config"
	"backend/routes/auth"
	"backend/routes/user"
	"github.com/gorilla/mux"
)

type Router struct {
	Router *mux.Router
}

func NewRouter() *Router {
	return &Router{Router: mux.NewRouter()}
}

func (r *Router) InitializeRoutes(cfg *config.LocalConfig) {
	// initialize routes
	auth.RegisterRoutes(cfg, (*auth.Router)(r))
	user.RegisterRoutes(cfg, (*user.Router)(r))
}
