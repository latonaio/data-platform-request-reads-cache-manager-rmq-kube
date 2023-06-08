package dpfm_api_output_formatter

type OrdersDetailPagination struct {
	Paginations []OderDetailPage `json:"Paginations"`
}

type OderDetailPage struct {
	OrderID   int    `json:"OrderID"`
	OrderItem int    `json:"OrderItem"`
	Product   string `json:"Product"`
}
