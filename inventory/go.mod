module github.com/kirillmc/starShipsCompany/inventory

go 1.24.5

replace github.com/kirillmc/starShipsCompany/shared => ../shared

require (
	github.com/brianvoe/gofakeit/v7 v7.3.0
	github.com/kirillmc/starShipsCompany/shared v0.0.0-00010101000000-000000000000
	github.com/samber/lo v1.51.0
	github.com/stretchr/testify v1.11.1
	google.golang.org/grpc v1.73.0
	google.golang.org/protobuf v1.36.6
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/objx v0.5.2 // indirect
	golang.org/x/net v0.40.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.25.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250324211829-b45e905df463 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
