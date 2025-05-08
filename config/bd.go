package config

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	return sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/apprest?sslmode=disable")
}
