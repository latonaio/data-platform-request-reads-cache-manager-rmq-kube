package apiModuleRuntimesResponsesPointTransactionType

type PointTransactionTypeRes struct {
	Message PointTransactionTypeGlobal `json:"message,omitempty"`
}

type PointTransactionTypeGlobal struct {
	PointTransactionType   *[]PointTransactionType    `json:"PointTransactionType,omitempty"`
	Text                   *[]Text                    `json:"Text,omitempty"`
}

type PointTransactionType struct {
	PointTransactionType    string	`json:"PointTransactionType"`
	CreationDate		    string	`json:"CreationDate"`
	LastChangeDate		    string	`json:"LastChangeDate"`
	IsMarkedForDeletion	    *bool	`json:"IsMarkedForDeletion"`
}

type Text struct {
	PointTransactionType       string  `json:"PointTransactionType"`
	Language          	       string  `json:"Language"`
	PointTransactionTypeName   string  `json:"PointTransactionTypeName"`
	CreationDate		       string  `json:"CreationDate"`
	LastChangeDate		       string  `json:"LastChangeDate"`
	IsMarkedForDeletion	       *bool   `json:"IsMarkedForDeletion"`
}
