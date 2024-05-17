package apiInputReader

type BusinessPartner struct {
	BusinessPartnerGeneral			*BusinessPartnerGeneral
	BusinessPartnerPerson			*BusinessPartnerPerson
	BusinessPartnerAddress			*BusinessPartnerAddress
	BusinessPartnerGPS				*BusinessPartnerGPS
	BusinessPartnerRank				*BusinessPartnerRank
	BusinessPartnerDocGeneralDoc	*BusinessPartnerDocGeneralDoc
}

type BusinessPartnerGeneral struct {
	BusinessPartner            int        `json:"BusinessPartner"`
	IsMarkedForDeletion        *bool      `json:"IsMarkedForDeletion"`
}

type BusinessPartnerPerson struct {
	BusinessPartner            int        `json:"BusinessPartner"`
	IsMarkedForDeletion        *bool      `json:"IsMarkedForDeletion"`
}

type BusinessPartnerAddress struct {
	BusinessPartner            int        `json:"BusinessPartner"`
	AddressID	               int        `json:"AddressID"`
	IsMarkedForDeletion        *bool      `json:"IsMarkedForDeletion"`
}

type BusinessPartnerGPS struct {
	BusinessPartner            int        `json:"BusinessPartner"`
	IsMarkedForDeletion        *bool      `json:"IsMarkedForDeletion"`
}

type BusinessPartnerRank struct {
	BusinessPartner            int        `json:"BusinessPartner"`
	RankType				   string     `json:"RankType"`
	IsMarkedForDeletion        *bool      `json:"IsMarkedForDeletion"`
}

type BusinessPartnerDocGeneralDoc struct {
	BusinessPartner				int     `json:"BusinessPartner"`
	DocType						*string `json:"DocType"`
	DocIssuerBusinessPartner	*int    `json:"DocIssuerBusinessPartner"`
}
