package apiOutputFormatter

type ProductionOrder struct {
	ProductionOrderHeader         []ProductionOrderHeader         `json:"Header"`
	ProductionOrderHeaderWithItem []ProductionOrderHeaderWithItem `json:"HeaderWithItem"`
	// todo Header という名前は競合するため使用不可
	ProductionOrderHeaderSingleUnit []ProductionOrderHeaderSingleUnit `json:"HeaderSingleUnit"`
	ProductionOrderItemSingleUnit   []ProductionOrderItemSingleUnit   `json:"ItemSingleUnit"`
	ProductionOrderItem             []ProductionOrderItem             `json:"Item"`
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
	MRPArea                                            *string `json:"MRPArea"`
	Product                                            string  `json:"Product"`
	ProductDescription                                 string  `json:"ProductDescription"`
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
	Images                                             Images  `json:"Images"`
}

type ProductionOrderHeaderSingleUnit struct {
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

type ProductionOrderItem struct {
	ProductionOrderItem                int      `json:"ProductionOrderItem"`
	MRPArea                            *string  `json:"MRPArea"`
	Product                            string   `json:"Product"`
	ProductDescription                 string   `json:"ProductDescription"`
	ProductionPlantBusinessPartner     int      `json:"ProductionPlantBusinessPartner"`
	ProductionPlantBusinessPartnerName string   `json:"ProductionPlantBusinessPartnerName"`
	ProductionPlant                    string   `json:"ProductionPlant"`
	ProductionPlantName                string   `json:"ProductionPlantName"`
	ProductionOrderQuantityInBaseUnit  float32  `json:"ProductionOrderQuantityInBaseUnit"`
	ConfirmedYieldQuantityInBaseUnit   *float32 `json:"ConfirmedYieldQuantityInBaseUnit"`
	IsPartiallyConfirmed               *bool    `json:"IsPartiallyConfirmed"`
	IsReleased                         *bool    `json:"IsReleased"`
	IsConfirmed                        *bool    `json:"IsConfirmed"`
	IsCancelled                        *bool    `json:"IsCancelled"`
	IsMarkedForDeletion                *bool    `json:"IsMarkedForDeletion"`
	Images                             Images   `json:"Images"`
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

type ProductionOrderItemSingleUnit struct {
	SizeOrDimensionText                         *string  `json:"SizeOrDimensionText"`
	SafetyStockQuantityInBaseUnit               *float32 `json:"SafetyStockQuantityInBaseUnit"`
	InternalCapacityQuantity                    *float32 `json:"InternalCapacityQuantity"`
	ReorderThresholdQuantityInBaseUnit          *float32 `json:"ReorderThresholdQuantityInBaseUnit"`
	StandardProductionLotSizeQuantityInBaseUnit *float32 `json:"StandardProductionLotSizeQuantityInBaseUnit"`
	Images                                      Images   `json:"Images"`
}
