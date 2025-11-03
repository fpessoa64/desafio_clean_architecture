package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/fpessoa64/desafio_clean_arch/docs"
	_ "github.com/mattn/go-sqlite3"

	"github.com/fpessoa64/desafio_clean_arch/internal/repository/sqlite"
	"github.com/fpessoa64/desafio_clean_arch/internal/servers"
	"github.com/fpessoa64/desafio_clean_arch/internal/usecase"
)

func startServerGraphQL(uc *usecase.OrderUsecase) {
	// Implementação do servidor GraphQL aqui
}

var db *sql.DB

func init() {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}
	var err error
	db, err = sql.Open("sqlite3", dbURL)
	if err != nil {
		log.Fatal("failed to connect to db:", err)
	}
	// Migração automática
	sqlBytes, err := os.ReadFile("migrations/create_orders_table.up.sql")
	if err != nil {
		log.Fatal("failed to read migration:", err)
	}
	if _, err = db.Exec(string(sqlBytes)); err != nil {
		log.Fatal("failed to migrate:", err)
	}
}

func main() {
	defer db.Close()

	repo := sqlite.NewOrderRepositorySqlite(db)
	uc := usecase.NewOrderUsecase(repo)

	grpcServer := servers.NewGrpc()
	go grpcServer.Start(uc)

	restServer := servers.NewRest()
	go restServer.Start(uc)

	graphServer := servers.NewGraphQL()
	go graphServer.Start(uc)

	select {}
}
