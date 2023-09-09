package apiOutputFormatter

type ProductStock struct {
	ProductStockHeader                    []ProductStockHeader                    `json:"ProductStockHeader"`
	ProductStockDetailHeader              []ProductStockDetailHeader              `json:"ProductStockDetailHeader"`
	ProductStockSingleUnit                []ProductStockSingleUnit                `json:"ProductStockSingleUnit"`
	ProductStockAvailabilityHeader        []ProductStockAvailabilityHeader        `json:"ProductStockAvailabilityHeader"`
	ProductStockAvailabilityDetailHeader  []ProductStockAvailabilityDetailHeader  `json:"ProductStockAvailabilityDetailHeader"`
	ProductStockByStorageBinByBatchHeader []ProductStockByStorageBinByBatchHeader `json:"ProductStockByStorageBinByBatchHeader"`
}

type ProductStockHeader struct {
	SupplyChainRelationshipID int     `json:"SupplyChainRelationshipID"`
	Buyer                     int     `json:"Buyer"`
	BuyerName                 string  `json:"BuyerName"`
	Seller                    int     `json:"Seller"`
	SellerName                string  `json:"SellerName"`
	BusinessPartner           int     `json:"BusinessPartner"`
	BusinessPartnerName       string  `json:"BusinessPartnerName"`
	Plant                     string  `json:"Plant"`
	PlantName                 string  `json:"PlantName"`
	Product                   string  `json:"Product"`
	ProductStock              float32 `json:"ProductStock"`
	Images                    Images  `json:"Images"`
}

type ProductStockDetailHeader struct {
	BusinessPartner      int     `json:"BusinessPartner"`
	BusinessPartnerName  string  `json:"BusinessPartnerName"`
	Plant                string  `json:"Plant"`
	PlantName            string  `json:"PlantName"`
	Product              string  `json:"Product"`
	ProductDescription   string  `json:"ProductDescription"`
	ProductStock         float32 `json:"ProductStock"`
	DeliverToPlant       string  `json:"DeliverToPlant"`
	DeliverToPlantName   string  `json:"DeliverToPlantName"`
	DeliverFromPlant     string  `json:"DeliverFromPlant"`
	DeliverFromPlantName string  `json:"DeliverFromPlantName"`
	InventoryStockType   string  `json:"InventoryStockType"`
	Images               Images  `json:"Images"`
}

type ProductStockSingleUnit struct {
	BusinessPartner     int     `json:"BusinessPartner"`
	BusinessPartnerName string  `json:"BusinessPartnerName"`
	Plant               string  `json:"Plant"`
	PlantName           string  `json:"PlantName"`
	Product             string  `json:"Product"`
	ProductStock        float32 `json:"ProductStock"`
	Images              Images  `json:"Images"`
}

type ProductStockAvailabilityHeader struct {
	SupplyChainRelationshipID int    `json:"SupplyChainRelationshipID"`
	Buyer                     int    `json:"Buyer"`
	BuyerName                 string `json:"BuyerName"`
	Seller                    int    `json:"Seller"`
	SellerName                string `json:"SellerName"`
}

type ProductStockAvailabilityDetailHeader struct {
	BusinessPartner              int     `json:"BusinessPartner"`
	BusinessPartnerName          string  `json:"BusinessPartnerName"`
	Plant                        string  `json:"Plant"`
	PlantName                    string  `json:"PlantName"`
	Product                      string  `json:"Product"`
	ProductDescription           string  `json:"ProductDescription"`
	AvailableProductStock        float32 `json:"AvailableProductStock"`
	DeliverToPlant               string  `json:"DeliverToPlant"`
	DeliverToPlantName           string  `json:"DeliverToPlantName"`
	DeliverFromPlant             string  `json:"DeliverFromPlant"`
	DeliverFromPlantName         string  `json:"DeliverFromPlantName"`
	ProductStockAvailabilityDate string  `json:"ProductStockAvailabilityDate"`
	Images                       Images  `json:"Images"`
}

type ProductStockByStorageBinByBatchHeader struct {
	BusinessPartner      int     `json:"BusinessPartner"`
	BusinessPartnerName  string  `json:"BusinessPartnerName"`
	Plant                string  `json:"Plant"`
	PlantName            string  `json:"PlantName"`
	StorageLocation      string  `json:"StorageLocation"`
	StorageBin           string  `json:"StorageBin"`
	Batch                string  `json:"Batch"`
	Product              string  `json:"Product"`
	ProductDescription   string  `json:"ProductDescription"`
	ProductStock         float32 `json:"ProductStock"`
	DeliverToPlant       string  `json:"DeliverToPlant"`
	DeliverToPlantName   string  `json:"DeliverToPlantName"`
	DeliverFromPlant     string  `json:"DeliverFromPlant"`
	DeliverFromPlantName string  `json:"DeliverFromPlantName"`
	ValidityStartDate    string  `json:"ValidityStartDate"`
	ValidityStartTime    string  `json:"ValidityStartTime"`
	ValidityEndDate      string  `json:"ValidityEndDate"`
	ValidityEndTime      string  `json:"ValidityEndTime"`
	InventoryStockType   string  `json:"InventoryStockType"`
	Images               Images  `json:"Images"`
}
