package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

var DB *sql.DB

func InitDB() {
	connStr := "host=localhost port=5432 user=postgres password=2005b dbname=postgres sslmode=disable"
	var err error
	DB, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error connecting to DB: %s", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("DB ping error: %s", err)
	}

	fmt.Println("DB successfully connected to postgres")

}
