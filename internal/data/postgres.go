package data

import (
	"database/sql"
	"os"

	// registering database driver
	_ "github.com/lib/pq"
)

func getConnection() (*sql.DB, error) {
	uri := os.Getenv("DATABASE_URI")
	return sql.Open("postgres", uri)
}
