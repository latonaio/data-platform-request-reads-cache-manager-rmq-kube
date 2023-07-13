package apiInputReader

type InvoiceDocument struct {
	InvoiceDocumentHeader *InvoiceDocumentHeader
	InvoiceDocumentItems  *InvoiceDocumentItems
	InvoiceDocumentItem   *InvoiceDocumentItem
}

type InvoiceDocumentHeader struct {
	InvoiceDocument                 int     `json:"InvoiceDocument"`
	IsCancelled						*bool	`json:"IsCancelled"`
	IsMarkedForDeletion      		*bool   `json:"IsMarkedForDeletion"`
}

type InvoiceDocumentItems struct {
	InvoiceDocument          int    `json:"InvoiceDocument"`
	IsCancelled				 *bool	`json:"IsCancelled"`
	IsMarkedForDeletion      *bool  `json:"IsMarkedForDeletion"`
}

type InvoiceDocumentItem struct {
	InvoiceDocument      int   `json:"InvoiceDocument"`
	InvoiceDocumentItem  int   `json:"InvoiceDocumentItem"`
	IsCancelled			 *bool `json:"IsCancelled"`
	IsMarkedForDeletion  *bool `json:"IsMarkedForDeletion"`
}
