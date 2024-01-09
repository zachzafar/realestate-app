package server

import (
	"fmt"
	"net/http"
	"strconv"

	"openlettings.com/types"
)

func (s *Server) CreateProperty(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)

	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "unable to parse form", http.StatusBadRequest)
		return
	}

	userId := r.Context().Value("user-id").(int)
	title := r.PostFormValue("title")
	description := r.PostFormValue("description")
	property_type := r.PostFormValue("type")
	price, _ := strconv.ParseFloat(r.PostFormValue("price"), 64)
	address := r.PostFormValue("address")
	size, _ := strconv.Atoi(r.PostFormValue("size"))
	bedrooms, _ := strconv.Atoi(r.PostFormValue("bedrooms"))
	bathrooms, _ := strconv.Atoi(r.PostFormValue("bathrooms"))
	year := r.PostFormValue("year")
	city := r.PostFormValue("city")
	files := r.MultipartForm.File["images"]

	property := types.NewProperty(title, description, property_type, address, city, year, bathrooms, size, bedrooms, userId, price)

	id, err := s.db.CreateProperty(property)
	err = s.db.UploadPropertyPhotos(files, id)

	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/listings", http.StatusSeeOther)
}

func (s *Server) SearchProperties(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	address := r.PostFormValue("address")
	route := fmt.Sprintf("/listings?page=1&pageSize=10&address=%s", address)
	http.Redirect(w, r, route, http.StatusSeeOther)
}
