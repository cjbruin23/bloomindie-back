package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	pgConnString, exists := os.LookupEnv("POSTGRES_CONNECTION_STRING")

	if exists {
		fmt.Println(pgConnString)
	}

	connStr := pgConnString

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	var output int

	err = db.QueryRow("SELECT 1").Scan(&output)
	if err != nil {
		log.Fatalf("QueryRow failed: %v", err)
	}

	fmt.Printf("Output: %d\n", output)
}
