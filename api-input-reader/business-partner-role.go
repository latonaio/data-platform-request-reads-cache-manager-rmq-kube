package apiInputReader

type BusinessPartnerRoleGlobal struct {
	BusinessPartnerRole     *BusinessPartnerRole
	BusinessPartnerRoleText *BusinessPartnerRoleText
}

type BusinessPartnerRole struct {
	BusinessPartnerRole string `json:"BusinessPartnerRole"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}

type BusinessPartnerRoleText struct {
	BusinessPartnerRole string `json:"BusinessPartnerRole"`
	Language            string `json:"Language"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}
