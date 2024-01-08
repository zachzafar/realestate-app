package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) HandleRoutes(r *mux.Router) {
	fs := http.FileServer(http.Dir("./svelte/public/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	ar := r.PathPrefix("/admin").Subrouter()
	ar.Use(s.AuthorizeUser)
	ar.HandleFunc("/listings/{id}", s.GetListingDetails).Methods("GET")
	ar.HandleFunc("/listings", s.GetAdminListings)
	ar.HandleFunc("/newProperty", s.GetNewPropertyForm).Methods("GET")
	ar.HandleFunc("/create-property/", s.CreateProperty).Methods("POST")
	// views
	r.HandleFunc("/", s.GetHomePage).Methods("GET")
	r.HandleFunc("/search-properties", s.SearchProperties).Methods("GET")
	r.HandleFunc("/listings", s.GetListings).Methods("GET")
	r.HandleFunc("/listings/{id}", s.GetListingDetails)
	r.HandleFunc("/login", s.GetLogin).Methods("GET")
	r.HandleFunc("/register", s.GetRegistration).Methods("GET")

	// actions
	r.HandleFunc("/create-user/", s.CreateUser).Methods("POST")
	r.HandleFunc("/auth-user/", s.LoginUser).Methods("POST")
}