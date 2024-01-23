package db

import (
	"database/sql"
	"fmt"
	"os"
)

type Database struct {
	db *sql.DB
}

func InitDB() (*Database, error) {

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	if err != nil {
		return nil, err
	}

	return &Database{db: db}, nil

}
