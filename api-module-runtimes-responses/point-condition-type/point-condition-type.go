package apiModuleRuntimesResponsesPointConditionType

type PointConditionTypeRes struct {
	Message PointConditionTypeGlobal `json:"message,omitempty"`
}

type PointConditionTypeGlobal struct {
	PointConditionType     *[]PointConditionType    `json:"PointConditionType,omitempty"`
	Text                   *[]Text                  `json:"Text,omitempty"`
}

type PointConditionType struct {
	PointConditionType  string	`json:"PointConditionType"`
	CreationDate		string	`json:"CreationDate"`
	LastChangeDate		string	`json:"LastChangeDate"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}

type Text struct {
	PointConditionType         string  `json:"PointConditionType"`
	Language          	       string  `json:"Language"`
	PointConditionTypeName	   string  `json:"PointConditionTypeName"`
	CreationDate		       string  `json:"CreationDate"`
	LastChangeDate		       string  `json:"LastChangeDate"`
	IsMarkedForDeletion	       *bool   `json:"IsMarkedForDeletion"`
}
