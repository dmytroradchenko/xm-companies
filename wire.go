//go:build wireinject

package main

import (
	"github.com/google/wire"
	"xm-companies/config"
	"xm-companies/internal/xm-companies/db"
	"xm-companies/internal/xm-companies/security/otp"
	"xm-companies/internal/xm-companies/security/otp/cotp"
	"xm-companies/internal/xm-companies/server"
	"xm-companies/internal/xm-companies/store"
	"xm-companies/internal/xm-companies/store/postgresql"
)

func BuildServerCompileTime() (*server.Server, error) {
	wire.Build(config.NewConfigProvider,
		db.New,
		postgresql.NewStoreProvider,
		wire.Bind(new(store.Store), new(*postgresql.Store)),
		postgresql.NewUsersRepositoryProvider,
		wire.Bind(new(store.UsersRepository), new(*postgresql.UsersRepository)),
		postgresql.NewCompaniesRepositoryProvider,
		wire.Bind(new(store.CompaniesRepository), new(*postgresql.CompaniesRepository)),
		cotp.NewServiceProvider,
		wire.Bind(new(otp.Service), new(*cotp.Service)),
		server.NewServerProvider,
	)
	return &server.Server{}, nil
}
