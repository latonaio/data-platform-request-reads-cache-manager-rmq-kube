package dpfm_api_output_formatter

type SupplyChainRelationshipDetailList struct {
	Header     SupplyChainRelationshipDetailHeader   `json:"Header"`
	Existences []SupplyChainRelationshipDetailExconf `json:"Existences"`
}

type SupplyChainRelationshipDetailExconf struct {
	Content string      `json:"Content"`
	Exist   *bool       `json:"Exist"`
	Param   interface{} `json:"Param"`
}

type SupplyChainRelationshipDetailHeader struct {
	Index int    `json:"Index"`
	Key   string `json:"Key"`
}
