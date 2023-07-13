package apiInputReader

type BillOfMaterial struct {
	BillOfMaterialHeader *BillOfMaterialHeader
	BillOfMaterialItems  *BillOfMaterialItems
	BillOfMaterialItem   *BillOfMaterialItem
}

type BillOfMaterialHeader struct {
	BillOfMaterial      int   `json:"BillOfMaterial"`
	IsMarkedForDeletion *bool `json:"IsMarkedForDeletion"`
}

type BillOfMaterialItems struct {
	BillOfMaterial      int   `json:"BillOfMaterial"`
	IsMarkedForDeletion *bool `json:"IsMarkedForDeletion"`
}

type BillOfMaterialItem struct {
	BillOfMaterial      int   `json:"BillOfMaterial"`
	BillOfMaterialItem  int   `json:"BillOfMaterialItem"`
	IsMarkedForDeletion *bool `json:"IsMarkedForDeletion"`
}
