package dpfm_api_output_formatter

type SupplyChainRelationshipDetail struct {
	Header   SupplyChainRelationshipDetailHeader    `json:"Header"`
	Contents []SupplyChainRelationshipDetailContent `json:"Contents"`
}

type SupplyChainRelationshipDetailContent struct {
	Content string      `json:"Content"`
	Param   interface{} `json:"Param"`
}
