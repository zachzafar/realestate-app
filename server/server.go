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
	InfoStore     db.InfoStore
}

func NewServer(addr string, db *db.Database, logger *logrus.Logger, InfoStore db.InfoStore) *Server {
	return &Server{
		listenAddress: addr,
		db:            db,
		logger:        logger,
		InfoStore:     InfoStore,
	}
}

func (s *Server) Start() error {

	r := mux.NewRouter()

	s.HandleRoutes(r)
	fmt.Println("Starting server....")

	return http.ListenAndServe(s.listenAddress, r)
}
