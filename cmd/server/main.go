package main

import (
	"database/sql"
	"log"
	"net"
	"net/http"
	"os"

	_ "github.com/fpessoa64/desafio_clean_arch/docs"
	"github.com/go-chi/chi/v5"
	_ "github.com/mattn/go-sqlite3"
	httpSwagger "github.com/swaggo/http-swagger"
	"google.golang.org/grpc"

	orderpb "github.com/fpessoa64/desafio_clean_arch/internal/handlers/grpc/proto"
	grpcorder "github.com/fpessoa64/desafio_clean_arch/internal/handlers/grpc/proto/service"
	"github.com/fpessoa64/desafio_clean_arch/internal/handlers/rest"
	"github.com/fpessoa64/desafio_clean_arch/internal/handlers/rest/routes"
	"github.com/fpessoa64/desafio_clean_arch/internal/repository/sqlite"
	"github.com/fpessoa64/desafio_clean_arch/internal/usecase"
	"google.golang.org/grpc/reflection"
)

func startServerREST(uc *usecase.OrderUsecase) {
	handler := rest.NewHandler(uc)
	r := chi.NewRouter()
	routes.RegisterOrderRoutes(r, handler)
	r.Get("/swagger/*", httpSwagger.WrapHandler)
	restPort := os.Getenv("REST_PORT")
	if restPort == "" {
		restPort = "8080"
	}
	log.Printf("REST API running on :%s", restPort)
	if err := http.ListenAndServe(":"+restPort, r); err != nil {
		log.Fatal(err)
	}
}

func startServerGrpc(uc *usecase.OrderUsecase) {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	orderpb.RegisterOrderServiceServer(grpcServer, grpcorder.NewOrderServiceServer(uc))
	grpcPort := os.Getenv("GRPC_PORT")
	if grpcPort == "" {
		grpcPort = "50051"
	}
	reflection.Register(grpcServer)
	log.Printf("gRPC server running on :%s", grpcPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC: %v", err)
	}
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

	go startServerGrpc(uc)
	startServerREST(uc)
}

// 	repo := sqlite.NewOrderRepositorySqlite(db)
// 	uc := usecase.NewOrderUsecase(repo)
// 	handler := rest.NewHandler(uc)

// 	r := chi.NewRouter()
// 	routes.RegisterOrderRoutes(r, handler)

// 	// Swagger docs endpoint
// 	r.Get("/swagger/*", httpSwagger.WrapHandler)

// 	// gRPC server
// 	go func() {
// 		lis, err := net.Listen("tcp", ":50051")
// 		if err != nil {
// 			log.Fatalf("failed to listen: %v", err)
// 		}
// 		grpcServer := grpc.NewServer()
// 		orderpb.RegisterOrderServiceServer(grpcServer, grpcorder.NewOrderServiceServer(uc))
// 		log.Println("gRPC server running on :50051")
// 		if err := grpcServer.Serve(lis); err != nil {
// 			log.Fatalf("failed to serve gRPC: %v", err)
// 		}
// 	}()
// 	log.Println("REST API running on :8080")
// 	if err := http.ListenAndServe(":8080", r); err != nil {
// 		log.Fatal(err)
// 	}
// }
