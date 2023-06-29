package dpfm_api_output_formatter

type SupplyChainRelationshipGeneralExconfList struct {
	General    SupplyChainRelationshipExconfGeneral `json:"General"`
	Existences []SupplyChainRelationshipExconfList  `json:"Existences"`
}

type SupplyChainRelationshipExconfList struct {
	Content string      `json:"Content"`
	Exist   bool        `json:"Exist"`
	Param   interface{} `json:"Param"`
}

type SupplyChainRelationshipExconfGeneral struct {
	Index int    `json:"Index"`
	Key   string `json:"Key"`
}
