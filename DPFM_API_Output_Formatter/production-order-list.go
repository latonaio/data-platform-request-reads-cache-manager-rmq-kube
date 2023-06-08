package dpfm_api_output_formatter

type ProductionOrderList struct {
	ProductionOrders []ProductionOrder `json:"ProductionOrders"`
}
type ProductionOrder struct {
	ProductionOrder                     int      `json:"ProductionOrder"`
	MRPArea                             *string  `json:"MRPArea"`
	Product                             *string  `json:"Product"`
	ProductName                         *string  `json:"ProductName"`
	OwnerProductionPlantBusinessPartner *string  `json:"OwnerProductionPlantBusinessPartner"`
	OwnerProductionPlant                *string  `json:"OwnerProductionPlant"`
	TotalQuantity                       *float32 `json:"TotalQuantity"`

	ProductionOrderPlannedStartDate *string `json:"ProductionOrderPlannedStartDate"`
	ProductionOrderPlannedStartTime *string `json:"ProductionOrderPlannedStartTime"`
	ProductionOrderPlannedEndDate   *string `json:"ProductionOrderPlannedEndDate"`
	ProductionOrderPlannedEndTime   *string `json:"ProductionOrderPlannedEndTime"`

	HeaderIsConfirmed          *bool `json:"HeaderIsConfirmed"`
	HeaderIsPartiallyConfirmed *bool `json:"HeaderIsPartiallyConfirmed"`
	HeaderIsReleased           *bool `json:"HeaderIsReleased"`
	IsCancelled                *bool `json:"IsCancelled"`
	IsMarkedForDeletion        *bool `json:"IsMarkedForDeletion"`
}
