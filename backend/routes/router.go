package routes

import (
	"backend/config"
	"backend/routes/auth"
	"backend/routes/category"
	"backend/routes/transaction"
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
	transaction.RegisterRoutes(cfg, (*transaction.Router)(r))
	category.RegisterRoutes(cfg, (*category.Router)(r))
}
