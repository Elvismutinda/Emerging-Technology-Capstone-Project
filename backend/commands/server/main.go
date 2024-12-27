package main

import (
	"backend/config"
	"backend/routes"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
)

type Server struct {
	Configuration *config.LocalConfig
	Router        *routes.Router
}

// New Router
func NewRouter() *routes.Router {
	return &routes.Router{Router: mux.NewRouter()}
}

// NewServer ...
func NewServer(config *config.LocalConfig) *Server {
	server := &Server{
		Configuration: config,
		Router:        NewRouter(),
	}

	return server
}

func main() {
	// connect DB
	config.ConnectDB()

	// ping DB
	db, err := config.DB.DB()
	if err != nil {
		log.Fatal("Error getting database instance:", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	// Continue with other setups like routes
	localCfg, err := config.FromEnv()
	if err != nil {
		logrus.Error(err)
		return
	}
	server := NewServer(localCfg)
	server.Router.RegisterRoutes(localCfg)

	var handler http.Handler
	handler = server.Router.Router

	logrus.Infoln("Starting server on Port", localCfg.Port)
	err = http.ListenAndServe(fmt.Sprintf("%v:%v", "", localCfg.Port),
		handler)
	if err != nil {
		logrus.Error(err)
		return
	}

}
