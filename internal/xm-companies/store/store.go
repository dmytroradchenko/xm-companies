package store

type Store interface {
	Companies() *CompaniesRepository
	Users() *UsersRepository
}
