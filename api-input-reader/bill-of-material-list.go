package apiInputReader

type BillOfMaterialListParams struct {
	//IsCancelled         *bool `json:"IsCancelled"`
	IsMarkedForDeletion *bool `json:"IsMarkedForDeletion"`
}
