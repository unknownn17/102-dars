package database

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectPostgres() (*sqlx.DB, error) {
	dsn := os.Getenv("DATABASE_URL") // .env dan olish
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatal("Could not connect to the database:", err)
		return nil, err
	}
	return db, nil
}
