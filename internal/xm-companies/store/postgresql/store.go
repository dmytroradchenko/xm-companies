package postgresql

import (
	"database/sql"
	"xm-companies/internal/xm-companies/store"
)

type Store struct {
	db        *sql.DB
	users     store.UsersRepository
	companies store.CompaniesRepository
}

func NewStoreProvider(db *sql.DB, users store.UsersRepository, companies store.CompaniesRepository) *Store {
	return &Store{
		db:        db,
		users:     users,
		companies: companies,
	}
}

func (s Store) Companies() store.CompaniesRepository {
	return s.companies
}

func (s Store) Users() store.UsersRepository {
	return s.users
}
