package apiModuleRuntimesResponsesPointConsumptionType

type PointConsumptionTypeRes struct {
	Message PointConsumptionTypeGlobal `json:"message,omitempty"`
}

type PointConsumptionTypeGlobal struct {
	PointConsumptionType   *[]PointConsumptionType    `json:"PointConsumptionType,omitempty"`
	Text                   *[]Text                    `json:"Text,omitempty"`
}

type PointConsumptionType struct {
	PointConsumptionType    string	`json:"PointConsumptionType"`
	CreationDate		    string	`json:"CreationDate"`
	LastChangeDate		    string	`json:"LastChangeDate"`
	IsMarkedForDeletion	    *bool	`json:"IsMarkedForDeletion"`
}

type Text struct {
	PointConsumptionType       string  `json:"PointConsumptionType"`
	Language          	       string  `json:"Language"`
	PointConsumptionTypeName   string  `json:"PointConsumptionTypeName"`
	CreationDate		       string  `json:"CreationDate"`
	LastChangeDate		       string  `json:"LastChangeDate"`
	IsMarkedForDeletion	       *bool   `json:"IsMarkedForDeletion"`
}
