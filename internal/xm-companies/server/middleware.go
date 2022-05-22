package server

import "net/http"

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
