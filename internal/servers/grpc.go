package servers

import (
	"log"
	"net"
	"os"

	orderpb "github.com/fpessoa64/desafio_clean_arch/internal/handlers/grpc/proto"
	grpcorder "github.com/fpessoa64/desafio_clean_arch/internal/handlers/grpc/proto/service"
	"github.com/fpessoa64/desafio_clean_arch/internal/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Grpc struct {
	port string
}

func NewGrpc() *Grpc {
	grpcPort := os.Getenv("GRPC_PORT")
	if grpcPort == "" {
		grpcPort = "50051"
	}
	return &Grpc{port: grpcPort}
}

func (g *Grpc) Start(uc *usecase.OrderUsecase) error {

	lis, err := net.Listen("tcp", ":"+g.port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	orderpb.RegisterOrderServiceServer(grpcServer, grpcorder.NewOrderServiceServer(uc))

	reflection.Register(grpcServer)
	log.Printf("gRPC server running on :%s", g.port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC: %v", err)
		return err
	}
	return nil
}
