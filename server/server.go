package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"openlettings.com/db"
)

type Server struct {
	listenAddress string
	db            *db.Database
	logger        *logrus.Logger
}

func NewServer(addr string, db *db.Database, logger *logrus.Logger) *Server {
	return &Server{
		listenAddress: addr,
		db:            db,
		logger:        logger,
	}
}

func (s *Server) Start() error {

	r := mux.NewRouter()

	s.HandleRoutes(r)
	fmt.Println("Starting server....")

	return http.ListenAndServe(s.listenAddress, r)
}
