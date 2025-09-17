package model

type (
	OrderPartID = uint64
)

type OrderPart struct {
	ID       OrderPartID
	OrderID  OrderID
	PartUUID PartUUID
}
