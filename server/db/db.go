package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func NewDatabase() (*Database, error) {
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	port := os.Getenv("POSTGRES_PORT")
	host := "localhost" // os.Getenv("POSTGRES_HOST")
	dsn := "postgresql://" + user + ":" + password + "@" + host + ":" + port + "/" + dbname + "?sslmode=disable"

	db, err := sql.Open("postgres", dsn)
	log.Println("dsn", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &Database{db: db}, nil
}

func (d *Database) Close() error {
	return d.db.Close()
}

func (d *Database) GetDB() *sql.DB {
	return d.db
}
