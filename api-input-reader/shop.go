package apiInputReader

type Shop struct {
	ShopHeader       *ShopHeader
	ShopPartner      *ShopPartner
	ShopAddress      *ShopAddress
	ShopDocHeaderDoc *ShopDocHeaderDoc
}

type ShopHeader struct {
	Shop                int   `json:"Shop"`
	ShopOwner           int   `json:"ShopOwner"`
	IsReleased          *bool `json:"IsReleased"`
	IsMarkedForDeletion *bool `json:"IsMarkedForDeletion"`
}

type ShopPartner struct {
	Shop            int    `json:"Shop"`
	PartnerFunction string `json:"PartnerFunction"`
	BusinessPartner int    `json:"BusinessPartner"`
}

type ShopAddress struct {
	Shop           int     `json:"Shop"`
	AddressID      int     `json:"AddressID"`
	LocalSubRegion *string `json:"LocalSubRegion"`
	LocalRegion    *string `json:"LocalRegion"`
}

type ShopDocHeaderDoc struct {
	Shop                     int     `json:"Shop"`
	DocType                  *string `json:"DocType"`
	DocIssuerBusinessPartner *int    `json:"DocIssuerBusinessPartner"`
}
