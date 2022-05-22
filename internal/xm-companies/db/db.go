package db

import (
	"database/sql"
	"fmt"

	"xm-companies/config"

	_ "github.com/lib/pq"
)

func New(cfg *config.Config) (*sql.DB, error) {
	url := fmt.Sprintf("jdbc:postgresql://%s:%s/%s?%s=&password=%s&ssl=false",
		cfg.DbHost, cfg.DbPort, cfg.DbName, cfg.DbUser, cfg.DbPassword)
	db, err := sql.Open(cfg.DbDriverName, url)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
