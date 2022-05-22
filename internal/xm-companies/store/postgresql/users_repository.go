package postgresql

import (
	"context"
	"database/sql"
	"errors"

	"github.com/leporo/sqlf"
	"xm-companies/internal/xm-companies/model"
)

type UsersRepository struct {
	db *sql.DB
}

const (
	TableUsers     = "users"
	ColumnId       = "id"
	ColumnUsername = "username"
	ColumnPassword = "password"
)

var NoSuchUserError = errors.New("no user with such username")

func NewUsersRepositoryProvider(db *sql.DB) *UsersRepository {
	return &UsersRepository{
		db: db,
	}
}

func (ur *UsersRepository) Create(ctx context.Context, user *model.User) (string, error) {
	err := sqlf.InsertInto(TableUsers).
		Set(ColumnUsername, user.Username).
		Set(ColumnPassword, user.Password).
		Returning(ColumnId).To(&user.Id).
		QueryRowAndClose(ctx, ur.db)
	return user.Id, err
}

func (ur *UsersRepository) Find(ctx context.Context, username string) (*model.User, error) {
	user := &model.User{}
	err := sqlf.From(TableUsers).
		Select(ColumnId).To(&user.Id).
		Select(ColumnUsername).To(&user.Username).
		Select(ColumnPassword).To(&user.Password).
		Where("username = ?", username).QueryRowAndClose(ctx, ur.db)
	if err == sql.ErrNoRows {
		return nil, NoSuchUserError
	} else if err != nil {
		return nil, err
	}
	return user, nil
}
