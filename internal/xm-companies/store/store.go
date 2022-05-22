//go:generate mockery --name Store --output mocks --case underscore
package store

type Store interface {
	Companies() CompaniesRepository
	Users() UsersRepository
}
