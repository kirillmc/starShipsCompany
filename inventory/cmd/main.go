package main

import (
	"fmt"
	partRepo "github.com/kirillmc/starShipsCompany/inventory/internal/repository/mongo/part"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	inventoryV1API "github.com/kirillmc/starShipsCompany/inventory/internal/api/inventory/v1"
	partService "github.com/kirillmc/starShipsCompany/inventory/internal/service/part"
	inventoryV1 "github.com/kirillmc/starShipsCompany/shared/pkg/proto/inventory/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const grpcPort = 50051

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Printf("failed to listen: %v\n", err)
		return
	}

	s := grpc.NewServer()
	reflection.Register(s)

	repo := partRepo.NewRepository()
	service := partService.NewService(repo)
	api := inventoryV1API.NewAPI(service)

	inventoryV1.RegisterInventoryServiceServer(s, api)

	go func() {
		log.Printf("Starting gRPC server at port %d", grpcPort)
		err = s.Serve(lis)
		if err != nil {
			log.Printf("failed to serve: %v\n", err)
			return
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("🛑 Shutting down gRPC server...")
	s.GracefulStop()
	log.Println("✅ Server stopped")
}
