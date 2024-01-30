package server

import (
	"context"
	"net/http"

	"openlettings.com/template"
	"openlettings.com/types"
)

func (s *Server) GetHomePage(w http.ResponseWriter, r *http.Request) {
	Home := template.Home()
	template.MainLayout(Home).Render(r.Context(), w)
}

func (s *Server) GetListingsPage(w http.ResponseWriter, r *http.Request) {
	propertyFilter, _ := types.ParseListingParams(r)

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

	Listings := template.ListingsPage(*propertyFilter, r.URL.RawQuery)
	template.MainLayout(Listings).Render(ctx, w)
}

func (s *Server) GetAdminPropertiesPage(w http.ResponseWriter, r *http.Request) {
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

	Properties := template.AdminMainProperties()
	template.NewAdminLayout(Properties).Render(ctx, w)
}

func (s *Server) GetAdminMessagesPage(w http.ResponseWriter, r *http.Request) {

}

// func (s *Server) GetAdminListings(w http.ResponseWriter, r *http.Request) {
// 	userData, ok := r.Context().Value("user-id").(*types.SessionData)
// 	if !ok {
// 		fmt.Print("not okay")
// 		http.Error(w, "Unauthorized", http.StatusForbidden)
// 		return
// 	}
// 	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
// 	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))

// 	emptyPriceRange := types.NewRange(0, 0)

// 	propertyFilter := types.NewPropertyFilter(0, 0, 0, *emptyPriceRange, 0, 0, userData.UserId)
// 	properties, err := s.db.GetProperties(propertyFilter, page, pageSize)

// 	if err != nil {
// 		s.logger.Error(err, " error at line 44 view_handler.go")
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	Listings := template.MyListingsPage(properties)
// 	template.AdminLayout(Listings).Render(r.Context(), w)
// }
