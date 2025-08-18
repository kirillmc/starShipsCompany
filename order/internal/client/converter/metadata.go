package converter

import inventoryV1 "github.com/kirillmc/starShipsCompany/shared/pkg/proto/inventory/v1"

func MetadataToModel(metadata map[string]*inventoryV1.Value) map[string]interface{} {
	var metadataProto = make(map[string]interface{}, len(metadata))
	for key, value := range metadata {
		metadataProto[key] = value
	}

	return metadataProto
}
