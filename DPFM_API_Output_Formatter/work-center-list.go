package dpfm_api_output_formatter

type WorkCenterList struct {
	WorkCenters []WorkCenters `json:"WorkCenters"`
}

type WorkCenters struct {
	WorkCenter          int     `json:"WorkCenter"`
	PlantName           string  `json:"PlantName"`
	WorkCenterName      string  `json:"WorkCenterName"`
	WorkCenterLocation  *string `json:"WorkCenterLocation"`
	ValidityStartDate   *string `json:"ValidityStartDate"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
}
