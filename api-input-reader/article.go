package apiInputReader

type Article struct {
	ArticleHeader        *ArticleHeader
	ArticlePartner       *ArticlePartner
	ArticleAddress       *ArticleAddress
	ArticleCounter       *ArticleCounter
	ArticleDocHeaderDoc  *ArticleDocHeaderDoc
}

type ArticleHeader struct {
	Article                  int     `json:"Article"`
	ArticleOwner             *int    `json:"ArticleOwner"`
	IsReleased               *bool   `json:"IsReleased"`
	IsMarkedForDeletion      *bool   `json:"IsMarkedForDeletion"`
}

type ArticlePartner struct {
	Article				int		`json:"Article"`
	PartnerFunction		string	`json:"PartnerFunction"`
	BusinessPartner		int		`json:"BusinessPartner"`
}

type ArticleAddress struct {
	Article			int		`json:"Article"`
	AddressID		int		`json:"AddressID"`
	LocalSubRegion 	*string `json:"LocalSubRegion"`
	LocalRegion 	*string `json:"LocalRegion"`
}

type ArticleCounter struct {
	Article                  int     `json:"Article"`
}

type ArticleDocHeaderDoc struct {
	Article            		 int     `json:"Article"`
	DocType                  *string `json:"DocType"`
	DocIssuerBusinessPartner *int    `json:"DocIssuerBusinessPartner"`
}
