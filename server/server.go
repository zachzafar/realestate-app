package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"openlettings.com/db"
)

type Server struct {
	listenAddress string
	db            *db.Database
}

func NewServer(addr string, db *db.Database) *Server {
	return &Server{
		listenAddress: addr,
		db:            db,
	}
}

func (s *Server) Start() error {

	r := mux.NewRouter()

	s.HandleRoutes(r)
	fmt.Println("Starting server....")

	return http.ListenAndServe(s.listenAddress, r)
}
