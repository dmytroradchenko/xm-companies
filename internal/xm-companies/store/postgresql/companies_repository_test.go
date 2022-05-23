package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"xm-companies/internal/xm-companies/model"
)

func TestCompaniesRepository_Create(t *testing.T) {
	db, mock, target := createMockedCompaniesRepository(t)
	defer db.Close()
	comp := &model.Company{
		Name:    "Test Company",
		Code:    "1234",
		Country: "Ukraine",
		Phone:   "0000",
		Website: "web.site",
	}

	mock.ExpectExec("INSERT INTO companies (.+)").
		WithArgs("1234", "Test Company", "Ukraine", "0000", "web.site").
		WillReturnResult(sqlmock.NewResult(0, 1))

	if err := target.Create(context.Background(), comp); err != nil {
		t.Errorf("returned error: %v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

func TestCompaniesRepository_Update(t *testing.T) {
	db, mock, target := createMockedCompaniesRepository(t)
	defer db.Close()
	comp := &model.Company{
		Name:    "Test Company",
		Code:    "1234",
		Country: "Ukraine",
		Phone:   "0000",
		Website: "web.site",
	}
	query := "UPDATE companies SET name=(.+), country=(.+), phone=(.+), website=(.+) WHERE code = (.+)"

	mock.ExpectExec(query).
		WithArgs("Test Company", "Ukraine", "0000", "web.site", "1234").
		WillReturnResult(sqlmock.NewResult(0, 1))

	if err := target.Update(context.Background(), comp); err != nil {
		t.Errorf("should update company but returns err: %v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

func TestCompaniesRepository_Delete(t *testing.T) {
	db, mock, target := createMockedCompaniesRepository(t)
	defer db.Close()

	mock.ExpectExec("DELETE FROM companies WHERE code = (.+)").
		WithArgs("test").
		WillReturnResult(sqlmock.NewResult(0, 1))

	if err := target.Delete(context.Background(), "test"); err != nil {
		t.Errorf("expected to remove company but returned error: %v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

func TestCompaniesRepository_FindBy_Name(t *testing.T) {
	expected := &model.Company{
		Name:    "Test",
		Code:    "1234",
		Country: "Ukraine",
		Phone:   "0000",
		Website: "web.site",
	}
	db, mock, target := createMockedCompaniesRepository(t)
	defer db.Close()

	columns := []string{"code", "name", "country", "phone", "website"}

	mock.ExpectQuery("SELECT (.+) FROM companies WHERE name LIKE (.+)").
		WithArgs("%Test%").
		WillReturnRows(sqlmock.NewRows(columns).AddRow("1234", "Test", "Ukraine", "0000", "web.site"))

	actual, err := target.FindBy(context.Background(), model.SearchFilter{Name: "Test"})
	if err != nil {
		t.Errorf("should return company, but returns error: %v", err)
	} else if fmt.Sprint(expected) != fmt.Sprint(actual[0]) {
		t.Fatalf("Expected: %v. Actual: %v", expected, actual[0])
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

func TestCompaniesRepository_FindBy_NameAndCountry(t *testing.T) {
	expected := &model.Company{
		Name:    "Test",
		Code:    "1234",
		Country: "Ukraine",
		Phone:   "0000",
		Website: "web.site",
	}
	db, mock, target := createMockedCompaniesRepository(t)
	defer db.Close()

	columns := []string{"code", "name", "country", "phone", "website"}

	mock.ExpectQuery("SELECT (.+) FROM companies WHERE name LIKE (.+) AND country LIKE (.+)").
		WithArgs("%Test%", "%Uk%").
		WillReturnRows(sqlmock.NewRows(columns).AddRow("1234", "Test", "Ukraine", "0000", "web.site"))

	actual, err := target.FindBy(context.Background(), model.SearchFilter{Name: "Test", Country: "Uk"})
	if err != nil {
		t.Errorf("should return company, but returns error: %v", err)
	} else if fmt.Sprint(expected) != fmt.Sprint(actual[0]) {
		t.Fatalf("Expected: %v. Actual: %v", expected, actual[0])
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

func createMockedCompaniesRepository(t *testing.T) (*sql.DB, sqlmock.Sqlmock, *CompaniesRepository) {
	db, mock, err := sqlmock.New()
	if err == nil {
		return db, mock, &CompaniesRepository{
			db: db,
		}
	} else {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	return db, mock, nil
}
