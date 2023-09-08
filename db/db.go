package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

var db *pgxpool.Pool

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
	db, err = pgxpool.Connect(context.Background(), psqlInfo)
	if err != nil {
		panic(err)
	}

	log.Println("Successfully connected to the database!")
	InitRedis()  // Depois de conectar com o banco de dados, inicializar o Redis
}

func InitSchema() {
	conn, err := db.Acquire(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Release()

	_, err = conn.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS pessoas (
			id UUID PRIMARY KEY,
			apelido TEXT NOT NULL,
			nome TEXT NOT NULL,
			nascimento TEXT NOT NULL,
			stack TEXT[]
		)
	`)
	if err != nil {
		log.Fatal(err)
	}
}

func GetDB() *pgxpool.Pool {
	return db
}
