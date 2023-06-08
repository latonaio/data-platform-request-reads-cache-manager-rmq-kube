package dpfm_api_output_formatter

type DeliveryDocumentDetailList struct {
	DeliveryDocumentDetailHeader DeliveryDocumentDetailHeader    `json:"DeliveryDocumentDetailHeader"`
	DeliveryDocumentDetail       []DeliveryDocumentDetailSummary `json:"DeliveryDocumentDetail"`
}
type DeliveryDocumentDetailSummary struct {
	DeliveryDocumentItem           int                `json:"DeliveryDocumentItem"`
	Product                        *string            `json:"Product"`
	DeliveryDocumentItemText       string             `json:"DeliveryDocumentItemText"`
	OriginalQuantityInDeliveryUnit int                `json:"OriginalQuantityInDeliveryUnit"`
	DeliveryUnit                   *string            `json:"DeliveryUnit"`
	ActualGoodsIssueDate           *string            `json:"ActualGoodsIssueDate"`
	ActualGoodsIssueTime           *string            `json:"ActualGoodsIssueTime"`
	ActualGoodsReceiptDate         *string            `json:"ActualGoodsReceiptDate"`
	ActualGoodsReceiptTime         *string            `json:"ActualGoodsReceiptTime"`
	IsCancelled                    *bool              `json:"IsCancelled"`
	IsMarkedForDeletion            *bool              `json:"IsMarkedForDeletion"`
	OrderDetailJumpReq             OrderDetailJumpReq `json:"OrdersDetailJumpReq"`
}

type OrderDetailJumpReq struct {
	OrderID   *int    `json:"OrderID"`
	OrderItem *int    `json:"OrderItem"`
	Product   *string `json:"Product"`
	Payer     *int    `json:"Payer"`
}

type DeliveryDocumentDetailHeader struct {
	Index int    `json:"Index"`
	Key   string `json:"Key"`
}
