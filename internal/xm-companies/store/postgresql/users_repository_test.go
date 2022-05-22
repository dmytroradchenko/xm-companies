package postgresql

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"testing"
	"xm-companies/internal/xm-companies/model"
)

func TestUsersRepository_Create(t *testing.T) {
	db, mock, target := createMockedUsersRepository(t)
	defer db.Close()
	user := &model.User{
		Username: "test",
		Password: "test",
	}

	mock.ExpectQuery("INSERT INTO users (.+) VALUES (.+) RETURNING id").
		WithArgs("test", "test").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("1"))

	id, err := target.Create(context.Background(), user)
	if err != nil {
		t.Fatalf("failed to create user: '%s'", err)
	}
	if id != "1" {
		t.Errorf("returned user ID is '%s'. Expected: '1'", id)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

func TestUsersRepository_Create_ReturnsError(t *testing.T) {
	db, mock, target := createMockedUsersRepository(t)
	defer db.Close()
	user := &model.User{}

	mock.ExpectQuery("INSERT INTO users (.+) VALUES (.+) RETURNING id").WillReturnError(sql.ErrNoRows)

	if _, err := target.Create(context.Background(), user); err == nil {
		t.Fatalf("should return an error: %s", sql.ErrNoRows)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

func TestUsersRepository_Find(t *testing.T) {
	db, mock, target := createMockedUsersRepository(t)
	defer db.Close()

	columns := []string{"id", "username", "password"}

	mock.ExpectQuery("SELECT id, username, password FROM users").
		WithArgs("test").
		WillReturnRows(sqlmock.NewRows(columns).AddRow("1", "test", "test"))

	if u, err := target.Find(context.Background(), "test"); err != nil {
		t.Errorf("should return user with 'test' usersname but error accured: %s", err)
	} else if u.Username != "test" {
		t.Fatalf("expected username is 'test'. Actual: %s", u.Username)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

func TestUsersRepository_Find_ReturnsNoSuchUserError(t *testing.T) {
	db, mock, target := createMockedUsersRepository(t)
	defer db.Close()

	mock.ExpectQuery("SELECT id, username, password FROM users").
		WithArgs("test").
		WillReturnError(sql.ErrNoRows)

	if _, err := target.Find(context.Background(), "test"); err != NoSuchUserError {
		t.Fatalf("expected error: %v. Actual: %v", NoSuchUserError, err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

func createMockedUsersRepository(t *testing.T) (*sql.DB, sqlmock.Sqlmock, *UsersRepository) {
	db, mock, err := sqlmock.New()
	if err == nil {
		return db, mock, &UsersRepository{
			db: db,
		}
	} else {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	return db, mock, nil
}
