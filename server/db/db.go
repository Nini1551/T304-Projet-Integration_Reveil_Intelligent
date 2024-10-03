package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
)

type Database struct {
	db *gorm.DB
}

func NewDatabase() (*Database, error) {
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	port := os.Getenv("POSTGRES_PORT")
	host := "localhost" // os.Getenv("POSTGRES_HOST")
	dsn := "postgresql://" + user + ":" + password + "@" + host + ":" + port + "/" + dbname + "?sslmode=disable&TimeZone=Europe/Paris"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	log.Println("dsn", dsn)
	if err != nil {
		fmt.Println("Could not initialize DB connection :", err)
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	err = sqlDB.Ping()
	if err != nil {
		return nil, err
	}
	return &Database{db: db}, nil
}

func (d *Database) Close() error {
	sqlDB, err := d.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func (d *Database) GetDB() *sql.DB {
	sqlDB, err := d.db.DB()
	if err != nil {
		return nil
	}
	return sqlDB
}
