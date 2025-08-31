package v1

import (
	"context"
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/kirillmc/starShipsCompany/inventory/internal/converter"
	"github.com/kirillmc/starShipsCompany/inventory/internal/model"
	inventoryV1 "github.com/kirillmc/starShipsCompany/shared/pkg/proto/inventory/v1"
	"github.com/samber/lo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ServiceSuite) TestGetPartSuccess() {
	var (
		ctx = context.Background()

		partUUID = gofakeit.UUID()

		req = &inventoryV1.GetPartRequest{
			Uuid: partUUID,
		}

		foundedPart = &model.Part{
			UUID:          partUUID,
			Name:          gofakeit.Name(),
			Description:   gofakeit.Slogan(),
			Price:         gofakeit.Price(0, 101),
			StockQuantity: int64(gofakeit.UintRange(0, 101)),
			Category:      model.ENGINE,
			Dimensions: &model.Dimensions{
				Length: 11,
				Width:  11,
				Height: 11,
				Weight: 11,
			},
			Manufacturer: &model.Manufacturer{
				Name:    gofakeit.FirstName(),
				Country: gofakeit.Country(),
				Website: gofakeit.URL(),
			},
			Tags:      []model.Tag{gofakeit.EmojiTag(), gofakeit.EmojiTag(), gofakeit.EmojiTag()},
			Metadata:  nil,
			CreatedAt: lo.ToPtr(time.Now()),
			UpdatedAt: nil,
		}
		foundedErr error = nil

		expectedPart = converter.PartToProto(foundedPart)
	)

	s.service.On("Get", ctx, req.GetUuid()).Return(foundedPart, foundedErr).Once()

	resp, err := s.api.GetPart(ctx, req)
	s.Assert().NoError(err)
	s.Assert().NotNil(resp)
	s.Assert().Equal(expectedPart, resp.Part)
}

func (s *ServiceSuite) TestFailedGetPart() {
	var (
		ctx = context.Background()

		partUUID = gofakeit.UUID()

		req = &inventoryV1.GetPartRequest{
			Uuid: partUUID,
		}

		foundedPart = &model.Part{}
		foundedErr  = fmt.Errorf("failed to execute Gwt method: part with PartUUID %s not found", partUUID)

		expectedErr = status.Errorf(codes.NotFound, "failed to find part: %s", foundedErr)
	)

	s.service.On("Get", ctx, req.GetUuid()).Return(foundedPart, foundedErr).Once()

	_, err := s.api.GetPart(ctx, req)
	s.Assert().Error(err)
	s.Assert().Equal(expectedErr, err)
}
