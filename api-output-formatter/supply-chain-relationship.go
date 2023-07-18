package apiOutputFormatter

type SupplyChainRelationship struct {
	SupplyChainRelationshipGeneral       []SupplyChainRelationshipGeneral       `json:"Generals"`
	SupplyChainRelationshipDetailGeneral []SupplyChainRelationshipDetailGeneral `json:"DetailGeneral"`
}

type SupplyChainRelationshipGeneral struct {
	SupplyChainRelationshipID int    `json:"SupplyChainRelationshipID"`
	Buyer                     int    `json:"Buyer"`
	BuyerName                 string `json:"BuyerName"`
	Seller                    int    `json:"Seller"`
	SellerName                string `json:"SellerName"`
	IsMarkedForDeletion       *bool  `json:"IsMarkedForDeletion"`
}

type SupplyChainRelationshipDetailGeneral struct {
	SupplyChainRelationshipID int    `json:"SupplyChainRelationshipID"`
	CreationDate              string `json:"CreationDate"`
	LastChangeDate            string `json:"LastChangeDate"`
	IsMarkedForDeletion       *bool  `json:"IsMarkedForDeletion"`
}
