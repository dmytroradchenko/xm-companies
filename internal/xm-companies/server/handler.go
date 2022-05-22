package server

import (
	"encoding/json"
	"errors"
	"net/http"
	"xm-companies/internal/xm-companies/model"
	"xm-companies/internal/xm-companies/util"
)

type HandlerFunc func(responseWriter http.ResponseWriter, request *http.Request) error

var CodeNotSpecifiedError = errors.New("field 'code' is mandatory")
var DownstreamError = errors.New("downstream error")
var SignUpValidationError = errors.New("'username' should by longer 4 and 'password' longer 6 symbols")
var InvalidCredentialsError = errors.New("invalid user credentials")

func (s *Server) SignUp() HandlerFunc {
	type request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	return func(rw http.ResponseWriter, r *http.Request) error {
		form := &request{}
		if err := json.NewDecoder(r.Body).Decode(form); err != nil {
			s.respond(rw, http.StatusBadRequest, err)
			return err
		}

		if len(form.Username) <= 4 || len(form.Password) <= 6 {
			s.respond(rw, http.StatusBadRequest, SignUpValidationError)
			return nil
		}

		if _, err := s.store.Users().Create(r.Context(), &model.User{
			Username: form.Username,
			Password: util.Hash(form.Password),
		}); err != nil {
			s.respond(rw, http.StatusInternalServerError, DownstreamError)
			return err
		}
		s.respond(rw, http.StatusOK, nil)
		return nil
	}
}

func (s *Server) SignIn() HandlerFunc {
	type request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	type response struct {
		Token string `json:"token"`
	}
	return func(rw http.ResponseWriter, r *http.Request) error {
		form := &request{}
		if err := json.NewDecoder(r.Body).Decode(form); err != nil {
			s.respond(rw, http.StatusBadRequest, err)
			return err
		}

		u, err := s.store.Users().Find(r.Context(), form.Username)
		if err != nil {
			s.respond(rw, http.StatusInternalServerError, nil)
			return err
		}
		if !util.CheckHash(form.Password, u.Password) {
			s.respond(rw, http.StatusUnauthorized, InvalidCredentialsError)
			return nil
		}
		s.respond(rw, http.StatusOK, &response{
			Token: s.auth.CreateToken(u.Username),
		})
		return nil
	}
}

func (s *Server) GetCompanies() HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) error {
		search := model.SearchFilter{}
		if err := json.NewDecoder(r.Body).Decode(&search); err != nil {
			s.respond(rw, http.StatusBadRequest, err)
			return err
		}

		result, err := s.store.Companies().FindBy(r.Context(), search)
		if err != nil {
			s.respond(rw, http.StatusInternalServerError, DownstreamError)
			return err
		}
		s.respond(rw, http.StatusOK, result)
		return nil
	}
}

func (s *Server) DeleteCompany() HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) error {
		code := r.URL.Query().Get("code")
		if code == "" {
			s.respond(rw, http.StatusBadRequest, CodeNotSpecifiedError)
			return nil
		}

		if err := s.store.Companies().Delete(r.Context(), code); err != nil {
			s.respond(rw, http.StatusInternalServerError, DownstreamError)
			return err
		}
		s.respond(rw, http.StatusOK, nil)
		return nil
	}
}

func (s *Server) CreateOrUpdateCompany() HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) error {
		comp := model.Company{}
		if err := json.NewDecoder(r.Body).Decode(&comp); err != nil {
			s.respond(rw, http.StatusBadRequest, err)
			return err
		}

		if comp.Code == "" {
			s.respond(rw, http.StatusBadRequest, CodeNotSpecifiedError)
			return nil
		}

		result, err := s.store.Companies().FindBy(r.Context(), model.SearchFilter{Code: comp.Code})
		if err != nil {
			s.respond(rw, http.StatusInternalServerError, DownstreamError)
			return err
		}

		if len(result) > 0 {
			err = s.store.Companies().Update(r.Context(), &comp)
		} else {
			err = s.store.Companies().Create(r.Context(), &comp)
		}

		if err != nil {
			s.respond(rw, http.StatusInternalServerError, DownstreamError)
			return err
		}
		s.respond(rw, http.StatusOK, "")
		return nil
	}
}
