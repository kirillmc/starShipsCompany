package part

import (
	"github.com/brianvoe/gofakeit/v7"
	inventoryV1 "github.com/kirillmc/starShipsCompany/inventory/internal/repository/model"
	"github.com/samber/lo"
	"time"
)

func setDefaultPartsMap() map[string]*inventoryV1.Part {
	const defaultPartsCoount = 11
	defaultMap := make(map[string]*inventoryV1.Part)
	for range defaultPartsCoount {
		UUID := gofakeit.UUID()
		defaultMap[UUID] = &inventoryV1.Part{
			UUID:          UUID,
			Name:          gofakeit.Name(),
			Description:   gofakeit.Slogan(),
			Price:         gofakeit.Price(1, 101),
			StockQuantity: gofakeit.Int64(),
			Category:      inventoryV1.Category(gofakeit.Int32() % 4),
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
			CreatedAt: lo.ToPtr(time.Now()),
			UpdatedAt: lo.ToPtr(time.Now()),
		}
	}

	return defaultMap
}
