package model

type PartsFilter struct {
	UUIDs                 []UUID
	Names                 []string
	Categories            []Category
	ManufacturerCountries []string
	Tags                  []string
}
