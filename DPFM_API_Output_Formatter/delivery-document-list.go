package dpfm_api_output_formatter

type DeliveryDocumentList struct {
	Pulldown          DeliveryDocumentListPullDown `json:"Pulldown"`
	DeliveryDocuments []DeliveryDocument           `json:"DeliveryDocuments"`
}
type DeliveryDocument struct {
	DeliveryDocument          int `json:"DeliveryDocument"`
	SupplyChainRelationshipID int `json:"SupplyChainRelationshipID"`

	DeliverToParty     int     `json:"DeliverToParty"`
	DeliverToPartyName *string `json:"DeliverToPartyName"`
	DeliverToPlant     string  `json:"DeliverToPlant"`
	DeliverToPlantName string  `json:"DeliverToPlantName"`

	DeliverFromParty     int     `json:"DeliverFromParty"`
	DeliverFromPartyName *string `json:"DeliverFromPartyName"`
	DeliverFromPlant     string  `json:"DeliverFromPlant"`
	DeliverFromPlantName string  `json:"DeliverFromPlantName"`

	HeaderDeliveryStatus    *string `json:"HeaderDeliveryStatus"`
	HeaderBillingStatus     *string `json:"HeaderBillingStatus"`
	PlannedGoodsReceiptDate *string `json:"PlannedGoodsReceiptDate"`
	PlannedGoodsReceiptTime *string `json:"PlannedGoodsReceiptTime"`

	IsCancelled         *bool `json:"IsCancelled"`
	IsMarkedForDeletion *bool `json:"IsMarkedForDeletion"`
}

type DeliveryDocumentListPullDown struct {
	SupplyChains map[int]Deliver `json:"SupplyChains"`
}

type DeliveryPlant struct {
	SupplyChainRelationshipDeliveryID      int `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID int `json:"SupplyChainRelationshipDeliveryPlantID"`

	DeliverToPlantName   *string `json:"DeliverToPlantName,omitempty"`
	DeliverFromPlantName *string `json:"DeliverFromPlantName,omitempty"`
	DeliverToPlant       string  `json:"DeliverToPlant,omitempty"`
	DeliverFromPlant     string  `json:"DeliverFromPlant,omitempty"`
	DeliverToParty       int     `json:"DeliverToParty,omitempty"`
	DeliverToPartyName   *string `json:"DeliverToPartyName,omitempty"`
	DeliverFromParty     int     `json:"DeliverFromParty,omitempty"`
	DeliverFromPartyName *string `json:"DeliverFromPartyName,omitempty"`

	DefaultRelation bool `json:"DefaultRelation"`
}
type Deliver struct {
	DeliverFromParty []DeliveryPlant `json:"DeliverFromParty"`
	DeliverToParty   []DeliveryPlant `json:"DeliverToParty"`
}
