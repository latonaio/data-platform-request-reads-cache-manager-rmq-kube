package apiModuleRuntimesResponsesWorkCenter

type WorkCenterRes struct {
	Message WorkCenter `json:"message,omitempty"`
}

type WorkCenter struct {
	General *[]General `json:"Generals,omitempty"`
}

type General struct {
	WorkCenter                   	int     `json:"WorkCenter"`
	WorkCenterType               	string  `json:"WorkCenterType"`
	WorkCenterName               	string  `json:"WorkCenterName"`
	BusinessPartner              	int     `json:"BusinessPartner"`
	Plant                        	string  `json:"Plant"`
	WorkCenterCategory           	*string `json:"WorkCenterCategory"`
	WorkCenterResponsible        	*string `json:"WorkCenterResponsible"`
	SupplyArea                   	*string `json:"SupplyArea"`
	WorkCenterUsage              	*string `json:"WorkCenterUsage"`
	ComponentIsMarkedForBackflush	*bool   `json:"ComponentIsMarkedForBackflush"`
	WorkCenterLocation           	*string `json:"WorkCenterLocation"`
	CapacityCategory         		string  `json:"CapacityCategory"`
	CapacityQuantityUnit         	string  `json:"CapacityQuantityUnit"`
	CapacityQuantity         		float32 `json:"CapacityQuantity"`
	ValidityStartDate            	string  `json:"ValidityStartDate"`
	ValidityEndDate              	string  `json:"ValidityEndDate"`
	CreationDate            		string  `json:"CreationDate"`
	LastChangeDate              	string  `json:"LastChangeDate"`
	IsMarkedForDeletion          	*bool   `json:"IsMarkedForDeletion"`
}
