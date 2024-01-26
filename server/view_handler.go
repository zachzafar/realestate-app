package server

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"openlettings.com/db"
	"openlettings.com/template"
	"openlettings.com/types"
)

func (s *Server) GetLogin(w http.ResponseWriter, r *http.Request) {
	Login := template.Login()
	template.MainLayout(Login).Render(r.Context(), w)
}

func (s *Server) GetHomePage(w http.ResponseWriter, r *http.Request) {
	Home := template.Home()
	template.MainLayout(Home).Render(r.Context(), w)
}

func (s *Server) GetRegistration(w http.ResponseWriter, r *http.Request) {
	Register := template.RegisterPage()
	template.MainLayout(Register).Render(r.Context(), w)
}

func (s *Server) GetAdminListings(w http.ResponseWriter, r *http.Request) {
	userData, ok := r.Context().Value("user-id").(*db.SessionData)
	if !ok {
		fmt.Print("not okay")
		http.Error(w, "Unauthorized", http.StatusForbidden)
		return
	}
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))

	emptyPriceRange := types.NewRange(0, 0)

	propertyFilter := types.NewPropertyFilter(0, 0, 0, *emptyPriceRange, 0, 0, userData.UserId)
	properties, err := s.db.GetProperties(propertyFilter, page, pageSize)

	if err != nil {
		s.logger.Error(err, " error at line 44 view_handler.go")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	Listings := template.MyListingsPage(properties)
	template.AdminLayout(Listings).Render(r.Context(), w)
}

func (s *Server) GetListings(w http.ResponseWriter, r *http.Request) {

	propertyFilter, page := types.ParseListingParams(r)

	property_count, err := s.db.GetPropertyCount(propertyFilter)

	if err != nil {
		s.logger.Error(err, " Line 59 view_handle.go")
	}

	hasNextPage := false
	nextPage := ""

	if property_count-(page*10) > 0 {
		hasNextPage = true

		values := r.URL.Query()
		values.Set("page", fmt.Sprint(page+1))
		r.URL.RawQuery = values.Encode()
		nextPage = r.URL.String()

	}

	var countryKey types.ContextKey = "countries"
	var propertyKey types.ContextKey = "property_types"

	var ctx context.Context
	countries, err := s.db.GetAllCountries()
	if err != nil {
		s.logger.Error(err, " line 32 middleware.go")
	}
	property_types, err := s.db.GetAllPropertyTypes()
	if err != nil {
		s.logger.Error(err, " line 36 middleware.go")
	}

	ctx = context.WithValue(context.Background(), countryKey, countries)
	ctx = context.WithValue(ctx, propertyKey, property_types)

	properties, err := s.db.GetProperties(propertyFilter, page, 10)

	if err != nil {
		s.logger.Error(err, " Line 79 view_handler.go")

	}

	if page == 1 {

		Listings := template.ListingsPage(*propertyFilter, properties, hasNextPage, nextPage)
		template.MainLayout(Listings).Render(ctx, w)
		return
	}

	template.Listings(properties, hasNextPage, nextPage).Render(ctx, w)

}

func (s *Server) GetListingDetails(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user-id")
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	property, err := s.db.GetPropertyDetails(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if userId == nil {
		ListingDetails := template.ListingDetails(*property)
		template.MainLayout(ListingDetails).Render(r.Context(), w)
		return
	}

	template.AdminListingDetails().Render(r.Context(), w)

}

func (s *Server) GetNewPropertyForm(w http.ResponseWriter, r *http.Request) {
	CreateNewPropertyForm := template.NewPropertyForm()
	template.AdminLayout(CreateNewPropertyForm).Render(r.Context(), w)
}
