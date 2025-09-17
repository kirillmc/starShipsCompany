package integration

import (
	"context"
	"os"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	integrationModel "github.com/kirillmc/starShipsCompany/inventory/tests/integration/model"
	inventoryV1 "github.com/kirillmc/starShipsCompany/shared/pkg/proto/inventory/v1"
	"github.com/samber/lo"
	"go.mongodb.org/mongo-driver/bson"
)

// ClearPartsCollection — удаляет все записи из коллекции parts
func (env *TestEnvironment) ClearPartsCollection(ctx context.Context) error {
	// Используем базу данных из переменной окружения MONGO_DATABASE
	databaseName := os.Getenv("MONGO_DATABASE")
	if databaseName == "" {
		databaseName = dbName // fallback значение
	}

	_, err := env.Mongo.Client().Database(databaseName).Collection(partsCollectionName).DeleteMany(ctx, bson.M{})
	if err != nil {
		return err
	}

	return nil
}

// InsertToPartsCollection — добавляет запись в коллекцию parts
func (env *TestEnvironment) InsertToPartsCollection(ctx context.Context) (integrationModel.PartUUID, error) {
	// Используем базу данных из переменной окружения MONGO_DATABASE
	databaseName := os.Getenv("MONGO_DATABASE")
	if databaseName == "" {
		databaseName = dbName // fallback значение
	}

	partUUID := gofakeit.UUID()
	tempPart := integrationModel.Part{
		UUID:          partUUID,
		Name:          gofakeit.Name(),
		Description:   gofakeit.Slogan(),
		Price:         gofakeit.Price(1, 101),
		StockQuantity: gofakeit.Int64(),
		Category:      integrationModel.Category(gofakeit.Int32() % 4),
		Dimensions: &integrationModel.Dimensions{
			Length: gofakeit.Float64(),
			Width:  gofakeit.Float64(),
			Height: gofakeit.Float64(),
			Weight: gofakeit.Float64(),
		},
		Manufacturer: &integrationModel.Manufacturer{
			Name:    gofakeit.Name(),
			Country: gofakeit.Country(),
			Website: gofakeit.Word(),
		},
		Tags:      []string{gofakeit.Word(), gofakeit.Word(), gofakeit.Word()},
		Metadata:  nil,
		CreatedAt: lo.ToPtr(time.Now()),
		UpdatedAt: lo.ToPtr(time.Now()),
	}

	_, err := env.Mongo.Client().Database(databaseName).Collection(partsCollectionName).InsertOne(ctx, tempPart)
	if err != nil {
		return "", err
	}

	return partUUID, nil
}

// GetWrongListPartsParams — получить фильтр на неудачный поиск в коллекции parts
func (env *TestEnvironment) GetWrongListPartsParams() *inventoryV1.ListPartsRequest {
	return &inventoryV1.ListPartsRequest{Filter: &inventoryV1.PartsFilter{Uuids: []string{"IS NOT UUID"}}}
}
