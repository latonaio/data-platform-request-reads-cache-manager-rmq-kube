package dpfm_api_output_formatter

type ProductionVersionDetailList struct {
	Header             ProductionVersionDetailHeader `json:"Header"`
	ProductionVersions []ProductionVersionDetail     `json:"Details"`
}

type ProductionVersionDetail struct {
	ProductionVersion     *int    `json:"ProductionVersion"`
	ProductionVersionItem *int    `json:"ProductionVersionItem"`
	Product               *string `json:"Product"`
	ProductDescription    *string `json:"ProductDescription"`
	OperationsText        *string `json:"OperationsText"`
	Plant                 *string `json:"Plant"`
	PlantName             *string `json:"PlantName"`
	BillOfMaterial        *int    `json:"BillOfMaterial"`
	Operations            *int    `json:"Operations"`
	ValidityStartDate     *string `json:"ValidityStartDate"`
	IsMarkedForDeletion   *bool   `json:"IsMarkedForDeletion"`
	// OperationDetailJumpReq OrdersDetailJumpReq   `json:"OperationDetailJumpReq"`
	// BoMDetailJumpReq       DeliveryDetailJumpReq `json:"BoMDetailJumpReq"`
}

type ProductionVersionDetailHeader struct {
	Index int    `json:"Index"`
	Key   string `json:"Key"`
}
