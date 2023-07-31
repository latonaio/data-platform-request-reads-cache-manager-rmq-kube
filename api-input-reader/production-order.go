package apiInputReader

type ProductionOrder struct {
	ProductionOrderHeader *ProductionOrderHeader
	ProductionOrderItems  *ProductionOrderItems
	ProductionOrderItem   *ProductionOrderItem
}

type ProductionOrderHeader struct {
	ProductionOrder     int   `json:"ProductionOrder"`
	IsReleased          *bool `json:"IsReleased"`
	IsMarkedForDeletion *bool `json:"IsMarkedForDeletion"`
}

type ProductionOrderItems struct {
	ProductionOrder     int   `json:"ProductionOrder"`
	IsMarkedForDeletion *bool `json:"IsMarkedForDeletion"`
}

type ProductionOrderItem struct {
	ProductionOrder     int   `json:"ProductionOrder"`
	ProductionOrderItem int   `json:"ProductionOrderItem"`
	IsMarkedForDeletion *bool `json:"IsMarkedForDeletion"`
}
