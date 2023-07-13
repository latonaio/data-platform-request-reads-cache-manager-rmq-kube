package apiInputReader

type SupplyChainRelationship struct {
	SupplyChainRelationshipGeneral *SupplyChainRelationshipGeneral
}

type SupplyChainRelationshipGeneral struct {
	SupplyChainRelationshipID int     `json:"SupplyChainRelationshipID"`
	Buyer                     *int    `json:"Buyer"`
	Seller                    *int    `json:"Seller"`
	CreationDate              *string `json:"CreationDate"`
	LastChangeDate            *string `json:"LastChangeDate"`
	IsMarkedForDeletion       *bool   `json:"IsMarkedForDeletion"`
}
