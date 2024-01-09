package server

import (
	"net/http"

	"openlettings.com/types"
	"openlettings.com/utils"
)

func (s *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")
	role := r.PostFormValue("role")

	_, err := s.db.GetUser(email)

	if err.Error() != "sql: no rows in result set" {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	hashedPassword, err := utils.GenerateHash(password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newUser := types.NewUser(0, username, hashedPassword, role, email)

	err = s.db.CreateUser(newUser)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)

}

func (s *Server) LoginUser(w http.ResponseWriter, r *http.Request) {

	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	user, err := s.db.GetUser(email)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if user == nil {
		http.Error(w, "Sorry those credentials do not match", http.StatusBadRequest)
		return
	}
	isPassword := utils.CheckPassword(password, user.Password_hash)

	if !isPassword {
		http.Error(w, "Sorry those credentials do not match", http.StatusBadRequest)
		return
	}

	err = s.db.CreateSession(w, r, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	http.Redirect(w, r, "/admin/listings", http.StatusFound)
}
