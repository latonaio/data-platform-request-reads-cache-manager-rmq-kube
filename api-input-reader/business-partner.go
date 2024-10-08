package apiInputReader

type BusinessPartner struct {
	BusinessPartnerGeneral				*BusinessPartnerGeneral
	BusinessPartnerBPRole				*BusinessPartnerBPRole
	BusinessPartnerPerson				*BusinessPartnerPerson
	BusinessPartnerAddress				*BusinessPartnerAddress
	BusinessPartnerSNS					*BusinessPartnerSNS
	BusinessPartnerGPS					*BusinessPartnerGPS
	BusinessPartnerRank					*BusinessPartnerRank
	BusinessPartnerPersonOrganization	*BusinessPartnerPersonOrganization
	BusinessPartnerDocGeneralDoc		*BusinessPartnerDocGeneralDoc
}

type BusinessPartnerGeneral struct {
	BusinessPartner            int        `json:"BusinessPartner"`
	Withdrawal           	   *bool	  `json:"Withdrawal"`
	IsReleased           	   *bool	  `json:"IsReleased"`
	IsMarkedForDeletion        *bool      `json:"IsMarkedForDeletion"`
}

type BusinessPartnerBPRole struct {
	BusinessPartner            int        `json:"BusinessPartner"`
	BusinessPartnerRole    	   string	  `json:"BusinessPartnerRole"`
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

type BusinessPartnerSNS struct {
	BusinessPartner            int        `json:"BusinessPartner"`
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

type BusinessPartnerPersonOrganization struct {
	BusinessPartner            int        `json:"BusinessPartner"`
	IsMarkedForDeletion        *bool      `json:"IsMarkedForDeletion"`
}

type BusinessPartnerDocGeneralDoc struct {
	BusinessPartner				int     `json:"BusinessPartner"`
	DocType						*string `json:"DocType"`
	DocIssuerBusinessPartner	*int    `json:"DocIssuerBusinessPartner"`
}
