package converter

import inventoryV1 "github.com/kirillmc/starShipsCompany/shared/pkg/proto/inventory/v1"

func ToModelMetadata(metadata map[string]*inventoryV1.Value) map[string]interface{} {
	metadataMapped := make(map[string]interface{}, len(metadata))
	for key, value := range metadata {
		metadataMapped[key] = value
	}
	return metadataMapped
}
