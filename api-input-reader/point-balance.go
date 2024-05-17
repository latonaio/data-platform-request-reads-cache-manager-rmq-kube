package apiInputReader

type PointBalanceGlobal struct {
	PointBalance    *PointBalance
}

type PointBalance struct {
	BusinessPartner		int			`json:"BusinessPartner"`
	PointSymbol			string		`json:"PointSymbol"`
}
