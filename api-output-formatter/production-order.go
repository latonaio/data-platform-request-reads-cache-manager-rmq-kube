package apiOutputFormatter

type ProductionOrder struct {
	ProductionOrderHeader           []ProductionOrderHeader           `json:"Header"`
	ProductionOrderHeaderWithItem   []ProductionOrderHeaderWithItem   `json:"HeaderWithItem"`
	ProductionOrderHeaderSingleUnit []ProductionOrderHeaderSingleUnit `json:"HeaderSingleUnit"`
	ProductionOrderItem             []ProductionOrderItem             `json:"Item"`
	ProductionOrderItemSingleUnit   []ProductionOrderItemSingleUnit   `json:"ItemSingleUnit"`
	ProductionOrderItemComponent    []ProductionOrderItemComponent    `json:"ItemComponent"`
	ProductionOrderItemOperation    []ProductionOrderItemOperation    `json:"ItemOperation"`
}

type ProductionOrderHeader struct {
	ProductionOrder                         int     `json:"ProductionOrder"`
	MRPArea                                 *string `json:"MRPArea"`
	Product                                 string  `json:"Product"`
	ProductDescription                      string  `json:"ProductDescription"`
	OwnerProductionPlantBusinessPartner     int     `json:"OwnerProductionPlantBusinessPartner"`
	OwnerProductionPlantBusinessPartnerName string  `json:"OwnerProductionPlantBusinessPartnerName"`
	OwnerProductionPlant                    string  `json:"OwnerProductionPlant"`
	OwnerProductionPlantName                string  `json:"OwnerProductionPlantName"`
	ProductionOrderQuantityInBaseUnit       float32 `json:"ProductionOrderQuantityInBaseUnit"`
	IsReleased                              *bool   `json:"IsReleased"`
	IsPartiallyConfirmed                    *bool   `json:"IsPartiallyConfirmed"`
	IsConfirmed                             *bool   `json:"IsConfirmed"`
	IsCancelled                             *bool   `json:"IsCancelled"`
	IsMarkedForDeletion                     *bool   `json:"IsMarkedForDeletion"`
	Images                                  Images  `json:"Images"`
}

type ProductionOrderHeaderWithItem struct {
	ProductionOrder                                    int     `json:"ProductionOrder"`
	ProductionOrderDate                                string  `json:"ProductionOrderDate"`
	Product                                            string  `json:"Product"`
	ProductDescription                                 string  `json:"ProductDescription"`
	Buyer                                              int     `json:"Buyer"`
	BuyerName                                          string  `json:"BuyerName"`
	Seller                                             int     `json:"Seller"`
	SellerName                                         string  `json:"SellerName"`
	MRPArea                                            *string `json:"MRPArea"`
	OwnerProductionPlantBusinessPartner                int     `json:"OwnerProductionPlantBusinessPartner"`
	OwnerProductionPlantBusinessPartnerName            string  `json:"OwnerProductionPlantBusinessPartnerName"`
	OwnerProductionPlant                               string  `json:"OwnerProductionPlant"`
	OwnerProductionPlantName                           string  `json:"OwnerProductionPlantName"`
	ProductionOrderQuantityInBaseUnit                  float32 `json:"ProductionOrderQuantityInBaseUnit"`
	ProductionOrderQuantityInDestinationProductionUnit float32 `json:"ProductionOrderQuantityInDestinationProductionUnit"`
	IsReleased                                         *bool   `json:"IsReleased"`
	IsPartiallyConfirmed                               *bool   `json:"IsPartiallyConfirmed"`
	IsConfirmed                                        *bool   `json:"IsConfirmed"`
	IsCancelled                                        *bool   `json:"IsCancelled"`
	IsMarkedForDeletion                                *bool   `json:"IsMarkedForDeletion"`
	ProductionOrderPlannedStartDate                    string  `json:"ProductionOrderPlannedStartDate"`
	ProductionOrderPlannedStartTime                    string  `json:"ProductionOrderPlannedStartTime"`
	ProductionOrderPlannedEndDate                      string  `json:"ProductionOrderPlannedEndDate"`
	ProductionOrderPlannedEndTime                      string  `json:"ProductionOrderPlannedEndTime"`
	InspectionLot                                      *int    `json:"InspectionLot"`
	Images                                             Images  `json:"Images"`
}

type ProductionOrderHeaderSingleUnit struct {
	ProductionOrder                                    int     `json:"ProductionOrder"`
	ProductionOrderDate                                string  `json:"ProductionOrderDate"`
	MRPArea                                            *string `json:"MRPArea"`
	Product                                            string  `json:"Product"`
	Buyer                                              int     `json:"Buyer"`
	BuyerName                                          string  `json:"BuyerName"`
	Seller                                             int     `json:"Seller"`
	SellerName                                         string  `json:"SellerName"`
	ProductDescription                                 string  `json:"ProductDescription"`
	OwnerProductionPlantBusinessPartner                int     `json:"OwnerProductionPlantBusinessPartner"`
	OwnerProductionPlantBusinessPartnerName            string  `json:"OwnerProductionPlantBusinessPartnerName"`
	OwnerProductionPlant                               string  `json:"OwnerProductionPlant"`
	OwnerProductionPlantName                           string  `json:"OwnerProductionPlantName"`
	ProductionOrderQuantityInBaseUnit                  float32 `json:"ProductionOrderQuantityInBaseUnit"`
	ProductionOrderQuantityInDestinationProductionUnit float32 `json:"ProductionOrderQuantityInDestinationProductionUnit"`
	InspectionLot                                      *int    `json:"InspectionLot"`
	IsReleased                                         *bool   `json:"IsReleased"`
	IsPartiallyConfirmed                               *bool   `json:"IsPartiallyConfirmed"`
	IsConfirmed                                        *bool   `json:"IsConfirmed"`
	IsCancelled                                        *bool   `json:"IsCancelled"`
	IsMarkedForDeletion                                *bool   `json:"IsMarkedForDeletion"`
	Images                                             Images  `json:"Images"`
}

type ProductionOrderItem struct {
	ProductionOrder                         int      `json:"ProductionOrder"`
	ProductionOrderItem                     int      `json:"ProductionOrderItem"`
	Product                                 string   `json:"Product"`
	ProductDescription                      string   `json:"ProductDescription"`
	Buyer                                   int      `json:"Buyer"`
	BuyerName                               string   `json:"BuyerName"`
	Seller                                  int      `json:"Seller"`
	SellerName                              string   `json:"SellerName"`
	MRPArea                                 *string  `json:"MRPArea"`
	ProductBaseUnit                         string   `json:"ProductBaseUnit"`
	ProductProductionUnit                   string   `json:"ProductProductionUnit"`
	ProductionPlantBusinessPartner          int      `json:"ProductionPlantBusinessPartner"`
	ProductionPlantBusinessPartnerName      string   `json:"ProductionPlantBusinessPartnerName"`
	ProductionPlant                         string   `json:"ProductionPlant"`
	ProductionPlantName                     string   `json:"ProductionPlantName"`
	ProductionOrderQuantityInBaseUnit       float32  `json:"ProductionOrderQuantityInBaseUnit"`
	ProductionOrderQuantityInProductionUnit float32  `json:"ProductionOrderQuantityInProductionUnit"`
	ProductionOrderPlannedStartDate         string   `json:"ProductionOrderPlannedStartDate"`
	ProductionOrderPlannedStartTime         string   `json:"ProductionOrderPlannedStartTime"`
	ProductionOrderPlannedEndDate           string   `json:"ProductionOrderPlannedEndDate"`
	ProductionOrderPlannedEndTime           string   `json:"ProductionOrderPlannedEndTime"`
	ConfirmedYieldQuantityInBaseUnit        *float32 `json:"ConfirmedYieldQuantityInBaseUnit"`
	InspectionLot                           *int     `json:"InspectionLot"`
	IsPartiallyConfirmed                    *bool    `json:"IsPartiallyConfirmed"`
	IsReleased                              *bool    `json:"IsReleased"`
	IsConfirmed                             *bool    `json:"IsConfirmed"`
	IsCancelled                             *bool    `json:"IsCancelled"`
	IsMarkedForDeletion                     *bool    `json:"IsMarkedForDeletion"`
	Images                                  Images   `json:"Images"`
}

type ProductionOrderItemSingleUnit struct {
	ProductionOrder                             int      `json:"ProductionOrder"`
	ProductionOrderItem                         int      `json:"ProductionOrderItem"`
	ProductionOrderItemDate                     string   `json:"ProductionOrderItemDate"`
	Product                                     string   `json:"Product"`
	Buyer                                       int      `json:"Buyer"`
	BuyerName                                   string   `json:"BuyerName"`
	Seller                                      int      `json:"Seller"`
	SellerName                                  string   `json:"SellerName"`
	ProductionPlantBusinessPartner              int      `json:"ProductionPlantBusinessPartner"`
	ProductionPlantBusinessPartnerName          string   `json:"ProductionPlantBusinessPartnerName"`
	ProductionPlant                             string   `json:"ProductionPlant"`
	ProductionPlantName                         string   `json:"ProductionPlantName"`
	ProductionOrderQuantityInBaseUnit           float32  `json:"ProductionOrderQuantityInBaseUnit"`
	ProductionOrderQuantityInProductionUnit     float32  `json:"ProductionOrderQuantityInProductionUnit"`
	InspectionLot                               *int     `json:"InspectionLot"`
	SizeOrDimensionText                         *string  `json:"SizeOrDimensionText"`
	SafetyStockQuantityInBaseUnit               *float32 `json:"SafetyStockQuantityInBaseUnit"`
	InternalCapacityQuantity                    *float32 `json:"InternalCapacityQuantity"`
	ReorderThresholdQuantityInBaseUnit          *float32 `json:"ReorderThresholdQuantityInBaseUnit"`
	StandardProductionLotSizeQuantityInBaseUnit *float32 `json:"StandardProductionLotSizeQuantityInBaseUnit"`
	Images                                      Images   `json:"Images"`
}

type ProductionOrderItemComponent struct {
	ProductionOrder                                int     `json:"ProductionOrder"`
	ProductionOrderItem                            int     `json:"ProductionOrderItem"`
	BillOfMaterial                                 int     `json:"BillOfMaterial"`
	BillOfMaterialItem                             int     `json:"BillOfMaterialItem"`
	ComponentProduct                               string  `json:"ComponentProduct"`
	ComponentProductBuyer                          int     `json:"ComponentProductBuyer"`
	ComponentProductBuyerName                      string  `json:"ComponentProductBuyerName"`
	ComponentProductSeller                         int     `json:"ComponentProductSeller"`
	ComponentProductSellerName                     string  `json:"ComponentProductSellerName"`
	ComponentProductBaseUnit                       string  `json:"ComponentProductBaseUnit"`
	ComponentProductDeliveryUnit                   string  `json:"ComponentProductDeliveryUnit"`
	ComponentProductRequiredQuantityInBaseUnit     float32 `json:"ComponentProductRequiredQuantityInBaseUnit"`
	ComponentProductRequiredQuantityInDeliveryUnit float32 `json:"ComponentProductRequiredQuantityInDeliveryUnit"`
	Images                                         Images  `json:"Images"`
}

type ProductionOrderItemOperation struct {
	ProductionOrder                          int    `json:"ProductionOrder"`
	ProductionOrderItem                      int    `json:"ProductionOrderItem"`
	Operations                               int    `json:"Operations"`
	OperationsItem                           int    `json:"OperationsItem"`
	OperationID                              int    `json:"OperationID"`
	SupplyChainRelationshipID                int    `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID        int    `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID   int    `json:"SupplyChainRelationshipDeliveryPlantID"`
	SupplyChainRelationshipProductionPlantID int    `json:"SupplyChainRelationshipProductionPlantID"`
	Product                                  string `json:"Product"`
	Buyer                                    int    `json:"Buyer"`
	Seller                                   int    `json:"Seller"`
	SellerName                               string `json:"SellerName"`
	DeliverToParty                           int    `json:"DeliverToParty"`
	DeliverToPlant                           string `json:"DeliverToPlant"`
	DeliverFromParty                         int    `json:"DeliverFromParty"`
	DeliverFromPlant                         string `json:"DeliverFromPlant"`
	ProductionPlantBusinessPartner           int    `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                          string `json:"ProductionPlant"`
	OperationText                            string `json:"OperationText"`
	WorkCenter                               int    `json:"WorkCenter"`
}
