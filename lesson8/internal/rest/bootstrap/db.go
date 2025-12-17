package bootstrap

import (
	"database/sql"
	"mediasoft/lesson8/internal/rest/config"

	_ "github.com/jackc/pgx/stdlib"
)

func InitDB(cfg *config.Config) (*sql.DB, error) {
	db, err := sql.Open("pgx", cfg.PG)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
