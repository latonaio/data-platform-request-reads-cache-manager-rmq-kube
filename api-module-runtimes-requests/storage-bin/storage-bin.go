package apiModuleRuntimesRequestsStorageBin

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type StorageBinReq struct {
	General  General   `json:"StorageBin"`
	Generals []General `json:"StorageBins"`
	Accepter []string  `json:"accepter"`
}

type General struct {
	BusinessPartner       int      `json:"BusinessPartner"`
	Plant                 string   `json:"Plant"`
	StorageLocation       string   `json:"StorageLocation"`
	StorageBin            string   `json:"StorageBin"`
	StorageBinDescription *string  `json:"StorageType"`
	XCoordinates          *float32 `json:"XCoordinates"`
	YCoordinates          *float32 `json:"YCoordinates"`
	ZCoordinates          *float32 `json:"ZCoordinates"`
	Capacity              *float32 `json:"Capacity"`
	CapacityUnit          *string  `json:"CapacityUnit"`
	CapacityWeight        *float32 `json:"CapacityWeight"`
	CapacityWeightUnit    *string  `json:"CapacityWeightUnit"`
	CreationDate          *string  `json:"CreationDate"`
	LastChangeDate        *string  `json:"LastChangeDate"`
	IsMarkedForDeletion   *bool    `json:"IsMarkedForDeletion"`
}

func CreateStorageBinRequestGeneralsByStorageBins(
	requestPram *apiInputReader.Request,
	input []General,
) StorageBinReq {
	req := StorageBinReq{
		Generals: input,
		Accepter: []string{
			"GeneralsByStorageBins",
		},
	}
	return req
}

func CreateStorageBinRequestGenerals(
	requestPram *apiInputReader.Request,
	input apiInputReader.StorageBin,
) StorageBinReq {
	isMarkedForDeletion := false

	req := StorageBinReq{
		General: General{
			IsMarkedForDeletion: &isMarkedForDeletion,
			//IsMarkedForDeletion: requestPram.IsMarkedForDeletion,
		},
		Accepter: []string{
			"Generals",
		},
	}
	return req
}

func StorageBinReadsGenerals(
	requestPram *apiInputReader.Request,
	input apiInputReader.StorageBin,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_STORAGE_BIN_SRV"
	aPIType := "reads"

	var request StorageBinReq

	request = CreateStorageBinRequestGenerals(
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
	)

	return responseBody
}

func StorageBinReadsByStorageBins(
	requestPram *apiInputReader.Request,
	input []General,
	controller *beego.Controller,
	accepter string,
) []byte {
	aPIServiceName := "DPFM_API_STORAGE_BIN_SRV"
	aPIType := "reads"

	var request StorageBinReq

	request = CreateStorageBinRequestGeneralsByStorageBins(
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
	)

	return responseBody
}
