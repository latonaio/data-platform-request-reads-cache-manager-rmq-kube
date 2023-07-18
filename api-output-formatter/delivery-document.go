package apiOutputFormatter

type DeliveryDocument struct {
	DeliveryDocumentHeader         []DeliveryDocumentHeader         `json:"Header"`
	DeliveryDocumentHeaderWithItem []DeliveryDocumentHeaderWithItem `json:"HeaderWithItem"`
	DeliveryDocumentItem           []DeliveryDocumentItem           `json:"Item"`
}

type DeliveryDocumentHeader struct {
	DeliveryDocument        int     `json:"DeliveryDocument"`
	DeliverToParty          int     `json:"DeliverToParty"`
	DeliverToPartyName      string  `json:"DeliverToPartyName"`
	DeliverToPlant          string  `json:"DeliverToPlant"`
	DeliverToPlantName      string  `json:"DeliverToPlantName"`
	DeliverFromParty        int     `json:"DeliverFromParty"`
	DeliverFromPartyName    string  `json:"DeliverFromPartyName"`
	DeliverFromPlant        string  `json:"DeliverFromPlant"`
	DeliverFromPlantName    string  `json:"DeliverFromPlantName"`
	HeaderDeliveryStatus    *string `json:"HeaderDeliveryStatus"`
	HeaderBillingStatus     *string `json:"HeaderBillingStatus"`
	PlannedGoodsReceiptDate string  `json:"PlannedGoodsReceiptDate"`
	PlannedGoodsReceiptTime string  `json:"PlannedGoodsReceiptTime"`
	IsCancelled             *bool   `json:"IsCancelled"`
	IsMarkedForDeletion     *bool   `json:"IsMarkedForDeletion"`
}

type DeliveryDocumentHeaderWithItem struct {
	DeliveryDocument        int    `json:"DeliveryDocument"`
	DeliverToParty          int    `json:"DeliverToParty"`
	DeliverToPartyName      string `json:"DeliverToPartyName"`
	DeliverToPlant          string `json:"DeliverToPlant"`
	DeliverToPlantName      string `json:"DeliverToPlantName"`
	DeliverFromParty        int    `json:"DeliverFromParty"`
	DeliverFromPartyName    string `json:"DeliverFromPartyName"`
	DeliverFromPlant        string `json:"DeliverFromPlant"`
	DeliverFromPlantName    string `json:"DeliverFromPlantName"`
	PlannedGoodsReceiptDate string `json:"PlannedGoodsReceiptDate"`
	PlannedGoodsReceiptTime string `json:"PlannedGoodsReceiptTime"`
}

type DeliveryDocumentItem struct {
	DeliveryDocumentItem                 int     `json:"DeliveryDocumentItem"`
	Product                              string  `json:"Product"`
	DeliveryDocumentItemItemTextByBuyer  string  `json:"DeliveryDocumentItemItemTextByBuyer"`
	DeliveryDocumentItemItemTextBySeller string  `json:"DeliveryDocumentItemItemTextBySeller"`
	PlannedGoodsIssueQuantity            float32 `json:"PlannedGoodsIssueQuantity"`
	DeliveryUnit                         string  `json:"DeliveryUnit"`
	PlannedGoodsIssueDate                string  `json:"PlannedGoodsIssueDate"`
	PlannedGoodsIssueTime                string  `json:"PlannedGoodsIssueTime"`
	PlannedGoodsReceiptDate              string  `json:"PlannedGoodsReceiptDate"`
	PlannedGoodsReceiptTime              string  `json:"PlannedGoodsReceiptTime"`
	IsCancelled                          *bool   `json:"IsCancelled"`
	IsMarkedForDeletion                  *bool   `json:"IsMarkedForDeletion"`
	Images                               Images  `json:"Images"`
}
