package apiOutputFormatter

type BusinessPartnerRole struct {
	BusinessPartnerRoleBusinessPartnerRole    []BusinessPartnerRoleBusinessPartnerRole    `json:"BusinessPartnerRoleBusinessPartnerRole"`
	BusinessPartnerRoleText                   []BusinessPartnerRoleText                   `json:"BusinessPartnerRoleText"`
	Accepter                                  []string                                    `json:"Accepter"`
}

type BusinessPartnerRoleBusinessPartnerRole struct {
	BusinessPartnerRole    string    `json:"BusinessPartnerRole"`
}

type BusinessPartnerRoleText struct {
	BusinessPartnerRole      string `json:"BusinessPartnerRole"`
	Language                 string `json:"Language"`
	BusinessPartnerRoleName  string `json:"BusinessPartnerRoleName"`
}
