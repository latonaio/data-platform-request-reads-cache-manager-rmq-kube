package apiInputReader

type PointBalanceGlobal struct {
	PointBalance *PointBalance
	ByShop       *ByShop
}

type PointBalance struct {
	BusinessPartner int      `json:"BusinessPartner"`
	PointSymbol     string   `json:"PointSymbol"`
	ByShop          []ByShop `json:"ByShop"`
}

type ByShop struct {
	BusinessPartner int    `json:"BusinessPartner"`
	PointSymbol     string `json:"PointSymbol"`
	Shop            int    `json:"Shop"`
}
