package apiOutputFormatter

type BatchMasterRecord struct {
	BatchMasterRecordHeader       []BatchMasterRecordHeader       `json:"Header"`
}

type BatchMasterRecordHeader struct {
	Product				string  `json:"Product"`
	BusinessPartner		int		`json:"BusinessPartner"`
	Plant				string	`json:"Plant"`
    PlantName           string  `json:"PlantName"`
	Batch				string	`json:"Batch"`
	ValidityStartDate	string	`json:"ValidityStartDate"`
	ValidityStartTime	string	`json:"ValidityStartTime"`
	ValidityEndDate		string	`json:"ValidityEndDate"`
	ValidityEndTime		string	`json:"ValidityEndTime"`
	ManufactureDate		*string	`json:"ManufactureDate"`
	CreationDate         string  `json:"CreationDate"`
	CreationTime         string  `json:"CreationTime"`
	LastChangeDate       string  `json:"LastChangeDate"`
	LastChangeTime       string  `json:"LastChangeTime"`	
	IsMarkedForDeletion  *bool   `json:"IsMarkedForDeletion"`
}
