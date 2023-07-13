package apiOutputFormatter

type DeliveryDocument struct {
	DeliveryDocumentHeader  []DeliveryDocumentHeader  `json:"Header"`
	DeliveryDocumentItem    []DeliveryDocumentItem    `json:"Item"`
}

type DeliveryDocumentHeader struct {
	DeliveryDocument          int     `json:"DeliveryDocument"`
	DeliverToParty            int     `json:"DeliverToParty"`
	DeliverToPartyName        string  `json:"DeliverToPartyName"`
	DeliverToPlant            string  `json:"DeliverToPlant"`
	DeliverToPlantName        string  `json:"DeliverToPlantName"`
    DeliverFromParty          int     `json:"DeliverFromParty"`
	DeliverFromPartyName      string  `json:"DeliverFromPartyName"`
	DeliverFromPlant          string  `json:"DeliverFromPlant"`
	DeliverFromPlantName      string  `json:"DeliverFromPlantName"`
    HeaderDeliveryStatus      *string `json:"HeaderDeliveryStatus"`
    HeaderBillingStatus       *string `json:"HeaderBillingStatus"`
    IsCancelled               *bool   `json:"IsCancelled"`
	IsMarkedForDeletion       *bool   `json:"IsMarkedForDeletion"`
}

type DeliveryDocumentItem struct {
	DeliveryDocumentItem                      int     `json:"DeliveryDocumentItem"`
	Product                                   string  `json:"Product"`
    DeliveryDocumentItemItemTextByBuyer       string  `json:"DeliveryDocumentItemItemTextByBuyer"`
    DeliveryDocumentItemItemTextBySeller      string  `json:"DeliveryDocumentItemItemTextBySeller"`
    ActualGoodsIssueQtyInBaseUnit             float32 `json:"ActualGoodsIssueQtyInBaseUnit"`
    DeliveryUnit                              string  `json:"DeliveryUnit"`
    ActualGoodsIssueDate                      *string `json:"ActualGoodsIssueDate"`
    ActualGoodsIssueTime                      *string `json:"ActualGoodsIssueTime"`
    ActualGoodsReceiptDate                    *string `json:"ActualGoodsReceiptDate"`
    ActualGoodsReceiptTime                    *string `json:"ActualGoodsReceiptTime"`
    IsCancelled                               *bool   `json:"IsCancelled"`
	IsMarkedForDeletion                       *bool   `json:"IsMarkedForDeletion"`
	Images                                    Images  `json:"Images"`
}
