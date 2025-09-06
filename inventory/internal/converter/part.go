package converter

import (
	"time"

	"github.com/kirillmc/starShipsCompany/inventory/internal/model"
	inventoryV1 "github.com/kirillmc/starShipsCompany/shared/pkg/proto/inventory/v1"
	"github.com/samber/lo"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToProtoPart(part *model.Part) *inventoryV1.Part {
	partMapped := &inventoryV1.Part{
		Uuid:          part.UUID,
		Name:          part.Name,
		Description:   part.Description,
		Price:         part.Price,
		StockQuantity: part.StockQuantity,
		Category:      ToProtoCategory(part.Category),
		Dimensions:    ToProtoDimensions(part.Dimensions),
		Manufacturer:  ToProtoManufacturer(part.Manufacturer),
		Tags:          part.Tags,
		Metadata:      ToProtoMetadata(part.Metadata),
		CreatedAt:     timestamppb.New(lo.FromPtrOr(part.CreatedAt, time.Time{})),
		UpdatedAt:     timestamppb.New(lo.FromPtrOr(part.UpdatedAt, time.Time{})),
	}

	return partMapped
}
