package apiresponses

import (
	"encoding/json"

	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"golang.org/x/xerrors"
)

type PlantRes struct {
	ConnectionKey       string       `json:"connection_key,omitempty"`
	Result              bool         `json:"result,omitempty"`
	RedisKey            string       `json:"redis_key,omitempty"`
	Filepath            string       `json:"filepath,omitempty"`
	APIStatusCode       int          `json:"api_status_code,omitempty"`
	RuntimeSessionID    string       `json:"runtime_session_id,omitempty"`
	BusinessPartnerID   *int         `json:"business_partner,omitempty"`
	ServiceLabel        string       `json:"service_label,omitempty"`
	APIType             string       `json:"api_type,omitempty"`
	Message             PlantMessage `json:"message,omitempty"`
	APISchema           string       `json:"api_schema,omitempty"`
	Accepter            []string     `json:"accepter,omitempty"`
	Deleted             bool         `json:"deleted,omitempty"`
	SQLUpdateResult     *bool        `json:"sql_update_result,omitempty"`
	SQLUpdateError      string       `json:"sql_update_error,omitempty"`
	SubfuncResult       *bool        `json:"subfunc_result,omitempty"`
	SubfuncError        string       `json:"subfunc_error,omitempty"`
	ExconfResult        *bool        `json:"exconf_result,omitempty"`
	ExconfError         string       `json:"exconf_error,omitempty"`
	APIProcessingResult *bool        `json:"api_processing_result,omitempty"`
	APIProcessingError  string       `json:"api_processing_error,omitempty"`
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

func CreatePlantRes(msg rabbitmq.RabbitmqMessage) (*PlantRes, error) {
	res := PlantRes{}
	err := json.Unmarshal(msg.Raw(), &res)
	if err != nil {
		return nil, xerrors.Errorf("unmarshal error: %w", err)
	}
	return &res, nil
}
