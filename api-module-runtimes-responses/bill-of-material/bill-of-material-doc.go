package apiModuleRuntimesResponsesBillOfMaterial

type BillOfMaterialDocRes struct {
	Message  *ResponseGeneralDoc `json:"message,omitempty"`
	Accepter []string            `json:"accepter,omitempty"`
}

type ResponseGeneralDoc struct {
	HeaderDoc *[]HeaderDoc `json:"HeaderDoc"`
	ItemDoc   *[]ItemDoc   `json:"ItemDoc"`
}

type HeaderDoc struct {
	BillOfMaterial           int    `json:"BillOfMaterial"`
	DocType                  string `json:"DocType"`
	FileExtension            string `json:"FileExtension"`
	DocVersionID             int    `json:"DocVersionID"`
	DocID                    string `json:"DocID"`
	DocIssuerBusinessPartner int    `json:"DocIssuerBusinessPartner"`
	FilePath                 string `json:"FilePath"`
	FileName                 string `json:"FileName"`
}

type ItemDoc struct {
	BillOfMaterial           int    `json:"BillOfMaterial"`
	BillOfMaterialItem       int    `json:"BillOfMaterialItem"`
	DocType                  string `json:"DocType"`
	FileExtension            string `json:"FileExtension"`
	DocVersionID             int    `json:"DocVersionID"`
	DocID                    string `json:"DocID"`
	DocIssuerBusinessPartner int    `json:"DocIssuerBusinessPartner"`
	FilePath                 string `json:"FilePath"`
	FileName                 string `json:"FileName"`
}
