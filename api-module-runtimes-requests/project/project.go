package apiModuleRuntimesRequestsProject

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type ProjectReq struct {
	Project  Project   `json:"Project"`
	Projects []Project `json:"Projects"`
	Accepter []string  `json:"accepter"`
}

type Project struct {
	Project               int          `json:"Project"`
	ProjectCode           string       `json:"ProjectCode"`
	ProjectDescription    string       `json:"ProjectDescription"`
	OwnerBusinessPartner  int          `json:"OwnerBusinessPartner"`
	OwnerPlant            string       `json:"OwnerPlant"`
	ProjectProfile        *string      `json:"ProjectProfile"`
	ResponsiblePerson     *int         `json:"ResponsiblePerson"`
	ResponsiblePersonName *string      `json:"ResponsiblePersonName"`
	PlannedStartDate      *string      `json:"PlannedStartDate"`
	PlannedEndDate        *string      `json:"PlannedEndDate"`
	ActualStartDate       *string      `json:"ActualStartDate"`
	ActualEndDate         *string      `json:"ActualEndDate"`
	CreationDate          string       `json:"CreationDate"`
	LastChangeDate        string       `json:"LastChangeDate"`
	IsMarkedForDeletion   *bool        `json:"IsMarkedForDeletion"`
	WBSElement            []WBSElement `json:"WBSElement"`
}

type WBSElement struct {
	Project               int     `json:"Project"`
	WBSElement            int     `json:"WBSElement"`
	WBSElementCode        string  `json:"WBSElementCode"`
	WBSElementDescription string  `json:"WBSElementDescription"`
	BusinessPartner       int     `json:"BusinessPartner"`
	Plant                 string  `json:"Plant"`
	ResponsiblePerson     *int    `json:"ResponsiblePerson"`
	ResponsiblePersonName *string `json:"ResponsiblePersonName"`
	PlannedStartDate      *string `json:"PlannedStartDate"`
	PlannedEndDate        *string `json:"PlannedEndDate"`
	ActualStartDate       *string `json:"ActualStartDate"`
	ActualEndDate         *string `json:"ActualEndDate"`
	CreationDate          string  `json:"CreationDate"`
	LastChangeDate        string  `json:"LastChangeDate"`
	IsMarkedForDeletion   *bool   `json:"IsMarkedForDeletion"`
}

func CreateProjectRequestProjectsByProjects(
	requestPram *apiInputReader.Request,
	input []Project,
) ProjectReq {
	req := ProjectReq{
		Projects: input,
		Accepter: []string{
			"ProjectsByProjects",
		},
	}
	return req
}

func CreateProjectRequestProjects(
	requestPram *apiInputReader.Request,
	input apiInputReader.Project,
) ProjectReq {
	isMarkedForDeletion := false

	req := ProjectReq{
		Projects: []Project{
			{
				Project:             input.Project,
				IsMarkedForDeletion: &isMarkedForDeletion,
			},
		},
		Accepter: []string{
			"Projects",
		},
	}
	return req
}

func CreateProjectRequestWBSElement(
	requestPram *apiInputReader.Request,
	input WBSElement,
) ProjectReq {
	isMarkedForDeletion := false

	req := ProjectReq{
		Project: Project{
			Project: input.Project,
			WBSElement: []WBSElement{
				{
					WBSElement:          input.WBSElement,
					IsMarkedForDeletion: &isMarkedForDeletion,
				},
			},
		},
		Accepter: []string{
			"WBSElement",
		},
	}
	return req
}

func CreateProjectRequestWBSElements(
	requestPram *apiInputReader.Request,
	input Project,
) ProjectReq {
	isMarkedForDeletion := false

	req := ProjectReq{
		Project: Project{
			Project: input.Project,
			WBSElement: []WBSElement{
				{
					IsMarkedForDeletion: &isMarkedForDeletion,
				},
			},
		},
		Accepter: []string{
			"WBSElements",
		},
	}
	return req
}

func ProjectReadsProjects(
	requestPram *apiInputReader.Request,
	input apiInputReader.Project,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_PROJECT_SRV"
	aPIType := "reads"

	var request ProjectReq

	request = CreateProjectRequestProjects(
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

func ProjectReadsProjectsByProjects(
	requestPram *apiInputReader.Request,
	input []Project,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_PROJECT_SRV"
	aPIType := "reads"

	var request ProjectReq

	request = CreateProjectRequestProjectsByProjects(
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

func ProjectReadsWBSElement(
	requestPram *apiInputReader.Request,
	input WBSElement,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_PROJECT_SRV"
	aPIType := "reads"

	var request ProjectReq

	request = CreateProjectRequestWBSElement(
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

func ProjectReadsWBSElements(
	requestPram *apiInputReader.Request,
	input Project,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_PROJECT_SRV"
	aPIType := "reads"

	var request ProjectReq

	request = CreateProjectRequestWBSElements(
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
