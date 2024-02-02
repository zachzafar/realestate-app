package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) HandleRoutes(r *mux.Router) {
	r.Use(s.SerialiseUser)
	fs := http.FileServer(http.Dir("./svelte/public/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	is := http.FileServer(http.Dir("./media/properties/"))
	r.PathPrefix("/images/").Handler(http.StripPrefix("/images/", is))

	ar := r.PathPrefix("/admin").Subrouter()
	ar.Use(s.AuthorizeUser)
	// views
	ar.HandleFunc("/listings/{id}", s.GetAdminListingDetails)
	ar.HandleFunc("/", s.GetAdminPropertiesPage)

	//actions
	ar.HandleFunc("/create-property/", s.CreateProperty).Methods("POST")
	ar.HandleFunc("/delete-property/{id}", s.DeleteProperty).Methods("DELETE")
	ar.HandleFunc("/update-property/{id}", s.UpdateProperty).Methods("PUT")
	ar.HandleFunc("/logout", s.Logout)

	// views
	r.HandleFunc("/", s.GetHomePage).Methods("GET")
	r.HandleFunc("/search-properties", s.SearchProperties).Methods("GET")

	r.HandleFunc("/homes", s.GetListingsPage).Methods("GET")
	r.HandleFunc("/listings", s.GetListings).Methods("GET")
	r.HandleFunc("/listings/{id}", s.GetListingDetails)
	r.HandleFunc("/login", s.GetLogin).Methods("GET")
	r.HandleFunc("/register", s.GetRegistration).Methods("GET")

	// actions
	r.HandleFunc("/create-user/", s.CreateUser).Methods("POST")
	r.HandleFunc("/auth-user/", s.LoginUser).Methods("POST")

	r.HandleFunc("/parishes", s.GetParishes).Methods("GET")

}
