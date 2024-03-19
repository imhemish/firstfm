package repository

import (
	"database/sql"
	"fmt"

	"os"

	_ "github.com/lib/pq"
)

func ConnectDb() *sql.DB {
	DATABASE_URL := fmt.Sprintf("postgres://%v:%v@localhost:5432/%v?sslmode=disable", "postgres", "postgres", "auth")
	db, err := sql.Open("postgres", DATABASE_URL)
	if err != nil {
		println("failed")
		os.Exit(1)
	}
	return db
}

var connectedDb = ConnectDb()
