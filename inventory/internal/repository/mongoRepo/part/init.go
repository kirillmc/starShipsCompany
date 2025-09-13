package part

import (
	"context"
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	repoModel "github.com/kirillmc/starShipsCompany/inventory/internal/repository/mongoRepo/model"
	"github.com/kirillmc/starShipsCompany/inventory/internal/serviceErrors"
	"github.com/samber/lo"
	"go.mongodb.org/mongo-driver/mongo"
)

func setDefaultPartsMap(ctx context.Context, collection *mongo.Collection) error {
	const defaultPartsCount = 11
	defaultParts := make([]interface{}, 0, defaultPartsCount)

	for range defaultPartsCount {
		tempPart := repoModel.Part{
			UUID:          gofakeit.UUID(),
			Name:          gofakeit.Name(),
			Description:   gofakeit.Slogan(),
			Price:         gofakeit.Price(1, 101),
			StockQuantity: gofakeit.Int64(),
			Category:      repoModel.Category(gofakeit.Int32() % 4),
			Dimensions: &repoModel.Dimensions{
				Length: gofakeit.Float64(),
				Width:  gofakeit.Float64(),
				Height: gofakeit.Float64(),
				Weight: gofakeit.Float64(),
			},
			Manufacturer: &repoModel.Manufacturer{
				Name:    gofakeit.Name(),
				Country: gofakeit.Country(),
				Website: gofakeit.Word(),
			},
			Tags:      []string{gofakeit.Word(), gofakeit.Word(), gofakeit.Word()},
			Metadata:  nil,
			CreatedAt: lo.ToPtr(time.Now()),
			UpdatedAt: lo.ToPtr(time.Now()),
		}

		defaultParts = append(defaultParts, tempPart)
	}

	_, err := collection.InsertMany(ctx, defaultParts)
	if err != nil {
		return fmt.Errorf("%w: ошибка при базовом заполнении храниллища деталей: %s",
			serviceErrors.ErrInternalServer, err.Error())
	}

	return nil
}
