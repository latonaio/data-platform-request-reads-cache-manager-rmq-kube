package apiInputReader

type BusinessPartner struct {
	BusinessPartnerGeneral     *BusinessPartnerGeneral
}

type BusinessPartnerGeneral struct {
	BusinessPartner            int        `json:"BusinessPartner"`
	IsMarkedForDeletion        *bool      `json:"IsMarkedForDeletion"`
}
