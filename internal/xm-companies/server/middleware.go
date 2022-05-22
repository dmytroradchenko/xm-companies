package server

import (
	"errors"
	"net/http"
)

var InvalidTokenError = errors.New("invalid or expired token")

func (s *Server) authenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("auth-token")
		if ok, _ := s.auth.ValidateToken(token); ok {
			next.ServeHTTP(w, r)
		} else {
			s.respond(w, http.StatusUnauthorized, InvalidTokenError)
		}
	})
}

func (s *Server) loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.logger.Info(r)
		next.ServeHTTP(w, r)
	})
}

func (s *Server) errorLoggerMiddleware(handler HandlerFunc) http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		if err := handler(responseWriter, request); err != nil {
			s.logger.Error(err)
		}
	})
}
