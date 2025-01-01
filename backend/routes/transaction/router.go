package transaction

import (
	"backend/config"
	"backend/handlers/transactions"
	"backend/utils/middlewares"
	"github.com/gorilla/mux"
	"net/http"
)

type Router struct {
	Router *mux.Router
}

func RegisterRoutes(cfg *config.LocalConfig, r *Router) {
	transactionRoutes := r.Router.PathPrefix("/transaction").Subrouter()
	transactionRoutes.Use(middlewares.AuthenticateUserMiddleware(cfg))

	transactionRoutes.HandleFunc("/create/{transactionType}", transactions.CreateTransactionHandler).Methods(http.MethodPost)

	transactionRoutes.HandleFunc("/get/{transactionId}", transactions.GetTransactionHandler).Methods(http.MethodGet)
	transactionRoutes.HandleFunc("/get-all", transactions.GetPaginatedTransactionsHandler).Methods(http.MethodGet)

	transactionRoutes.HandleFunc("/update/{transactionId}", transactions.UpdateTransactionHandler).Methods(http.MethodPatch)

	transactionRoutes.HandleFunc("/delete/{transactionId}", transactions.DeleteTransactionHandler).Methods(http.MethodDelete)

	// analytics
	transactionRoutes.HandleFunc("/get-overview", transactions.GetTransactionsOverviewHandler).Methods(http.MethodGet)
	transactionRoutes.HandleFunc("/get-category-stats/{categoryName}", transactions.GetTransactionStatsByCategoryHandler).Methods(http.MethodGet)

}
