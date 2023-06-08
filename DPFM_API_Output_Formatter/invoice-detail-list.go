package dpfm_api_output_formatter

type InvoiceDocumentDetailList struct {
	InvoiceDocumentDetailHeader InvoiceDocumentDetailHeader    `json:"InvoiceDocumentDetailHeader"`
	InvoiceDetail               []InvoiceDocumentDetailSummary `json:"InvoiceDocumentDetails"`
}

type InvoiceDocumentDetailSummary struct {
	InvoiceDocument           int                   `json:"InvoiceDocument"`
	InvoiceDocumentItem       int                   `json:"InvoiceDocumentItem"`
	Product                   *string               `json:"Product"`
	InvoiceDocumentItemText   *string               `json:"InvoiceDocumentItemText"`
	InvoiceQuantityInBaseUnit *float32              `json:"InvoiceQuantityInBaseUnit"`
	InvoiceQuantityUnit       *string               `json:"InvoiceQuantityUnit"`
	ActualGoodsIssueDate      *string               `json:"ActualGoodsIssueDate"`
	ActualGoodsIssueTime      *string               `json:"ActualGoodsIssueTime"`
	ActualGoodsReceiptDate    *string               `json:"ActualGoodsReceiptDate"`
	ActualGoodsReceiptTime    *string               `json:"ActualGoodsReceiptTime"`
	ItemBillingIsConfirmed    *bool                 `json:"ItemBillingIsConfirmed"`
	IsCancelled               *bool                 `json:"IsCancelled"`
	IsMarkedForDeletion       *bool                 `json:"IsMarkedForDeletion"`
	OrdersDetailJumpReq       OrdersDetailJumpReq   `json:"OrdersDetailJumpReq"`
	DeliveryDetailJumpReq     DeliveryDetailJumpReq `json:"DeliveryDocumentDetailJumpReq"`
}

type OrdersDetailJumpReq struct {
	OrderID   *int    `json:"OrderID"`
	OrderItem *int    `json:"OrderItem"`
	Product   *string `json:"Product"`
	Buyer     *int    `json:"Buyer"`
}

type DeliveryDetailJumpReq struct {
	DeliveryDocument     int     `json:"DeliveryDocument"`
	DeliveryDocumentItem int     `json:"DeliveryDocumentItem"`
	DeliverToParty       *int    `json:"DeliverToParty"`
	DeliverFromParty     *int    `json:"DeliverFromParty"`
	Product              *string `json:"Product"`
	Buyer                *int    `json:"Buyer"`
}

type InvoiceDocumentDetailHeader struct {
	Index int    `json:"Index"`
	Key   string `json:"Key"`
}
