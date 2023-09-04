package db

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "videocoding"
	dbname   = "rinha"
)

var db *sql.DB

func InitDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

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
