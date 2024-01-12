package server

import (
	"fmt"
	"net/http"

	"openlettings.com/types"
)

func (s *Server) CreateProperty(w http.ResponseWriter, r *http.Request) {
	property, err := types.ParsePropertyBody(r)

	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	files := r.MultipartForm.File["images"]

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

	route := fmt.Sprintf("/listings?page=1&pageSize=10&address=%s", r.URL.Query().Get("address"))
	fmt.Println(route)
	http.Redirect(w, r, route, http.StatusSeeOther)
}
