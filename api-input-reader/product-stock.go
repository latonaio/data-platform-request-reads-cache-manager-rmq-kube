package apiInputReader

type ProductStock struct {
	ProductStockHeader *ProductStockHeader
}

type ProductStockHeader struct {
	Product                                string  `json:"Product"`
	BusinessPartner                        int     `json:"BusinessPartner"`
	Plant                                  string  `json:"Plant"`
	SupplyChainRelationshipID              int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID      int     `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID int     `json:"SupplyChainRelationshipDeliveryPlantID"`
	Buyer                                  *int    `json:"Buyer"`
	Seller                                 *int    `json:"Seller"`
	DeliverToParty                         int     `json:"DeliverToParty"`
	DeliverFromParty                       int     `json:"DeliverFromParty"`
	DeliverToPlant                         string  `json:"DeliverToPlant"`
	DeliverFromPlant                       string  `json:"DeliverFromPlant"`
	InventoryStockType                     string  `json:"InventoryStockType"`
	ProductStock                           float32 `json:"ProductStock"`
}
