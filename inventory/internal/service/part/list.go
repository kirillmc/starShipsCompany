package part

import (
	"context"
	"github.com/kirillmc/starShipsCompany/inventory/internal/model"
)

func (s *service) List(ctx context.Context, filter *model.PartFilter) []*model.Part {
	parts := s.repo.List(ctx)
	filteredParts := filterValues(filter, parts)

	return filteredParts
}

func filterValues(filter *model.PartFilter, partsSet map[model.UUID]*model.Part) []*model.Part {
	parts := make([]*model.Part, 0, len(partsSet))
	for _, part := range partsSet {
		parts = append(parts, part)
	}

	if filter == nil {
		return parts
	}

	if len(filter.UUIDs) > 0 {
		parts = applySimpleFieldFilter(filter.UUIDs, parts, func(p *model.Part) model.UUID { return p.UUID })
	}

	if len(filter.Names) > 0 {
		parts = applySimpleFieldFilter(filter.Names, parts, func(p *model.Part) string { return p.Name })
	}

	if len(filter.Categories) > 0 {
		parts = applySimpleFieldFilter(filter.Categories, parts, func(p *model.Part) model.Category { return p.Category })
	}

	if len(filter.ManufacturerCountries) > 0 {
		parts = applySimpleFieldFilter(filter.ManufacturerCountries, parts, func(p *model.Part) string { return p.Manufacturer.Country })
	}

	if len(filter.Tags) > 0 {
		parts = applyTagsFilter(filter.Tags, parts)
	}

	return parts
}

func applySimpleFieldFilter[T allowedFilterTypes](filterValues []T, allowedParts []*model.Part,
	getField func(*model.Part) T) []*model.Part {
	if filterValues == nil {
		return allowedParts
	}

	valueSet := getFilterValuesSet[T](filterValues)
	var res []*model.Part
	for _, part := range allowedParts {
		if _, ok := valueSet[getField(part)]; ok {
			res = append(res, part)
		}
	}
	return res
}

func applyTagsFilter(filterValues []string, allowedParts []*model.Part) []*model.Part {
	if filterValues == nil {
		return allowedParts
	}

	tagsSet := getFilterValuesSet[string](filterValues)
	var res []*model.Part

	for _, part := range allowedParts {
		for _, tag := range part.Tags {
			if _, ok := tagsSet[tag]; ok {
				res = append(res, part)
				break
			}
		}
	}

	return res
}

type allowedFilterTypes interface {
	string | model.Category
}

func getFilterValuesSet[T allowedFilterTypes](filterValues []T) map[T]struct{} {
	if len(filterValues) == 0 {
		return nil
	}

	filterSet := make(map[T]struct{})

	for _, filterValue := range filterValues {
		if _, ok := filterSet[filterValue]; !ok {
			filterSet[filterValue] = struct{}{}
		}
	}

	return filterSet
}
