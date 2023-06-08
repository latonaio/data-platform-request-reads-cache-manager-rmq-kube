package dpfm_api_output_formatter

type ProductionOrderDetailList struct {
	Header  ProductionOrderDetailHeader  `json:"Header"`
	Details []ProductionOrderItemSummary `json:"Details"`
}
type ProductionOrderItemSummary struct {
	ProductionOrderItem      int      `json:"ProductionOrderItem"`
	Product                  string   `json:"Product"`
	ProductName              *string  `json:"ProductName"`
	MRPArea                  *string  `json:"MRPArea"`
	OrderItemTextBySeller    string   `json:"OrderItemTextBySeller"`
	TotalQuantity            *float32 `json:"TotalQuantity"`
	ConfirmedYieldQuantity   *float32 `json:"ConfirmedYieldQuantity"`
	ItemIsConfirmed          *bool    `json:"ItemIsConfirmed"`
	ItemIsPartiallyConfirmed *bool    `json:"ItemIsPartiallyConfirmed"`
	ItemIsReleased           *bool    `json:"ItemIsReleased"`
}

type ProductionOrderDetailHeader struct {
	Index int    `json:"Index"`
	Key   string `json:"Key"`
}
