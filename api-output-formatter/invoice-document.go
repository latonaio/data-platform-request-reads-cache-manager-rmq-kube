package apiOutputFormatter

type InvoiceDocument struct {
	InvoiceDocumentHeader  []InvoiceDocumentHeader  `json:"Header"`
	InvoiceDocumentItem    []InvoiceDocumentItem    `json:"Item"`
}

type InvoiceDocumentHeader struct {
	InvoiceDocument           int     `json:"InvoiceDocument"`
	BillToParty               int     `json:"BillToParty"`
	BillToPartyName           string  `json:"BillToPartyName"`
	BillFromParty             int     `json:"BillFromParty"`
	BillFromPartyName         string  `json:"BillFromPartyName"`
    InvoiceDocumentDate       string  `json:"InvoiceDocumentDate"`
    PaymentDueDate            *string `json:"PaymentDueDate"`
    HeaderBillingIsConfirmed  *bool   `json:"HeaderBillingIsConfirmed"`
    IsCancelled               *bool   `json:"IsCancelled"`
}

type InvoiceDocumentItem struct {
	InvoiceDocumentItem                 int     `json:"InvoiceDocumentItem"`
	Product                             string  `json:"Product"`
    InvoiceDocumentItemTextByBuyer      string  `json:"InvoiceDocumentItemTextByBuyer"`
    InvoiceDocumentItemTextBySeller     string  `json:"InvoiceDocumentItemTextBySeller"`
    InvoiceQuantityInBaseUnit           float32 `json:"InvoiceQuantityInBaseUnit"`
    InvoiceQuantityUnit                 string  `json:"InvoiceQuantityUnit"`
    ActualGoodsIssueDate                *string `json:"ActualGoodsIssueDate"`
    ActualGoodsIssueTime                *string `json:"ActualGoodsIssueTime"`
    ActualGoodsReceiptDate              *string `json:"ActualGoodsReceiptDate"`
    ActualGoodsReceiptTime              *string `json:"ActualGoodsReceiptTime"`
	ItemBillingIsConfirmed              *bool   `json:"ItemBillingIsConfirmed"`
    IsCancelled                         *bool   `json:"IsCancelled"`
	Images                              Images  `json:"Images"`
}
