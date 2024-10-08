package apiOutputFormatter

type PointBalance struct {
	PointBalancePointBalance []PointBalancePointBalance `json:"PointBalancePointBalance"`
	PointBalanceByShop       []PointBalanceByShop       `json:"PointBalanceByShop"`
	MountPath                *string                    `json:"mount_path"`
	Accepter                 []string                   `json:"Accepter"`
}

type PointBalancePointBalance struct {
	BusinessPartner int      `json:"BusinessPartner"`
	PointSymbol     string   `json:"PointSymbol"`
	CurrentBalance  float32  `json:"CurrentBalance"`
	LimitBalance    *float32 `json:"LimitBalance"`
}

type PointBalanceByShop struct {
	BusinessPartner int      `json:"BusinessPartner"`
	PointSymbol     string   `json:"PointSymbol"`
	Shop            int      `json:"Shop"`
	CurrentBalance  float32  `json:"CurrentBalance"`
	LimitBalance    *float32 `json:"LimitBalance"`
}
