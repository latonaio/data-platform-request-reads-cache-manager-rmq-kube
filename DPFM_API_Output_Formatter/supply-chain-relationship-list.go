package dpfm_api_output_formatter

type SupplyChainRelationshipList struct {
	SupplyChainRelationship []SupplyChainRelationship `json:"SupplyChainRelationship"`
}

type SupplyChainRelationship struct {
	SupplyChainRelationshipID int    `json:"SupplyChainRelationshipID"`
	SellerName                string `json:"SellerName"`
	Seller                    *int   `json:"Seller"`
	BuyerName                 string `json:"BuyerName"`
	Buyer                     *int   `json:"Buyer"`
	IsMarkedForDeletion       *bool  `json:"IsMarkedForDeletion"`
}
