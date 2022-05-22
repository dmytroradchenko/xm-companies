package server

import (
	"encoding/json"
	"errors"
	"net/http"
	"xm-companies/internal/xm-companies/model"
)

type HandlerFunc func(responseWriter http.ResponseWriter, request *http.Request) error

var CodeNotSpecifiedError = errors.New("field 'code' is mandatory")
var DownstreamError = errors.New("downstream error")

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
