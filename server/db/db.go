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
	sslmode := "disable"
	timezone := "Europe/Paris"
	dsn := "postgresql://" + user + ":" + password + "@" + host + ":" + port + "/" + dbname + "?sslmode=" + sslmode + "&TimeZone=" + timezone

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}) // Ouverture, via la requête DSN et l'ORM Go GORM, d'une connexion à la base de données PostgreSQL
	log.Println("dsn", dsn)
	if err != nil { // Vérification d'une erreur lors de l'ouverture de la base de données'
		fmt.Println("Could not initialize DB connection :", err)
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil { // Vérification d'une erreur lors de l'accès à la base de données
		return nil, err
	}
	err = sqlDB.Ping()
	if err != nil { // Vérification d'une erreur lors de la connexion à la base de données
		return nil, err
	}
	return &Database{db: db}, nil
}

func (d *Database) Close() error { // Fermer la connexion à la base de données
	sqlDB, err := d.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func (d *Database) GetDB() *sql.DB { // Récupérer la base de données
	sqlDB, err := d.db.DB()
	if err != nil {
		return nil
	}
	return sqlDB
}
