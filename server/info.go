package server

import (
	"net/http"
	"strconv"

	"openlettings.com/template"
)

func (s *Server) GetParishes(w http.ResponseWriter, r *http.Request) {
	countryID, err := strconv.Atoi(r.URL.Query().Get("country"))

	if err != nil {
		s.logger.Error(err, " error occurred while converting id to int line 23 countries.go")
	}

	parishes, err := s.db.GetParishes(countryID)

	if err != nil {
		s.logger.Error(err, "error occured while fethcing parishes on Line 21 info.go")
	}

	template.Options(parishes).Render(r.Context(), w)
}
