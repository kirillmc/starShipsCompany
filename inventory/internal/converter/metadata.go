package converter

import inventoryV1 "github.com/kirillmc/starShipsCompany/shared/pkg/proto/inventory/v1"

func ToProtoMetadata(metadata map[string]interface{}) map[string]*inventoryV1.Value {
	metadataMapped := make(map[string]*inventoryV1.Value, len(metadata))
	for key, value := range metadata {
		switch v := value.(type) {
		case string:
			metadataMapped[key] = &inventoryV1.Value{
				Kind: &inventoryV1.Value_StringValue{StringValue: v},
			}
		case int64:
			metadataMapped[key] = &inventoryV1.Value{
				Kind: &inventoryV1.Value_Int64Value{Int64Value: v},
			}
		case float64:
			metadataMapped[key] = &inventoryV1.Value{
				Kind: &inventoryV1.Value_DoubleValue{DoubleValue: v},
			}
		case bool:
			metadataMapped[key] = &inventoryV1.Value{
				Kind: &inventoryV1.Value_BoolValue{BoolValue: v},
			}
		}
	}

	return metadataMapped
}
