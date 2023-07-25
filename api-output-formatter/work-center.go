package apiOutputFormatter

type WorkCenter struct {
	WorkCenterGeneral []WorkCenterGeneral `json:"Generals"`
}

type WorkCenterGeneral struct {
	WorkCenter                    int     `json:"WorkCenter"`
	WorkCenterName                string  `json:"WorkCenterName"`
	Plant                         string  `json:"Plant"`
	PlantName                     string  `json:"PlantName"`
	WorkCenterLocation			  *string `json:"WorkCenterLocation"`
	CapacityCategory			  string  `json:"CapacityCategory"`
	ValidityStartDate             string  `json:"ValidityStartDate"`
	IsMarkedForDeletion           *bool   `json:"IsMarkedForDeletion"`
}
