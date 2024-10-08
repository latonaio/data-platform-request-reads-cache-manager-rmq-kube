package apiModuleRuntimesRequestsPlant

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type PlantReq struct {
	General  General   `json:"Plant"`
	Generals []General `json:"Plants"`
	Accepter []string  `json:"accepter"`
}

type General struct {
	BusinessPartner      int               `json:"BusinessPartner"`
	Plant                string            `json:"Plant"`
	PlantName            *string           `json:"PlantName"`
	PlantFullName        *string           `json:"PlantFullName"`
	Language             *string           `json:"Language"`
	PlantFoundationDate  *string           `json:"PlantFoundationDate"`
	PlantLiquidationDate *string           `json:"PlantLiquidationDate"`
	PlantDeathDate       *string           `json:"PlantDeathDate"`
	AddressID            *int              `json:"AddressID"`
	Country              *string           `json:"Country"`
	TimeZone             *string           `json:"TimeZone"`
	PlantIDByExtSystem   *string           `json:"PlantIDByExtSystem"`
	CreationDate         *string           `json:"CreationDate"`
	LastChangeDate       *string           `json:"LastChangeDate"`
	IsMarkedForDeletion  *bool             `json:"IsMarkedForDeletion"`
	StorageLocation      []StorageLocation `json:"StorageLocation"`
}

type StorageLocation struct {
	BusinessPartner              int     `json:"BusinessPartner"`
	Plant                        string  `json:"Plant"`
	StorageLocation              string  `json:"StorageLocation"`
	StorageLocationFullName      *string `json:"StorageLocationFullName"`
	StorageLocationName          *string `json:"StorageLocationName"`
	StorageLocationIDByExtSystem *string `json:"StorageLocationIDByExtSystem"`
	CreationDate                 *string `json:"CreationDate"`
	LastChangeDate               *string `json:"LastChangeDate"`
	IsMarkedForDeletion          *bool   `json:"IsMarkedForDeletion"`
}

func CreatePlantRequestGeneralsByPlants(
	requestPram *apiInputReader.Request,
	input []General,
) PlantReq {
	req := PlantReq{
		Generals: input,
		Accepter: []string{
			"GeneralsByPlants",
		},
	}
	return req
}

func CreatePlantRequestGenerals(
	requestPram *apiInputReader.Request,
	input apiInputReader.Plant,
) PlantReq {
	isMarkedForDeletion := false

	req := PlantReq{
		General: General{
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
		Accepter: []string{
			"Generals",
		},
	}
	return req
}

func CreatePlantRequestStorageLocations(
	requestPram *apiInputReader.Request,
	input apiInputReader.Plant,
) PlantReq {
	isMarkedForDeletion := false

	req := PlantReq{
		General: General{
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
		Accepter: []string{
			"StorageLocations",
		},
	}
	return req
}

func PlantReadsStorageLocations(
	requestPram *apiInputReader.Request,
	input apiInputReader.Plant,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_PLANT_SRV"
	aPIType := "reads"

	var request PlantReq

	request = CreatePlantRequestStorageLocations(
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

func PlantReadsGenerals(
	requestPram *apiInputReader.Request,
	input apiInputReader.Plant,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_PLANT_SRV"
	aPIType := "reads"

	var request PlantReq

	request = CreatePlantRequestGenerals(
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

func PlantReadsGeneralsByPlants(
	requestPram *apiInputReader.Request,
	input []General,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_PLANT_SRV"
	aPIType := "reads"

	var request PlantReq

	request = CreatePlantRequestGeneralsByPlants(
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
