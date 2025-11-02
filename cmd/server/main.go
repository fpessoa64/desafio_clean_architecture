package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	_ "github.com/mattn/go-sqlite3"

	"github.com/fpessoa64/desafio_clean_arch/internal/handlers/rest"
	"github.com/fpessoa64/desafio_clean_arch/internal/repository/sqlite"
	"github.com/fpessoa64/desafio_clean_arch/internal/usecase"
)

const (
	DATABASE_URL_ENV       = "DATABASE_URL"
	ErrorDatabaseURLNotSet = "DATABASE_URL is not set"
)

var db *sql.DB

func init() {
	dbURL := os.Getenv(DATABASE_URL_ENV)
	if dbURL == "" {
		panic(ErrorDatabaseURLNotSet)
	}
	var err error
	db, err = sql.Open("sqlite3", dbURL)
	if err != nil {
		log.Fatal("failed to connect to db:", err)
	}

	if err := migrateDB(db); err != nil {
		log.Fatal("failed to migrate:", err)
	}
}

func migrateDB(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS orders (
	       id INTEGER PRIMARY KEY AUTOINCREMENT,
	       name TEXT NOT NULL,
	       amount REAL NOT NULL DEFAULT 0,
	       status TEXT NOT NULL DEFAULT 'pending',
	       created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
       )`)
	return err
}

func main() {
	defer db.Close()

	repo := sqlite.NewOrderRepositorySqlite(db)
	uc := usecase.NewOrderUsecase(repo)
	handler := rest.NewHandler(uc)

	r := chi.NewRouter()
	r.Post("/order", handler.CreateOrder)
	r.Get("/order", handler.ListOrders)

	log.Println("REST API running on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
