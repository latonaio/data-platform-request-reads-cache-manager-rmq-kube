package apiModuleRuntimesRequests

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego"
)

type PlantReq struct {
	General  PlantGeneral  `json:"Plant"`
	Generals PlantGenerals `json:"Plants"`
	Accepter []string      `json:"accepter"`
}

type PlantGeneral struct {
	BusinessPartner      int             `json:"BusinessPartner"`
	Plant                string          `json:"Plant"`
	PlantFullName        *string         `json:"PlantFullName"`
	PlantName            *string         `json:"PlantName"`
	Language             *string         `json:"Language"`
	CreationDate         *string         `json:"CreationDate"`
	CreationTime         *string         `json:"CreationTime"`
	LastChangeDate       *string         `json:"LastChangeDate"`
	LastChangeTime       *string         `json:"LastChangeTime"`
	PlantFoundationDate  *string         `json:"PlantFoundationDate"`
	PlantLiquidationDate *string         `json:"PlantLiquidationDate"`
	SearchTerm1          *string         `json:"SearchTerm1"`
	SearchTerm2          *string         `json:"SearchTerm2"`
	PlantDeathDate       *string         `json:"PlantDeathDate"`
	PlantIsBlocked       *bool           `json:"PlantIsBlocked"`
	GroupPlantName1      *string         `json:"GroupPlantName1"`
	GroupPlantName2      *string         `json:"GroupPlantName2"`
	AddressID            *int            `json:"AddressID"`
	Country              *string         `json:"Country"`
	TimeZone             *string         `json:"TimeZone"`
	PlantIDByExtSystem   *string         `json:"PlantIDByExtSystem"`
	IsMarkedForDeletion  *bool           `json:"IsMarkedForDeletion"`
	StorageLocation      StorageLocation `json:"StorageLocation"`
}

type PlantGenerals []struct {
	BusinessPartner *int    `json:"BusinessPartner"`
	Plant           *string `json:"Plant"`
	Language        *string `json:"Language"`
}

type StorageLocation struct {
	BusinessPartner              int     `json:"BusinessPartner"`
	Plant                        string  `json:"Plant"`
	StorageLocation              string  `json:"StorageLocation"`
	StorageLocationFullName      *string `json:"StorageLocationFullName"`
	StorageLocationName          *string `json:"StorageLocationName"`
	CreationDate                 *string `json:"CreationDate"`
	CreationTime                 *string `json:"CreationTime"`
	LastChangeDate               *string `json:"LastChangeDate"`
	LastChangeTime               *string `json:"LastChangeTime"`
	SearchTerm1                  *string `json:"SearchTerm1"`
	SearchTerm2                  *string `json:"SearchTerm2"`
	StorageLocationIsBlocked     *bool   `json:"StorageLocationIsBlocked"`
	GroupStorageLocationName1    *string `json:"GroupStorageLocationName1"`
	GroupStorageLocationName2    *string `json:"GroupStorageLocationName2"`
	StorageLocationIDByExtSystem *string `json:"StorageLocationIDByExtSystem"`
	IsMarkedForDeletion          *bool   `json:"IsMarkedForDeletion"`
}

func CreatePlantRequestGenerals(
	requestPram *apiInputReader.Request,
	plantGenerals PlantGenerals,
) PlantReq {
	req := PlantReq{
		Generals: plantGenerals,
		Accepter: []string{
			"Generals",
		},
	}

	return req
}

func PlantReads(
	requestPram *apiInputReader.Request,
	plantGenerals PlantGenerals,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_PLANT_SRV"
	aPIType := "reads"

	request := CreatePlantRequestGenerals(
		requestPram,
		plantGenerals,
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
