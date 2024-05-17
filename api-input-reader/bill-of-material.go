package apiInputReader

type BillOfMaterial struct {
	BillOfMaterialHeader	*BillOfMaterialHeader
	BillOfMaterialItems		*BillOfMaterialItems
	BillOfMaterialItem		*BillOfMaterialItem
	BillOfMaterialHeaderDoc	*BillOfMaterialHeaderDoc
	BillOfMaterialItemDoc	*BillOfMaterialItemDoc
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

type BillOfMaterialHeaderDoc struct {
	BillOfMaterial           int     `json:"BillOfMaterial"`
	DocType                  *string `json:"DocType"`
	DocIssuerBusinessPartner *int    `json:"DocIssuerBusinessPartner"`
}

type BillOfMaterialItemDoc struct {
	BillOfMaterial           int     `json:"BillOfMaterial"`
	BillOfMaterialItem       int     `json:"BillOfMaterialItem"`
	DocType                  *string `json:"DocType"`
	DocIssuerBusinessPartner *int    `json:"DocIssuerBusinessPartner"`
}
