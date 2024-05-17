package apiModuleRuntimesResponsesPointBalance

type PointBalanceRes struct {
	Message PointBalanceGlobal `json:"message,omitempty"`
}

type PointBalanceGlobal struct {
	PointBalance    *[]PointBalance    `json:"PointBalance,omitempty"`
}

type PointBalance struct {
	BusinessPartner		int			`json:"BusinessPartner"`
	PointSymbol			string		`json:"PointSymbol"`
	CurrentBalance		float32		`json:"CurrentBalance"`
	LimitBalance		*float32	`json:"LimitBalance"`
	CreationDate		string		`json:"CreationDate"`
	CreationTime		string		`json:"CreationTime"`
	LastChangeDate		string		`json:"LastChangeDate"`
	LastChangeTime		string		`json:"LastChangeTime"`
}
