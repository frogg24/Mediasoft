package database

import "database/sql"

type DB struct {
	*sql.DB
}

func NewDatabase(db *sql.DB) *DB {
	return &DB{db}
}
