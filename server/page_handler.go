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

	ctx = context.WithValue(context.Background(), countryKey, s.InfoStore.Countries)
	ctx = context.WithValue(ctx, propertyKey, s.InfoStore.Property_types)

	Listings := template.ListingsPage(*propertyFilter, r.URL.RawQuery)
	template.MainLayout(Listings).Render(ctx, w)
}

func (s *Server) GetAdminPropertiesPage(w http.ResponseWriter, r *http.Request) {
	var countryKey types.ContextKey = "countries"
	var propertyKey types.ContextKey = "property_types"

	var ctx context.Context

	ctx = context.WithValue(context.Background(), countryKey, s.InfoStore.Countries)
	ctx = context.WithValue(ctx, propertyKey, s.InfoStore.Property_types)

	Properties := template.AdminMainProperties()
	template.NewAdminLayout(Properties).Render(ctx, w)
}

func (s *Server) GetAdminMessagesPage(w http.ResponseWriter, r *http.Request) {

}
