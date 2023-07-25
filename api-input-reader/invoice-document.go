package apiInputReader

type InvoiceDocument struct {
	InvoiceDocumentHeader *InvoiceDocumentHeader
	InvoiceDocumentItems  *InvoiceDocumentItems
	InvoiceDocumentItem   *InvoiceDocumentItem
}

type InvoiceDocumentHeader struct {
	InvoiceDocument int   `json:"InvoiceDocument"`
	BillToParty     *int  `json:"BillToParty"`
	BillFromParty   *int  `json:"BillFromParty"`
	IsCancelled     *bool `json:"IsCancelled"`
}

type InvoiceDocumentItems struct {
	InvoiceDocument     int   `json:"InvoiceDocument"`
	IsCancelled         *bool `json:"IsCancelled"`
}

type InvoiceDocumentItem struct {
	InvoiceDocument     int   `json:"InvoiceDocument"`
	InvoiceDocumentItem int   `json:"InvoiceDocumentItem"`
	IsCancelled         *bool `json:"IsCancelled"`
}
