package apiInputReader

type Quotations struct {
	QuotationsHeader *QuotationsHeader
	QuotationsItems  *QuotationsItems
	QuotationsItem   *QuotationsItem
}

type QuotationsHeader struct {
	Quotation            int   `json:"Quotation"`
	Buyer                *int  `json:"Buyer"`
	Seller               *int  `json:"Seller"`
	HeaderOrderIsDefined *bool `json:"HeaderOrderIsDefined"`
	HeaderIsClosed       *bool `json:"HeaderIsClosed"`
	HeaderBlockStatus    *bool `json:"HeaderBlockStatus"`
	IsCancelled          *bool `json:"IsCancelled"`
	IsMarkedForDeletion  *bool `json:"IsMarkedForDeletion"`
}

type QuotationsItems struct {
	Quotation           int   `json:"Quotation"`
	IsCancelled         *bool `json:"IsCancelled"`
	IsMarkedForDeletion *bool `json:"IsMarkedForDeletion"`
}

type QuotationsItem struct {
	Quotation           int   `json:"Quotation"`
	QuotationItem       int   `json:"QuotationItem"`
	IsCancelled         *bool `json:"IsCancelled"`
	IsMarkedForDeletion *bool `json:"IsMarkedForDeletion"`
}
