package apiInputReader

type StorageBin struct {
	StorageBinGeneral	    *StorageBinGeneral
}

type StorageBinGeneral struct {
	BusinessPartner            int        `json:"BusinessPartner"`
	Plant                      string     `json:"Plant"`
	StorageLocation            string     `json:"StorageLocation"`
    StorageBin                 string     `json:"StorageLocation"`
	IsMarkedForDeletion        *bool      `json:"IsMarkedForDeletion"`
}
