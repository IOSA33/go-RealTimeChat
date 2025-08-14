package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func NewDataBase() (*Database, error) {
	const op = "server.db.NewDataBase"

	db, err := sql.Open("postgres", "postgresql://root:password@localhost:5433/go-chat?sslmode=disabled")
	if err != nil {
		return nil, fmt.Errorf("op: %s, error: %w", op, err)
	}

	return &Database{db: db}, nil
}

func (db *Database) CloseDB() error {
	err := db.db.Close()
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) GetDB() *sql.DB {
	return db.db
}
