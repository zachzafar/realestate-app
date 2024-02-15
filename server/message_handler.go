package server

import (
	"net/http"

	"openlettings.com/types"
)

func (s *Server) CreateMessage(w http.ResponseWriter, r *http.Request) {
	message, err := types.ParseMessageBody(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = s.db.CreateMessage(message)

	if err != nil {

	}
}

func (s *Server) GetMessageDetails(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) GetMessages(w http.ResponseWriter, r *http.Request) {

}
