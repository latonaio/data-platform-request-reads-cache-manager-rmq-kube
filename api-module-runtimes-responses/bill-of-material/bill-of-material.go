package apiModuleRuntimesResponsesBillOfMaterial

type BillOfMaterialRes struct {
	Message BillOfMaterial `json:"message,omitempty"`
}

type BillOfMaterial struct {
	Header *[]Header `json:"Header,omitempty"`
	Item   *[]Item   `json:"Item,omitempty"`
}

type Header struct {
	BillOfMaterial                           int     `json:"BillOfMaterial"`
	BillOfMaterialType                       string  `json:"BillOfMaterialType"`
	SupplyChainRelationshipID                int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID        int     `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID   int     `json:"SupplyChainRelationshipDeliveryPlantID"`
	SupplyChainRelationshipProductionPlantID int     `json:"SupplyChainRelationshipProductionPlantID"`
	Product                                  string  `json:"Product"`
	Buyer                                    int     `json:"Buyer"`
	BuyerName                                string  `json:"BuyerName"`
	Seller                                   int     `json:"Seller"`
	SellerName                               string  `json:"SellerName"`
	DepartureDeliverFromParty                int     `json:"DepartureDeliverFromParty"`
	DepartureDeliverFromPartyName            string  `json:"DepartureDeliverFromPartyName"`
	DepartureDeliverFromPlant                string  `json:"DepartureDeliverFromPlant"`
	DepartureDeliverFromPlantName            string  `json:"DepartureDeliverFromPlantName"`
	DestinationDeliverToParty                int     `json:"DestinationDeliverToParty"`
	DestinationDeliverToPartyName            string  `json:"DestinationDeliverToPartyName"`
	DestinationDeliverToPlant                string  `json:"DestinationDeliverToPlant"`
	DestinationDeliverToPlantName            string  `json:"DestinationDeliverToPlantName"`
	OwnerProductionPlantBusinessPartner      int     `json:"OwnerProductionPlantBusinessPartner"`
	OwnerProductionPlantBusinessPartnerName  string  `json:"OwnerProductionPlantBusinessPartnerName"`
	OwnerProductionPlant                     string  `json:"OwnerProductionPlant"`
	OwnerProductionPlantName                 string  `json:"OwnerProductionPlantName"`
	ProductBaseUnit                          string  `json:"ProductBaseUnit"`
	ProductDeliveryUnit                      string  `json:"ProductDeliveryUnit"`
	ProductProductionUnit                    string  `json:"ProductProductionUnit"`
	ProductStandardQuantityInBaseUnit        float32 `json:"ProductStandardQuantityInBaseUnit"`
	ProductStandardQuantityInDeliveryUnit    float32 `json:"ProductStandardQuantityInDeliveryUnit"`
	ProductStandardQuantityInProductionUnit  float32 `json:"ProductStandardQuantityInProductionUnit"`
	BillOfMaterialHeaderText                 *string `json:"BillOfMaterialHeaderText"`
	ValidityStartDate                        *string `json:"ValidityStartDate"`
	ValidityEndDate                          *string `json:"ValidityEndDate"`
	CreationDate                             string  `json:"CreationDate"`
	LastChangeDate                           string  `json:"LastChangeDate"`
	IsMarkedForDeletion                      *bool   `json:"IsMarkedForDeletion"`
}

type Item struct {
	BillOfMaterial                                  int      `json:"BillOfMaterial"`
	BillOfMaterialItem                              int      `json:"BillOfMaterialItem"`
	SupplyChainRelationshipID                       int      `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID               int      `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID          int      `json:"SupplyChainRelationshipDeliveryPlantID"`
	SupplyChainRelationshipStockConfPlantID         int      `json:"SupplyChainRelationshipStockConfPlantID"`
	Product                                         string   `json:"Product"`
	ProductionPlantBusinessPartner                  int      `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                                 string   `json:"ProductionPlant"`
	ComponentProduct                                string   `json:"ComponentProduct"`
	ComponentProductBuyer                           int      `json:"ComponentProductBuyer"`
	ComponentProductSeller                          int      `json:"ComponentProductSeller"`
	ComponentProductDeliverToParty                  int      `json:"ComponentProductDeliverToParty"`
	ComponentProductDeliverToPlant                  string   `json:"ComponentProductDeliverToPlant"`
	ComponentProductDeliverFromParty                int      `json:"ComponentProductDeliverFromParty"`
	ComponentProductDeliverFromPlant                string   `json:"ComponentProductDeliverFromPlant"`
	StockConfirmationBusinessPartner                int      `json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant                          string   `json:"StockConfirmationPlant"`
	ComponentProductStandardQuantityInBaseUnit      float32  `json:"ComponentProductStandardQuantityInBaseUnit"`
	ComponentProductStandardQuantityInDeliveryUnit  float32  `json:"ComponentProductStandardQuantityInDeliveryUnit"`
	ComponentProductBaseUnit                        string   `json:"ComponentProductBaseUnit"`
	ComponentProductDeliveryUnit                    string   `json:"ComponentProductDeliveryUnit"`
	ComponentProductStandardScrapInPercent          *float32 `json:"ComponentProductStandardScrapInPercent"`
	IsMarkedForBackflush                            *bool    `json:"IsMarkedForBackflush"`
	BillOfMaterialItemText                          *string  `json:"BillOfMaterialItemText"`
	ValidityStartDate                               *string  `json:"ValidityStartDate"`
	ValidityEndDate                                 *string  `json:"ValidityEndDate"`
	CreationDate                                    string   `json:"CreationDate"`
	LastChangeDate                                  string   `json:"LastChangeDate"`
	IsMarkedForDeletion                             *bool    `json:"IsMarkedForDeletion"`
}
