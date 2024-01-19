package apiModuleRuntimesResponsesInspectionLot

type InspectionLotRes struct {
	Message InspectionLot `json:"message,omitempty"`
}

type InspectionLot struct {
	Header               *[]Header               `json:"Header,omitempty"`
	SpecDetail           *[]SpecDetail           `json:"SpecDetail,omitempty"`
	ComponentComposition *[]ComponentComposition `json:"ComponentComposition,omitempty"`
	Inspection           *[]Inspection           `json:"Inspection,omitempty"`
	Operation            *[]Operation            `json:"Operation,omitempty"`
}

type Header struct {
	InspectionLot                  int     `json:"InspectionLot"`
	InspectionLotDate              string  `json:"InspectionLotDate"`
	InspectionPlan                 int     `json:"InspectionPlan"`
	InspectionPlantBusinessPartner int     `json:"InspectionPlantBusinessPartner"`
	InspectionPlant                string  `json:"InspectionPlant"`
	Product                        string  `json:"Product"`
	ProductSpecification           *string `json:"ProductSpecification"`
	InspectionSpecification        *string `json:"InspectionSpecification"`
	ProductionOrder                *int    `json:"ProductionOrder"`
	ProductionOrderItem            *int    `json:"ProductionOrderItem"`
	InspectionLotHeaderText        *string `json:"InspectionLotHeaderText"`
	ExternalReferenceDocument      *string `json:"ExternalReferenceDocument"`
	CertificateAuthorityChain      *string `json:"CertificateAuthorityChain"`
	UsageControlChain              *string `json:"UsageControlChain"`
	CreationDate                   string  `json:"CreationDate"`
	CreationTime                   string  `json:"CreationTime"`
	LastChangeDate                 string  `json:"LastChangeDate"`
	LastChangeTime                 string  `json:"LastChangeTime"`
	IsReleased                     *bool   `json:"IsReleased"`
	IsPartiallyConfirmed           *bool   `json:"IsPartiallyConfirmed"`
	IsConfirmed                    *bool   `json:"IsConfirmed"`
	IsLocked                       *bool   `json:"IsLocked"`
	IsCancelled                    *bool   `json:"IsCancelled"`
	IsMarkedForDeletion            *bool   `json:"IsMarkedForDeletion"`
}

type SpecDetail struct {
	InspectionLot       int     `json:"InspectionLot"`
	SpecType            string  `json:"SpecType"`
	UpperLimitValue     float32 `json:"UpperLimitValue"`
	LowerLimitValue     float32 `json:"LowerLimitValue"`
	StandardValue       float32 `json:"StandardValue"`
	SpecTypeUnit        string  `json:"SpecTypeUnit"`
	Formula             *string `json:"Formula"`
	CreationDate        string  `json:"CreationDate"`
	CreationTime        string  `json:"CreationTime"`
	LastChangeDate      string  `json:"LastChangeDate"`
	LastChangeTime      string  `json:"LastChangeTime"`
	IsReleased          *bool   `json:"IsReleased"`
	IsLocked            *bool   `json:"IsLocked"`
	IsCancelled         *bool   `json:"IsCancelled"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
}

type ComponentComposition struct {
	InspectionLot                              int     `json:"InspectionLot"`
	ComponentCompositionType                   string  `json:"ComponentCompositionType"`
	ComponentCompositionUpperLimitInPercent    float32 `json:"ComponentCompositionUpperLimitInPercent"`
	ComponentCompositionLowerLimitInPercent    float32 `json:"ComponentCompositionLowerLimitInPercent"`
	ComponentCompositionStandardValueInPercent float32 `json:"ComponentCompositionStandardValueInPercent"`
	CreationDate                               string  `json:"CreationDate"`
	CreationTime                               string  `json:"CreationTime"`
	LastChangeDate                             string  `json:"LastChangeDate"`
	LastChangeTime                             string  `json:"LastChangeTime"`
	IsReleased                                 *bool   `json:"IsReleased"`
	IsLocked                                   *bool   `json:"IsLocked"`
	IsCancelled                                *bool   `json:"IsCancelled"`
	IsMarkedForDeletion                        *bool   `json:"IsMarkedForDeletion"`
}

type Inspection struct {
	InspectionLot                            int      `json:"InspectionLot"`
	Inspection                               int      `json:"Inspection"`
	InspectionDate                           string   `json:"InspectionDate"`
	InspectionType                           string   `json:"InspectionType"`
	InspectionTypeValueUnit                  *string  `json:"InspectionTypeValueUnit"`
	InspectionTypePlannedValue               *float32 `json:"InspectionTypePlannedValue"`
	InspectionTypeCertificateType            *string  `json:"InspectionTypeCertificateType"`
	InspectionTypeCertificateValueInText     *string  `json:"InspectionTypeCertificateValueInText"`
	InspectionTypeCertificateValueInQuantity *float32 `json:"InspectionTypeCertificateValueInQuantity"`
	InspectionLotInspectionText              *string  `json:"InspectionLotInspectionText"`
	CreationDate                             string   `json:"CreationDate"`
	CreationTime                             string   `json:"CreationTime"`
	LastChangeDate                           string   `json:"LastChangeDate"`
	LastChangeTime                           string   `json:"LastChangeTime"`
	IsReleased                               *bool    `json:"IsReleased"`
	IsLocked                                 *bool    `json:"IsLocked"`
	IsCancelled                              *bool    `json:"IsCancelled"`
	IsMarkedForDeletion                      *bool    `json:"IsMarkedForDeletion"`
}

type Operation struct {
	InspectionLot                                   int      `json:"InspectionLot"`
	Operations                                      int      `json:"Operations"`
	OperationsItem                                  int      `json:"OperationsItem"`
	OperationID                                     int      `json:"OperationID"`
	Inspection                                      int      `json:"Inspection"`
	OperationType                                   string   `json:"OperationType"`
	SupplyChainRelationshipID                       int      `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID               int      `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID          int      `json:"SupplyChainRelationshipDeliveryPlantID"`
	SupplyChainRelationshipProductionPlantID        int      `json:"SupplyChainRelationshipProductionPlantID"`
	Product                                         string   `json:"Product"`
	Buyer                                           int      `json:"Buyer"`
	Seller                                          int      `json:"Seller"`
	DeliverToParty                                  int      `json:"DeliverToParty"`
	DeliverToPlant                                  string   `json:"DeliverToPlant"`
	DeliverFromParty                                int      `json:"DeliverFromParty"`
	DeliverFromPlant                                string   `json:"DeliverFromPlant"`
	InspectionPlantBusinessPartner                  int      `json:"InspectionPlantBusinessPartner"`
	InspectionPlant                                 string   `json:"InspectionPlant"`
	Sequence                                        int      `json:"Sequence"`
	SequenceText                                    *string  `json:"SequenceText"`
	OperationText                                   string   `json:"OperationText"`
	OperationStatus                                 *string  `json:"OperationStatus"`
	ResponsiblePlannerGroup                         *string  `json:"ResponsiblePlannerGroup"`
	OperationUnit                                   string   `json:"OperationUnit"`
	StandardLotSizeQuantity                         *float32 `json:"StandardLotSizeQuantity"`
	MinimumLotSizeQuantity                          *float32 `json:"MinimumLotSizeQuantity"`
	MaximumLotSizeQuantity                          *float32 `json:"MaximumLotSizeQuantity"`
	PlainLongText                                   *string  `json:"PlainLongText"`
	WorkCenter                                      *int     `json:"WorkCenter"`
	CapacityCategoryCode                            *string  `json:"CapacityCategoryCode"`
	OperationCostingRelevancyType                   *string  `json:"OperationCostingRelevancyType"`
	OperationSetupType                              *string  `json:"OperationSetupType"`
	OperationSetupGroupCategory                     *string  `json:"OperationSetupGroupCategory"`
	OperationSetupGroup                             *string  `json:"OperationSetupGroup"`
	OperationReferenceQuantity                      *float32 `json:"OperationReferenceQuantity"`
	MaximumWaitDuration                             *float32 `json:"MaximumWaitDuration"`
	StandardWaitDuration                            *float32 `json:"StandardWaitDuration"`
	MinimumWaitDuration                             *float32 `json:"MinimumWaitDuration"`
	WaitDurationUnit                                *string  `json:"WaitDurationUnit"`
	MaximumQueueDuration                            *float32 `json:"MaximumQueueDuration"`
	StandardQueueDuration                           *float32 `json:"StandardQueueDuration"`
	MinimumQueueDuration                            *float32 `json:"MinimumQueueDuration"`
	QueueDurationUnit                               *string  `json:"QueueDurationUnit"`
	MaximumMoveDuration                             *float32 `json:"MaximumMoveDuration"`
	StandardMoveDuration                            *float32 `json:"StandardMoveDuration"`
	MinimumMoveDuration                             *float32 `json:"MinimumMoveDuration"`
	MoveDurationUnit                                *string  `json:"MoveDurationUnit"`
	StandardDeliveryDuration                        *float32 `json:"StandardDeliveryDuration"`
	StandardDeliveryDurationUnit                    *string  `json:"StandardDeliveryDurationUnit"`
	StandardOperationScrapPercent                   *float32 `json:"StandardOperationScrapPercent"`
	PlannedOperationStandardValue                   *float32 `json:"PlannedOperationStandardValue"`
	PlannedOperationLowerValue                      *float32 `json:"PlannedOperationLowerValue"`
	PlannedOperationUpperValue                      *float32 `json:"PlannedOperationUpperValue"`
	PlannedOperationValueUnit                       *string  `json:"PlannedOperationValueUnit"`
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
	ValidityStartDate                               *string  `json:"ValidityStartDate"`
	ValidityEndDate                                 *string  `json:"ValidityEndDate"`
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
