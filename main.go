package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
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

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		// Make user object from body
		io.Copy(os.Stdout, r.Body)
		w.Write([]byte("Hello World!"))
	})
	http.ListenAndServe(":8080", r)
}
