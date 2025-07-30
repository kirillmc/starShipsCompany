package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/brianvoe/gofakeit/v7"
	inventoryV1 "github.com/kirillmc/starShipsCompany/shared/pkg/proto/inventory/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
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
	service := &inventoryService{}
	service = service.newService()

	inventoryV1.RegisterInventoryServiceServer(s, service)

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

type inventoryService struct {
	inventoryV1.UnimplementedInventoryServiceServer

	mu    sync.RWMutex
	parts map[string]*inventoryV1.Part
}

func (s *inventoryService) newService() *inventoryService {
	return &inventoryService{
		parts: setDefaultPartsMap(),
	}
}

func (s *inventoryService) GetPart(_ context.Context, req *inventoryV1.GetPartRequest) (*inventoryV1.GetPartResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	part, ok := s.parts[req.GetUuid()]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "part with UUID %s not found", req.GetUuid())
	}

	return &inventoryV1.GetPartResponse{
		Part: part,
	}, nil
}

func (s *inventoryService) ListParts(_ context.Context, req *inventoryV1.ListPartsRequest) (*inventoryV1.ListPartsResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	filteredParts := s.filterValues(req.GetFilter())

	return &inventoryV1.ListPartsResponse{
		Parts: filteredParts,
	}, nil
}

func setDefaultPartsMap() map[string]*inventoryV1.Part {
	const defaultPartsCoount = 11
	defaultMap := make(map[string]*inventoryV1.Part)
	for range defaultPartsCoount {
		UUID := gofakeit.UUID()
		defaultMap[UUID] = &inventoryV1.Part{
			Uuid:          UUID,
			Name:          gofakeit.Name(),
			Description:   gofakeit.Slogan(),
			Price:         gofakeit.Price(1, 101),
			StockQuantity: gofakeit.Int64(),
			Category:      inventoryV1.CATEGORY(gofakeit.Int32() % 4),
			Dimensions: &inventoryV1.Dimensions{
				Length: gofakeit.Float64(),
				Width:  gofakeit.Float64(),
				Height: gofakeit.Float64(),
				Weight: gofakeit.Float64(),
			},
			Manufacturer: &inventoryV1.Manufacturer{
				Name:    gofakeit.Name(),
				Country: gofakeit.Country(),
				Website: gofakeit.Word(),
			},
			Tags:      []string{gofakeit.Word(), gofakeit.Word(), gofakeit.Word()},
			Metadata:  nil,
			CreatedAt: timestamppb.Now(),
			UpdatedAt: timestamppb.Now(),
		}
	}

	return defaultMap
}

func (s *inventoryService) filterValues(filter *inventoryV1.PartsFilter) []*inventoryV1.Part {
	parts := make([]*inventoryV1.Part, 0, len(s.parts))
	for _, part := range s.parts {
		parts = append(parts, part)
	}

	if filter == nil {
		return parts
	}

	if len(filter.Uuids) > 0 {
		parts = applySimpleFieldFilter(filter.Uuids, parts, func(p *inventoryV1.Part) string { return p.Uuid })
	}

	if len(filter.Names) > 0 {
		parts = applySimpleFieldFilter(filter.Names, parts, func(p *inventoryV1.Part) string { return p.Name })
	}

	if len(filter.Categories) > 0 {
		parts = applySimpleFieldFilter(filter.Categories, parts, func(p *inventoryV1.Part) inventoryV1.CATEGORY { return p.Category })
	}

	if len(filter.ManufacturerCountries) > 0 {
		parts = applySimpleFieldFilter(filter.ManufacturerCountries, parts, func(p *inventoryV1.Part) string { return p.Manufacturer.Country })
	}

	if len(filter.Tags) > 0 {
		parts = applyTagsFilter(filter.Tags, parts)
	}

	return parts
}

func applySimpleFieldFilter[T AllowedFilterTypes](filterValues []T, allowedParts []*inventoryV1.Part,
	getField func(*inventoryV1.Part) T,
) []*inventoryV1.Part {
	if filterValues == nil {
		return allowedParts
	}

	valueSet := getFilterValuesSet[T](filterValues)
	var res []*inventoryV1.Part
	for _, part := range allowedParts {
		if _, ok := valueSet[getField(part)]; ok {
			res = append(res, part)
		}
	}
	return res
}

func applyTagsFilter(filterValues []string, allowedParts []*inventoryV1.Part) []*inventoryV1.Part {
	if filterValues == nil {
		return allowedParts
	}

	tagsSet := getFilterValuesSet[string](filterValues)
	var res []*inventoryV1.Part

	for _, part := range allowedParts {
		for _, tag := range part.Tags {
			if _, ok := tagsSet[tag]; ok {
				res = append(res, part)
				break
			}
		}
	}

	return res
}

type AllowedFilterTypes interface {
	string | inventoryV1.CATEGORY
}

func getFilterValuesSet[T AllowedFilterTypes](filterValues []T) map[T]struct{} {
	if len(filterValues) == 0 {
		return nil
	}

	filterSet := make(map[T]struct{})

	for _, filterValue := range filterValues {
		if _, ok := filterSet[filterValue]; !ok {
			filterSet[filterValue] = struct{}{}
		}
	}

	return filterSet
}
