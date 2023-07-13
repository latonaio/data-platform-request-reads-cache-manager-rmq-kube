package apiInputReader

type Operations struct {
	OperationsHeader *OperationsHeader
	OperationsItems  *OperationsItems
	OperationsItem   *OperationsItem
}

type OperationsHeader struct {
	Operations          int   `json:"Operations"`
	IsMarkedForDeletion *bool `json:"IsMarkedForDeletion"`
}

type OperationsItems struct {
	Operations          int   `json:"Operations"`
	IsMarkedForDeletion *bool `json:"IsMarkedForDeletion"`
}

type OperationsItem struct {
	Operations          int   `json:"Operations"`
	OperationsItem      int   `json:"OperationsItem"`
	IsMarkedForDeletion *bool `json:"IsMarkedForDeletion"`
}
