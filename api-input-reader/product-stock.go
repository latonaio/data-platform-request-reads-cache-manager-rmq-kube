package apiInputReader

type ProductStock struct {
	ProductStockHeader                    *ProductStockHeader
	ProductStockAvailabilityHeader        *ProductStockAvailabilityHeader
	ProductStockByStorageBinByBatchHeader *ProductStockByStorageBinByBatchHeader
	ProductStockDocProductStockDoc        *ProductStockDocProductStockDoc
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
	CreationDate                           string  `json:"CreationDate"`
	CreationTime                           string  `json:"CreationTime"`
	LastChangeDate                         string  `json:"LastChangeDate"`
	LastChangeTime                         string  `json:"LastChangeTime"`
}

type ProductStockAvailabilityHeader struct {
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
	ProductStockAvailabilityDate           string  `json:"ProductStockAvailabilityDate"`
	AvailableProductStock                  float32 `json:"AvailableProductStock"`
	CreationDate                           string  `json:"CreationDate"`
	CreationTime                           string  `json:"CreationTime"`
	LastChangeDate                         string  `json:"LastChangeDate"`
	LastChangeTime                         string  `json:"LastChangeTime"`
}

type ProductStockByStorageBinByBatchHeader struct {
	Product                                string  `json:"Product"`
	BusinessPartner                        int     `json:"BusinessPartner"`
	Plant                                  string  `json:"Plant"`
	StorageLocation                        string  `json:"StorageLocation"`
	StorageBin                             string  `json:"StorageBin"`
	Batch                                  string  `json:"Batch"`
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
	CreationDate                           string  `json:"CreationDate"`
	CreationTime                           string  `json:"CreationTime"`
	LastChangeDate                         string  `json:"LastChangeDate"`
	LastChangeTime                         string  `json:"LastChangeTime"`
}

type ProductStockDocProductStockDoc struct {
	Product                  *string `json:"Product"`
	BusinessPartner          *int    `json:"BusinessPartner"`
	Plant                    *string `json:"Plant"`
	DocType                  string  `json:"DocType"`
	DocIssuerBusinessPartner *int    `json:"DocIssuerBusinessPartner"`
}
