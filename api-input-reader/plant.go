package apiInputReader

type Plant struct {
	PlantGeneral			*PlantGeneral
	PlantStorageLocations	*PlantStorageLocations
	PlantStorageLocation	*PlantStorageLocation
}

type PlantGeneral struct {
	BusinessPartner            int        `json:"BusinessPartner"`
	Plant                      string     `json:"Plant"`
	IsMarkedForDeletion        *bool      `json:"IsMarkedForDeletion"`
}

type PlantStorageLocations struct {
	BusinessPartner            int        `json:"BusinessPartner"`
	Plant                      string     `json:"Plant"`
	IsMarkedForDeletion        *bool      `json:"IsMarkedForDeletion"`
}

type PlantStorageLocation struct {
	BusinessPartner            int        `json:"BusinessPartner"`
	Plant                      string     `json:"Plant"`
	StorageLocation            string     `json:"StorageLocation"`
	IsMarkedForDeletion        *bool      `json:"IsMarkedForDeletion"`
}
