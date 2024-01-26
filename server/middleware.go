package server

import (
	"context"
	"net/http"

	"openlettings.com/types"
)

func (s *Server) AuthorizeUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		sessionCookie, err := r.Cookie("session-id")
		if err != nil {

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sessionID := sessionCookie.Value

		sessionData, err := s.db.GetSessionData(sessionID)
		if err != nil {
			s.logger.Error(err.Error(), " error on Line 24 middleware.go")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		ctx := context.WithValue(r.Context(), "user-id", sessionData)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func (s *Server) PassDataToCTX(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

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

		ctx = context.WithValue(r.Context(), countryKey, countries)
		ctx = context.WithValue(r.Context(), propertyKey, property_types)

		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
