package server

import (
	"encoding/json"
	"net/http"
	"xm-companies/config"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"xm-companies/internal/xm-companies/security/otp"
	"xm-companies/internal/xm-companies/store"
)

type Server struct {
	store  store.Store
	auth   otp.Service
	logger *logrus.Logger
	router *mux.Router
	cfg    *config.Config
}

type ErrorBody struct {
	Message string `json:"message"`
}

func NewServerProvider(cfg *config.Config, store store.Store, auth otp.Service) *Server {
	s := &Server{
		store:  store,
		auth:   auth,
		logger: logrus.New(),
		router: mux.NewRouter(),
		cfg:    cfg,
	}
	s.configureRouter()
	return s
}

func (s *Server) StartServer() error {
	s.logger.Info("Server is starting on port: " + s.cfg.Port)
	return http.ListenAndServe(":"+s.cfg.Port, s)
}

func (s *Server) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(rw, r)
}

func (s *Server) configureRouter() {
	s.router.Handle(
		"/search",
		s.loggerMiddleware(s.errorLoggerMiddleware(s.GetCompanies()))).
		Methods(http.MethodPost)
	s.router.Handle(
		"/company",
		s.loggerMiddleware(s.authenticationMiddleware(s.errorLoggerMiddleware(s.CreateOrUpdateCompany())))).
		Methods(http.MethodPost)
	s.router.Handle(
		"/company",
		s.loggerMiddleware(s.authenticationMiddleware(s.errorLoggerMiddleware(s.DeleteCompany())))).
		Methods(http.MethodDelete)
	s.router.Handle(
		"/sign-up",
		s.loggerMiddleware(s.errorLoggerMiddleware(s.SignUp()))).
		Methods(http.MethodPost)
	s.router.Handle(
		"/sign-in",
		s.loggerMiddleware(s.errorLoggerMiddleware(s.SignIn()))).
		Methods(http.MethodPost)
}

func (s *Server) respond(rw http.ResponseWriter, code int, data interface{}) {
	rw.WriteHeader(code)
	if data != nil {
		if code != http.StatusOK {
			data = ErrorBody{Message: data.(error).Error()}
		}
		if err := json.NewEncoder(rw).Encode(data); err != nil {
			s.logger.Error(err)
		}
	}
}
