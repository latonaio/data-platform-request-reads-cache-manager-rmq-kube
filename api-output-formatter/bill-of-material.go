package apiOutputFormatter

type BillOfMaterial struct {
	BillOfMaterialHeader []BillOfMaterialHeader `json:"Header"`
	BillOfMaterialItem   []BillOfMaterialItem   `json:"Item"`
}

type BillOfMaterialHeader struct {
	BillOfMaterial                          int     `json:"BillOfMaterial"`
	Product                                 string  `json:"Product"`
	ProductDescription                      *string `json:"ProductDescription"`
	Seller                                  int     `json:"Seller"`
	SellerName                              string  `json:"SellerName"`
	Buyer                                   int     `json:"Buyer"`
	BuyerName                               string  `json:"BuyerName"`
	OwnerProductionPlant                    string  `json:"OwnerProductionPlant"`
	OwnerProductionPlantName                string  `json:"OwnerProductionPlantName"`
	OwnerBusinessPartner                    int     `json:"OwnerBusinessPartner"`
	OwnerBusinessPartnerName                string  `json:"OwnerBusinessPartnerName"`
	ProductBaseUnit                         string  `json:"ProductBaseUnit"`
	ProductProductionUnit                   string  `json:"ProductProductionUnit"`
	ProductStandardQuantityInBaseUnit       float32 `json:"ProductStandardQuantityInBaseUnit"`
	ProductStandardQuantityInProductionUnit float32 `json:"ProductStandardQuantityInProductionUnit"`
	ValidityStartDate                       *string `json:"ValidityStartDate"`
	IsMarkedForDeletion                     *bool   `json:"IsMarkedForDeletion"`
	Images                                  Images  `json:"Images"`
}

type BillOfMaterialItem struct {
	BillOfMaterial                                 int     `json:"BillOfMaterial"`
	BillOfMaterialItem                             int     `json:"BillOfMaterialItem"`
	ComponentProduct                               string  `json:"ComponentProduct"`
	ComponentProductBuyer                          int     `json:"ComponentProductBuyer"`
	ComponentProductBuyerName                      string  `json:"ComponentProductBuyerName"`
	ComponentProductSeller                         int     `json:"ComponentProductSeller"`
	ComponentProductSellerName                     string  `json:"ComponentProductSellerName"`
	StockConfirmationPlant                         string  `json:"StockConfirmationPlant"`
	StockConfirmationPlantName                     string  `json:"StockConfirmationPlantName"`
	ComponentProductStandardQuantityInBaseUnit     float32 `json:"ComponentProductStandardQuantityInBaseUnit"`
	ComponentProductStandardQuantityInDeliveryUnit float32 `json:"ComponentProductStandardQuantityInDeliveryUnit"`
	ComponentProductBaseUnit                       string  `json:"ComponentProductBaseUnit"`
	ComponentProductDeliveryUnit                   string  `json:"ComponentProductDeliveryUnit"`
	BillOfMaterialItemText                         *string `json:"BillOfMaterialItemText"`
	ValidityStartDate                              *string `json:"ValidityStartDate"`
	IsMarkedForDeletion                            *bool   `json:"IsMarkedForDeletion"`
	Images                                         Images  `json:"Images"`
}
