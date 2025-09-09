package part

import (
	model2 "github.com/kirillmc/starShipsCompany/inventory/internal/repository/mongo/model"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/samber/lo"
)

func setDefaultPartsMap() map[model2.PartUUID]*model2.Part {
	const defaultPartsCoount = 11
	defaultMap := make(map[string]*model2.Part)
	for range defaultPartsCoount {
		UUID := gofakeit.UUID()
		defaultMap[UUID] = &model2.Part{
			UUID:          UUID,
			Name:          gofakeit.Name(),
			Description:   gofakeit.Slogan(),
			Price:         gofakeit.Price(1, 101),
			StockQuantity: gofakeit.Int64(),
			Category:      model2.Category(gofakeit.Int32() % 4),
			Dimensions: &model2.Dimensions{
				Length: gofakeit.Float64(),
				Width:  gofakeit.Float64(),
				Height: gofakeit.Float64(),
				Weight: gofakeit.Float64(),
			},
			Manufacturer: &model2.Manufacturer{
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
