package apiModuleRuntimesResponsesProductMaster

type ProductMasterDocRes struct {
	Message  *ResponseGeneralDoc `json:"message,omitempty"`
	Accepter []string            `json:"accepter,omitempty"`
}

type ResponseGeneralDoc struct {
	GeneralDoc *[]GeneralDoc `json:"GeneralDoc"`
}

type GeneralDoc struct {
	Product                  string `json:"Product"`
	DocType                  string `json:"DocType"`
	FileExtension            string `json:"FileExtension"`
	DocVersionID             int    `json:"DocVersionID"`
	DocID                    string `json:"DocID"`
	DocIssuerBusinessPartner int    `json:"DocIssuerBusinessPartner"`
	FilePath                 string `json:"FilePath"`
	FileName                 string `json:"FileName"`
}
