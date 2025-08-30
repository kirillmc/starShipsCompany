package part

import (
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/kirillmc/starShipsCompany/inventory/internal/repository/model"
	"github.com/samber/lo"
)

func setDefaultPartsMap() map[model.PartUUID]*model.Part {
	const defaultPartsCoount = 11
	defaultMap := make(map[string]*model.Part)
	for range defaultPartsCoount {
		UUID := gofakeit.UUID()
		defaultMap[UUID] = &model.Part{
			UUID:          UUID,
			Name:          gofakeit.Name(),
			Description:   gofakeit.Slogan(),
			Price:         gofakeit.Price(1, 101),
			StockQuantity: gofakeit.Int64(),
			Category:      model.Category(gofakeit.Int32() % 4),
			Dimensions: &model.Dimensions{
				Length: gofakeit.Float64(),
				Width:  gofakeit.Float64(),
				Height: gofakeit.Float64(),
				Weight: gofakeit.Float64(),
			},
			Manufacturer: &model.Manufacturer{
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
