package apiresponses

import (
	"encoding/json"

	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"golang.org/x/xerrors"
)

type EquipmentRes struct {
	ConnectionKey       string           `json:"connection_key"`
	Result              bool             `json:"result"`
	RedisKey            string           `json:"redis_key"`
	Filepath            string           `json:"filepath"`
	APIStatusCode       int              `json:"api_status_code"`
	RuntimeSessionID    string           `json:"runtime_session_id"`
	BusinessPartnerID   *int             `json:"business_partner"`
	ServiceLabel        string           `json:"service_label"`
	APIType             string           `json:"api_type"`
	Message             EquipmentMessage `json:"message"`
	APISchema           string           `json:"api_schema"`
	Accepter            []string         `json:"accepter"`
	Deleted             bool             `json:"deleted"`
	SQLUpdateResult     *bool            `json:"sql_update_result"`
	SQLUpdateError      string           `json:"sql_update_error"`
	SubfuncResult       *bool            `json:"subfunc_result"`
	SubfuncError        string           `json:"subfunc_error"`
	ExconfResult        *bool            `json:"exconf_result"`
	ExconfError         string           `json:"exconf_error"`
	APIProcessingResult *bool            `json:"api_processing_result"`
	APIProcessingError  string           `json:"api_processing_error"`
}

type EquipmentMessage struct {
	General              *[]EquipmentGeneral         `json:"General"`
	Address              *[]Address                  `json:"Address"`
	BusinessPartner      *[]EquipmentBusinessPartner `json:"BusinessPartner"`
	GeneralDoc           *[]EquipmentGeneralDoc      `json:"GeneralDoc"`
	OwnerBusinessPartner *[]OwnerBusinessPartner     `json:"OwnerBusinessPartner"`
}

type EquipmentGeneral struct {
	Equipment                       int      `json:"Equipment"`
	BusinessPartner                 *int     `json:"BusinessPartner"`
	ValidityStartDate               *string  `json:"ValidityStartDate"`
	ValidityEndDate                 *string  `json:"ValidityEndDate"`
	EquipmentName                   *string  `json:"EquipmentName"`
	EquipmentType                   *string  `json:"EquipmentType"`
	EquipmentCategory               *string  `json:"EquipmentCategory"`
	TechnicalObjectType             *string  `json:"TechnicalObjectType"`
	GrossWeight                     *float32 `json:"GrossWeight"`
	NetWeight                       *float32 `json:"NetWeight"`
	WeightUnit                      *string  `json:"WeightUnit"`
	SizeOrDimensionText             *string  `json:"SizeOrDimensionText"`
	InventoryNumber                 *string  `json:"InventoryNumber"`
	OperationStartDate              *string  `json:"OperationStartDate"`
	OperationStartTime              *string  `json:"OperationStartTime"`
	OperationEndDate                *string  `json:"OperationEndDate"`
	OperationEndTime                *string  `json:"OperationEndTime"`
	EquipmentStandardID             *string  `json:"EquipmentStandardID"`
	EquipmentIndustryStandardName   *string  `json:"EquipmentIndustryStandardName"`
	CountryOfOrigin                 *string  `json:"CountryOfOrigin"`
	CountryOfOriginLanguage         *string  `json:"CountryOfOriginLanguage"`
	BarcodeType                     *string  `json:"BarcodeType"`
	AcquisitionDate                 *string  `json:"AcquisitionDate"`
	Manufacturer                    *int     `json:"Manufacturer"`
	ManufacturedCountry             *string  `json:"ManufacturedCountry"`
	ConstructionYear                *int     `json:"ConstructionYear"`
	ConstructionMonth               *int     `json:"ConstructionMonth"`
	ManufacturerPartNmbr            *string  `json:"ManufacturerPartNmbr"`
	ManufacturerSerialNumber        *string  `json:"ManufacturerSerialNumber"`
	MaintenancePlantBusinessPartner int      `json:"MaintenancePlantBusinessPartner"`
	MaintenancePlant                string   `json:"MaintenancePlant"`
	AssetLocation                   *string  `json:"AssetLocation"`
	AssetRoom                       *string  `json:"AssetRoom"`
	PlantSection                    *string  `json:"PlantSection"`
	WorkCenter                      *string  `json:"WorkCenter"`
	Project                         *string  `json:"Project"`
	MaintenancePlannerGroup         *string  `json:"MaintenancePlannerGroup"`
	CatalogProfile                  *string  `json:"CatalogProfile"`
	FunctionalLocation              *string  `json:"FunctionalLocation"`
	SuperordinateEquipment          *string  `json:"SuperordinateEquipment"`
	EquipInstallationPositionNmbr   *string  `json:"EquipInstallationPositionNmbr"`
	EquipmentIsAvailable            *bool    `json:"EquipmentIsAvailable"`
	EquipmentIsInstalled            *bool    `json:"EquipmentIsInstalled"`
	EquipIsAllocToSuperiorEquip     *bool    `json:"EquipIsAllocToSuperiorEquip"`
	EquipHasSubOrdinateEquipment    *string  `json:"EquipHasSubOrdinateEquipment"`
	MasterFixedAsset                *string  `json:"MasterFixedAsset"`
	FixedAsset                      *string  `json:"FixedAsset"`
	CreationDate                    *string  `json:"CreationDate"`
	LastChangeDateTime              *string  `json:"LastChangeDateTime"`
	IsMarkedForDeletion             *bool    `json:"IsMarkedForDeletion"`
}

type Address struct {
	Equipment   int     `json:"Equipment"`
	AddressID   int     `json:"AddressID"`
	PostalCode  *string `json:"PostalCode"`
	LocalRegion *string `json:"LocalRegion"`
	Country     *string `json:"Country"`
	District    *string `json:"District"`
	StreetName  *string `json:"StreetName"`
	CityName    *string `json:"CityName"`
	Building    *string `json:"Building"`
	Floor       *int    `json:"Floor"`
	Room        *int    `json:"Room"`
}

type EquipmentBusinessPartner struct {
	Equipment                  int    `json:"Equipment"`
	EquipmentPartnerObjectNmbr int    `json:"EquipmentPartnerObjectNmbr"`
	BusinessPartner            int    `json:"BusinessPartner"`
	PartnerFunction            string `json:"PartnerFunction"`
	ValidityStartDate          string `json:"ValidityStartDate"`
	ValidityEndDate            string `json:"ValidityEndDate"`
	CreationDate               string `json:"CreationDate"`
	IsMarkedForDeletion        *bool  `json:"IsMarkedForDeletion"`
}

type EquipmentGeneralDoc struct {
	Equipment                int     `json:"Equipment"`
	DocType                  string  `json:"DocType"`
	DocVersionID             int     `json:"DocVersionID"`
	DocID                    string  `json:"DocID"`
	FileExtension            string  `json:"FileExtension"`
	FileName                 *string `json:"FileName"`
	FilePath                 *string `json:"FilePath"`
	DocIssuerBusinessPartner *int    `json:"DocIssuerBusinessPartner"`
}

type OwnerBusinessPartner struct {
	Equipment                int     `json:"Equipment"`
	OwnerBusinessPartner     int     `json:"OwnerBusinessPartner"`
	ValidityStartDate        string  `json:"ValidityStartDate"`
	ValidityEndDate          string  `json:"ValidityEndDate"`
	CreationDate             *string `json:"CreationDate"`
	BusinessPartnerEquipment *int    `json:"BusinessPartnerEquipment"`
	IsMarkedForDeletion      *bool   `json:"IsMarkedForDeletion"`
}

func CreateEquipmentRes(msg rabbitmq.RabbitmqMessage) (*EquipmentRes, error) {
	res := EquipmentRes{}
	err := json.Unmarshal(msg.Raw(), &res)
	if err != nil {
		return nil, xerrors.Errorf("unmarshal error: %w", err)
	}
	return &res, nil
}
