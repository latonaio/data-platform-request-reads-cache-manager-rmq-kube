package apiOutputFormatter

type Plant struct {
	PlantGeneral                    []PlantGeneral                    `json:"Generals"`
	PlantGeneralWithStorageLocation []PlantGeneralWithStorageLocation `json:"GeneralWithStorageLocation"`
	PlantStorageLocation            []PlantStorageLocation            `json:"StorageLocation"`
}

type PlantGeneral struct {
	BusinessPartner     int    `json:"BusinessPartner"`
	BusinessPartnerName string `json:"BusinessPartnerName"`
	Plant               string `json:"Plant"`
	PlantName           string `json:"PlantName"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}

type PlantGeneralWithStorageLocation struct {
	BusinessPartner     int    `json:"BusinessPartner"`
	BusinessPartnerName string `json:"BusinessPartnerName"`
	Plant               string `json:"Plant"`
	PlantName           string `json:"PlantName"`
}

type PlantStorageLocation struct {
	StorageLocation     string `json:"StorageLocation"`
	StorageLocationName string `json:"StorageLocationName"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}
