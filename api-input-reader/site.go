package apiInputReader

type Site struct {
	SiteHeader              *SiteHeader
	SitePartner             *SitePartner
	SiteAddress             *SiteAddress
	SiteDocHeaderDoc        *SiteDocHeaderDoc
}

type SiteHeader struct {
	Site                    int    `json:"Site"`
	IsReleased              *bool   `json:"IsReleased"`
	IsMarkedForDeletion     *bool  `json:"IsMarkedForDeletion"`
}

type SitePartner struct {
	Site				int		`json:"Site"`
	PartnerFunction		string	`json:"PartnerFunction"`
	BusinessPartner		int		`json:"BusinessPartner"`
}

type SiteAddress struct {
	Site			int		`json:"Site"`
	AddressID		int		`json:"AddressID"`
	LocalSubRegion 	*string `json:"LocalSubRegion"`
	LocalRegion 	*string `json:"LocalRegion"`
}

type SiteDocHeaderDoc struct {
	Site            		 int     `json:"Site"`
	DocType                  *string `json:"DocType"`
	DocIssuerBusinessPartner *int    `json:"DocIssuerBusinessPartner"`
}
