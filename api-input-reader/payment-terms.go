package apiInputReader

type PaymentTermsGlobal struct {
	PaymentTerms     *PaymentTerms
	PaymentTermsText *PaymentTermsText
}

type PaymentTerms struct {
	PaymentTerms        string `json:"PaymentTerms"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}

type PaymentTermsText struct {
	PaymentTerms        string `json:"PaymentTerms"`
	Language            string `json:"Language"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}
