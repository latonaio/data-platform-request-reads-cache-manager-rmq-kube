package apiOutputFormatter

type DeliveryDocument struct {
	DeliveryDocumentHeader         []DeliveryDocumentHeader         `json:"Header"`
	DeliveryDocumentHeaderWithItem []DeliveryDocumentHeaderWithItem `json:"HeaderWithItem"`
	DeliveryDocumentItem           []DeliveryDocumentItem           `json:"Item"`
	DeliveryDocumentSingleUnit     []DeliveryDocumentSingleUnit     `json:"SingleUnit"`
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
	DeliveryDocument        int     `json:"DeliveryDocument"`
	DeliveryDocumentDate    string  `json:"DeliveryDocumentDate"`
	DeliverToParty          int     `json:"DeliverToParty"`
	DeliverToPartyName      string  `json:"DeliverToPartyName"`
	DeliverToPlant          string  `json:"DeliverToPlant"`
	DeliverToPlantName      string  `json:"DeliverToPlantName"`
	DeliverFromParty        int     `json:"DeliverFromParty"`
	DeliverFromPartyName    string  `json:"DeliverFromPartyName"`
	DeliverFromPlant        string  `json:"DeliverFromPlant"`
	DeliverFromPlantName    string  `json:"DeliverFromPlantName"`
	PlannedGoodsIssueDate   string  `json:"PlannedGoodsIssueDate"`
	PlannedGoodsIssueTime   string  `json:"PlannedGoodsIssueTime"`
	PlannedGoodsReceiptDate string  `json:"PlannedGoodsReceiptDate"`
	PlannedGoodsReceiptTime string  `json:"PlannedGoodsReceiptTime"`
	HeaderGrossWeight       float32 `json:"HeaderGrossWeight"`
	HeaderNetWeight         float32 `json:"HeaderNetWeight"`
	HeaderWeightUnit        string  `json:"HeaderWeightUnit"`
}

type DeliveryDocumentItem struct {
	DeliveryDocumentItem                 int     `json:"DeliveryDocumentItem"`
	Product                              string  `json:"Product"`
	DeliveryDocumentItemText             string  `json:"DeliveryDocumentItemText"`
	DeliveryDocumentItemItemTextByBuyer  string  `json:"DeliveryDocumentItemItemTextByBuyer"`
	DeliveryDocumentItemItemTextBySeller string  `json:"DeliveryDocumentItemItemTextBySeller"`
	PlannedGoodsIssueQuantity            float32 `json:"PlannedGoodsIssueQuantity"`
	PlannedGoodsIssueQtyInBaseUnit       float32 `json:"PlannedGoodsIssueQtyInBaseUnit"`
	DeliveryUnit                         string  `json:"DeliveryUnit"`
	BaseUnit                             string  `json:"BaseUnit"`
	PlannedGoodsIssueDate                string  `json:"PlannedGoodsIssueDate"`
	PlannedGoodsIssueTime                string  `json:"PlannedGoodsIssueTime"`
	PlannedGoodsReceiptDate              string  `json:"PlannedGoodsReceiptDate"`
	PlannedGoodsReceiptTime              string  `json:"PlannedGoodsReceiptTime"`
	IsCancelled                          *bool   `json:"IsCancelled"`
	IsMarkedForDeletion                  *bool   `json:"IsMarkedForDeletion"`
	Images                               Images  `json:"Images"`
}

type DeliveryDocumentSingleUnit struct {
	DeliveryDocument        int    `json:"DeliveryDocument"`
	DeliveryDocumentItem    int    `json:"DeliveryDocumentItem"`
	PlannedGoodsIssueDate   string `json:"PlannedGoodsIssueDate"`
	PlannedGoodsIssueTime   string `json:"PlannedGoodsIssueTime"`
	PlannedGoodsReceiptDate string `json:"PlannedGoodsReceiptDate"`
	PlannedGoodsReceiptTime string `json:"PlannedGoodsReceiptTime"`
	DeliverToParty          int    `json:"DeliverToParty"`
	DeliverToPartyName      string `json:"DeliverToPartyName"`
	DeliverToPlant          string `json:"DeliverToPlant"`
	DeliverToPlantName      string `json:"DeliverToPlantName"`
	DeliverFromParty        int    `json:"DeliverFromParty"`
	DeliverFromPartyName    string `json:"DeliverFromPartyName"`
	DeliverFromPlant        string `json:"DeliverFromPlant"`
	DeliverFromPlantName    string `json:"DeliverFromPlantName"`
	Images                  Images `json:"Images"`
}
