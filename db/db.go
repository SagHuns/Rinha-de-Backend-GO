package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() {
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "5432"
	}
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		dbUser = "postgres"
	}
	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		dbPassword = "videocoding"
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "rinha"
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	log.Println("Successfully connected to the database!")
}

func InitSchema() {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS pessoas (
			id UUID PRIMARY KEY,
			apelido TEXT NOT NULL,
			nome TEXT NOT NULL,
			nascimento TEXT NOT NULL,
			stack TEXT[] NOT NULL
		)
	`)
	if err != nil {
		log.Fatal(err)
	}
}

func GetDB() *sql.DB {
	return db
}
