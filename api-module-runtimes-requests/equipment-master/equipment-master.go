package apiModuleRuntimesRequestsEquipmentMaster

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type EquipmentMasterReq struct {
	BusinessPartnerID *int      `json:"business_partner"`
	General           General   `json:"EquipmentMaster"`
	Generals          []General `json:"EquipmentMasters"`
	Accepter          []string  `json:"accepter"`
}

type General struct {
	Equipment                       int                    `json:"Equipment"`
	ValidityStartDate               *string                `json:"ValidityStartDate"`
	ValidityEndDate                 *string                `json:"ValidityEndDate"`
	EquipmentName                   *string                `json:"EquipmentName"`
	EquipmentType                   *string                `json:"EquipmentType"`
	EquipmentCategory               *string                `json:"EquipmentCategory"`
	TechnicalObjectType             *string                `json:"TechnicalObjectType"`
	GrossWeight                     *float32               `json:"GrossWeight"`
	NetWeight                       *float32               `json:"NetWeight"`
	WeightUnit                      *string                `json:"WeightUnit"`
	SizeOrDimensionText             *string                `json:"SizeOrDimensionText"`
	InventoryNumber                 *string                `json:"InventoryNumber"`
	OperationStartDate              *string                `json:"OperationStartDate"`
	OperationStartTime              *string                `json:"OperationStartTime"`
	OperationEndDate                *string                `json:"OperationEndDate"`
	OperationEndTime                *string                `json:"OperationEndTime"`
	EquipmentStandardID             *string                `json:"EquipmentStandardID"`
	EquipmentIndustryStandardName   *string                `json:"EquipmentIndustryStandardName"`
	BarcodeType                     *string                `json:"BarcodeType"`
	AcquisitionDate                 *string                `json:"AcquisitionDate"`
	Manufacturer                    *int                   `json:"Manufacturer"`
	ManufacturerCountry             *string                `json:"ManufacturerCountry"`
	ConstructionYear                *int                   `json:"ConstructionYear"`
	ConstructionMonth               *int                   `json:"ConstructionMonth"`
	ConstructionDate                *string                `json:"ConstructionDate"`
	ManufacturerPartNmbr            *string                `json:"ManufacturerPartNmbr"`
	ManufacturerSerialNumber        *string                `json:"ManufacturerSerialNumber"`
	MaintenancePlantBusinessPartner *int                   `json:"MaintenancePlantBusinessPartner"`
	MaintenancePlant                *string                `json:"MaintenancePlant"`
	AssetLocation                   *string                `json:"AssetLocation"`
	AssetRoom                       *string                `json:"AssetRoom"`
	PlantSection                    *string                `json:"PlantSection"`
	WorkCenter                      *int                   `json:"WorkCenter"`
	Project                         *int                   `json:"Project"`
	WBSElement                      *int                   `json:"WBSElement"`
	MaintenancePlannerGroup         *string                `json:"MaintenancePlannerGroup"`
	CatalogProfile                  *string                `json:"CatalogProfile"`
	FunctionalLocation              *string                `json:"FunctionalLocation"`
	SuperordinateEquipment          *int                   `json:"SuperordinateEquipment"`
	EquipInstallationPositionNmbr   *string                `json:"EquipInstallationPositionNmbr"`
	BillOfMaterial                  *int                   `json:"BillOfMaterial"`
	BillOfMaterialItem              *int                   `json:"BillOfMaterialItem"`
	EquipmentIsAvailable            *bool                  `json:"EquipmentIsAvailable"`
	EquipmentIsInstalled            *bool                  `json:"EquipmentIsInstalled"`
	EquipHasSubOrdinateEquipment    *bool                  `json:"EquipHasSubOrdinateEquipment"`
	MasterFixedAsset                *string                `json:"MasterFixedAsset"`
	FixedAsset                      *string                `json:"FixedAsset"`
	CreationDate                    *string                `json:"CreationDate"`
	LastChangeDate                  *string                `json:"LastChangeDate"`
	IsMarkedForDeletion             *bool                  `json:"IsMarkedForDeletion"`
	OwnerBusinessPartner            []OwnerBusinessPartner `json:"OwnerBusinessPartner"`
	BusinessPartner                 []BusinessPartner      `json:"BusinessPartner"`
}

type OwnerBusinessPartner struct {
	Equipment            int     `json:"Equipment"`
	OwnerBusinessPartner int     `json:"OwnerBusinessPartner"`
	ValidityStartDate    string  `json:"ValidityStartDate"`
	ValidityEndDate      string  `json:"ValidityEndDate"`
	CreationDate         *string `json:"CreationDate"`
	LastChangeDate       *string `json:"LastChangeDate"`
	IsMarkedForDeletion  *bool   `json:"IsMarkedForDeletion"`
}

type BusinessPartner struct {
	Equipment                  int     `json:"Equipment"`
	EquipmentPartnerObjectNmbr int     `json:"EquipmentPartnerObjectNmbr"`
	BusinessPartner            *int    `json:"BusinessPartner"`
	PartnerFunction            *string `json:"PartnerFunction"`
	ValidityStartDate          *string `json:"ValidityStartDate"`
	ValidityEndDate            *string `json:"ValidityEndDate"`
	CreationDate               *string `json:"CreationDate"`
	LastChangeDate             *string `json:"LastChangeDate"`
	IsMarkedForDeletion        *bool   `json:"IsMarkedForDeletion"`
}

func CreateEquipmentMasterRequestGeneralsByEquipmentMasters(
	requestPram *apiInputReader.Request,
	generals []General,
) EquipmentMasterReq {
	req := EquipmentMasterReq{
		Generals: generals,
		Accepter: []string{
			"GeneralsByEquipmentMasters",
		},
	}
	return req
}

func CreateEquipmentMasterRequestGenerals(
	requestPram *apiInputReader.Request,
	input apiInputReader.EquipmentMaster,
) EquipmentMasterReq {
	//isMarkedForDeletion := false

	req := EquipmentMasterReq{
		General: General{
			//IsMarkedForDeletion: requestPram.IsMarkedForDeletion,
			//IsMarkedForDeletion: &isMarkedForDeletion,
		},
		Accepter: []string{
			"Generals",
		},
	}
	return req
}

func EquipmentMasterReads(
	requestPram *apiInputReader.Request,
	generals []General,
	controller *beego.Controller,
	accepter string,
) []byte {
	aPIServiceName := "DPFM_API_EQUIPMENT_MASTER_SRV"
	aPIType := "reads"

	var request EquipmentMasterReq

	if accepter == "GeneralsByEquipmentMasters" {
		request = CreateEquipmentMasterRequestGeneralsByEquipmentMasters(
			requestPram,
			generals,
		)
	}

	marshaledRequest, err := json.Marshal(request)
	if err != nil {
		services.HandleError(
			controller,
			err,
			nil,
		)
	}

	responseBody := services.Request(
		aPIServiceName,
		aPIType,
		ioutil.NopCloser(strings.NewReader(string(marshaledRequest))),
		controller,
		requestPram,
	)

	return responseBody
}

func EquipmentMasterReadsGenerals(
	requestPram *apiInputReader.Request,
	input apiInputReader.EquipmentMaster,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_EQUIPMENT_MASTER_SRV"
	aPIType := "reads"

	var request EquipmentMasterReq

	request = CreateEquipmentMasterRequestGenerals(
		requestPram,
		input,
	)

	marshaledRequest, err := json.Marshal(request)
	if err != nil {
		services.HandleError(
			controller,
			err,
			nil,
		)
	}

	responseBody := services.Request(
		aPIServiceName,
		aPIType,
		ioutil.NopCloser(strings.NewReader(string(marshaledRequest))),
		controller,
		requestPram,
	)

	return responseBody
}
