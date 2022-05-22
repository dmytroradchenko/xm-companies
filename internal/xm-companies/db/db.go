package db

import (
	"database/sql"
	"fmt"

	"xm-companies/config"

	_ "github.com/lib/pq"
)

func New(cfg *config.Config) (*sql.DB, error) {
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DbName, cfg.DbSslMode)
	db, err := sql.Open(cfg.DbDriverName, url)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
