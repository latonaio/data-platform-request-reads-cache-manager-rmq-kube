package attendance

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type AttendanceReq struct {
	Header   Header   `json:"Attendance"`
	Accepter []string `json:"accepter"`
}

type Header struct {
	Attendance           int     `json:"Attendance"`
	AttendanceDate       *string `json:"AttendanceDate"`
	AttendanceTime       *string `json:"AttendanceTime"`
	Attender             *int    `json:"Attender"`
	AttendanceObjectType *string `json:"AttendanceObjectType"`
	AttendanceObject     *int    `json:"AttendanceObject"`
	Participation        *int    `json:"Participation"`
	Invitation           *int    `json:"Invitation"`
	CreationDate         *string `json:"CreationDate"`
	CreationTime         *string `json:"CreationTime"`
	IsCancelled          *bool   `json:"IsCancelled"`
}

func CreateAttendanceRequestHeader(
	requestPram *apiInputReader.Request,
	attendanceHeader *apiInputReader.AttendanceHeader,
) AttendanceReq {
	req := AttendanceReq{
		Header: Header{
			Attendance:  attendanceHeader.Attendance,
			IsCancelled: attendanceHeader.IsCancelled,
		},
		Accepter: []string{
			"Header",
		},
	}
	return req
}

func CreateAttendanceRequestHeadersByAttender(
	requestPram *apiInputReader.Request,
	attendanceHeader *apiInputReader.AttendanceHeader,
) AttendanceReq {
	req := AttendanceReq{
		Header: Header{
			Attender:    attendanceHeader.Attender,
			IsCancelled: attendanceHeader.IsCancelled,
		},
		Accepter: []string{
			"HeadersByAttender",
		},
	}
	return req
}

func CreateAttendanceRequestHeadersByEvent(
	requestPram *apiInputReader.Request,
	attendanceHeader *apiInputReader.AttendanceHeader,
) AttendanceReq {
	attendanceObjectType := "EVENT"

	req := AttendanceReq{
		Header: Header{
			AttendanceObjectType: &attendanceObjectType,
			AttendanceObject:     attendanceHeader.AttendanceObject,
			IsCancelled:          attendanceHeader.IsCancelled,
		},
		Accepter: []string{
			"HeadersByEvent",
		},
	}
	return req
}

func AttendanceReadsHeader(
	requestPram *apiInputReader.Request,
	input apiInputReader.Attendance,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_ATTENDANCE_SRV"
	aPIType := "reads"

	var request AttendanceReq

	request = CreateAttendanceRequestHeader(
		requestPram,
		&apiInputReader.AttendanceHeader{
			Attendance:  input.AttendanceHeader.Attendance,
			IsCancelled: input.AttendanceHeader.IsCancelled,
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
		requestPram,
	)

	return responseBody
}

func AttendanceReadsHeadersByAttender(
	requestPram *apiInputReader.Request,
	input apiInputReader.Attendance,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_ATTENDANCE_SRV"
	aPIType := "reads"

	var request AttendanceReq

	request = CreateAttendanceRequestHeadersByAttender(
		requestPram,
		&apiInputReader.AttendanceHeader{
			Attender:    input.AttendanceHeader.Attender,
			IsCancelled: input.AttendanceHeader.IsCancelled,
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
		requestPram,
	)

	return responseBody
}

func AttendanceReadsHeadersByEvent(
	requestPram *apiInputReader.Request,
	input apiInputReader.Attendance,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_ATTENDANCE_SRV"
	aPIType := "reads"

	var request AttendanceReq

	attendanceObjectType := "EVENT"

	request = CreateAttendanceRequestHeadersByEvent(
		requestPram,
		&apiInputReader.AttendanceHeader{
			AttendanceObjectType: &attendanceObjectType,
			AttendanceObject:     input.AttendanceHeader.AttendanceObject,
			IsCancelled:          input.AttendanceHeader.IsCancelled,
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
		requestPram,
	)

	return responseBody
}
