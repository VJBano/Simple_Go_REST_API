package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func DBConnection() *sql.DB {

	psqlCon := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"localhost", "5432", "postgres", "postgres123", "sampleuser")

	db, err := sql.Open("postgres", psqlCon)
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}

	// Verify the connection to the database
	if err := db.Ping(); err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	log.Println("Successfully connected to the database")
	return db
}
