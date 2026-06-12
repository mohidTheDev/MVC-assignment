package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	connStr := "host=localhost port=5432 user=admin password=1234 dbname=coclone_database sslmode=disable"

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to open a DB connection: ", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Failed to ping the database! Is Docker running? Error: ", err)
	}

	fmt.Println("Successfully connected to PostgreSQL!")
}
