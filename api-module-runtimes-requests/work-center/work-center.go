package apiInputReader

type WorkCenter struct {
	WorkCenterGeneral *WorkCenterGeneral
}

type WorkCenterGeneral struct {
	WorkCenter          int   `json:"WorkCenter"`
	IsMarkedForDeletion *bool `json:"IsMarkedForDeletion"`
}
apackage apiModuleRuntimesRequestsWorkCenter

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego"
)

type WorkCenterReq struct {
	General   General    `json:"WorkCenter"`
	Accepter []string    `json:"accepter"`
}

type General struct {
	WorkCenter                   	int      `json:"WorkCenter"`
	WorkCenterType               	*string  `json:"WorkCenterType"`
	WorkCenterName               	*string  `json:"WorkCenterName"`
	BusinessPartner              	*int     `json:"BusinessPartner"`
	Plant                        	*string  `json:"Plant"`
	WorkCenterCategory           	*string  `json:"WorkCenterCategory"`
	WorkCenterResponsible        	*string  `json:"WorkCenterResponsible"`
	SupplyArea                   	*string  `json:"SupplyArea"`
	WorkCenterUsage              	*string  `json:"WorkCenterUsage"`
	ComponentIsMarkedForBackflush	*bool    `json:"ComponentIsMarkedForBackflush"`
	WorkCenterLocation           	*string  `json:"WorkCenterLocation"`
	CapacityCategory         		*string  `json:"CapacityCategory"`
	CapacityQuantityUnit         	*string  `json:"CapacityQuantityUnit"`
	CapacityQuantity         		*float32 `json:"CapacityQuantity"`
	ValidityStartDate            	*string  `json:"ValidityStartDate"`
	ValidityEndDate              	*string  `json:"ValidityEndDate"`
	CreationDate            		*string  `json:"CreationDate"`
	LastChangeDate              	*string  `json:"LastChangeDate"`
	IsMarkedForDeletion          	*bool    `json:"IsMarkedForDeletion"`
}

func CreateWorkCenterRequestGeneralByBusinessPartner(
	requestPram *apiInputReader.Request,
	workCenterGeneral *apiInputReader.WorkCenterGeneral,
) WorkCenterReq {
	req := WorkCenterReq{
		General: General{
			BusinessPartner: 			requestPram.BusinessPartner,
			IsMarkedForDeletion:        workCenterGeneral.IsMarkedForDeletion,
		},
		Accepter: []string{
			"GeneralByBusinessPartner",
		},
	}
	return req
}

func CreateWorkCenterRequestItems(
	requestPram *apiInputReader.Request,
	workCenterItems *apiInputReader.WorkCenterItems,
) WorkCenterReq {
	req := WorkCenterReq{
		General: General{
			WorkCenter: workCenterItems.WorkCenter,
			Item: []Item{
				{
					IsMarkedForDeletion: workCenterItems.IsMarkedForDeletion,
				},
			},
		},
		Accepter: []string{
			"Items",
		},
	}
	return req
}

func WorkCenterReads(
	requestPram *apiInputReader.Request,
	input apiInputReader.WorkCenter,
	controller *beego.Controller,
	accepter string,
) []byte {
	aPIServiceName := "DPFM_API_WORK_CENTER_SRV"
	aPIType := "reads"

	var request WorkCenterReq

	if accepter == "GeneralByBusinessPartner" {
		request = CreateWorkCenterRequestGeneralByBusinessPartner(
			requestPram,
			&apiInputReader.WorkCenterGeneral{
				//IsMarkedForDeletion: input.WorkCenterGeneral.IsMarkedForDeletion,
			},
		)
	}

	if accepter == "Items" {
		request = CreateWorkCenterRequestItems(
			requestPram,
			&apiInputReader.WorkCenterItems{
				WorkCenter: input.WorkCenterItems.WorkCenter,
				//IsMarkedForDeletion: input.WorkCenterItems.IsMarkedForDeletion,
			},
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
