package apiOutputFormatter

type StorageBin struct {
	StorageBinGeneral []StorageBinGeneral `json:"Generals"`
}

type StorageBinGeneral struct {
	BusinessPartner     int    `json:"BusinessPartner"`
	BusinessPartnerName string `json:"BusinessPartnerName"`
	Plant               string `json:"Plant"`
	PlantName           string `json:"PlantName"`
	StorageLocation     string `json:"StorageLocation"`
	StorageLocationName string `json:"StorageLocationName"`
	StorageBin          string `json:"StorageBin"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}
