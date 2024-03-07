package apiModuleRuntimesResponsesPaymentTerms

type PaymentTermsRes struct {
	Message PaymentTermsGlobal `json:"message,omitempty"`
}

type PaymentTermsGlobal struct {
	PaymentTerms    *[]PaymentTerms    `json:"PaymentTerms,omitempty"`
	Text            *[]Text            `json:"Text,omitempty"`
}

type PaymentTerms struct {
	PaymentTerms		string	`json:"PaymentTerms"`
	CreationDate		string	`json:"CreationDate"`
	LastChangeDate		string	`json:"LastChangeDate"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}

type Text struct {
	PaymentTerms     	string  `json:"PaymentTerms"`
	Language          	string  `json:"Language"`
	PaymentTermsName	string 	`json:"PaymentTermsName"`
	CreationDate		string	`json:"CreationDate"`
	LastChangeDate		string	`json:"LastChangeDate"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}
