package dpfm_api_output_formatter

type BusinessPartnerList struct {
	BusinessPartners []BusinessPartner `json:"BusinessPartners"`
}

type BusinessPartner struct {
	BusinessPartner         int     `json:"BusinessPartner"`
	BusinessPartnerFullName *string `json:"BusinessPartnerFullName"`
	BusinessPartnerName     *string `json:"BusinessPartnerName"`
	IsMarkedForDeletion     *bool   `json:"IsMarkedForDeletion"`
}
