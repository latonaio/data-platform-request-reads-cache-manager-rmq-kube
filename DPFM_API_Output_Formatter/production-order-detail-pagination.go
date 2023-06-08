package dpfm_api_output_formatter

type ProductionOrderDetailPagination struct {
	Paginations []ProductionOrderDetailPage `json:"Paginations"`
}

type ProductionOrderDetailPage struct {
	ProductionOrder     int    `json:"ProductionOrder"`
	ProductionOrderItem int    `json:"ProductionOrderItem"`
	Product             string `json:"Product"`
}
