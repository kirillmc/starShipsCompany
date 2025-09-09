package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	partRepo "github.com/kirillmc/starShipsCompany/inventory/internal/repository/mongo/part"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

const (
	mongoURI       = "MONGO_URI"
	mongoDB        = "MONGO_INITDB_DATABASE"
	grpcPort       = 50051
	envPath        = ".env.example"
	connectionType = "tcp"
)

func main() {
	ctx := context.Background()

	lis, err := net.Listen(connectionType, fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Printf("failed to listen: %v\n", err)
		return
	}

	s := grpc.NewServer()
	reflection.Register(s)

	err = godotenv.Load(envPath)
	if err != nil {
		log.Printf("failed to load .env.example file: %v\n", err)
		return
	}

	dbURI := os.Getenv(mongoURI)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbURI))
	if err != nil {
		log.Printf("failed to connect to database: %v\n", err)
		return
	}
	defer func() {
		cerr := client.Disconnect(ctx)
		if cerr != nil {
			log.Printf("failed to disconnect: %v\n", cerr)
		}
	}()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Printf("failed to ping database: %v\n", err)
		return
	}

	mongoNameDB := os.Getenv(mongoDB)
	mongoInventoryDB := client.Database(mongoNameDB)
	repo, err := partRepo.NewRepository(ctx, mongoInventoryDB)
	if err != nil {
		log.Fatalf("ошибка инициализации репозитория: %s", err)
	}

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
