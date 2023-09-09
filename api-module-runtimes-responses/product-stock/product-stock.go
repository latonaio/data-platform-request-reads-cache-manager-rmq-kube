package apiModuleRuntimesResponsesProductStock

type ProductStockRes struct {
	Message ProductStockResType `json:"message,omitempty"`
}

type ProductStockResType struct {
	ProductStock                                *[]ProductStock                                `json:"ProductStock,omitempty"`
	ProductStockByStorageBinByBatch             *[]ProductStockByStorageBinByBatch             `json:"ProductStockByStorageBinByBatch,omitempty"`
	ProductStockAvailability                    *[]ProductStockAvailability                    `json:"ProductStockAvailability,omitempty"`
	ProductStockAvailabilityByStorageBinByBatch *[]ProductStockAvailabilityByStorageBinByBatch `json:"ProductStockAvailabilityByStorageBinByBatch,omitempty"`
}

type ProductStock struct {
	Product                                string  `json:"Product"`
	BusinessPartner                        int     `json:"BusinessPartner"`
	Plant                                  string  `json:"Plant"`
	SupplyChainRelationshipID              int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID      int     `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID int     `json:"SupplyChainRelationshipDeliveryPlantID"`
	Buyer                                  int     `json:"Buyer"`
	Seller                                 int     `json:"Seller"`
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

type ProductStockAvailability struct {
	Product                                string  `json:"Product"`
	BusinessPartner                        int     `json:"BusinessPartner"`
	Plant                                  string  `json:"Plant"`
	SupplyChainRelationshipID              int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID      int     `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID int     `json:"SupplyChainRelationshipDeliveryPlantID"`
	Buyer                                  int     `json:"Buyer"`
	Seller                                 int     `json:"Seller"`
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

type ProductStockByStorageBinByBatch struct {
	Product                                string  `json:"Product"`
	BusinessPartner                        int     `json:"BusinessPartner"`
	Plant                                  string  `json:"Plant"`
	StorageLocation                        string  `json:"StorageLocation"`
	StorageBin                             string  `json:"StorageBin"`
	Batch                                  string  `json:"Batch"`
	SupplyChainRelationshipID              int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID      int     `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID int     `json:"SupplyChainRelationshipDeliveryPlantID"`
	Buyer                                  int     `json:"Buyer"`
	Seller                                 int     `json:"Seller"`
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

type ProductStockAvailabilityByStorageBinByBatch struct {
	Product                                string  `json:"Product"`
	BusinessPartner                        int     `json:"BusinessPartner"`
	Plant                                  string  `json:"Plant"`
	StorageLocation                        string  `json:"StorageLocation"`
	StorageBin                             string  `json:"StorageBin"`
	Batch                                  string  `json:"Batch"`
	SupplyChainRelationshipID              int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID      int     `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID int     `json:"SupplyChainRelationshipDeliveryPlantID"`
	Buyer                                  int     `json:"Buyer"`
	Seller                                 int     `json:"Seller"`
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
