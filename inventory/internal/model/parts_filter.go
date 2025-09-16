package model

type PartsFilter struct {
	UUIDs                 []PartUUID
	Names                 []string
	Categories            []Category
	ManufacturerCountries []string
	Tags                  []string
}
