package apiModuleRuntimesResponsesSupplyChainRelationship

type SupplyChainRelationshipRes struct {
	Message SupplyChainRelationship `json:"message,omitempty"`
}

type SupplyChainRelationship struct {
	General *[]General `json:"Generals,omitempty"`
}

type General struct {
	SupplyChainRelationshipID int    `json:"SupplyChainRelationshipID"`
	Buyer                     int    `json:"Buyer"`
	Seller                    int    `json:"Seller"`
	CreationDate              string `json:"CreationDate"`
	LastChangeDate            string `json:"LastChangeDate"`
	IsMarkedForDeletion       *bool  `json:"IsMarkedForDeletion"`
}
