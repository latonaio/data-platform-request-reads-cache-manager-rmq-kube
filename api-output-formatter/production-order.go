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
	ProductionOrder                                 int      `json:"ProductionOrder"`
	ProductionOrderItem                             int      `json:"ProductionOrderItem"`
	Operations                                      int      `json:"Operations"`
	OperationsItem                                  int      `json:"OperationsItem"`
	OperationID                                     int      `json:"OperationID"`
	SupplyChainRelationshipID                       int      `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID               int      `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID          int      `json:"SupplyChainRelationshipDeliveryPlantID"`
	SupplyChainRelationshipProductionPlantID        int      `json:"SupplyChainRelationshipProductionPlantID"`
	Product                                         string   `json:"Product"`
	Buyer                                           int      `json:"Buyer"`
	Seller                                          int      `json:"Seller"`
	SellerName                                      string   `json:"SellerName"`
	DeliverToParty                                  int      `json:"DeliverToParty"`
	DeliverToPlant                                  string   `json:"DeliverToPlant"`
	DeliverFromParty                                int      `json:"DeliverFromParty"`
	DeliverFromPlant                                string   `json:"DeliverFromPlant"`
	ProductionPlantBusinessPartner                  int      `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                                 string   `json:"ProductionPlant"`
	MRPArea                                         *string  `json:"MRPArea"`
	MRPController                                   *string  `json:"MRPController"`
	ProductionVersion                               *int     `json:"ProductionVersion"`
	ProductionVersionItem                           *int     `json:"ProductionVersionItem"`
	Sequence                                        int      `json:"Sequence"`
	SequenceText                                    *string  `json:"SequenceText"`
	OperationText                                   string   `json:"OperationText"`
	ProductBaseUnit                                 string   `json:"ProductBaseUnit"`
	ProductProductionUnit                           string   `json:"ProductProductionUnit"`
	ProductOperationUnit                            string   `json:"ProductOperationUnit"`
	ProductDeliveryUnit                             string   `json:"ProductDeliveryUnit"`
	StandardLotSizeQuantity                         float32  `json:"StandardLotSizeQuantity"`
	MinimumLotSizeQuantity                          float32  `json:"MinimumLotSizeQuantity"`
	MaximumLotSizeQuantity                          float32  `json:"MaximumLotSizeQuantity"`
	OperationPlannedQuantityInBaseUnit              float32  `json:"OperationPlannedQuantityInBaseUnit"`
	OperationPlannedQuantityInProductionUnit        float32  `json:"OperationPlannedQuantityInProductionUnit"`
	OperationPlannedQuantityInOperationUnit         float32  `json:"OperationPlannedQuantityInOperationUnit"`
	OperationPlannedQuantityInDeliveryUnit          float32  `json:"OperationPlannedQuantityInDeliveryUnit"`
	OperationPlannedScrapInPercent                  *float32 `json:"OperationPlannedScrapInPercent"`
	ResponsiblePlannerGroup                         *string  `json:"ResponsiblePlannerGroup"`
	PlainLongText                                   *string  `json:"PlainLongText"`
	WorkCenter                                      int      `json:"WorkCenter"`
	CapacityCategory                                *string  `json:"CapacityCategory"`
	OperationCostingRelevancyType                   *string  `json:"OperationCostingRelevancyType"`
	OperationSetupType                              *string  `json:"OperationSetupType"`
	OperationSetupGroupCategory                     *string  `json:"OperationSetupGroupCategory"`
	OperationSetupGroup                             *string  `json:"OperationSetupGroup"`
	MaximumWaitDuration                             *float32 `json:"MaximumWaitDuration"`
	StandardWaitDuration                            *float32 `json:"StandardWaitDuration"`
	MinimumWaitDuration                             *float32 `json:"MinimumWaitDuration"`
	WaitDurationUnit                                *string  `json:"WaitDurationUnit"`
	MaximumQueDuration                              *float32 `json:"MaximumQueDuration"`
	StandardQueueDuration                           *float32 `json:"StandardQueueDuration"`
	MinimumQueueDuration                            *float32 `json:"MinimumQueueDuration"`
	QueDurationUnit                                 *string  `json:"QueDurationUnit"`
	MaximumMoveDuration                             *float32 `json:"MaximumMoveDuration"`
	StandardMoveDuration                            *float32 `json:"StandardMoveDuration"`
	MinimumMoveDuration                             *float32 `json:"MinimumMoveDuration"`
	MoveDurationUnit                                *string  `json:"MoveDurationUnit"`
	StandardDeliveryDuration                        *float32 `json:"StandardDeliveryDuration"`
	StandardDeliveryDurationUnit                    *string  `json:"StandardDeliveryDurationUnit"`
	CostElement                                     *string  `json:"CostElement"`
	OperationErlstSchedldExecStrtDte                *string  `json:"OperationErlstSchedldExecStrtDte"`
	OperationErlstSchedldExecStrtTme                *string  `json:"OperationErlstSchedldExecStrtTme"`
	OperationErlstSchedldExecEndDte                 *string  `json:"OperationErlstSchedldExecEndDte"`
	OperationErlstSchedldExecEndTme                 *string  `json:"OperationErlstSchedldExecEndTme"`
	OperationActualExecutionStartDate               *string  `json:"OperationActualExecutionStartDate"`
	OperationActualExecutionStartTime               *string  `json:"OperationActualExecutionStartTime"`
	OperationActualExecutionEndDate                 *string  `json:"OperationActualExecutionEndDate"`
	OperationActualExecutionEndTime                 *string  `json:"OperationActualExecutionEndTime"`
	OperationConfirmedYieldQuantityInBaseUnit       *float32 `json:"OperationConfirmedYieldQuantityInBaseUnit"`
	OperationConfirmedYieldQuantityInProductionUnit *float32 `json:"OperationConfirmedYieldQuantityInProductionUnit"`
	OperationConfirmedYieldQuantityInOperationUnit  *float32 `json:"OperationConfirmedYieldQuantityInOperationUnit"`
	CreationDate                                    string   `json:"CreationDate"`
	CreationTime                                    string   `json:"CreationTime"`
	LastChangeDate                                  string   `json:"LastChangeDate"`
	LastChangeTime                                  string   `json:"LastChangeTime"`
	IsReleased                                      *bool    `json:"IsReleased"`
	IsPartiallyConfirmed                            *bool    `json:"IsPartiallyConfirmed"`
	IsConfirmed                                     *bool    `json:"IsConfirmed"`
	IsLocked                                        *bool    `json:"IsLocked"`
	IsCancelled                                     *bool    `json:"IsCancelled"`
	IsMarkedForDeletion                             *bool    `json:"IsMarkedForDeletion"`
}

type ProductionOrderItemSingleUnit struct {
	SizeOrDimensionText                         *string  `json:"SizeOrDimensionText"`
	SafetyStockQuantityInBaseUnit               *float32 `json:"SafetyStockQuantityInBaseUnit"`
	InternalCapacityQuantity                    *float32 `json:"InternalCapacityQuantity"`
	ReorderThresholdQuantityInBaseUnit          *float32 `json:"ReorderThresholdQuantityInBaseUnit"`
	StandardProductionLotSizeQuantityInBaseUnit *float32 `json:"StandardProductionLotSizeQuantityInBaseUnit"`
	Images                                      Images   `json:"Images"`
}
