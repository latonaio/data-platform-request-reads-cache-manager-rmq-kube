package models

type OperationsReq struct {
	ConnectionKey     string           `json:"connection_key"`
	Result            bool             `json:"result"`
	RedisKey          string           `json:"redis_key"`
	Filepath          string           `json:"filepath"`
	APIStatusCode     int              `json:"api_status_code"`
	RuntimeSessionID  string           `json:"runtime_session_id"`
	BusinessPartnerID *int             `json:"business_partner"`
	ServiceLabel      string           `json:"service_label"`
	APIType           string           `json:"api_type"`
	Header            OperationsHeader `json:"Operations"`
	APISchema         string           `json:"api_schema"`
	Accepter          []string         `json:"accepter"`
	Deleted           bool             `json:"deleted"`
}

type OperationsHeader struct {
	Operations              int              `json:"Operations"`
	Product                 *string          `json:"Product"`
	OwnerBusinessPartner    *int             `json:"OwnerBusinessPartner"`
	OwnerPlant              *string          `json:"OwnerPlant"`
	OperationsText          *string          `json:"OperationsText"`
	OperationsStatus        *string          `json:"OperationsStatus"`
	ResponsiblePlannerGroup *string          `json:"ResponsiblePlannerGroup"`
	ValidityStartDate       *string          `json:"ValidityStartDate"`
	ValidityEndDate         *string          `json:"ValidityEndDate"`
	CreationDate            *string          `json:"CreationDate"`
	LastChangeDate          *string          `json:"LastChangeDate"`
	PlainLongText           *string          `json:"PlainLongText"`
	IsMarkedForDeletion     *bool            `json:"IsMarkedForDeletion"`
	Item                    []OperationsItem `json:"Item"`
}

type OperationsItem struct {
	Operations                    int                             `json:"Operations"`
	OperationssItem               int                             `json:"OperationssItem"`
	Sequence                      *int                            `json:"Sequence"`
	Product                       *string                         `json:"Product"`
	BusinessPartner               *int                            `json:"BusinessPartner"`
	Plant                         *string                         `json:"Plant"`
	OperationsText                *string                         `json:"OperationsText"`
	SequenceText                  *string                         `json:"SequenceText"`
	OperationsStatus              *string                         `json:"OperationsStatus"`
	ResponsiblePlannerGroup       *string                         `json:"ResponsiblePlannerGroup"`
	ValidityStartDate             *string                         `json:"ValidityStartDate"`
	ValidityEndDate               *string                         `json:"ValidityEndDate"`
	StandardLotSizeQuantity       *float32                        `json:"StandardLotSizeQuantity"`
	MinimumLotSizeQuantity        *float32                        `json:"MinimumLotSizeQuantity"`
	MaximumLotSizeQuantity        *float32                        `json:"MaximumLotSizeQuantity"`
	OperationsUnit                *string                         `json:"OperationsUnit"`
	CreationDate                  *string                         `json:"CreationDate"`
	LastChangeDate                *string                         `json:"LastChangeDate"`
	PlainLongText                 *string                         `json:"PlainLongText"`
	WorkCenterTypeCode            *string                         `json:"WorkCenterTypeCode"`
	WorkCenterInternalID          *string                         `json:"WorkCenterInternalID"`
	CapacityCategoryCode          *string                         `json:"CapacityCategoryCode"`
	OperationCostingRelevancyType *string                         `json:"OperationCostingRelevancyType"`
	NumberOfTimeTickets           *string                         `json:"NumberOfTimeTickets"`
	NumberOfConfirmationSlips     *string                         `json:"NumberOfConfirmationSlips"`
	OperationSetupType            *string                         `json:"OperationSetupType"`
	OperationSetupGroupCategory   *string                         `json:"OperationSetupGroupCategory"`
	OperationSetupGroup           *string                         `json:"OperationSetupGroup"`
	OperationReferenceQuantity    *float32                        `json:"OperationReferenceQuantity"`
	OpQtyToBaseQtyNmrtr           *float32                        `json:"OpQtyToBaseQtyNmrtr"`
	OpQtyToBaseQtyDnmntr          *float32                        `json:"OpQtyToBaseQtyDnmntr"`
	MaximumWaitDuration           *float32                        `json:"MaximumWaitDuration"`
	MaximumWaitDurationUnit       *string                         `json:"MaximumWaitDurationUnit"`
	MinimumWaitDuration           *float32                        `json:"MinimumWaitDuration"`
	MinimumWaitDurationUnit       *string                         `json:"MinimumWaitDurationUnit"`
	StandardQueueDuration         *float32                        `json:"StandardQueueDuration"`
	StandardQueueDurationUnit     *string                         `json:"StandardQueueDurationUnit"`
	MinimumQueueDuration          *float32                        `json:"MinimumQueueDuration"`
	MinimumQueueDurationUnit      *string                         `json:"MinimumQueueDurationUnit"`
	StandardMoveDuration          *float32                        `json:"StandardMoveDuration"`
	StandardMoveDurationUnit      *string                         `json:"StandardMoveDurationUnit"`
	MinimumMoveDuration           *float32                        `json:"MinimumMoveDuration"`
	MinimumMoveDurationUnit       *string                         `json:"MinimumMoveDurationUnit"`
	OpIsExtlyProcdWithSubcontrg   *bool                           `json:"OpIsExtlyProcdWithSubcontrg"`
	PlannedDeliveryDuration       *int                            `json:"PlannedDeliveryDuration"`
	Seller                        *int                            `json:"Seller"`
	NumberOfOperationPriceUnits   *int                            `json:"NumberOfOperationPriceUnits"`
	CostElement                   *string                         `json:"CostElement"`
	OpExternalProcessingPrice     *float32                        `json:"OpExternalProcessingPrice"`
	OpExternalProcessingCurrency  *string                         `json:"OpExternalProcessingCurrency"`
	OperationScrapPercent         *float32                        `json:"OperationScrapPercent"`
	IsMarkedForDeletion           *bool                           `json:"IsMarkedForDeletion"`
	ComponentAllocation           []OperationsComponentAllocation `json:"ComponentAllocation"`
}

type OperationsComponentAllocation struct {
	Operations           int     `json:"Operations"`
	OperationsItem       int     `json:"OperationsItem"`
	BillOfMaterial       int     `json:"BillOfMaterial"`
	BillOfMaterialItem   int     `json:"BillOfMaterialItem"`
	IsMarkedForBackflush *bool   `json:"IsMarkedForBackflush"`
	ValidityStartDate    *string `json:"ValidityStartDate"`
	ValidityEndDate      *string `json:"ValidityEndDate"`
	CreationDate         *string `json:"CreationDate"`
	LastChangeDate       *string `json:"LastChangeDate"`
}
