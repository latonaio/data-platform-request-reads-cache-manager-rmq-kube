package dpfm_api_output_formatter

type ProductionOrderDetail struct {
	ProductionOrder         int      `json:"ProductionOrder"`
	ProductionOrderItem     int      `json:"ProductionOrderItem"`
	OrderItemText           *string  `json:"OrderItemText"`
	Product                 *string  `json:"Product"`
	ProductName             *string  `json:"ProductName"`
	MRPArea                 *string  `json:"MRPArea"`
	ProductionVersion       *int     `json:"ProductionVersion"`
	MinimumLotSizeQuantity  *float32 `json:"MinimumLotSizeQuantity"`
	MaximumLotSizeQuantity  *float32 `json:"MaximumLotSizeQuantity"`
	StandardLotSizeQuantity *float32 `json:"StandardLotSizeQuantity"`
	LotSizeRoundingQuantity *float32 `json:"LotSizeRoundingQuantity"`

	ProductionOrderPlannedStartDate *string `json:"ProductionOrderPlannedStartDate"`
	ProductionOrderPlannedStartTime *string `json:"ProductionOrderPlannedStartTime"`
	ProductionOrderPlannedEndDate   *string `json:"ProductionOrderPlannedEndDate"`
	ProductionOrderPlannedEndTime   *string `json:"ProductionOrderPlannedEndTime"`

	ProductionOrderActualStartDate *string `json:"ProductionOrderActualStartDate"`
	ProductionOrderActualStartTime *string `json:"ProductionOrderActualStartTime"`
	ProductionOrderActualEndDate   *string `json:"ProductionOrderActualEndDate"`
	ProductionOrderActualEndTime   *string `json:"ProductionOrderActualEndTime"`

	TotalQuantity                  *float32 `json:"TotalQuantity"`
	PlannedScrapQuantity           *float32 `json:"PlannedScrapQuantity"`
	ConfirmedYieldQuantity         *float32 `json:"ConfirmedYieldQuantity"`
	ProductionUnit                 *string  `json:"ProductionUnit"`
	ProductionPlant                *string  `json:"ProductionPlant"`
	ProductionPlantStorageLocation *string  `json:"ProductionPlantStorageLocation"`
	BillOfMaterialItem             *string  `json:"BillOfMaterialItem"`

	Components []ComponentItem `json:"Components"`
	Operations []Operation     `json:"Operations"`
	Images     Images          `json:"Images"`
}

type ComponentItem struct {
	ComponentProduct                *string
	ComponentProductRequirementDate *string
	ComponentProductRequirementTime *string
	RequiredQuantity                *float32
	WithdrawnQuantity               *float32
	BaseUnit                        *string
	CostingPolicy                   *string
	StandardPrice                   *float32
	MovingAveragePrice              *float32
}

type Operation struct {
	OperationText                        *string
	WorkCenter                           *int
	OperationPlannedTotalQuantity        *float32 `json:"OperationPlannedTotalQuantity"`
	OperationTotalConfirmedYieldQuantity *float32 `json:"OperationTotalConfirmedYieldQuantity"`
	OperationErlstSchedldExecStrtDte     *string  `json:"OperationErlstSchedldExecStrtDte"`
	OperationErlstSchedldExecStrtTme     *string  `json:"OperationErlstSchedldExecStrtTme"`
	OperationErlstSchedldExecEndDate     *string  `json:"OperationErlstSchedldExecEndDate"`
	OperationErlstSchedldExecEndTime     *string  `json:"OperationErlstSchedldExecEndTime"`
	OperationActualExecutionStartDate    *string  `json:"OperationActualExecutionStartDate"`
	OperationActualExecutionStartTime    *string  `json:"OperationActualExecutionStartTime"`
	OperationActualExecutionEndDate      *string  `json:"OperationActualExecutionEndDate"`
	OperationActualExecutionEndTime      *string  `json:"OperationActualExecutionEndTime"`
}
