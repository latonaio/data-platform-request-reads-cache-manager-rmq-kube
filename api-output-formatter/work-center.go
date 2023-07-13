package apiOutputFormatter

type WorkCenter struct {
	WorkCenterGeneral []WorkCenterGeneral `json:"General"`
}

type WorkCenterGeneral struct {
	WorkCenter                    int     `json:"WorkCenter"`
	WorkCenterName                string  `json:"WorkCenterName"`
	Plant                         string  `json:"Plant"`
	PlantName                     string  `json:"PlantName"`
	ComponentIsMarkedForBackflush *string `json:"ComponentIsMarkedForBackflush"`
	CapacityInternalID            *int    `json:"CapacityInternalID"`
    CapacityCategoryCode          *string `json:"CapacityCategoryCode"`
	IsMarkedForDeletion           *bool   `json:"IsMarkedForDeletion"`
}
