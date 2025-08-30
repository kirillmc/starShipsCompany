package main

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	orderAPI "github.com/kirillmc/starShipsCompany/order/internal/api/order/v1"
	inventoryClient "github.com/kirillmc/starShipsCompany/order/internal/client/grpc/inventory/v1"
	paymentClient "github.com/kirillmc/starShipsCompany/order/internal/client/grpc/payment/v1"
	orderRepository "github.com/kirillmc/starShipsCompany/order/internal/repository/order"
	orderService "github.com/kirillmc/starShipsCompany/order/internal/service/order"
	orderV1 "github.com/kirillmc/starShipsCompany/shared/pkg/openapi/order/v1"
	inventoryV1 "github.com/kirillmc/starShipsCompany/shared/pkg/proto/inventory/v1"
	paymentV1 "github.com/kirillmc/starShipsCompany/shared/pkg/proto/payment/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	httpPort = "8080"
	httpHost = "localhost"

	inventoryServiceAddress = "localhost:50051"
	paymentServiceAddress   = "localhost:50052"
	readHeaderTimeout       = 5 * time.Second
	shutdownTimeout         = 10 * time.Second
)

func main() {
	connInventory, err := grpc.NewClient(
		inventoryServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Printf("failed to connect: %v\n", err)
		return
	}
	defer func() {
		if cerr := connInventory.Close(); cerr != nil {
			log.Printf("failed to close connect: %v", cerr)
		}
	}()
	inventoryClient := inventoryClient.NewClient(inventoryV1.NewInventoryServiceClient(connInventory))

	connPayment, err := grpc.NewClient(
		paymentServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Printf("failed to connect: %v\n", err)
		return
	}
	defer func() {
		if cerr := connPayment.Close(); cerr != nil {
			log.Printf("failed to close connect: %v", cerr)
		}
	}()
	paymentClient := paymentClient.NewClient(paymentV1.NewPaymentServiceClient(connPayment))

	storage := orderRepository.NewRepository()
	service := orderService.NewService(inventoryClient, paymentClient, storage)

	orderHandler := orderAPI.NewAPI(service)

	orderServer, err := orderV1.NewServer(orderHandler)
	if err != nil {
		log.Printf("ошибка создания сервера OpenAPI: %v", err)
		return
	}

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(10 * time.Second))

	r.Mount("/", orderServer)

	server := &http.Server{
		Addr:              net.JoinHostPort(httpHost, httpPort),
		Handler:           r,
		ReadHeaderTimeout: readHeaderTimeout,
	}

	go func() {
		log.Printf("🚀 HTTP-сервер запущен на порту %s\n", httpPort)
		err = server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("❌ Ошибка запуска сервера: %v\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("🛑 Завершение работы сервера...")

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	err = server.Shutdown(ctx)
	if err != nil {
		log.Printf("❌ Ошибка при остановке сервера: %v\n", err)
	}

	log.Println("✅ Сервер остановлен")
}
