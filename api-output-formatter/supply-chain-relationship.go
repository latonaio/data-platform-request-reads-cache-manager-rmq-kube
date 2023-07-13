package apiOutputFormatter

type SupplyChainRelationship struct {
	SupplyChainRelationshipGeneral []SupplyChainRelationshipGeneral `json:"General"`
}

type SupplyChainRelationshipGeneral struct {
	SupplyChainRelationshipID int    `json:"SupplyChainRelationshipID"`
	Buyer                     int    `json:"Buyer"`
	BuyerName                 string `json:"BuyerName"`
	Seller                    int    `json:"Seller"`
	SellerName                string `json:"SellerName"`
	IsMarkedForDeletion       *bool  `json:"IsMarkedForDeletion"`
}
