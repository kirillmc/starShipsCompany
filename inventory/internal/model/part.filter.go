package model

type PartFilter struct {
	UUIDs                 []UUID
	Names                 []string
	Categories            []Category
	ManufacturerCountries []string
	Tags                  []string
}
