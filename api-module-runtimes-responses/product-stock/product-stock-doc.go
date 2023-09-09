package apiModuleRuntimesResponsesProductStock

type ProductStockDocRes struct {
	Message  *ResponseProductStockDoc `json:"message,omitempty"`
	Accepter []string                 `json:"accepter,omitempty"`
}

type ResponseProductStockDoc struct {
	ProductStockDoc *[]ProductStockDoc `json:"ProductStockDoc"`
}

type ProductStockDoc struct {
	Product				 	 string `json:"Product"`
	BusinessPartner			 int    `json:"BusinessPartner"`
	Plant				 	 string `json:"Plant"`
	DocType                  string `json:"DocType"`
	FileExtension            string `json:"FileExtension"`
	DocVersionID             int    `json:"DocVersionID"`
	DocID                    string `json:"DocID"`
	DocIssuerBusinessPartner int    `json:"DocIssuerBusinessPartner"`
	FilePath                 string `json:"FilePath"`
	FileName                 string `json:"FileName"`
}
