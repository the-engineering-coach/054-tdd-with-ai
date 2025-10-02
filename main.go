package main

import (
	"database/sql"
	"net/http"
)

func NewServer(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[]`))
	})
}

func main() {
}
