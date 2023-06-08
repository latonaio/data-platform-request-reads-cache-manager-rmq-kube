package dpfm_api_output_formatter

type InvoiceList struct {
	Invoices []Invoices `json:"InvoiceDocuments"`
}

type Invoices struct {
	InvoiceDocument          int     `json:"InvoiceDocument"`
	BillToParty              *string `json:"BillToParty"`
	BillFromParty            *string `json:"BillFromParty"`
	InvoiceDocumentDate      *string `json:"InvoiceDocumentDate"`
	InvoiceDocumentTime      *string `json:"InvoiceDocumentTime"`
	PaymentDueDate           *string `json:"PaymentDueDate"`
	HeaderBillingIsConfirmed *bool   `json:"HeaderBillingIsConfirmed"`

	IsCancelled         *bool    `json:"IsCancelled"`
	IsMarkedForDeletion *bool    `json:"IsMarkedForDeletion"`
}
