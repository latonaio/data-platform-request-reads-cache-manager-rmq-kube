package apiModuleRuntimesRequestsInspectionLotDoc

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type InspectionLotDocReq struct {
	BusinessPartnerID *int      `json:"business_partner"`
	HeaderDoc         HeaderDoc `json:"InspectionLot"`
	Accepter          []string  `json:"accepter"`
}

type HeaderDoc struct {
	InspectionLot            *int         `json:"InspectionLot"`
	DocType                  *string      `json:"DocType"`
	FileExtension            *string      `json:"FileExtension"`
	DocVersionID             *int         `json:"DocVersionID"`
	DocID                    *string      `json:"DocID"`
	DocIssuerBusinessPartner *int         `json:"DocIssuerBusinessPartner"`
	FilePath                 *string      `json:"FilePath"`
	FileName                 *string      `json:"FileName"`
	OperationDoc             OperationDoc `json:"OperationDoc"`
}

type OperationDoc struct {
	InspectionLot            *int    `json:"InspectionLot"`
	InspectionLotOperation   *int    `json:"InspectionLotOperation"`
	Operations               *int    `json:"Operations"`
	OperationsItem           *int    `json:"OperationsItem"`
	OperationID              *int    `json:"OperationID"`
	DocType                  *string `json:"DocType"`
	FileExtension            *string `json:"FileExtension"`
	DocVersionID             *int    `json:"DocVersionID"`
	DocID                    *string `json:"DocID"`
	DocIssuerBusinessPartner *int    `json:"DocIssuerBusinessPartner"`
	FilePath                 *string `json:"FilePath"`
	FileName                 *string `json:"FileName"`
}

func CreateInspectionLotDocRequestHeaderDoc(
	requestPram *apiInputReader.Request,
	headerDoc apiInputReader.InspectionLotDocHeaderDoc,
) InspectionLotDocReq {
	docIssuerBusinessPartner := 201 // TODO 暫定対応

	req := InspectionLotDocReq{
		HeaderDoc: HeaderDoc{
			DocType:                  headerDoc.DocType,
			DocIssuerBusinessPartner: &docIssuerBusinessPartner,
			//DocIssuerBusinessPartner: requestPram.BusinessPartner,
		},
		Accepter: []string{
			"HeaderDoc",
		},
	}
	return req
}

func CreateInspectionLotDocRequestOperationDoc(
	requestPram *apiInputReader.Request,
	operationDoc apiInputReader.InspectionLotDocOperationDoc,
) InspectionLotDocReq {
	req := InspectionLotDocReq{
		HeaderDoc: HeaderDoc{
			InspectionLot: &operationDoc.InspectionLot,
			OperationDoc: OperationDoc{
				InspectionLot:            &operationDoc.InspectionLot,
				Operations:               &operationDoc.Operations,
				OperationsItem:           &operationDoc.OperationsItem,
				OperationID:              &operationDoc.OperationID,
				DocType:                  &operationDoc.DocType,
				DocIssuerBusinessPartner: &operationDoc.DocIssuerBusinessPartner,
			},
		},
		Accepter: []string{
			"OperationDoc",
		},
	}
	return req
}

func InspectionLotDocReads(
	requestPram *apiInputReader.Request,
	input apiInputReader.InspectionLot,
	controller *beego.Controller,
	accepter string,
) []byte {
	aPIServiceName := "DPFM_API_INSPECTION_LOT_DOC_SRV"
	aPIType := "reads"

	var request InspectionLotDocReq

	if accepter == "HeaderDoc" {
		request = CreateInspectionLotDocRequestHeaderDoc(
			requestPram,
			apiInputReader.InspectionLotDocHeaderDoc{
				InspectionLot:            input.InspectionLotDocHeaderDoc.InspectionLot,
				DocType:                  input.InspectionLotDocHeaderDoc.DocType,
				DocIssuerBusinessPartner: input.InspectionLotDocHeaderDoc.DocIssuerBusinessPartner,
			},
		)
	}

	if accepter == "OperationDoc" {
		request = CreateInspectionLotDocRequestOperationDoc(
			requestPram,
			apiInputReader.InspectionLotDocOperationDoc{
				InspectionLot:            input.InspectionLotDocOperationDoc.InspectionLot,
				Operations:               input.InspectionLotDocOperationDoc.Operations,
				OperationsItem:           input.InspectionLotDocOperationDoc.OperationsItem,
				OperationID:              input.InspectionLotDocOperationDoc.OperationID,
				DocType:                  input.InspectionLotDocOperationDoc.DocType,
				DocIssuerBusinessPartner: input.InspectionLotDocOperationDoc.DocIssuerBusinessPartner,
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
	)

	return responseBody
}
