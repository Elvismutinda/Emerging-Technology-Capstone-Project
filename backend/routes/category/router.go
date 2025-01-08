package category

import (
	"backend/config"
	"backend/handlers/categories"
	"backend/utils/middlewares"
	"github.com/gorilla/mux"
	"net/http"
)

type Router struct {
	Router *mux.Router
}

func RegisterRoutes(cfg *config.LocalConfig, r *Router) {
	categoryRoutes := r.Router.PathPrefix("/category").Subrouter()
	categoryRoutes.Use(middlewares.AuthenticateUserMiddleware(cfg))

	categoryRoutes.HandleFunc("", categories.CreateCategoryHandler).Methods(http.MethodPost)
	categoryRoutes.HandleFunc("", categories.GetPaginatedCategoriesHandler).Methods(http.MethodGet)
	categoryRoutes.HandleFunc("/{categoryId}", categories.GetCategoryByIdHandler).Methods(http.MethodGet)
	categoryRoutes.HandleFunc("/name/{categoryName}", categories.GetCategoryByNameHandler).Methods(http.MethodGet)
	categoryRoutes.HandleFunc("/type/{categoryType}", categories.GetPaginatedCategoriesByCategoryTypeHandler).Methods(http.MethodGet)
	categoryRoutes.HandleFunc("/{categoryId}", categories.UpdateCategoryHandler).Methods(http.MethodPatch)
	categoryRoutes.HandleFunc("/{categoryId}", categories.DeleteCategoryHandler).Methods(http.MethodDelete)

}
