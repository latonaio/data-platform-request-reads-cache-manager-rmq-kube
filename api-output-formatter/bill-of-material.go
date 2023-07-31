package apiOutputFormatter

type BillOfMaterial struct {
	BillOfMaterialHeader         []BillOfMaterialHeader         `json:"Header"`
	BillOfMaterialHeaderWithItem []BillOfMaterialHeaderWithItem `json:"HeaderWithItem"`
	BillOfMaterialItem           []BillOfMaterialItem           `json:"Item"`
}

type BillOfMaterialHeader struct {
	Product                  string  `json:"Product"`
	BillOfMaterial           int     `json:"BillOfMaterial"`
	ProductDescription       *string `json:"ProductDescription"`
	OwnerProductionPlant     string  `json:"OwnerProductionPlant"`
	OwnerProductionPlantName *string `json:"OwnerProductionPlantName"`
	ValidityStartDate        *string `json:"ValidityStartDate"`
	IsMarkedForDeletion      *bool   `json:"IsMarkedForDeletion"`
	Images                   Images  `json:"Images"`
}

type BillOfMaterialHeaderWithItem struct {
	BillOfMaterial                           int     `json:"BillOfMaterial"`
	BillOfMaterialType                       string  `json:"BillOfMaterialType"`
	SupplyChainRelationshipID                int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID        int     `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID   int     `json:"SupplyChainRelationshipDeliveryPlantID"`
	SupplyChainRelationshipProductionPlantID int     `json:"SupplyChainRelationshipProductionPlantID"`
	Buyer                                    int     `json:"Buyer"`
	Seller                                   int     `json:"Seller"`
	DepartureDeliverFromParty                int     `json:"DepartureDeliverFromParty"`
	DepartureDeliverFromPlant                string  `json:"DepartureDeliverFromPlant"`
	DestinationDeliverToParty                int     `json:"DestinationDeliverToParty"`
	DestinationDeliverToPlant                string  `json:"DestinationDeliverToPlant"`
	OwnerProductionPlantBusinessPartner      int     `json:"OwnerProductionPlantBusinessPartner"`
	OwnerProductionPlant                     string  `json:"OwnerProductionPlant"`
	OwnerProductionPlantName                 string  `json:"OwnerProductionPlantName"`
	Product                                  string  `json:"Product"`
	ProductBaseUnit                          string  `json:"ProductBaseUnit"`
	ProductDeliveryUnit                      string  `json:"ProductDeliveryUnit"`
	ProductProductionUnit                    string  `json:"ProductProductionUnit"`
	ProductStandardQuantityInBaseUnit        float32 `json:"ProductStandardQuantityInBaseUnit"`
	ProductStandardQuantityInDeliveryUnit    float32 `json:"ProductStandardQuantityInDeliveryUnit"`
	ProductStandardQuantityInProductionUnit  float32 `json:"ProductStandardQuantityInProductionUnit"`
	ProductDescription                       string  `json:"ProductDescription"`
	BillOfMaterialHeaderText                 *string `json:"BillOfMaterialHeaderText"`
	ValidityStartDate                        *string `json:"ValidityStartDate"`
	ValidityEndDate                          *string `json:"ValidityEndDate"`
	CreationDate                             string  `json:"CreationDate"`
	LastChangeDate                           string  `json:"LastChangeDate"`
	Images                                   Images  `json:"Images"`
}

type BillOfMaterialHeaderWithItemOtherInfo struct {
	Product                  string  `json:"Product"`
	BillOfMaterial           int     `json:"BillOfMaterial"`
	ProductDescription       *string `json:"ProductDescription"`
	OwnerProductionPlant     string  `json:"OwnerProductionPlant"`
	OwnerProductionPlantName *string `json:"OwnerProductionPlantName"`
	ValidityStartDate        *string `json:"ValidityStartDate"`
	Images                   Images  `json:"Images"`
}

type BillOfMaterialItem struct {
	ComponentProduct                           string   `json:"ComponentProduct"`
	BillOfMaterialItem                         int      `json:"BillOfMaterialItem"`
	BillOfMaterialItemText                     string   `json:"BillOfMaterialItemText"`
	StockConfirmationPlant                     *string  `json:"StockConfirmationPlant"`
	StockConfirmationPlantName                 *string  `json:"StockConfirmationPlantName"`
	ComponentProductStandardQuantityInBaseUnit *float32 `json:"ComponentProductStandardQuantityInBaseUnit"`
	ComponentProductBaseUnit                   *string  `json:"ComponentProductBaseUnit"`
	ValidityStartDate                          *string  `json:"ValidityStartDate"`
	IsMarkedForDeletion                        *bool    `json:"IsMarkedForDeletion"`
}
