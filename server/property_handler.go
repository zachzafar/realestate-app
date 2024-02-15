package server

import (
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
		s.logger.Error(err, " error on line 18 property_handler.go")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	files := r.MultipartForm.File["images"]

	id, err := s.db.CreateProperty(property)

	if len(files) > 0 {
		err = s.db.UploadPropertyPhotos(files, id)
	}

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/admin/", http.StatusSeeOther)
}

func (s *Server) UpdateProperty(w http.ResponseWriter, r *http.Request) {
	property, err := types.ParsePropertyBody(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	files := r.MultipartForm.File["images"]

	err = s.db.UpdateProperty(property, id)

	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(files) > 0 {
		err = s.db.UploadPropertyPhotos(files, id)
	}

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/admin/listings/%d", id), http.StatusSeeOther)
}

func (s *Server) DeleteProperty(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = s.db.DeleteProperty(id)
}

func (s *Server) SearchProperties(w http.ResponseWriter, r *http.Request) {

	route := fmt.Sprintf("/homes?page=1&pageSize=10&address=%s", r.URL.Query().Get("address"))

	http.Redirect(w, r, route, http.StatusSeeOther)
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

	properties, err := s.db.GetProperties(propertyFilter, page, 10)

	if err != nil {
		s.logger.Error(err, " Line 79 view_handler.go")

	}

	if r.URL.Query().Get("format") == "admin" {
		template.AdminListings(properties, hasNextPage, nextPage).Render(r.Context(), w)
	} else {
		template.Listings(properties, hasNextPage, nextPage).Render(r.Context(), w)
	}

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

func (s *Server) GetAdminListingDetails(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	property, err := s.db.GetPropertyDetails(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	ListingDetails := template.EditPropertyDetails(*property, id)
	template.NewAdminLayout(ListingDetails).Render(r.Context(), w)
	return

}
