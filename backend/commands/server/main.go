package main

import (
	"backend/config"
	"backend/routes"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Server struct {
	Configuration *config.LocalConfig
	Router        *routes.Router
}

// NewServer ...
func NewServer(config *config.LocalConfig) *Server {
	server := &Server{
		Configuration: config,
		Router:        routes.NewRouter(),
	}

	return server
}

func main() {
	// connect DB
	_, err := config.ConnectDB()
	if err != nil {
		logrus.Fatal(err)
		return
	}
	logrus.Infoln("Connected to database")

	// Continue with other setups like routes
	localCfg, err := config.FromEnv()
	if err != nil {
		logrus.Error(err)
		return
	}
	server := NewServer(localCfg)
	server.Router.InitializeRoutes(localCfg)

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
