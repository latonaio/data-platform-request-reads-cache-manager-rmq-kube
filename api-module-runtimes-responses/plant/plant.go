package apiModuleRuntimesResponsesPlant

type PlantRes struct {
	Message  PlantMessage `json:"message,omitempty"`
	Accepter []string     `json:"accepter,omitempty"`
}

type PlantMessage struct {
	General         *PlantGeneral         `json:"General,omitempty"`
	Generals        *[]PlantGeneral       `json:"Generals,omitempty"`
	StorageLocation *PlantStorageLocation `json:"StorageLocation,omitempty"`
}

type Plant struct {
	ConnectionKey string `json:"connection_key,omitempty"`
	Result        bool   `json:"result,omitempty"`
	RedisKey      string `json:"redis_key,omitempty"`
	Filepath      string `json:"filepath,omitempty"`
	Product       string `json:"Product,omitempty"`
	APISchema     string `json:"api_schema,omitempty"`
	MaterialCode  string `json:"material_code,omitempty"`
	Deleted       string `json:"deleted,omitempty"`
}

type PlantGeneral struct {
	BusinessPartner      int                  `json:"BusinessPartner,omitempty"`
	Plant                string               `json:"Plant,omitempty"`
	PlantFullName        *string              `json:"PlantFullName,omitempty"`
	PlantName            *string              `json:"PlantName,omitempty"`
	Language             *string              `json:"Language,omitempty"`
	CreationDate         *string              `json:"CreationDate,omitempty"`
	CreationTime         *string              `json:"CreationTime,omitempty"`
	LastChangeDate       *string              `json:"LastChangeDate,omitempty"`
	LastChangeTime       *string              `json:"LastChangeTime,omitempty"`
	PlantFoundationDate  *string              `json:"PlantFoundationDate,omitempty"`
	PlantLiquidationDate *string              `json:"PlantLiquidationDate,omitempty"`
	SearchTerm1          *string              `json:"SearchTerm1,omitempty"`
	SearchTerm2          *string              `json:"SearchTerm2,omitempty"`
	PlantDeathDate       *string              `json:"PlantDeathDate,omitempty"`
	PlantIsBlocked       *bool                `json:"PlantIsBlocked,omitempty"`
	GroupPlantName1      *string              `json:"GroupPlantName1,omitempty"`
	GroupPlantName2      *string              `json:"GroupPlantName2,omitempty"`
	AddressID            *int                 `json:"AddressID,omitempty"`
	Country              *string              `json:"Country,omitempty"`
	TimeZone             *string              `json:"TimeZone,omitempty"`
	PlantIDByExtSystem   *string              `json:"PlantIDByExtSystem,omitempty"`
	IsMarkedForDeletion  *bool                `json:"IsMarkedForDeletion,omitempty"`
	StorageLocation      PlantStorageLocation `json:"StorageLocation,omitempty"`
}

type PlantStorageLocation struct {
	BusinessPartner              int     `json:"BusinessPartner,omitempty"`
	Plant                        string  `json:"Plant,omitempty"`
	StorageLocation              string  `json:"StorageLocation,omitempty"`
	StorageLocationFullName      *string `json:"StorageLocationFullName,omitempty"`
	StorageLocationName          *string `json:"StorageLocationName,omitempty"`
	CreationDate                 *string `json:"CreationDate,omitempty"`
	CreationTime                 *string `json:"CreationTime,omitempty"`
	LastChangeDate               *string `json:"LastChangeDate,omitempty"`
	LastChangeTime               *string `json:"LastChangeTime,omitempty"`
	SearchTerm1                  *string `json:"SearchTerm1,omitempty"`
	SearchTerm2                  *string `json:"SearchTerm2,omitempty"`
	StorageLocationIsBlocked     *bool   `json:"StorageLocationIsBlocked,omitempty"`
	GroupStorageLocationName1    *string `json:"GroupStorageLocationName1,omitempty"`
	GroupStorageLocationName2    *string `json:"GroupStorageLocationName2,omitempty"`
	StorageLocationIDByExtSystem *string `json:"StorageLocationIDByExtSystem,omitempty"`
	IsMarkedForDeletion          *bool   `json:"IsMarkedForDeletion,omitempty"`
}
