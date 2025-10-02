package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	httpAdapter "the-engineering-coach/tdd-with-ai/internal/adapters/http"
	"the-engineering-coach/tdd-with-ai/internal/adapters/sqlite"
	"the-engineering-coach/tdd-with-ai/internal/services"
)

func NewServer(db *sql.DB) http.Handler {
	repo := sqlite.NewFlightRepository(db)
	service := services.NewFlightService(repo)
	handler := httpAdapter.NewFlightHandler(service)

	mux := http.NewServeMux()
	mux.Handle("/flights", handler)

	return mux
}

func main() {
	db, err := sql.Open("sqlite3", "flights.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	server := NewServer(db)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", server); err != nil {
		log.Fatal(err)
	}
}
