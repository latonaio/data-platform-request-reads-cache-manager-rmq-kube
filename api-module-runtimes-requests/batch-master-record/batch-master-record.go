package apiModuleRuntimesRequestsBatchMasterRecord

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego"
)

type BatchMasterRecordReq struct {
	Batch    Batch    `json:"Batch"`
	Batches  []Batch  `json:"Batches"`
	Accepter []string `json:"accepter"`
}

type Batch struct {
	Product             string  `json:"Product"`
	BusinessPartner     int     `json:"BusinessPartner"`
	Plant               string  `json:"Plant"`
	Batch               string  `json:"Batch"`
	ValidityStartDate   *string `json:"ValidityStartDate"`
	ValidityStartTime   *string `json:"ValidityStartTime"`
	ValidityEndDate     *string `json:"ValidityEndDate"`
	ValidityEndTime     *string `json:"ValidityEndTime"`
	ManufactureDate     *string `json:"ManufactureDate"`
	CreationDate        *string `json:"CreationDate"`
	CreationTime        *string `json:"CreationTime"`
	LastChangeDate      *string `json:"LastChangeDate"`
	LastChangeTime      *string `json:"LastChangeTime"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
}

func CreateBatchMasterRecordRequestHeader(
	requestPram *apiInputReader.Request,
	batches []Batch,
) BatchMasterRecordReq {
	req := BatchMasterRecordReq{
		Batches: batches,
		Accepter: []string{
			"Batches",
		},
	}
	return req
}

func BatchMasterRecordReads(
	requestPram *apiInputReader.Request,
	batches []Batch,
	controller *beego.Controller,
	accepter string,
) []byte {
	aPIServiceName := "DPFM_API_BATCH_MASTER_RECORD_SRV"
	aPIType := "reads"

	var request BatchMasterRecordReq

	if accepter == "Batches" {
		request = CreateBatchMasterRecordRequestHeader(
			requestPram,
			batches,
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
