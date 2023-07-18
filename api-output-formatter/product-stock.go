package apiOutputFormatter

type ProductStock struct {
	ProductStockHeader       []ProductStockHeader       `json:"Header"`
	ProductStockDetailHeader []ProductStockDetailHeader `json:"DetailHeader"`
}

type ProductStockHeader struct {
	SupplyChainRelationshipID int    `json:"SupplyChainRelationshipID"`
	Buyer                     int    `json:"Buyer"`
	BuyerName                 string `json:"BuyerName"`
	Seller                    int    `json:"Seller"`
	SellerName                string `json:"SellerName"`
}

type ProductStockDetailHeader struct {
	BusinessPartner           int     `json:"BusinessPartner"`
	BusinessPartnerName       string  `json:"BusinessPartnerName"`
	Plant                     string  `json:"Plant"`
	PlantName                 string  `json:"PlantName"`
	Product                   string  `json:"Product"`
	ProductDescription        string  `json:"ProductDescription"`
    ProductStock              float32 `json:"ProductStock"`
	DeliverToPlant            string  `json:"DeliverToPlant"`
	DeliverToPlantName        string  `json:"DeliverToPlantName"`
	DeliverFromPlant          string  `json:"DeliverFromPlant"`
	DeliverFromPlantName      string  `json:"DeliverFromPlantName"`
	InventoryStockType        string  `json:"InventoryStockType"`
	Images                    Images  `json:"Images"`
}
