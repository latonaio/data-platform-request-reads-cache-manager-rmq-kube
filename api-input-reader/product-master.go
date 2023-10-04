package apiInputReader

type ProductMaster struct {
	ProductMasterGeneral         *ProductMasterGeneral
	ProductMasterBusinessPartner *ProductMasterBusinessPartner
	ProductMasterBPPlant         *ProductMasterBPPlant
	ProductMasterGeneralDoc      *ProductMasterGeneralDoc
}

type ProductMasterGeneral struct {
	Product             string `json:"Product"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}

type ProductMasterBusinessPartner struct {
	Product             string `json:"Product"`
	BusinessPartner     int    `json:"BusinessPartner"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}

type ProductMasterBPPlant struct {
	Product             string `json:"Product"`
	BusinessPartner     int    `json:"BusinessPartner"`
	Plant               string `json:"Plant"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}

type ProductMasterGeneralDoc struct {
	Product                  string  `json:"Product"`
	BusinessPartner          *int    `json:"BusinessPartner"`
	DocType                  *string `json:"DocType"`
	DocIssuerBusinessPartner *int    `json:"DocIssuerBusinessPartner"`
}
