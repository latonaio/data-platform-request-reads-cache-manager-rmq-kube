package apiInputReader

type ProductionOrder struct {
	ProductionOrderHeader                 *ProductionOrderHeader
	ProductionOrderItem                   *ProductionOrderItem
	ProductionOrderItemComponent          *ProductionOrderItemComponent
	ProductionOrderItemOperation          *ProductionOrderItemOperation
	ProductionOrderItemOperationComponent *ProductionOrderItemOperationComponent
	ProductionOrderDocHeaderDoc           *ProductionOrderDocHeaderDoc
	ProductionOrderDocItemDoc             *ProductionOrderDocItemDoc
}

type ProductionOrderHeader struct {
	ProductionOrder     int   `json:"ProductionOrder"`
	IsReleased          *bool `json:"IsReleased"`
	IsCancelled         *bool `json:"IsCancelled"`
	IsMarkedForDeletion *bool `json:"IsMarkedForDeletion"`
}

type ProductionOrderItem struct {
	ProductionOrder     int   `json:"ProductionOrder"`
	ProductionOrderItem int   `json:"ProductionOrderItem"`
	IsReleased          *bool `json:"IsReleased"`
	IsCancelled         *bool `json:"IsCancelled"`
	IsMarkedForDeletion *bool `json:"IsMarkedForDeletion"`
}

type ProductionOrderItemComponent struct {
	ProductionOrder     int   `json:"ProductionOrder"`
	ProductionOrderItem int   `json:"ProductionOrderItem"`
	BillOfMaterial      int   `json:"BillOfMaterial"`
	BillOfMaterialItem  int   `json:"BillOfMaterialItem"`
	IsReleased          *bool `json:"IsReleased"`
	IsCancelled         *bool `json:"IsCancelled"`
	IsMarkedForDeletion *bool `json:"IsMarkedForDeletion"`
}

type ProductionOrderItemOperation struct {
	ProductionOrder     int   `json:"ProductionOrder"`
	ProductionOrderItem int   `json:"ProductionOrderItem"`
	Operations          int   `json:"Operations"`
	OperationsItem      int   `json:"OperationsItem"`
	OperationID         int   `json:"OperationID"`
	IsReleased          *bool `json:"IsReleased"`
	IsCancelled         *bool `json:"IsCancelled"`
	IsMarkedForDeletion *bool `json:"IsMarkedForDeletion"`
}

type ProductionOrderItemOperationComponent struct {
	ProductionOrder     int   `json:"ProductionOrder"`
	ProductionOrderItem int   `json:"ProductionOrderItem"`
	IsReleased          *bool `json:"IsReleased"`
	IsCancelled         *bool `json:"IsCancelled"`
	IsMarkedForDeletion *bool `json:"IsMarkedForDeletion"`
}

type ProductionOrderDocHeaderDoc struct {
	ProductionOrder          *int    `json:"ProductionOrder"`
	DocType                  *string `json:"DocType"`
	DocIssuerBusinessPartner *int    `json:"DocIssuerBusinessPartner"`
}

type ProductionOrderDocItemDoc struct {
	ProductionOrder          int    `json:"ProductionOrder"`
	ProductionOrderItem      int    `json:"ProductionOrderItem"`
	DocType                  string `json:"DocType"`
	DocIssuerBusinessPartner int    `json:"DocIssuerBusinessPartner"`
}
