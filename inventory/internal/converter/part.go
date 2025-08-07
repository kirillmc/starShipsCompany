package converter

import (
	"github.com/kirillmc/starShipsCompany/inventory/internal/model"
	inventoryV1 "github.com/kirillmc/starShipsCompany/shared/pkg/proto/inventory/v1"
	"github.com/samber/lo"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func PartToProto(part *model.Part) *inventoryV1.Part {
	partProto := &inventoryV1.Part{
		Uuid:          part.UUID,
		Name:          part.Name,
		Description:   part.Description,
		Price:         part.Price,
		StockQuantity: part.StockQuantity,
		Category:      CategoryToProto(part.Category),
		Dimensions:    DimensionsToProto(part.Dimensions),
		Manufacturer:  ManufacturerToProto(part.Manufacturer),
		Tags:          part.Tags,
		Metadata:      MetadataToProto(part.Metadata),
		CreatedAt:     timestamppb.New(lo.FromPtrOr(part.CreatedAt, time.Time{})),
		UpdatedAt:     timestamppb.New(lo.FromPtrOr(part.UpdatedAt, time.Time{})),
	}

	return partProto
}
