package apiInputReader

type ProductionOrder struct {
	ProductionOrderHeader                 *ProductionOrderHeader
	ProductionOrderItem                   *ProductionOrderItem
	ProductionOrderItemOperation          *ProductionOrderItemOperation
	ProductionOrderItemOperationComponent *ProductionOrderItemOperationComponent
}

type ProductionOrderHeader struct {
	ProductionOrder     int   `json:"ProductionOrder"`
	IsReleased          *bool `json:"IsReleased"`
	IsMarkedForDeletion *bool `json:"IsMarkedForDeletion"`
}

type ProductionOrderItem struct {
	ProductionOrder     int   `json:"ProductionOrder"`
	ProductionOrderItem int   `json:"ProductionOrderItem"`
	IsMarkedForDeletion *bool `json:"IsMarkedForDeletion"`
}

type ProductionOrderItemOperation struct {
	ProductionOrder     int   `json:"ProductionOrder"`
	ProductionOrderItem int   `json:"ProductionOrderItem"`
	Operations          int   `json:"Operations"`
	OperationsItem      int   `json:"OperationsItem"`
	OperationID         int   `json:"OperationID"`
	IsMarkedForDeletion *bool `json:"IsMarkedForDeletion"`
}

type ProductionOrderItemOperationComponent struct {
	ProductionOrder     int   `json:"ProductionOrder"`
	ProductionOrderItem int   `json:"ProductionOrderItem"`
	IsMarkedForDeletion *bool `json:"IsMarkedForDeletion"`
}
