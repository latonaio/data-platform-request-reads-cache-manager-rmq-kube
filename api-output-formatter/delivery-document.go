package apiOutputFormatter

type DeliveryDocument struct {
	DeliveryDocumentHeader         []DeliveryDocumentHeader         `json:"Header"`
	DeliveryDocumentHeaderWithItem []DeliveryDocumentHeaderWithItem `json:"HeaderWithItem"`
	DeliveryDocumentItem           []DeliveryDocumentItem           `json:"Item"`
	DeliveryDocumentSingleUnit     []DeliveryDocumentSingleUnit     `json:"SingleUnit"`
	MountPath                      *string                          `json:"mount_path"`
	Accepter                       []string                         `json:"Accepter"`
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
	IsExportImport          *bool   `json:"IsExportImport"`
	OrderID                 *int    `json:"OrderID"`
	OrderItem               *int    `json:"OrderItem"`
	Contract                *int    `json:"Contract"`
	ContractItem            *int    `json:"ContractItem"`
	Project                 *int    `json:"Project"`
	WBSElement              *int    `json:"WBSElement"`
	WBSElementDescription   *string `json:"WBSElementDescription"`
	ProductionOrder         *int    `json:"ProductionOrder"`
	ProductionOrderItem     *int    `json:"ProductionOrderItem"`
	Incoterms               *string `json:"Incoterms"`
	IncotermsText           *string `json:"IncotermsText"`
}

type DeliveryDocumentItem struct {
	DeliveryDocument                     int      `json:"DeliveryDocument"`
	DeliveryDocumentItem                 int      `json:"DeliveryDocumentItem"`
	DeliveryDocumentItemCategory         string   `json:"DeliveryDocumentItemCategory"`
	ProductSpecification                 *string  `json:"ProductSpecification"`
	SizeOrDimensionText                  *string  `json:"SizeOrDimensionText"`
	Buyer                                int      `json:"Buyer"`
	BuyerName                            string   `json:"BuyerName"`
	Seller                               int      `json:"Seller"`
	SellerName                           string   `json:"SellerName"`
	Product                              string   `json:"Product"`
	DeliveryDocumentItemText             string   `json:"DeliveryDocumentItemText"`
	DeliveryDocumentItemItemTextByBuyer  string   `json:"DeliveryDocumentItemItemTextByBuyer"`
	DeliveryDocumentItemItemTextBySeller string   `json:"DeliveryDocumentItemItemTextBySeller"`
	PlannedGoodsIssueQuantity            float32  `json:"PlannedGoodsIssueQuantity"`
	PlannedGoodsIssueQtyInBaseUnit       float32  `json:"PlannedGoodsIssueQtyInBaseUnit"`
	DeliveryUnit                         string   `json:"DeliveryUnit"`
	BaseUnit                             string   `json:"BaseUnit"`
	PlannedGoodsIssueDate                string   `json:"PlannedGoodsIssueDate"`
	PlannedGoodsIssueTime                string   `json:"PlannedGoodsIssueTime"`
	PlannedGoodsReceiptDate              string   `json:"PlannedGoodsReceiptDate"`
	PlannedGoodsReceiptTime              string   `json:"PlannedGoodsReceiptTime"`
	ActualGoodsIssueDate                 *string  `json:"ActualGoodsIssueDate"`
	ActualGoodsIssueTime                 *string  `json:"ActualGoodsIssueTime"`
	ActualGoodsReceiptDate               *string  `json:"ActualGoodsReceiptDate"`
	ActualGoodsReceiptTime               *string  `json:"ActualGoodsReceiptTime"`
	ProductNetWeight                     *float32 `json:"ProductNetWeight"`
	ItemWeightUnit                       *string  `json:"ItemWeightUnit"`
	ItemNetWeight                        *float32 `json:"ItemNetWeight"`
	ItemGrossWeight                      *float32 `json:"ItemGrossWeight"`
	OrderID                              *int     `json:"OrderID"`
	OrderItem                            *int     `json:"OrderItem"`
	Project                              *int     `json:"Project"`
	WBSElement                           *int     `json:"WBSElement"`
	IsCancelled                          *bool    `json:"IsCancelled"`
	IsMarkedForDeletion                  *bool    `json:"IsMarkedForDeletion"`
	Images                               Images   `json:"Images"`
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
	OrderID                 *int   `json:"OrderID"`
	OrderItem               *int   `json:"OrderItem"`
	Images                  Images `json:"Images"`
}
