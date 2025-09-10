package part

import (
	"context"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/kirillmc/starShipsCompany/inventory/internal/model"
)

func (s *ServiceSuite) TestListWithEmptyFilter() {
	var (
		ctx = context.Background()

		filter = &model.PartsFilter{}

		part1 = model.Part{
			UUID:        gofakeit.UUID(),
			Name:        gofakeit.Name(),
			Description: gofakeit.Slogan(),
		}
		part1N = part1
		part2  = model.Part{
			UUID:        gofakeit.UUID(),
			Name:        gofakeit.Name(),
			Description: gofakeit.Slogan(),
		}
		part2N = part2
		part3  = model.Part{
			UUID:        gofakeit.UUID(),
			Name:        gofakeit.Name(),
			Description: gofakeit.Slogan(),
		}
		part3N       = part3
		foundedParts = []*model.Part{&part1, &part2, &part3}

		expectedParts = []*model.Part{&part1N, &part2N, &part3N}
	)

	s.repository.On("List", ctx).Return(foundedParts, nil).Once()

	filteredParts, err := s.service.List(ctx, filter)
	s.Assert().NoError(err)
	s.Assert().NotNil(filteredParts)
	s.Assert().ElementsMatch(expectedParts, filteredParts)
}

func (s *ServiceSuite) TestListWithUUIDFilter() {
	var (
		ctx = context.Background()

		part1 = model.Part{
			UUID:        gofakeit.UUID(),
			Name:        gofakeit.Name(),
			Description: gofakeit.Slogan(),
		}
		part1N = part1
		part2  = model.Part{
			UUID:        gofakeit.UUID(),
			Name:        gofakeit.Name(),
			Description: gofakeit.Slogan(),
		}
		part3 = model.Part{
			UUID:        gofakeit.UUID(),
			Name:        gofakeit.Name(),
			Description: gofakeit.Slogan(),
		}
		part3N = part3

		filter = &model.PartsFilter{UUIDs: []model.PartUUID{part1.UUID, part3.UUID}}

		foundedParts = []*model.Part{&part1, &part2, &part3}

		expectedParts = []*model.Part{&part1N, &part3N}
	)

	s.repository.On("List", ctx).Return(foundedParts, nil).Once()

	filteredParts, err := s.service.List(ctx, filter)
	s.Assert().NoError(err)
	s.Assert().NotNil(filteredParts)
	s.Assert().ElementsMatch(expectedParts, filteredParts)
}

func (s *ServiceSuite) TestListWithUUIDAndNameFilter() {
	var (
		ctx = context.Background()

		part1 = model.Part{
			UUID:        gofakeit.UUID(),
			Name:        gofakeit.Name(),
			Description: gofakeit.Slogan(),
		}
		part1N = part1
		part2  = model.Part{
			UUID:        gofakeit.UUID(),
			Name:        gofakeit.Name(),
			Description: gofakeit.Slogan(),
		}
		part3 = model.Part{
			UUID:        gofakeit.UUID(),
			Name:        gofakeit.Name(),
			Description: gofakeit.Slogan(),
		}

		filter = &model.PartsFilter{
			UUIDs: []model.PartUUID{part1.UUID, part3.UUID},
			Names: []string{part1.Name},
		}

		foundedParts = []*model.Part{&part1, &part2, &part3}

		expectedParts = []*model.Part{&part1N}
	)

	s.repository.On("List", ctx).Return(foundedParts, nil).Once()

	filteredParts, err := s.service.List(ctx, filter)
	s.Assert().NoError(err)
	s.Assert().NotNil(filteredParts)
	s.Assert().ElementsMatch(expectedParts, filteredParts)
}
