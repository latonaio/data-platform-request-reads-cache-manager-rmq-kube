package apiOutputFormatter

type PurchaseRequisition struct {
	PurchaseRequisitionHeader         []PurchaseRequisitionHeader         `json:"Header"`
	PurchaseRequisitionHeaderWithItem []PurchaseRequisitionHeaderWithItem `json:"HeaderWithItem"`
	PurchaseRequisitionItem           []PurchaseRequisitionItem           `json:"Item"`
}

type PurchaseRequisitionHeader struct {
	PurchaseRequisition     int     `json:"PurchaseRequisition"`
	Buyer                   int     `json:"Buyer"`
	BuyerName               string  `json:"BuyerName"`
	HeaderOrdertatus        *string `json:"HeaderOrdertatus"`
	PurchaseRequisitionType string  `json:"PurchaseRequisitionType"`
	IsReleased              *bool   `json:"IsReleased"`
	IsCancelled             *bool   `json:"IsCancelled"`
	IsMarkedForDeletion     *bool   `json:"IsMarkedForDeletion"`
}

type PurchaseRequisitionHeaderWithItem struct {
	PurchaseRequisition     int    `json:"PurchaseRequisition"`
	PurchaseRequisitionDate string `json:"PurchaseRequisitionDate"`
	PurchaseRequisitionType string `json:"PurchaseRequisitionType"`
	Buyer                   int    `json:"Buyer"`
	BuyerName               string `json:"BuyerName"`
}

type PurchaseRequisitionItem struct {
	PurchaseRequisitionItem             int     `json:"PurchaseRequisition"`
	Product                             string  `json:"Product"`
	PurchaseRequisitionItemTextByBuyer  string  `json:"PurchaseRequisitionItemTextByBuyer"`
	PurchaseRequisitionItemTextBySeller string  `json:"PurchaseRequisitionItemTextBySeller"`
	RequestedQuantityInDeliveryUnit     float32 `json:"RequestedQuantityInDeliveryUnit"`
	DeliveryUnit                        string  `json:"DeliveryUnit"`
	PurchaseRequisitionItemPrice        float32 `json:"PurchaseRequisitionItemPrice"`
	RequestedDeliveryDate               string  `json:"RequestedDeliveryDate"`
	IsCancelled                         *bool   `json:"IsCancelled"`
	IsMarkedForDeletion                 *bool   `json:"IsMarkedForDeletion"`
	Images                              Images  `json:"Images"`
}
