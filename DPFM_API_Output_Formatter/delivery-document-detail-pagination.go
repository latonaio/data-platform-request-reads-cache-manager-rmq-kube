package dpfm_api_output_formatter

type DeliveryDocumentDetailPagination struct {
	Paginations []DeliveryDocumentDetailPage `json:"Paginations"`
}

type DeliveryDocumentDetailPage struct {
	DeliveryDocument     int    `json:"DeliveryDocument"`
	DeliveryDocumentItem int    `json:"DeliveryDocumentItem"`
	Product              string `json:"Product"`
}
