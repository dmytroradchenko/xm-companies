package server

import (
	"github.com/stretchr/testify/mock"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"xm-companies/internal/xm-companies/model"
	"xm-companies/internal/xm-companies/store/mocks"
)

func TestServer_GetCompanies(t *testing.T) {
	payload := `{"name":"Test"}`
	req, err := http.NewRequest(http.MethodPost, "/search", strings.NewReader(payload))
	if err != nil {
		t.Fatal(err)
	}

	store := new(mocks.Store)
	companies := new(mocks.CompaniesRepository)
	store.On("Companies").Return(companies)
	companies.On("FindBy", mock.Anything, mock.Anything).
		Return([]*model.Company{{Name: "Test", Code: "1"}}, nil).Once()

	s := Server{
		store: store,
	}

	rr := httptest.NewRecorder()
	handler := unwrapCustomHandler(s.GetCompanies(), nil, t)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned incorrect status code: returned %v expected %v",
			status, http.StatusOK)
	}

	expected := "[{\"code\":\"1\",\"name\":\"Test\",\"country\":\"\",\"phone\":\"\",\"website\":\"\"}]\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: returned %v expected %v",
			rr.Body.String(), expected)
	}
}

func TestServer_GetCompanies_ErrorWhenEmptyPayload(t *testing.T) {
	payload := ""
	req, err := http.NewRequest(http.MethodPost, "/search", strings.NewReader(payload))
	if err != nil {
		t.Fatal(err)
	}

	s := Server{}

	rr := httptest.NewRecorder()
	handler := unwrapCustomHandler(s.GetCompanies(), io.EOF, t)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned incorrect status code: returned %v expected %v",
			status, http.StatusBadRequest)
	}
}

func TestServer_CreateOrUpdateCompany_NewUser(t *testing.T) {
	payload := `{"name":"Test","code":"1"}`
	req, err := http.NewRequest(http.MethodPost, "/company", strings.NewReader(payload))
	if err != nil {
		t.Fatal(err)
	}

	store := new(mocks.Store)
	companies := new(mocks.CompaniesRepository)
	store.On("Companies").Return(companies)
	companies.On("Create", mock.Anything, mock.Anything).Return(nil).Once()
	companies.On("FindBy", mock.Anything, mock.Anything).Return([]*model.Company{}, nil).Once()

	s := Server{
		store: store,
	}

	rr := httptest.NewRecorder()
	handler := unwrapCustomHandler(s.CreateOrUpdateCompany(), nil, t)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned incorrect status code: returned %v expected %v",
			status, http.StatusOK)
	}
}

func TestServer_CreateOrUpdateCompany_UpdateUser(t *testing.T) {
	payload := `{"name":"Test","code":"1"}`
	req, err := http.NewRequest(http.MethodPost, "/company", strings.NewReader(payload))
	if err != nil {
		t.Fatal(err)
	}

	store := new(mocks.Store)
	companies := new(mocks.CompaniesRepository)
	store.On("Companies").Return(companies)
	companies.On("Update", mock.Anything, mock.Anything).
		Return(nil).Once()
	companies.On("FindBy", mock.Anything, mock.Anything).
		Return([]*model.Company{{Name: "test", Code: "1"}}, nil).Once()

	s := Server{
		store: store,
	}

	rr := httptest.NewRecorder()
	handler := unwrapCustomHandler(s.CreateOrUpdateCompany(), nil, t)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned incorrect status code: returned %v expected %v",
			status, http.StatusOK)
	}
}

func TestServer_CreateOrUpdateCompany_CodeNotSpecifiedError(t *testing.T) {
	payload := `{"name":"Test"}`
	req, err := http.NewRequest(http.MethodPost, "/company", strings.NewReader(payload))
	if err != nil {
		t.Fatal(err)
	}

	s := Server{}

	rr := httptest.NewRecorder()
	handler := unwrapCustomHandler(s.CreateOrUpdateCompany(), nil, t)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned incorrect status code: returned %v expected %v",
			status, http.StatusBadRequest)
	}
	expected := "{\"message\":\"field 'code' is mandatory\"}\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: returned %v expected %v",
			rr.Body.String(), expected)
	}
}

func TestServer_DeleteCompany(t *testing.T) {
	req, err := http.NewRequest(http.MethodDelete, "/company?code=1", nil)
	if err != nil {
		t.Fatal(err)
	}
	store := new(mocks.Store)
	companies := new(mocks.CompaniesRepository)
	store.On("Companies").Return(companies)
	companies.On("Delete", mock.Anything, "1").Return(nil).Once()

	s := Server{
		store: store,
	}

	rr := httptest.NewRecorder()
	handler := unwrapCustomHandler(s.DeleteCompany(), nil, t)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned incorrect status code: returned %v expected %v",
			status, http.StatusOK)
	}
}

func TestServer_DeleteCompany_CodeNotSpecifiedError(t *testing.T) {
	req, err := http.NewRequest(http.MethodDelete, "/company", nil)
	if err != nil {
		t.Fatal(err)
	}

	s := Server{}

	rr := httptest.NewRecorder()
	handler := unwrapCustomHandler(s.DeleteCompany(), nil, t)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned incorrect status code: returned %v expected %v",
			status, http.StatusOK)
	}

	expected := "{\"message\":\"field 'code' is mandatory\"}\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: returned %v expected %v",
			rr.Body.String(), expected)
	}
}

func unwrapCustomHandler(handler HandlerFunc, expectedError error, t *testing.T) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		err := handler(rw, r)
		if err != expectedError {
			t.Fatalf("Unexpected error. Expected %v, returned %v", expectedError, err)
		}
	}
}
