package apiOutputFormatter

type WorkCenter struct {
	WorkCenterGeneral []WorkCenterGeneral `json:"Generals"`
}

type WorkCenterGeneral struct {
	WorkCenter                    int     `json:"WorkCenter"`
	WorkCenterName                string  `json:"WorkCenterName"`
	Plant                         string  `json:"Plant"`
	PlantName                     string  `json:"PlantName"`
	ComponentIsMarkedForBackflush *bool   `json:"ComponentIsMarkedForBackflush"`
	CapacityID                    *int    `json:"CapacityID"`
	CapacityCategory              *string `json:"CapacityCategory"`
	IsMarkedForDeletion           *bool   `json:"IsMarkedForDeletion"`
}
