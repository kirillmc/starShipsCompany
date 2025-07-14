package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
	orderV1 "github.com/kirillmc/starShipsCompany/shared/pkg/openapi/order/v1"
	inventoryV1 "github.com/kirillmc/starShipsCompany/shared/pkg/proto/inventory/v1"
	paymentV1 "github.com/kirillmc/starShipsCompany/shared/pkg/proto/payment/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Order = orderV1.GetOrderResponse

const (
	httpPort = "8080"
	httpHost = "localhost"

	inventoryServiceAddress = "localhost:50051"
	paymentServiceAddress   = "localhost:50052"
	// Таймауты для HTTP-сервера
	readHeaderTimeout = 5 * time.Second
	shutdownTimeout   = 10 * time.Second
)

// OrderStorage представляет потокобезопасное хранилище данных о заказах
type OrderStorage struct {
	mu     sync.RWMutex
	orders map[string]*Order
}

func NewWeatherStorage() *OrderStorage {
	return &OrderStorage{
		orders: make(map[string]*Order),
	}
}

func (s *OrderHandler) CancelOrder(_ context.Context, params orderV1.CancelOrderParams) (orderV1.CancelOrderRes, error) {
	order, err := s.storage.getOrder(params.OrderUUID.String())
	if err != nil {
		return &orderV1.NotFoundError{
			Code:    404,
			Message: err.Error(),
		}, nil
	}

	if order.Status.Value == orderV1.OrderStatusPAID {
		return &orderV1.ConflictError{
			Code:    409,
			Message: "Заказ уже отменен",
		}, nil
	}

	s.storage.setOrderStatus(params.OrderUUID.String(), orderV1.OrderStatusCANCELLED)

	return &orderV1.CancelOrderNoContent{}, nil
}

func (s *OrderHandler) CreateOrder(ctx context.Context, req *orderV1.CreateOrderRequest) (orderV1.CreateOrderRes, error) {
	if req == nil {
		return &orderV1.CreateOrderResponse{}, nil
	}

	inventoryReq := &inventoryV1.ListPartsRequest{Filter: &inventoryV1.PartsFilter{Uuids: req.PartUuids}}
	resp, err := s.inventoryService.ListParts(ctx, inventoryReq)
	if err != nil {
		return &orderV1.InternalServerError{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("Не удалось получить список доступных деталей: %s", err),
		}, nil
	}
	if len(resp.Parts) < len(req.PartUuids) {
		return &orderV1.InternalServerError{
			Code:    http.StatusInternalServerError,
			Message: "В наличии меньше деталей, чем требуется",
		}, nil
	}

	var totalPrice float64
	partsUUIDS := make([]string, 0, len(resp.Parts))

	for _, part := range resp.Parts {
		totalPrice += part.Price
		partsUUIDS = append(partsUUIDS, part.Uuid)
	}

	orderUUID := uuid.NewString()

	s.storage.addOrder(orderUUID, req.UserUUID, partsUUIDS, totalPrice)

	return &orderV1.CreateOrderResponse{
		OrderUUID:  orderUUID,
		TotalPrice: totalPrice,
	}, nil
}

func (s *OrderHandler) GetOrder(_ context.Context, params orderV1.GetOrderParams) (orderV1.GetOrderRes, error) {
	order, err := s.storage.getOrder(params.OrderUUID.String())
	if err != nil {
		return &orderV1.NotFoundError{
			Code:    404,
			Message: err.Error(),
		}, nil
	}

	return order, nil
}

func (s *OrderHandler) PayOrder(ctx context.Context, req *orderV1.PayOrderRequest, params orderV1.PayOrderParams) (orderV1.PayOrderRes, error) {
	if req == nil {
		return &orderV1.PayOrderResponse{}, nil
	}

	order, err := s.storage.getOrder(params.OrderUUID.String())
	if err != nil {
		return &orderV1.NotFoundError{
			Code:    404,
			Message: err.Error(),
		}, nil
	}

	paymentReq := paymentV1.PayOrderRequest{
		OrderUuid:     params.OrderUUID.String(),
		UserUuid:      order.UserUUID,
		PaymentMethod: paymentMethodToPaymentV1(req.PaymentMethod),
	}

	resp, err := s.paymentService.PayOrder(ctx, &paymentReq)
	if err != nil || resp == nil {
		return &orderV1.InternalServerError{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("Не удалось провести транзакцию для заказа с UUID %s: %s", order.OrderUUID, err),
		}, nil
	}

	s.storage.setOrderStatus(order.OrderUUID, orderV1.OrderStatusPAID)

	return &orderV1.PayOrderResponse{TransactionUUID: resp.TransactionUuid}, nil
}

func paymentMethodToPaymentV1(method orderV1.PaymentMethod) paymentV1.PAYMENTMETHOD {
	switch method {
	case orderV1.PaymentMethodCARD:
		return paymentV1.PAYMENTMETHOD_CARD
	case orderV1.PaymentMethodSBP:
		return paymentV1.PAYMENTMETHOD_SBP
	case orderV1.PaymentMethodCREDITCARD:
		return paymentV1.PAYMENTMETHOD_CREDIT_CARD
	case orderV1.PaymentMethodINVESTORMONEY:
		return paymentV1.PAYMENTMETHOD_INVESTOR_MONEY
	}

	return paymentV1.PAYMENTMETHOD_UNKNOWN
}

// GetWeather возвращает информацию о погоде по имени города
func (s *OrderStorage) GetWeather(city string) *Order {
	s.mu.RLock()
	defer s.mu.RUnlock()

	weather, ok := s.orders[city]
	if !ok {
		return nil
	}

	return weather
}

func (s *OrderStorage) addOrder(orderUUID, userUUID string, partsUUIDS []string, totalPrice float64) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.orders[orderUUID] = &Order{
		OrderUUID:       orderUUID,
		UserUUID:        userUUID,
		PartUuids:       partsUUIDS,
		TotalPrice:      orderV1.NewOptFloat64(totalPrice),
		TransactionUUID: orderV1.OptString{},
		PaymentMethod:   orderV1.OptPaymentMethod{},
		Status:          orderV1.NewOptOrderStatus(orderV1.OrderStatusPENDINGPAYMENT),
	}
}

func (s *OrderStorage) setOrderStatus(orderUUID string, status orderV1.OrderStatus) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.orders[orderUUID].Status = orderV1.NewOptOrderStatus(status)
}

func (s *OrderStorage) getOrder(orderUUID string) (*Order, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if order, ok := s.orders[orderUUID]; ok {
		return order, nil
	}

	return nil, fmt.Errorf("заказ с UUID %s не найден", orderUUID)
}

type OrderHandler struct {
	storage          *OrderStorage
	inventoryService inventoryV1.InventoryServiceClient
	paymentService   paymentV1.PaymentServiceClient
}

// NewOrderHandler создает новый обработчик запросов к API погоды
func NewOrderHandler(storage *OrderStorage, inventoryService inventoryV1.InventoryServiceClient, paymentService paymentV1.PaymentServiceClient) *OrderHandler {
	return &OrderHandler{
		storage:          storage,
		inventoryService: inventoryService,
		paymentService:   paymentService,
	}
}

// NewError создает новую ошибку в формате GenericError
func (h *OrderHandler) NewError(_ context.Context, err error) *orderV1.GenericErrorStatusCode {
	return &orderV1.GenericErrorStatusCode{
		StatusCode: http.StatusInternalServerError,
		Response: orderV1.GenericError{
			Code:    orderV1.NewOptInt(http.StatusInternalServerError),
			Message: orderV1.NewOptString(err.Error()),
		},
	}
}

func main() {
	storage := NewWeatherStorage()

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
	inventoryClient := inventoryV1.NewInventoryServiceClient(connInventory)

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
	pymentClient := paymentV1.NewPaymentServiceClient(connPayment)

	orderHandler := NewOrderHandler(storage, inventoryClient, pymentClient)

	// Создаем OpenAPI сервер
	orderServer, err := orderV1.NewServer(orderHandler)
	if err != nil {
		log.Printf("ошибка создания сервера OpenAPI: %v", err)
		return
	}

	// Инициализируем роутер Chi
	r := chi.NewRouter()

	// Добавляем middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	// r.Use(middleware.Timeout(10 * time.Second))

	// Монтируем обработчики OpenAPI
	r.Mount("/", orderServer)

	// Запускаем HTTP-сервер
	server := &http.Server{
		Addr:              net.JoinHostPort(httpHost, httpPort),
		Handler:           r,
		ReadHeaderTimeout: readHeaderTimeout, // Защита от Slowloris атак - тип DDoS-атаки, при которой
		// атакующий умышленно медленно отправляет HTTP-заголовки, удерживая соединения открытыми и истощая
		// пул доступных соединений на сервере. ReadHeaderTimeout принудительно закрывает соединение,
		// если клиент не успел отправить все заголовки за отведенное время.
	}

	// Запускаем сервер в отдельной горутине
	go func() {
		log.Printf("🚀 HTTP-сервер запущен на порту %s\n", httpPort)
		err = server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("❌ Ошибка запуска сервера: %v\n", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("🛑 Завершение работы сервера...")

	// Создаем контекст с таймаутом для остановки сервера
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	err = server.Shutdown(ctx)
	if err != nil {
		log.Printf("❌ Ошибка при остановке сервера: %v\n", err)
	}

	log.Println("✅ Сервер остановлен")
}
