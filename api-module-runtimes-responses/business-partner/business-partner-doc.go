package apiModuleRuntimesResponsesBusinessPartner

type BusinessPartnerDocRes struct {
	Message  *ResponseGeneralDoc `json:"message,omitempty"`
	Accepter []string            `json:"accepter,omitempty"`
}

type ResponseGeneralDoc struct {
	GeneralDoc *[]GeneralDoc `json:"GeneralDoc"`
}

type GeneralDoc struct {
	BusinessPartner          int    `json:"BusinessPartner"`
	DocType                  string `json:"DocType"`
	FileExtension            string `json:"FileExtension"`
	DocVersionID             int    `json:"DocVersionID"`
	DocID                    string `json:"DocID"`
	DocIssuerBusinessPartner int    `json:"DocIssuerBusinessPartner"`
	FilePath                 string `json:"FilePath"`
	FileName                 string `json:"FileName"`
}
