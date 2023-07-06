package apiModuleRuntimesResponses

type BillOfMaterialRes struct {
	Message BillOfMaterial `json:"message,omitempty"`
}

type BillOfMaterial struct {
	Header *[]BillOfMaterialHeader `json:"Header,omitempty"`
	Item   *[]BillOfMaterialItem   `json:"Item,omitempty"`
}

type BillOfMaterialHeader struct {
	BillOfMaterial                           int     `json:"BillOfMaterial"`
	BillOfMaterialType                       string  `json:"BillOfMaterialType"`
	SupplyChainRelationshipID                int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID        int     `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID   int     `json:"SupplyChainRelationshipDeliveryPlantID"`
	SupplyChainRelationshipProductionPlantID int     `json:"SupplyChainRelationshipProductionPlantID"`
	Product                                  string  `json:"Product"`
	Buyer                                    int     `json:"Buyer"`
	Seller                                   int     `json:"Seller"`
	DepartureDeliverFromParty                int     `json:"DepartureDeliverFromParty"`
	DepartureDeliverFromPlant                string  `json:"DepartureDeliverFromPlant"`
	DestinationDeliverToParty                int     `json:"DestinationDeliverToParty"`
	DestinationDeliverToPlant                string  `json:"DestinationDeliverToPlant"`
	OwnerProductionPlantBusinessPartner      int     `json:"OwnerProductionPlantBusinessPartner"`
	OwnerProductionPlant                     string  `json:"OwnerProductionPlant"`
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

type BillOfMaterialItem struct {
	BillOfMaterial                                  int      `json:"BillOfMaterial"`
	BillOfMaterialItem                              int      `json:"BillOfMaterialItem"`
	SupplyChainRelationshipID                       int      `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID               int      `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID          int      `json:"SupplyChainRelationshipDeliveryPlantID"`
	SupplyChainRelationshipStockConfPlantID         int      `json:"SupplyChainRelationshipStockConfPlantID"`
	Product                                         *string  `json:"Product"`
	ProductionPlantBusinessPartner                  *int     `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                                 *string  `json:"ProductionPlant"`
	ComponentProduct                                *string  `json:"ComponentProduct"`
	ComponentProductBuyer                           *int     `json:"ComponentProductBuyer"`
	ComponentProductSeller                          *int     `json:"ComponentProductSeller"`
	ComponentProductDeliverFromParty                *int     `json:"ComponentProductDeliverFromParty"`
	ComponentProductDeliverFromPlant                *string  `json:"ComponentProductDeliverFromPlant"`
	ComponentProductDeliverToParty                  *int     `json:"ComponentProductDeliverToParty"`
	ComponentProductDeliverToPlant                  *string  `json:"ComponentProductDeliverToPlant"`
	StockConfirmationBusinessPartner                *int     `json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant                          *string  `json:"StockConfirmationPlant"`
	ComponentProductStandardQuantityInBaseUnuit     *float32 `json:"ComponentProductStandardQuantityInBaseUnuit"`
	ComponentProductStandardQuantityInDeliveryUnuit *float32 `json:"ComponentProductStandardQuantityInDeliveryUnuit"`
	ComponentProductBaseUnit                        *string  `json:"ComponentProductBaseUnit"`
	ComponentProductDeliveryUnit                    *string  `json:"ComponentProductDeliveryUnit"`
	ComponentProductStandardScrapInPercent          *float32 `json:"ComponentProductStandardScrapInPercent"`
	IsMarkedForBackflush                            *bool    `json:"IsMarkedForBackflush"`
	BillOfMaterialItemText                          *string  `json:"BillOfMaterialItemText"`
	ValidityStartDate                               *string  `json:"ValidityStartDate"`
	ValidityEndDate                                 *string  `json:"ValidityEndDate"`
	CreationDate                                    *string  `json:"CreationDate"`
	LastChangeDate                                  *string  `json:"LastChangeDate"`
	IsMarkedForDeletion                             *bool    `json:"IsMarkedForDeletion"`
}
