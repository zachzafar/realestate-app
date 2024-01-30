package server

import (
	"net/http"
	"time"

	"openlettings.com/template"
	"openlettings.com/types"
	"openlettings.com/utils"
)

func (s *Server) GetRegistration(w http.ResponseWriter, r *http.Request) {
	Register := template.RegisterPage()
	template.MainLayout(Register).Render(r.Context(), w)
}

func (s *Server) GetLogin(w http.ResponseWriter, r *http.Request) {
	Login := template.Login()
	template.MainLayout(Login).Render(r.Context(), w)
}

func (s *Server) CreateUser(w http.ResponseWriter, r *http.Request) {

	password := r.PostFormValue("password")
	email := r.PostFormValue("email")

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

	newUser := &types.User{Email: email, PasswordHash: hashedPassword}

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

	isPassword := utils.CheckPassword(password, user.PasswordHash)

	if !isPassword {
		http.Error(w, "Sorry those credentials do not match", http.StatusBadRequest)
		return
	}

	sessionID, err := s.db.CreateSession(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sessionCookie := http.Cookie{Name: "session-id", Path: "/", Value: sessionID, Expires: time.Now().Add(24 * time.Hour), HttpOnly: true}

	http.SetCookie(w, &sessionCookie)

	http.Redirect(w, r, "/admin/", http.StatusFound)
}

func (s *Server) Logout(w http.ResponseWriter, r *http.Request) {
	sessionCookie, err := r.Cookie("session-id")
	sessionID := sessionCookie.Value

	s.logger.Info(err, "  user doesn't have session cookie")

	if err := s.db.DeleteSession(sessionID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
