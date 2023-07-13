package apiOutputFormatter

type BusinessPartner struct {
	BusinessPartnerGeneral		 []BusinessPartnerGeneral 		`json:"General"`
	BusinessPartnerDetailGeneral []BusinessPartnerDetailGeneral `json:"General"`
}

type BusinessPartnerGeneral struct {
	BusinessPartner     int    `json:"BusinessPartner"`
	BusinessPartnerName string `json:"BusinessPartnerName"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}

type BusinessPartnerDetailGeneral struct {
	BusinessPartnerFullName		*string	`json:"BusinessPartnerFullName"`
	Industry					*string	`json:"Industry"`
	LegalEntityRegistration		*string	`json:"LegalEntityRegistration"`
	Country						string	`json:"Country"`
	Language					*string	`json:"Language"`
	Currency					*string	`json:"Currency"`
	AddressID					*int	`json:"AddressID"`
	BusinessPartnerIsBlocked	*bool	`json:"BusinessPartnerIsBlocked"`
	CreationDate				string	`json:"CreationDate"`
	LastChangeDate				string	`json:"LastChangeDate"`
	IsMarkedForDeletion			*bool	`json:"IsMarkedForDeletion"`
}
