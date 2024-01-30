package server

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"openlettings.com/template"
	"openlettings.com/types"
)

func (s *Server) CreateProperty(w http.ResponseWriter, r *http.Request) {
	property, err := types.ParsePropertyBody(r)

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	files := r.MultipartForm.File["images"]

	id, err := s.db.CreateProperty(property)
	err = s.db.UploadPropertyPhotos(files, id)

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/listings", http.StatusSeeOther)
}

func (s *Server) UpdateProperty(w http.ResponseWriter, r *http.Request) {
	property, err := types.ParsePropertyBody(r)

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	files := r.MultipartForm.File["images"]

	id, err := s.db.CreateProperty(property)
	err = s.db.UploadPropertyPhotos(files, id)

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/listings", http.StatusSeeOther)
}

func (s *Server) DeleteProperty(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) SearchProperties(w http.ResponseWriter, r *http.Request) {

	route := fmt.Sprintf("/homes?page=1&pageSize=10&address=%s", r.URL.Query().Get("address"))

	http.Redirect(w, r, route, http.StatusSeeOther)
}

func (s *Server) GetListings(w http.ResponseWriter, r *http.Request) {
	// isAdminRoute, _ := regexp.MatchString("admin/+", fmt.Sprint(r.URL))

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

	template.Listings(properties, hasNextPage, nextPage).Render(ctx, w)

}

func (s *Server) GetListingDetails(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	property, err := s.db.GetPropertyDetails(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	ListingDetails := template.ListingDetails(*property)
	template.MainLayout(ListingDetails).Render(r.Context(), w)
	return

}
