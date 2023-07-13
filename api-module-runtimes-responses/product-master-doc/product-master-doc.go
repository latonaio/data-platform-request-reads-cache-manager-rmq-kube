package apiModuleRuntimesResponsesProductMasterDoc

type ProductMasterDocRes struct {
	Message  *ResponseHeaderDoc `json:"message,omitempty"`
	Accepter []string           `json:"accepter,omitempty"`
}

type ResponseHeaderDoc struct {
	HeaderDoc *[]PMDHeaderDoc `json:"HeaderDoc"`
}

type PMDHeaderDoc struct {
	Product                  string `json:"Product"`
	DocType                  string `json:"DocType"`
	FileExtension            string `json:"FileExtension"`
	DocVersionID             int    `json:"DocVersionID"`
	DocID                    string `json:"DocID"`
	DocIssuerBusinessPartner int    `json:"DocIssuerBusinessPartner"`
	FilePath                 string `json:"FilePath"`
	FileName                 string `json:"FileName"`
}
