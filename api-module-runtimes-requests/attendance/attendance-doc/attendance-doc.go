package apiModuleRuntimesRequestsAttendanceDoc

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type AttendanceDocReq struct {
	BusinessPartnerID *int      `json:"business_partner"`
	HeaderDoc         HeaderDoc `json:"Attendance"`
	Accepter          []string  `json:"accepter"`
}

type HeaderDoc struct {
	Attendance               *int    `json:"Attendance"`
	DocType                  *string `json:"DocType"`
	FileExtension            *string `json:"FileExtension"`
	DocVersionID             *int    `json:"DocVersionID"`
	DocID                    *string `json:"DocID"`
	DocIssuerBusinessPartner *int    `json:"DocIssuerBusinessPartner"`
	FilePath                 *string `json:"FilePath"`
	FileName                 *string `json:"FileName"`
}

func CreateAttendanceDocRequestHeaderDoc(
	requestPram *apiInputReader.Request,
	headerDoc apiInputReader.AttendanceDocHeaderDoc,
) AttendanceDocReq {
	docIssuerBusinessPartner := 1001 // TODO 暫定対応
	docType := "IMAGE"

	req := AttendanceDocReq{
		HeaderDoc: HeaderDoc{
			Attendance: &headerDoc.Attendance,
			//DocType:                   headerDoc.DocType,
			DocType:                  &docType,
			DocIssuerBusinessPartner: &docIssuerBusinessPartner,
			//DocIssuerBusinessPartner:  headerDoc.DocIssuerBusinessPartner,
		},
		Accepter: []string{
			"HeaderDoc",
		},
	}
	return req
}

func AttendanceDocReads(
	requestPram *apiInputReader.Request,
	input apiInputReader.Attendance,
	controller *beego.Controller,
	accepter string,
) []byte {
	aPIServiceName := "DPFM_API_PARTICIPATION_DOC_SRV"
	aPIType := "reads"

	var request AttendanceDocReq

	if accepter == "HeaderDoc" {
		request = CreateAttendanceDocRequestHeaderDoc(
			requestPram,
			apiInputReader.AttendanceDocHeaderDoc{
				Attendance:               input.AttendanceDocHeaderDoc.Attendance,
				DocType:                  input.AttendanceDocHeaderDoc.DocType,
				DocIssuerBusinessPartner: input.AttendanceDocHeaderDoc.DocIssuerBusinessPartner,
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
