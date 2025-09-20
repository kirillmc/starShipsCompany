package app

import (
	"context"
	"fmt"
	"github.com/kirillmc/starShipsCompany/platform/pkg/logger"
	"go.uber.org/zap"

	v1 "github.com/kirillmc/starShipsCompany/inventory/internal/api/inventory/v1"
	"github.com/kirillmc/starShipsCompany/inventory/internal/config"
	"github.com/kirillmc/starShipsCompany/inventory/internal/repository/mongoRepo"
	partRepo "github.com/kirillmc/starShipsCompany/inventory/internal/repository/mongoRepo/part"
	"github.com/kirillmc/starShipsCompany/inventory/internal/service"
	partService "github.com/kirillmc/starShipsCompany/inventory/internal/service/part"
	"github.com/kirillmc/starShipsCompany/platform/pkg/closer"
	inventoryV1 "github.com/kirillmc/starShipsCompany/shared/pkg/proto/inventory/v1"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type diContainer struct {
	inventoryV1API inventoryV1.InventoryServiceServer

	inventoryService    service.Service
	inventoryRepository mongoRepo.Repository

	mongoDBClient *mongo.Client
	mongoDBHandle *mongo.Database
}

func NewDIContainer() *diContainer {
	return &diContainer{}
}

func (d *diContainer) InventoryV1API(ctx context.Context) inventoryV1.InventoryServiceServer {
	if d.inventoryV1API == nil {
		d.inventoryV1API = v1.NewAPI(d.InventoryService(ctx))
	}

	return d.inventoryV1API
}

func (d *diContainer) InventoryService(ctx context.Context) service.Service {
	if d.inventoryService == nil {
		d.inventoryService = partService.NewService(d.InventoryRepository(ctx))
	}

	return d.inventoryService
}

func (d *diContainer) InventoryRepository(ctx context.Context) mongoRepo.Repository {
	if d.inventoryRepository == nil {
		var err error
		d.inventoryRepository, err = partRepo.NewRepository(ctx, d.MongoDBHandle(ctx))
		if err != nil {
			logger.Error(ctx, "failed to create new repository", zap.Error(err))
			panic(fmt.Sprintf("failed to create new repository: %s\n", err.Error()))
		}
	}

	return d.inventoryRepository
}

func (d *diContainer) MongoDBClient(ctx context.Context) *mongo.Client {
	if d.mongoDBClient == nil {
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.AppConfig().Mongo.URI()))
		if err != nil {
			logger.Error(ctx, "failed to connect to MongoDB", zap.Error(err))
			panic(fmt.Sprintf("failed to connect to MongoDB: %s\n", err.Error()))
		}

		err = client.Ping(ctx, readpref.Primary())
		if err != nil {

			logger.Error(ctx, "failed to ping MongoDB", zap.Error(err))
			panic(fmt.Sprintf("failed to ping MongoDB: %v\n", err))
		}

		closer.AddNamed("MongoDB client", func(ctx context.Context) error {
			return client.Disconnect(ctx)
		})

		d.mongoDBClient = client
	}

	return d.mongoDBClient
}

func (d *diContainer) MongoDBHandle(ctx context.Context) *mongo.Database {
	if d.mongoDBHandle == nil {
		d.mongoDBHandle = d.MongoDBClient(ctx).Database(config.AppConfig().Mongo.DatabaseName())
	}

	return d.mongoDBHandle
}
