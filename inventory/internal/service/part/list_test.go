package part

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/kirillmc/starShipsCompany/inventory/internal/model"
)

func (s *ServiceSuite) TestListWithEmptyFilter() {
	var (
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
		foundedParts = map[model.PartUUID]*model.Part{
			part1.UUID: &part1,
			part2.UUID: &part2,
			part3.UUID: &part3,
		}

		expectedParts = []*model.Part{&part1N, &part2N, &part3N}
	)

	s.repository.On("List", s.ctx).Return(foundedParts).Once()

	filteredParts := s.service.List(s.ctx, filter)
	s.Assert().NotNil(filteredParts)
	s.Assert().ElementsMatch(expectedParts, filteredParts)
}

func (s *ServiceSuite) TestListWithUUIDFilter() {
	var (
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

		foundedParts = map[model.PartUUID]*model.Part{
			part1.UUID: &part1,
			part2.UUID: &part2,
			part3.UUID: &part3,
		}

		expectedParts = []*model.Part{&part1N, &part3N}
	)

	s.repository.On("List", s.ctx).Return(foundedParts).Once()

	filteredParts := s.service.List(s.ctx, filter)
	s.Assert().NotNil(filteredParts)
	s.Assert().ElementsMatch(expectedParts, filteredParts)
}

func (s *ServiceSuite) TestListWithUUIDAndNameFilter() {
	var (
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

		foundedParts = map[model.PartUUID]*model.Part{
			part1.UUID: &part1,
			part2.UUID: &part2,
			part3.UUID: &part3,
		}

		expectedParts = []*model.Part{&part1N}
	)

	s.repository.On("List", s.ctx).Return(foundedParts).Once()

	filteredParts := s.service.List(s.ctx, filter)
	s.Assert().NotNil(filteredParts)
	s.Assert().ElementsMatch(expectedParts, filteredParts)
}
