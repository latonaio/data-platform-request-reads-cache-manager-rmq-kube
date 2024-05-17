package apiModuleRuntimesRequestsPointBalance

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type PointBalanceReq struct {
	PointBalance PointBalance `json:"PointBalance"`
	Accepter     []string     `json:"accepter"`
}

type PointBalance struct {
	BusinessPartner int      `json:"BusinessPartner"`
	PointSymbol     string   `json:"PointSymbol"`
	CurrentBalance  *float32 `json:"CurrentBalance"`
	LimitBalance    *float32 `json:"LimitBalance"`
	CreationDate    *string  `json:"CreationDate"`
	CreationTime    *string  `json:"CreationTime"`
	LastChangeDate  *string  `json:"LastChangeDate"`
	LastChangeTime  *string  `json:"LastChangeTime"`
}

func CreatePointBalanceRequestPointBalance(
	requestPram *apiInputReader.Request,
	pointBalance *apiInputReader.PointBalance,
) PointBalanceReq {
	req := PointBalanceReq{
		PointBalance: PointBalance{
			BusinessPartner: pointBalance.BusinessPartner,
			PointSymbol:     pointBalance.PointSymbol,
		},
		Accepter: []string{
			"PointBalance",
		},
	}
	return req
}

func CreatePointBalanceRequestPointBalancesByBusinessPartner(
	requestPram *apiInputReader.Request,
	pointBalance *apiInputReader.PointBalance,
) PointBalanceReq {
	req := PointBalanceReq{
		PointBalance: PointBalance{
			BusinessPartner: pointBalance.BusinessPartner,
		},
		Accepter: []string{
			"PointBalancesByBusinessPartner",
		},
	}
	return req
}

func PointBalanceReadsPointBalance(
	requestPram *apiInputReader.Request,
	input apiInputReader.PointBalanceGlobal,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_POINT_BALANCE_SRV"
	aPIType := "reads"

	var request PointBalanceReq

	request = CreatePointBalanceRequestPointBalance(
		requestPram,
		&apiInputReader.PointBalance{
			BusinessPartner: input.PointBalance.BusinessPartner,
			PointSymbol:     input.PointBalance.PointSymbol,
		},
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

func PointBalanceReadsPointBalancesByBusinessPartner(
	requestPram *apiInputReader.Request,
	input apiInputReader.PointBalanceGlobal,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_POINT_BALANCE_SRV"
	aPIType := "reads"

	var request PointBalanceReq

	request = CreatePointBalanceRequestPointBalancesByBusinessPartner(
		requestPram,
		&apiInputReader.PointBalance{
			BusinessPartner: input.PointBalance.BusinessPartner,
		},
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
