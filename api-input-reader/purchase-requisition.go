package apiInputReader

type PurchaseRequisition struct {
	PurchaseRequisitionHeader *PurchaseRequisitionHeader
	PurchaseRequisitionItems  *PurchaseRequisitionItems
	PurchaseRequisitionItem   *PurchaseRequisitionItem
}

type PurchaseRequisitionHeader struct {
	PurchaseRequisition             int     `json:"PurchaseRequisition"`
	Buyer                           *int    `json:"Buyer"`
	HeaderOrderStatus				*string `json:"HeaderOrderStatus"`
	HeaderCompleteOrderIsDefined	*bool   `json:"HeaderCompleteOrderIsDefined"`
	IsCancelled                     *bool   `json:"IsCancelled"`
	IsMarkedForDeletion             *bool   `json:"IsMarkedForDeletion"`
}

type PurchaseRequisitionItems struct {
	PurchaseRequisition           	int     `json:"PurchaseRequisition"`
	ItemCompleteOrderIsDefined	    *bool   `json:"ItemCompleteOrderIsDefined"`
	IsCancelled                   	*bool   `json:"IsCancelled"`
	IsMarkedForDeletion           	*bool   `json:"IsMarkedForDeletion"`
}

type PurchaseRequisitionItem struct {
	PurchaseRequisition           int     `json:"PurchaseRequisition"`
	PurchaseRequisitionItem       int     `json:"PurchaseRequisitionItem"`
	ItemCompleteOrderIsDefined	  *bool   `json:"ItemCompleteOrderIsDefined"`
	IsCancelled                   *bool   `json:"IsCancelled"`
	IsMarkedForDeletion           *bool   `json:"IsMarkedForDeletion"`
}
